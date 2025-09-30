/*
 *@author Zoe Wrubleski
 *@version 1.0.0
 *@date 2025-09-30
 *@fileoverview This program calculates the volume of a sphere with a radius of 4 cm
 */

package main

import "fmt"

func main() {
	//text asking the question
	fmt.Println("What is the volume of a sphere with a radius of 4 cm?")

	//calculation
	fmt.Println("V ≈ " + fmt.Sprint(4 / 3 * 3.141592653589 * 4 * 4 * 4)+ "cm³")

	fmt.Println("\nDone.")

}
