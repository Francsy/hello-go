package main

import "fmt"

type gasEngine struct {
	kmpliter uint8
	liters   uint8
}

type electricEngine struct {
	kmpkwh uint8
	kwh    uint8
}

func (e gasEngine) kmLeft() uint8 {
	return e.kmpliter * e.liters
}

func (e electricEngine) kmLeft() uint8 {
	return e.kmpkwh * e.kwh
}

type engine interface { // An interface can take the function of every type
	kmLeft() uint8
}

func canMakeIt(e engine, kilometers uint8) { // we can use interface engine parameter to make the func available for two types oof vehicles
	if kilometers <= e.kmLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You need to fuel the vehicle")
	}
}

func mainExample() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)
}
