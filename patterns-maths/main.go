package main

import (
	"fmt"
)

func main() {
	//pattern1()
	//pattern2()
	pattern3()
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

/*
1
12
123
1234
12345
*/

func pattern3(){
	// outer loop to iterate  5 times.

	for i:=0;i<5;i++{
		// inner loop to print the numbers.
		// j start from 0 and will go upto i
		for j:=0;j<=i;j++{
			fmt.Print(j+1)
		}
		fmt.Println()
	}
}

