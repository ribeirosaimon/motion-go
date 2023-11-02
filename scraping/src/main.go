package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/magiconair/properties"
	"github.com/ribeirosaimon/motion-go/confighub/pb"
	"github.com/ribeirosaimon/motion-go/confighub/util"
	"github.com/ribeirosaimon/motion-go/scraping/internal/scraping"

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
	file := getProperties()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", file.GetString("server.port.src", "")))
	if err != nil {
		log.Fatalf("Canot create listener :%s ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterScrapingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func getProperties() *properties.Properties {
	propertiesFile := "config.properties"
	dir, _ := util.FindRootDir()

	return properties.MustLoadFile(fmt.Sprintf("%s/%s", dir, propertiesFile), properties.UTF8)

}
