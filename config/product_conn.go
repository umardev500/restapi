package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductRPC() (conn *grpc.ClientConn) {
	host := os.Getenv("PRODUCT_HOST")
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msg(err.Error())
		conn.Close()
	}

	return
}
