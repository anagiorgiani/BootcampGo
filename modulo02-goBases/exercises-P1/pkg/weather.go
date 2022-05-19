package main
import "fmt"

func weather(){
	var temperature int
	var moisture int
	var pressure int
	temperature, moisture, pressure = 8, 60, 14
	fmt.Println(temperature, moisture, pressure)
}