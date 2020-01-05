package main

import "testing"

func BenchmarkCocktailSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{1, 7, 4, 2, -9, 15, 90, 500, 22, 1}
		CocktailSort(slice)
	}
}