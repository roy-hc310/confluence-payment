package main

import (
	"confluence-payment/core"
	"confluence-payment/core-internal/utils"
	"sync"

	payment_proto "confluence-payment/core-internal/protos/payment"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	err := utils.LoadGlobalEnv(".")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	application := core.NewApplication(utils.GlobalEnv)
	

	var wg sync.WaitGroup
    wg.Add(2)

	// gRPC Server
	go func () {
		defer wg.Done()

		lis, err := net.Listen("tcp", ":"+utils.GlobalEnv.GRPCPort)
		if err != nil {
			log.Fatal().Msg(err.Error())
			return
		}

		gRPCserver := grpc.NewServer()

		payment_proto.RegisterPaymentProtoServer(gRPCserver, application.PaymentProtoService)

		log.Info().Msg("gRPC server started on port " + utils.GlobalEnv.GRPCPort)
		err = gRPCserver.Serve(lis)
		if err != nil {
			log.Fatal().Msg(err.Error())
			return
		}
	}()
	
	// HTTP Server
	go func () {
		defer wg.Done()
		
		g := gin.Default()
		err = core.Router(g, application)
		if err != nil {
			log.Error().Msg(err.Error())
			return
		}

		g.Run(":" + utils.GlobalEnv.Port)
	}()

	wg.Wait()

}

// protoc --plugin=protoc-gen-ts_proto=".\\node_modules\\.bin\\protoc-gen-ts_proto.cmd" --ts_proto_out=./src ./proto/invoicer.proto
