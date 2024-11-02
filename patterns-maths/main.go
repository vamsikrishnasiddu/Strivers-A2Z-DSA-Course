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
	//pattern9()
	//pattern10()
	//pattern11()
	//pattern12()
	//pattern13()
	pattern14()

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

/*
*
**
***
****
*****
*****
****
***
**
*

 */

func pattern10() {
	//combination of pattern2 and pattern5
	pattern2()
	pattern5()
}

/*
1
01
101
0101
10101

*/

func pattern11() {
	var k int

	// outer loop to iterate 5 times.
	for i := 0; i < 5; i++ {
		k = 1

		// inner loop to go until i.
		if i%2 == 0 {
			k = 0
		}

		for j := 0; j <= i; j++ {
			k = 1 - k
			fmt.Print(k)
		}
		fmt.Println()
	}

}

/*
1      1
22    21
333  321
44444321

*/

func pattern12(){
	n:=10

	// outer loop to iterate 4 times
	 for i:=0;i<n;i++{
		 for j:=0;j<=2*n-1;j++{

			 // print the numbers and spaces

			 if (j>=0 && j<=i){
				 fmt.Print(j+1)
			 }else if (j>=2*n-1-i){
				 fmt.Print(2*n-j)
			 }else{
				 fmt.Print(" ")
			 }
		 }
		 fmt.Println()
	 }
}


/*
1 
2 3 
4 5 6 
7 8 9 10 
11 12 13 14 15 

*/


func pattern13(){
	n:=5
	k:=1
	// outer loop to iterate 5 times.

	for i:=0;i<n;i++{
		// inner loop will iterate upto i

		for j:=0;j<=i;j++{
			fmt.Print(k," ")
			k=k+1
		}
		fmt.Println()
	}
}

/*
A
AB
ABC
ABCD
ABCDE
*/


func pattern14(){
	var a rune

	// outer loop to iterate 5 times.

	for i:=0;i<5;i++{
		 a = 'A'

		// inner loop to iterate upto i.

		for j:=0;j<=i;j++{
			fmt.Print(string(a+rune(j)))
		}
		fmt.Println()
	}
  
}
