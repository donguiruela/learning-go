package main

import (
	"fmt"
	"strings"
)

func main() {
	var option, search string
	var b, n int

	fmt.Println("[B]inary Search\n[L]inear Search")
	fmt.Printf("Enter your choice of search: ")
	fmt.Scan(&option)

	switch strings.ToUpper(option) {
	case "B":
		b = 2
		search = "Binary"
	default:
		b = -1
		search = "Invalid"
	}

	fmt.Printf("Enter number of elements: ")
	fmt.Scan(&n)

	x := GetLogBS(&b, &n)

	fmt.Println(search, "Search")
	fmt.Println("Maximum Worst Comparison: ", x)
}

/*Binary Search

Worst-case performance	O(log n)
Best-case performance	O(1)
Average performance	O(log n)
Worst-case space complexity	O(1)
*/
func GetLogBS(b *int, n *int) int {
	var l int = 0

	for *n > 1 {
		*n /= *b

		if *n > 0 {
			l += 1
		}

		if *n == 1 {
			l = (l * 2) + 1
		}
	}

	return l
}
