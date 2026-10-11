package main
import "fmt"

func main() {
    store := newStore()
    store.Set("KEY1", "hihry")
	store.Delete("KEY1")
    val, err := store.Get("KEY1")

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Value:", val)
    }
    fmt.Println("Hello there")
}