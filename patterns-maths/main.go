package main

import (
	"fmt"
	"math"
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
	//pattern14()
	//pattern15()
	//pattern16()
	//pattern17()
	//pattern18()
	//pattern19()
	pattern20()
	//pattern21()
	//pattern22()

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

func pattern12() {
	n := 10

	// outer loop to iterate 4 times
	for i := 0; i < n; i++ {
		for j := 0; j <= 2*n-1; j++ {

			// print the numbers and spaces

			if j >= 0 && j <= i {
				fmt.Print(j + 1)
			} else if j >= 2*n-1-i {
				fmt.Print(2*n - j)
			} else {
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

func pattern13() {
	n := 5
	k := 1
	// outer loop to iterate 5 times.

	for i := 0; i < n; i++ {
		// inner loop will iterate upto i

		for j := 0; j <= i; j++ {
			fmt.Print(k, " ")
			k = k + 1
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

func pattern14() {
	var a rune

	// outer loop to iterate 5 times.

	for i := 0; i < 5; i++ {
		a = 'A'

		// inner loop to iterate upto i.

		for j := 0; j <= i; j++ {
			fmt.Print(string(a + rune(j)))
		}
		fmt.Println()
	}

}

/*
ABCDE
ABCD
ABC
AB
A
*/

func pattern15() {
	n := 5
	var a rune

	// outer loop to iterate 5 times.

	for i := 0; i < 5; i++ {
		a = 'A'

		// inner loop to iterate upto n-i
		for j := 0; j < n-i; j++ {
			fmt.Print(string(a + rune(j)))
		}
		fmt.Println()
	}
}

/*
A
BB
CCC
DDDD
EEEEE
*/

func pattern16() {
	var a rune

	// outer loop to iterate 5 times.

	for i := 0; i < 5; i++ {
		a = 'A'

		// inner loop to iterate upto i.

		for j := 0; j <= i; j++ {
			fmt.Print(string(a + rune(i)))
		}
		fmt.Println()
	}

}

/*
	 A
	ABA
 ABCBA
ABCDCBA
*/

func pattern17() {
	n := 4

	for i := 0; i < n; i++ {
		spaces := n - i - 1
		//print spaces

		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// print stars
		var a rune = 'A'

		breakpoint := (2*i + 1) / 2

		for j := 0; j < 2*i+1; j++ {

			fmt.Print(string(a))
			if j < breakpoint {
				a = a + 1
			} else {
				a = a - 1
			}
		}

		fmt.Println()

	}
}

/*
E
D E
C D E
B C D E
A B C D E
*/

func pattern18() {

	n := 5
	var a rune = 'A' + rune(n-1)
	// outer loop to iterate n times.
	for i := 0; i < n; i++ {
		// inner loop iterates upto i times.
		for j := 0; j <= i; j++ {
			fmt.Print(string(a-rune(i-j)), " ")
		}
		fmt.Println()
	}
}


/*
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
*/

func pattern19() {
	n := 5
	// outer loop iterates n times.
	spaces := 0
	for i := 0; i < n; i++ {

		// print stars
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}

		// print the spaces

		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// print the stars

		for j := 0; j < n-i; j++ {
			fmt.Print("*")

		}
		spaces = spaces + 2

		fmt.Println()
	}

	spaces = spaces - 2

	for i := 0; i < n; i++ {
		// print stars

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		// print spaces

		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// print stars

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
		spaces = spaces - 2
	}

}
/*
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **

*/
func pattern20() {
	n := 5
	spaces := 2*n - 2
	for i := 1; i <= 2*n-1; i++ {
		stars := i

		if i > n {
			stars = 2*n - i
		}

		// print stars

		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}

		// print spaces
		//fmt.Println(spaces)

		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}

		// print stars

		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}

		fmt.Println()

		if i < n {
			spaces = spaces - 2
		} else {
			
			spaces = spaces + 2
		
		}

	}
}

/*

*****
*   *
*   *
*   *
*****

*/

func pattern21() {

	n := 5
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == (n-1) || j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}


/*
4444444
4333334
4322234
4321234
4322234
4333334
4444444

*/


/*
The logic is not straight forward.
Explanation:

This matrix is generated by the following way
n - 
4444444
4333334
4322234
4321234
4322234
4333334
4444444

=
0000000
0111110
0122210
0123210
0122210
0111110
0000000


So what we will do 

n - currentValue = NewValue
So we get the current value by this way
currentmatrix = n - newMatrix

If we carefully observe the newMatrix
The each value will be the min of distance b/w(top,left,right,bottom)

so during the iteration we will do like this
n - min(top,bottom,right,left)
then we will achieve the above matrix.


*/

func pattern22() {
	n := 4
	for i := 0; i < 2*n-1; i++ {

		for j := 0; j < 2*n-1; j++ {
			top := i
			left := j
			right := 2*n - 2 - j
			bottom := 2*n - 2 - i

			fmt.Print(n-int(math.Min(math.Min(float64(top),float64(left)),math.Min(float64(right),float64(bottom)))))
		}
		fmt.Println()
	}
}
