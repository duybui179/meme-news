package main

import "fmt"

// Hàm Sum để chúng ta chạy unit test
func Sum(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println("GitHub Actions Go Sample")
    fmt.Printf("Sum of 5 + 5 is: %d\n", Sum(5, 5))
}