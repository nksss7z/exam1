package main

import "fmt"

func Countletters(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			c += 1
		}
	}
	return (c)
}

func main() {
	fmt.Println(Countletters("MOrning"))
	fmt.Println(Countletters("246810"))
	fmt.Println(Countletters("game over"))
	fmt.Println(Countletters("school"))
}
