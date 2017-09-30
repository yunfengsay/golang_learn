package main

import "fmt"

func panicAndRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("unable to connect to database")
}

type User struct {
	name string
}

func print(s interface{}) {
	fmt.Println(s)
}
func (u User) addString(s string) {
	print(u.name + s)
	// return u.name + s
}

type SetName interface {
	Change
}

func (u User) Change(name string) {
	u.name = name
}
func Change() {
	print("asdf")
}

func main() {
	panicAndRecover()
	u := User{}
	u.name = "yunfeng"
	print(u)
	// change()
	// u.addString("s")
	print("I need to run the statement")
}
