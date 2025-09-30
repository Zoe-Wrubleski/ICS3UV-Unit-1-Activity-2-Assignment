/*
 *@author Zoe Wrubleski
 *@version 1.0.0
 *@date 2025-09-30
 *@fileoverview This program converts 34 degrees Celsius to its Fahrenheit equivalent
 */

package main

import "fmt"

func main() {
	//text explaining calculation
	fmt.Println("What is 34 degrees Celsius in Fahrenheit?")

	//conversion of 34 degrees C to F
	fmt.Println("34 C = " + fmt.Sprint(9 / 5 * 34 + 32) + " F")

	fmt.Println("\nDone.")
}
