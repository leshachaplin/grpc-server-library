package server

import (
	"context"
	"fmt"
	"github.com/leshachaplin/grpc-server-library/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_GetBookByAuthor(t *testing.T) {
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnInterface.Close()

	client := protocol.NewBookServiceClient(clientConnInterface)

	requestGetBookByAuthor := &protocol.GetBooksByAuthorRequest{Author: "leshiy"}

	responseBookByAuthor, err := client.GetBookByAuthor(context.Background(), requestGetBookByAuthor)
	if err == nil {
		fmt.Printf("add book %s\n", responseBookByAuthor.Book.Name)
	} else {
		t.Errorf("add claims is failed, got:%s  , want:%s ", err, responseBookByAuthor)
	}

}

