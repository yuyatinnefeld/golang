package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(5, 15)
    if total != 20 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}