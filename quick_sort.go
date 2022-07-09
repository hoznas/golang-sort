package main

import "fmt"

func main() {
	fmt.Println("hello")
	a := []int{0, 9, 1, 8, 2, 7, 3, 6, 4, 5}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quicksort(a []int, s int, e int) {
	fmt.Println("qsort(", s, e, ")")
	if !(s < e) {
		return
	}
	pivot := a[(s+e)/2]
	fmt.Println("pivot", pivot)
	i := s
	j := e
	for true {
		for a[i] < pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(a, i, j)
		i++
		j--

	}
	x := make(chan struct{}, 0)
	y := make(chan struct{}, 0)
	go func() {
		quicksort(a, s, i-1)
		defer close(x)
	}()
	go func() {
		quicksort(a, j+1, e)
		defer close(y)
	}()
	<-x
	<-y

}

func swap(a []int, i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}
