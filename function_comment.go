package main

import "fmt"

// This funtion has been written by ويكيبيديا!‏
// The name should be "wikipedia" transliterated to arabic
// This is for reference only
func main() {
	var isAdmin = false
	var isSuperAdmin = false
	isAdmin = isAdmin || isSuperAdmin

	// go ask ويكيبيديا‏ if this breaks
	if isAdmin {
		fmt.Println("You are an admin.")
	}

	if isAdmin { // go ask ويكيبيديا if this breaks
		fmt.Println("You are an admin.")
	}
}
