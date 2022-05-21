package main

import (
	"fmt"
)

const (
	dog    = "dog"
	cat    = "cat"
	mouse  = "mouse"
	spider = "spider"
)
const (
	dogA    = 10
	catE    = 5
	mouseC  = 250
	spiderD = 150
)

func calcFoodDog(qtAnimal int) (float64, error) {
	var result float64 = dogA * float64(qtAnimal)
	return float64(result), nil
}

func calcFoodCat(qtAnimal int) (float64, error) {
	var result float64 = catE * float64(qtAnimal)
	return float64(result), nil
}

func calcFoodMouse(qtAnimal int) (float64, error) {
	var result float64 = mouseC * float64(qtAnimal)
	return float64(result), nil
}

func calcFoodSpider(qtAnimal int) (float64, error) {
	var result float64 = spiderD * float64(qtAnimal)
	return float64(result), nil
}

func calculateFood(animal string) func(qtAnimal int) (float64, error) {

	switch animal {
	case "dog":
		return calcFoodDog
	case "cat":
		return calcFoodCat
	case "spider":
		return calcFoodSpider
	case "mouse":
		return calcFoodMouse
	default:
		fmt.Println("Animal not exist")
	}
	return nil
}
