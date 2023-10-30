package grpcconnection

import (
	"context"
	"fmt"
	"time"

	"github.com/ribeirosaimon/motion-go/config/pb"
	"github.com/ribeirosaimon/motion-go/internal/config"

	"google.golang.org/grpc"
)

type connection struct {
	clientConn *grpc.ClientConn
}

var myGrpcConn *connection

func NewConnection() (*connection, error) {
	configurations := config.GetConfigurations()

	if myGrpcConn == nil {
		myGrpcConn = &connection{}
	}
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			myGrpcConn.clientConn, err = grpc.DialContext(ctx,
				fmt.Sprintf("%s:%s", configurations.GetString("grpc.host", ""),
					configurations.GetString("grpc.port", "")), grpc.WithInsecure())
			if err == nil {
				return myGrpcConn, nil
			}
		}
	}
}

func GetStock(code string, national bool) (pb.SummaryStock, error) {
	newConnection, err := NewConnection()
	if err != nil {
		panic(err)
	}
	client := pb.NewScrapingServiceClient(newConnection.clientConn)
	defer func() {
		newConnection.Close()
	}()
	stockCode := pb.StockCode{Code: code, National: national}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
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
