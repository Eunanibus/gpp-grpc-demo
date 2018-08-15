package main

import (
	"net"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"github.com/eunanibus/gpp-grpc-demo/internal/app/gpp-grpc-demo/db"
	"github.com/eunanibus/gpp-grpc-demo/internal/app/gpp-grpc-demo/service"
	"github.com/eunanibus/gpp-grpc-demo/proto/auth"
)

func main() {
	grpcServer := grpc.NewServer()

	dbm := db.NewDatabaseManager()
	usrService := service.NewUsersService(dbm)

	// ProtoBuf generated files will allow you to register your service with generated functions
	auth.RegisterMyExampleServiceServer(grpcServer, usrService)
	defer grpcServer.GracefulStop()

	gRPCServerPort := ":8000"
	listener, err := net.Listen("tcp", gRPCServerPort)
	log.Infof("Server listening on port: %s...", gRPCServerPort)
	if err != nil {
		log.WithError(err).Panic("failed to listen on port - %s", gRPCServerPort)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.WithError(err).Panic("gRPC server stopped serving requests")
	}
}
