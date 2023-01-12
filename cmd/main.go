package main

import (
	"log"
	"os"

	"github.com/maan19/hex/internal/adapters/app/api"
	"github.com/maan19/hex/internal/adapters/core/arithmetic"
	rpc "github.com/maan19/hex/internal/adapters/framework/left/grpc"
	"github.com/maan19/hex/internal/adapters/framework/right/db"
	"github.com/maan19/hex/internal/ports"
)

func main() {
	var err error
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initialize database driver: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbaseAdapter)
	gRPCAdapter = rpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()

}
