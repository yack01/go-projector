/*
https://github.com/yack01/go-projector/blob/main/hw03.go

Програма реалізує алгоритм пішохода при перетині дороги на світлофорі.
Спочатку пропонує вказати, яке світло світофора горить і для незелоного світла повертає рекомендації з очікування.
Для зеленого свтла пропонує все ж переконатись у відсутності автомобілів: подивитись наліво, а потім направо.В кінці, якщо автівак немає, - перехід дозволено
*/

package main

import (
	"fmt"
	"strings"
)

type RoadSide struct {
	side     string
	hasCars  bool
	question string
}

func ChooseTrafficLightColor() string {
	var input string
	fmt.Print("Enter the traffic light color (R for red, Y for yellow, G for green): ")
	fmt.Scanln(&input)
	return strings.ToLower(input)
}

func LookTheSide() {
	leftSide := RoadSide{side: "left", question: "Look to the left. Are there cars?"}
	rightSide := RoadSide{side: "right", question: "Look to the right. Are there cars?"}

	leftSide.askAndCheck()
	rightSide.askAndCheck()

	fmt.Println("OK, you should cross the road!")
}

func (r *RoadSide) askAndCheck() {
	for {
		fmt.Printf("%s (y/n): ", r.question)
		var input string
		fmt.Scanln(&input)

		response := strings.ToLower(input)

		switch response {
		case "y":
			fmt.Println("Be careful, look again.")
			r.hasCars = true
		case "n":
			return
		default:
			fmt.Println("Incorrect input. Please enter 'y' for yes or 'n' for no.")
		}
	}
}

func main() {
	for {
		lightColor := ChooseTrafficLightColor()

		switch lightColor {
		case "r":
			fmt.Println("Traffic light is Red - stay and wait for yellow!")
		case "y":
			fmt.Println("Traffic light is Yellow - stay and wait for green!")
		case "g":
			LookTheSide()
			return
		default:
			fmt.Println("Incorrect input. Please enter R for red, Y for yellow, or G for green.")
		}
	}
}
