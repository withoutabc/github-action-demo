package main

import "testing"

// 🐱
func TestCat(t *testing.T) {
	get := Cat()
	want := "Miao~~~~~~~"
	if get != want {
		t.Errorf("Cat saying %s", get)
	}
}
