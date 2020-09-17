package server

import (
	"context"
	"fmt"
	"github.com/leshachaplin/grpc-server-library/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_GetAllBooks(t *testing.T) {
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnInterface.Close()

	client := protocol.NewBookServiceClient(clientConnInterface)

	requestGetAllBooks := &protocol.EmptyRequest{}

	responseBooks, err := client.GetAllBooks(context.Background(), requestGetAllBooks)
	if err == nil {
		fmt.Printf("add book %s\n", responseBooks.Books[2].Author)
	} else {
		t.Errorf("add claims is failed, got:%s  , want:%s ", err, responseBooks.Books[2].Author)
	}

}


