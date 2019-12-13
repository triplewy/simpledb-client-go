package main

import (
	"fmt"
	"strings"

	pb "github.com/triplewy/simpledb/grpc"
)

func inputToValues(input string) (values []*pb.Attribute, err error) {
	arr := strings.Split(input, ",")
	for _, item := range arr {
		itemArr := strings.Split(item, ":")
		if len(itemArr) != 2 {
			return nil, fmt.Errorf("invalid value format: %v", item)
		}
		value := &pb.Attribute{
			Name:  itemArr[0],
			Type:  pb.Attribute_STRING,
			Value: []byte(itemArr[1]),
		}
		values = append(values, value)
	}
	return values, nil
}
