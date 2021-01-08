package main

import (
	"fmt"
	"github.com/DanielHauge/goSpace"
	"math/rand"
	"time"
)

func ProducerConsumer(){
	s := gospace.NewSpace("concurrency")

	go producer(&s)

	go consumer("1",&s)
	go consumer("2",&s)
	go consumer("3",&s)
	go consumer("4",&s)
	<- time.After(20*time.Second)
}

func producer (space *gospace.Space){
	for {
		<- time.After(time.Second)
		t := rand.Intn(10)
		space.Put(t)
	}
}

func consumer (name string,space *gospace.Space){
	for {
		var g int
		t, _ := space.Get(&g)
		fmt.Printf("consumer '%v' will work for: %v seconds \n", name, t.GetFieldAt(0).(int))
		<- time.After(time.Duration(int64(t.GetFieldAt(0).(int))*int64(time.Second)))
		fmt.Printf("consumer: '%v' done! \n", name)
	}
}