package main

import (
	"log"
	"net"

	"github.com/MaximKlimenko/gw-exchanger/internal/config"
	"github.com/MaximKlimenko/gw-exchanger/internal/exchanger"
	"github.com/MaximKlimenko/gw-exchanger/internal/storages/postgres"
	"google.golang.org/grpc"

	pb "github.com/MaximKlimenko/proto-exchange/exchange"
)

func main() {
	cfg := config.LoadConfig()

	db, err := postgres.NewConnection(cfg)
	if err != nil {
		log.Fatal("could not load the database")
	}

	server := &exchanger.ExchangeServiceServer{DB: db}

	grpcServer := grpc.NewServer()
	pb.RegisterExchangeServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}

	log.Printf("\033[1;32mgRPC сервер запущен на порту %s\033[0m", cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка работы gRPC сервера: %v", err)
	}
}
