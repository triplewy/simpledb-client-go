package main

import (
	"log"
	"path/filepath"
	"time"

	pb "github.com/triplewy/simpledb/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Client consists of client config and rpc client
type Client struct {
	Config *Config
	client pb.SimpleDbClient
}

// NewClient creates a new client to db
func NewClient(config *Config) *Client {
	c := new(Client)
	c.Config = config
	c.connect()
	return c
}

func (c *Client) connect() {
	creds, err := credentials.NewClientTLSFromFile(filepath.Join(c.Config.sslDir, "cert.pem"), "")
	if err != nil {
		log.Fatalf("could not create credentials: %v", err)
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = pb.NewSimpleDbClient(conn)
}
