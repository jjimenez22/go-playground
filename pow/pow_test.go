package main

import "testing"

func TestPow(t *testing.T) {
    expected := 100
    res := pow(10)
    if res != expected {
        t.Errorf("Failed pow func. Expected: %d but got: %d.", expected, res)
    }
}