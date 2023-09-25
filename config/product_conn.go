package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductRPC() (conn *grpc.ClientConn) {
	port := os.Getenv("PRODUCT_PORT")
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msg(err.Error())
		conn.Close()
	}

	return
}
