package server

import (
	"context"
	"fmt"
	"github.com/leshachaplin/grpc-server-library/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_AddBook(t *testing.T) {
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnInterface.Close()

	client := protocol.NewBookServiceClient(clientConnInterface)

	requestAddBook := &protocol.AddBookRequest{Book: &protocol.Book{
		Name:   "little prince1",
		Author: "st ekz",
		Genre:  "pos",
		Year:   110,
	}}

	responceAddBook, err := client.AddBook(context.Background(), requestAddBook)
	if err == nil {
		fmt.Printf("add book %s\n", requestAddBook.Book.Name)
	} else {
		t.Errorf("add claims is failed, got:%s  , want:%s ", err, responceAddBook )
	}

}
