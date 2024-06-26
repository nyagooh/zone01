// the function should have as paremeters a slice a[]int and a function f  func(int,int)int.For each element of the slice ,it should apply the function
// f func(int,int)int.save the result and then print it

// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice,
// it should apply the function f func(int, int) int, save the result and then print it.
// Expected function

package main

import "github.com/01-edu/z01"

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2,}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
func ReduceInt(a []int, f func(int, int) int) {
	if len(a) > 0 {
		z01.PrintRune(' ')
	}
	input := a[0]
	for i := 1; i < len(a); i++ {
		input = f(input, a[i])
	}
	//itoa
	output := itoa(input)
	for _, ch := range output {
		z01.PrintRune(ch)
	}
	z01.PrintRune(10)
}
func itoa (s int)string {
	sign := 1
	str := ""
	if s < 1 {
		sign*=-1
		s*=-1
	}
	for s > 0{
		digit := s % 10
		str = string(digit+'0') + str
		s /= 10
	}
	if sign < 0 {
		str = "-" + str
	}
	return str
}

