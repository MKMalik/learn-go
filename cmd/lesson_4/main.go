package main

import "fmt"


type engine interface {
	kmLeft() uint8;
}

type gasEngine struct {
	mileage uint8
	liter  uint8
}

func (engine gasEngine) kmLeft() uint8 {
	return engine.liter * engine.mileage
}

type electricEngine struct {
	kmpkWh uint8 // km per kWh
	kWh uint8
}

func (e electricEngine) kmLeft() uint8 {
	return e.kWh * e.kmpkWh;
}

func canMakeIt(e engine, km uint8) {
	if km <= e.kmLeft() {
		fmt.Println("Yes, it can make it there!")
	} else {
		fmt.Println("Can't make it there. Need to fuel/charge")
	}
}

func main() {

	var myGasEngine gasEngine
	myGasEngine.mileage = 60
	myGasEngine.liter = 2

	var myElectricEngine electricEngine
	myElectricEngine.kWh = 5
	myElectricEngine.kmpkWh = 10


	canMakeIt(myGasEngine, 100)
	canMakeIt(myElectricEngine, 50)

}
