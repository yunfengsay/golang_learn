package main

import (
	"fmt"
)

type Animal interface {
	run()
	jump()
}

type Dog struct {
}

func (d Dog) run() {
	fmt.Println("A dog is runing")
}

func (d Dog) jump() {
	fmt.Println("A dog is jumping")
}

type Cat struct {
}

func (c Cat) run() {
	fmt.Println("A cat is running")
}
func (c Cat) jump() {
	fmt.Println("A cat is jumping")
}

func main() {
	var i Animal
	i = Dog{}
	i.jump()
	i = Cat{}
	i.run()
}
