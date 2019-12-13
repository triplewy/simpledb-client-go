package main

import (
	"context"
	"time"

	pb "github.com/triplewy/simpledb/grpc"
)

func (c *Client) Read(key string, attributes []string) (*pb.Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	return c.client.ReadRPC(ctx, &pb.ReadMsg{Key: key, Attributes: attributes})
}

func (c *Client) Scan(key string, attributes []string) (*pb.EntriesMsg, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	return c.client.ScanRPC(ctx, &pb.ReadMsg{Key: key, Attributes: attributes})
}

func (c *Client) Write(key string, values []*pb.Attribute) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	_, err := c.client.InsertRPC(ctx, &pb.Entry{Key: key, Attributes: values})
	return err
}

func (c *Client) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	_, err := c.client.DeleteRPC(ctx, &pb.KeyMsg{Key: key})
	return err
}
