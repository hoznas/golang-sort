package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ary := []int{}
	rand.Seed(time.Now().UnixNano())
	n := 20
	for i := 0; i < n; i++ {
		ary = append(ary, rand.Intn(n))
	}

	fmt.Println("-------not sorted---------")
	fmt.Println(ary)
	fmt.Println("-------  sorted  ---------")
	shellsort(ary)
	fmt.Println(ary)
}

func shellsort(a []int) {
	for d := make_dist(a); d > 0; d = next_dist(d) {
		for i := 0; i < d; i++ {
			dist_sort(a, i, d)
		}
	}
}
func dist_sort(a []int, n int, d int) {
	for i := d + n; i < len(a); i += d {
		dist_insert(a, i, d)
	}
}
func dist_insert(a []int, n int, d int) {
	for i := n; i-d >= 0 && a[i-d] > a[i]; i -= d {
		swap(a, i-d, i)
	}
}

func next_dist(n int) int {
	return (n - 1) / 3
}
func make_dist(a []int) int {
	x := len(a) / 9
	var i int
	for i = 1; i < x; i = 3*i + 1 {
	}
	return i
}

func insertsort(a []int) {
	for i := 1; i < len(a); i++ {
		insert(a, i)
	}
}

func insert(a []int, n int) {
	for i := n; i-1 >= 0 && a[i-1] > a[i]; i-- {
		swap(a, i-1, i)
	}
}

func swap(a []int, x int, y int) {
	t := a[x]
	a[x] = a[y]
	a[y] = t
}
