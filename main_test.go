package main

import "testing"

// func TestSum(t *testing.T) {
//     total := Sum(5, 5)
//     if total != 10 {
//         t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
//     }
// }

func TestHello(t *testing.T){
	result := 200
	if (result == 200){
		t.Printf("The content: {}", result);
	}
}