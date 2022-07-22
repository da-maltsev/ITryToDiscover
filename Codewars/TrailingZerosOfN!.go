package Codewars

/*
Number of trailing zeros of N!

Write a program that will calculate the number of trailing zeros in a factorial of a given number.

N! = 1 * 2 * 3 *  ... * N

Be careful 1000! has 2568 digits...

The following are examples of expected output values:

zeros(6) = 1
# 6! = 1 * 2 * 3 * 4 * 5 * 6 = 720 --> 1 trailing zero

zeros(12) = 2
# 12! = 479001600 --> 2 trailing zeros
*/

func Zeros(n int) int {  
  count := 0
  for n > 0 {
    n = n/5
    count += n
  }
  return count
}
