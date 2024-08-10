package main

import "fmt"

type FlyWithWind struct{}

func (f FlyWithWind) Fly() {
	fmt.Println("I am flying with Wind")
}

type FlyWithNoWind struct{}

func (f FlyWithNoWind) Fly() {
	fmt.Println("I am flying with No Wind")
}
