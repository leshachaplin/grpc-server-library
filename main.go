package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/jmoiron/sqlx"
	"github.com/leshachaplin/grpc-server-library/internal/config"
	"github.com/leshachaplin/grpc-server-library/internal/repository"
	"github.com/leshachaplin/grpc-server-library/internal/server"
	"github.com/leshachaplin/grpc-server-library/internal/service"
	"github.com/leshachaplin/grpc-server-library/protocol"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfig()
	connStr := "user=su password=su dbname=book sslmode=disable"
	session, err := mgo.Dial("mongodb://127.0.0.1")
	db, err := sqlx.Open(cfg.DbDriver, connStr)
	if err != nil {
		log.Fatal("Canno't connect to database", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(cfg.GrpcPort))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", cfg.Port)

	done, cnsl := context.WithCancel(context.Background())
	bookRepo := repository.NewBookRepository(*db)

	r := service.New(*bookRepo)
	servicesRepo := repository.NewRepositoryOfServiceState(session.DB("services"))
	err = servicesRepo.AddService(done, "books", cfg.GrpcPort)
	if err != nil {
		log.Error(err)
	}
	s := grpc.NewServer()
	srv := &server.Server{Rpc: *r}
	authServiceService := protocol.BookServiceService{
		GetAllBooks:     srv.GetAllBooks,
		GetBookByName:   srv.GetBookByName,
		GetBookByAuthor: srv.GetBookByAuthor,
		AddBook:         srv.AddBook,
		DeleteBook:      srv.DeleteBook,
	}
	protocol.RegisterBookServiceService(s, &authServiceService)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("server is not connect %s", err)
		}
	}()

	for {
		select {
		case <-c:
			cnsl()
			if err := db.Close(); err != nil {
				log.Errorf("database not closed %s", err)
			}

			log.Info("Cansel is succesful")
			close(c)
			return
		}
	}
}
