package main

import "fmt"

func main() {
	sum := 0
	fibseq := []int{1, 1, 2}
	for fibseq[2] < 4000000 {
		if fibseq[2] % 2 == 0 {
			sum += fibseq[2]
		}
		fibseq[0] = fibseq[1]
		fibseq[1] = fibseq[2]
		fibseq[2] = fibseq[0] + fibseq[1]
	}
	fmt.Printf("%d\n", sum)
}
