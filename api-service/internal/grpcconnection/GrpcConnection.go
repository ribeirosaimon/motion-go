package grpcconnection

import (
	"context"
	"fmt"
	"github.com/ribeirosaimon/motion-go/config/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

type connection struct {
	clientConn *grpc.ClientConn
}

var myGrpcConn *connection

func NewConnection(serverAddress, port string) (*connection, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", serverAddress, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the gRPC server: %v", err)
		return nil, err
	}
	if myGrpcConn == nil {
		myGrpcConn = &connection{clientConn: conn}
		return myGrpcConn, nil
	}
	return myGrpcConn, nil
}

func GetStock(code string, national bool) (pb.SummaryStock, error) {
	client := pb.NewScrapingServiceClient(myGrpcConn.clientConn)
	defer myGrpcConn.Close()

	stockCode := pb.StockCode{Code: code, National: national}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	company, err := client.GetCompany(ctx, &stockCode)
	if err != nil {
		return pb.SummaryStock{}, err
	}
	return *company, nil
}

func (c *connection) Close() {
	if c.clientConn != nil {
		c.clientConn.Close()
	}
}
