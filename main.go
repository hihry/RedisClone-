package main

import (
	"fmt"
)

func main() {
	store := newStore()
	store.Set("KEY1", "hihry")
	val, ok := store.Get("KEY1")
	fmt.Println("%w , %w", val, ok)
	fmt.Println("Hello there")
}
