package server

import (
	"context"
	"fmt"
	"github.com/leshachaplin/grpc-server-library/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_GetBookByName(t *testing.T) {
	opts := grpc.WithInsecure()
	clientConnInterface, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Error(err)
	}
	defer clientConnInterface.Close()

	client := protocol.NewBookServiceClient(clientConnInterface)

	requestGetBookByName := &protocol.GetBookByNameRequest{Name: "little prince"}

	responseBookByName, err := client.GetBookByName(context.Background(), requestGetBookByName)
	if err == nil {
		fmt.Printf("add book %v\n", responseBookByName.Book.Year)
	} else {
		t.Errorf("add claims is failed, got:%s  , want:%s ", err, responseBookByName)
	}

}


