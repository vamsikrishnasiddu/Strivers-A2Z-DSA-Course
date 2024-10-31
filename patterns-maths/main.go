package main

import (
	"fmt"
)

func main() {
	//pattern1()
	pattern2()
}

/*
 *****
 *****
 *****
 *****
 *****
 */

func pattern1() {

	// outer loop repeats the inner loop 5 times.
	for i := 0; i < 5; i++ {
		//inner loop which prints the '*'
		for i := 0; i < 5; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

/*
*
**
***
****
*****
*/

func pattern2(){

	// outer to loop to repeat the inner loop 5 times.

	for i:=0;i<5;i++{
		// inner loop to print the '*'

		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}

