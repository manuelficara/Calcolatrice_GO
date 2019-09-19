package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for{
		fmt.Println("Inserisci il primo numero: ")
		var num1 float64
		if _, err := fmt.Scanf("%f\n", &num1); err == nil{
			var sel string = ""
			for{
				fmt.Println("Inserisci il secondo numero: ")
				var num2 float64
				if _, err := fmt.Scanf("%f\n", &num2); err == nil{
					fmt.Println("Inserisci l'operazione da effettuare (+, -, *, /, c): ")
					if _, err := fmt.Scanf("%s\n", &sel); err == nil{
						var num1Formatted = strconv.FormatFloat(num1, 'E', -1, 64)
						var num2Formatted = strconv.FormatFloat(num2, 'E', -1, 64)
						if sel == "+"{
							output := strings.Join([]string{num1Formatted, " + ", num2Formatted, " = ", strconv.FormatFloat(num1 + num2, 'E', -1, 64)}, "")
							num1 += num2
							fmt.Println(output)
						}else if sel == "-"{
							output := strings.Join([]string{num1Formatted, " - ", num2Formatted, " = ", strconv.FormatFloat(num1 - num2, 'E', -1, 64)}, "")
							num1 -= num2
							fmt.Println(output)
						}else if sel == "*"{
							output := strings.Join([]string{num1Formatted, " * ", num2Formatted, " = ", strconv.FormatFloat(num1 * num2, 'E', -1, 64)}, "")
							num1 *= num2
							fmt.Println(output)
						}else if sel == "/"{
							output := strings.Join([]string{num1Formatted, " / ", num2Formatted, " = ", strconv.FormatFloat(num1 / num2, 'E', -1, 64)}, "")
							num1 /= num2
							fmt.Println(output)
						}
					}
				}
				if sel == "c"{
					break
				}
			}
		}
	}
}
