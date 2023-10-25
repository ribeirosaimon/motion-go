package main

import (
	"context"
	"log"
	"net"

	"github.com/ribeirosaimon/motion-go/scraping-service/internal/scraping"
	"github.com/ribeirosaimon/motion-go/scraping-service/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedScrapingServiceServer
}

func (s server) GetCompany(ctx context.Context, code *pb.StockCode) (*pb.SummaryStock, error) {
	stock := scraping.GetStockSummary(code.Code)
	return &stock, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Canot create listener :%s ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterScrapingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
