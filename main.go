package main

import (
	"fmt"
)

func main() {
	var n, k, ma, mi int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&k)

		if i == 0 {
			mi = k
			ma = k
		}

		if k < mi {
			mi = k
		} else if k > ma {
			ma = k
		}
	}
	fmt.Println(ma - mi)
}
