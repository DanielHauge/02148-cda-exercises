package main

import (
	"fmt"
	"github.com/DanielHauge/goSpace"
)

func TestExamples() {
	firstSpace := gospace.NewSpace("space")
	firstSpace.Put("Hello there")
	var message string
	t, _ := firstSpace.Get(&message)
	//fmt.Println(t.Fields())
	fmt.Println(t.String())
	fmt.Println(message)


}

func TestExamples2(){
	sSpace := gospace.NewSpace("tuplesSpace")
	sSpace.Put(gospace.CreateTuple("hejsa", 123))
	//var message gospace.Tuple
	//t, _ := sSpace.Get(&message)
	//fmt.Println(t.Fields())

	sSpace.Put(gospace.CreateTuple("test", 123))
	sSpace.Put(gospace.CreateTuple("test", 123))
	sSpace.Put(gospace.CreateTuple("test", 123))
	sSpace.Put(gospace.CreateTuple("test", 123))
	sSpace.Put(gospace.CreateTuple("test", 123))

	/*
	v, _ := sSpace.GetAll(gospace.CreateTuple("test", 123))
	for _, tup := range v {
		fmt.Println(tup.Fields())
	}

	 */
}