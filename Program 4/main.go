/*
 *@author Zoe Wrubleski
 *@version 1.0.0
 *@date 2025-09-30
 *@fileoverview This program calculates the gross pay for an employee named Fred, who worked 40 hours at an hourly wage of $8.20
 */

package main

import "fmt"

func main() {
	//text asking the question
	fmt.Println("If Fred worked for 40 hours at an hourly wage of $8.20, what is his gross pay?")

	//calculation
	fmt.Println("Gross pay = $" + fmt.Sprint(40 * 8.20))

	fmt.Println("\nDone.")

}
