package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/abiosoft/ishell.v1"
)

var addr string

func init() {
	flag.StringVar(&addr, "c", "", "Set join address, if any")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <node-address> \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func printHelp(shell *ishell.Shell) {
	shell.Println("Commands:")
	shell.Println(" - help                    Prints this help message")
	shell.Println(" - read <key> <fields>     Read")
	shell.Println(" - scan <key> <fields>     Scan")
	shell.Println(" - update <key> <entry>    Update")
	shell.Println(" - write <key> <entry>     Write")
	shell.Println(" - delete <key>            Delete")
	shell.Println(" - exit                    Exit CLI")
}

func main() {
	flag.Parse()
	client := NewClient(addr)
	shell := ishell.New()
	printHelp(shell)

	shell.Register("help", func(args ...string) (string, error) {
		printHelp(shell)
		return "", nil
	})

	shell.Register("read", func(args ...string) (string, error) {
		if len(args) != 2 {
			return "", errors.New("Args should be length 2")
		}
		key := args[0]
		attributes := strings.Split(args[1], ",")
		reply, err := client.Read(key, attributes)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Key: %v, Attributes: %v", reply.Key, reply.Attributes), nil
	})

	shell.Register("scan", func(args ...string) (string, error) {
		if len(args) != 2 {
			return "", errors.New("Args should be length 2")
		}
		key := args[0]
		attributes := strings.Split(args[1], ",")
		reply, err := client.Scan(key, attributes)
		if err != nil {
			return "", err
		}
		result := []string{}
		for _, entry := range reply.Entries {
			result = append(result, fmt.Sprintf("Key: %v, Attributes: %v", entry.Key, entry.Attributes))
		}
		return strings.Join(result, "\n"), nil
	})

	// shell.Register("update", func(args ...string) (string, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// 	defer cancel()
	// 	if len(args) != 2 {
	// 		return "", errors.New("Args should be length 2")
	// 	}
	// 	fields := strings.Split(args[1], ",")
	// 	values := []*pb.Field{}
	// 	for _, field := range fields {
	// 		value := strings.Split(field, ":")
	// 		values = append(values, &pb.Field{
	// 			Name:  value[0],
	// 			Value: []byte(value[1]),
	// 		})
	// 	}
	// 	reply, err := c.UpdateRPC(ctx, &pb.Entry{Key: args[0], Values: values})
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	out, err := json.Marshal(reply)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	return string(out), nil
	// })

	shell.Register("write", func(args ...string) (string, error) {
		if len(args) != 2 {
			return "", errors.New("Args should be length 2")
		}
		key := args[0]
		values, err := inputToValues(args[1])
		if err != nil {
			return "", err
		}
		err = client.Write(key, values)
		if err != nil {
			return "fail", err
		}
		return "success", nil
	})

	shell.Register("delete", func(args ...string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("Args should be length 1")
		}
		key := args[0]
		err := client.Delete(key)
		if err != nil {
			return "fail", err
		}
		return "success", nil
	})

	shell.Register("exit", func(args ...string) (string, error) {
		shell.Stop()
		return "", nil
	})

	shell.Start()
}
