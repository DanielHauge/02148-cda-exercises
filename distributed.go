package main

import (
	"fmt"
	space "github.com/DanielHauge/goSpace"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)



func DistributedServer(uri string){


	go func() {
		http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Printf("ping from client\n")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()


	s := space.NewSpace(uri)

	fmt.Printf("Server initialized for url: %v \n", uri)

	var caller string
	var funcName string
	var arg1t1 string
	var arg2t1 string
	var arg1t2 int
	var arg2t2 int
	for {
		fmt.Printf("blocking for get rpc-call\n")
		s.Get(&caller, "rpc-call", &funcName)
		fmt.Printf("server got rpc call request from: %v with function: %v \n", caller, funcName)

		switch funcName {
			case "concat":
				fmt.Printf("waiting for concat args\n")
				t, _ := s.Get(caller, "rpc-args", &arg1t1, &arg2t1)
				arg1t1 = t.GetFieldAt(2).(string)
				arg1t1 = t.GetFieldAt(3).(string)
				s.Put(caller, "rpc-result", arg1t1+" <=> "+arg2t1)
			case "plus":
				fmt.Printf("waiting for plus args.\n")
				t, _ := s.Get(caller, "rpc-args", &arg1t2, &arg2t2)
				arg1t2 = t.GetFieldAt(2).(int)
				arg2t2 = t.GetFieldAt(3).(int)
				s.Put(caller, "rpc-result", arg1t2+arg2t2)
		}
	}
}

func DistributedClient(uri string) {

	go func() {
		host,_,_ := NetworkArgs()
		for {
			resp, e := http.Get("http://"+host+":8080/ping")
			if e != nil {panic(e)}
			fmt.Printf("Pinged server and got back: %v\n",resp.StatusCode)
			<- time.After(10 * time.Second)
		}
	}()



	callerName := "Client"+strconv.Itoa(rand.Intn(1000))
	s := space.NewRemoteSpace(uri)
	fmt.Printf("%v initialized and connected to: %v \n", callerName, uri)



	a, b := rand.Intn(1000), rand.Intn(1000)
	s.Put(callerName, "rpc-call", "plus")
	s.Put(callerName, "rpc-args", a, b)
	var plusResult int
	fmt.Printf("Sent out rpc-call and args, now waiting for rpc-result.\n")
	s.Get(callerName, "rpc-result", &plusResult)
	fmt.Printf("%v called plus with: %v and %v and got %v back.\n", callerName, a, b, plusResult)

	x,y := "hejsa med dejsa fra", callerName
	s.Put(callerName, "rpc-call", "concat")
	s.Put(callerName, "rpc-args", x,y)
	var concatResult string
	s.Get(callerName, "rpc-result", &concatResult)
	fmt.Printf("%v called concat with %v and %v and got %v back.\n", callerName, x, y, concatResult)

}
