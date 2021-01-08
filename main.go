package main

import (
	"fmt"
	"os"
)

func main() {

	//l1.TestExamples()
	//l1.TestExamples2()
	//l2.ProducerConsumer()

	uri := SpaceUri()

	if os.Getenv("type") == "server" {
		DistributedServer(uri)
	} else {
		DistributedClient(uri)
	}

	fmt.Println("end of execution")
	os.Exit(0)
}

func NetworkArgs() (string,string,string) {
	host := withDefault(os.Getenv("host"), "localhost")
	port := withDefault(os.Getenv("port"), "31415")
	space := withDefault(os.Getenv("space"), "distributed")
	return host, port, space
}

func SpaceUri() string{
	host, port, space := NetworkArgs()
	return "tcp://" + host + ":" + port + "/" + space
}

func withDefault(s string, def string) string{
	if len(s) == 0 { return def }
	return s
}
