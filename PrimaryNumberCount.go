package main

import "fmt"

func primeNumber(c chan int, num int, n int) {
	sum := 0
	for i := num; i <= num+(n); i++ {
		f := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				f = 1
				break
			}
		}
		if f == 0 && i != 1 && i != 0 {
			sum++
		}
	}
	c <- sum
}

func getStart(n, m int) int {
	return n * m
}

func maxNumber(n int) int {
	num := n / 5
	return num
}

func main() {
	sum := 0
	channel := make(chan int)
	fmt.Print("Enter the range for Prime Number: ")
	var num int
	fmt.Scan(&num)

	endNo := maxNumber(num)

	go primeNumber(channel, getStart(0, endNo), endNo)
	go primeNumber(channel, getStart(1, endNo), endNo)
	go primeNumber(channel, getStart(2, endNo), endNo)
	go primeNumber(channel, getStart(3, endNo), endNo)
	go primeNumber(channel, getStart(4, endNo), endNo)
	for i := 0; i < 5; i++ {
		sum = sum + <-channel
	}
	fmt.Printf("Prime Number from 1 to %d is %d", num, sum)

}
