package main

import (
	"log"
	"reflect"
)

func main() {
	a := 10
	b := &a
	log.Println(a,b,&b)
	of := reflect.TypeOf(b).String()
	log.Printf(of)
}
