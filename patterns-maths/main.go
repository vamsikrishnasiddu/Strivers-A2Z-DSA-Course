package main

import (
	"fmt"
)

func main() {
	//pattern1()
	//pattern2()
	//pattern3()
	//pattern4()
	//pattern5()
	//pattern6()
	//pattern7()
	//pattern8()
	pattern9()
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

func pattern2() {

	// outer to loop to repeat the inner loop 5 times.

	for i := 0; i < 5; i++ {
		// inner loop will iterate upto i
		// it will print the '*' in each iteration.

		for j := 0; j <= i; j++ {
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

func pattern3() {
	// outer loop to iterate  5 times.

	for i := 0; i < 5; i++ {
		// inner loop to iterate upto i.
		// it will print the numbers using j+1
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

/*
1
22
333
4444
55555
*/

func pattern4() {
	// outer loop iterates 5 times.
	for i := 0; i < 5; i++ {
		// inner loop will iterate upto i
		for j := 0; j <= i; j++ {
			fmt.Print(i + 1)
		}
		fmt.Println()
	}
}

/*

*****
****
***
**
*
 */

func pattern5() {
	n := 5
	// outer loop iterates 5 times.
	for i := 0; i < n; i++ {
		// inner loop goes upto n-i
		for j := 0; j < n-i; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

/*
12345
1234
123
12
1
*/

func pattern6() {
	n := 5
	// outer loop iterates 5 times.
	for i := 0; i < n; i++ {
		// inner loop goes upto n-i
		// it will print the numbers using j+1
		for j := 0; j < n-i; j++ {
			fmt.Print(j + 1)
		}

		fmt.Println()

	}
}

/*
        *
			 ***
      *****
     *******
		*********
*/

func pattern7() {
	n := 5
	// outer loop iterates 5 times.
	for i := 0; i < n; i++ {
		// inner loop iterates upto 2*n-1
		for j := 0; j < 2*n-1; j++ {
			if j >= n-i-1 && j <= n-1+i {
				print("*")
			} else {
				print(" ")
			}
		}
		fmt.Println()
	}
}

/*
*********
 *******
	*****
	 ***
		*
*/

func pattern8() {
	n := 5
	// outer loop iterates 5 times.
	for i := 0; i < n; i++ {
		for j := 0; j < 2*n-1; j++ {
			// inner loop iterates upto 2*n-1
			// * will be printed from i to 2*n-1-i
			if j >= i && j < 2*n-1-i {
				print("*")
			} else {
				print(" ")
			}
		}
		fmt.Println()
	}
}

/*
		*
	 ***
	*****
 *******
*********
*********
 *******
	*****
	 ***
		*

*/

func pattern9() {
	//combination of both pattern7 and pattern8
	pattern7()
	pattern8()
}
