package main

import "fmt"

func firstUniqChar(s string) int {
	dict := make([]int, 26)

	for i := 0; i < len(s); i++ {
		dict[s[i] - 'a']++
	}

	for i := 0; i < len(s); i++ {
		if dict[s[i] - 'a'] == 1 {
			return i
		}
	}

	return -1
}

func main(){
	fmt.Printf("Unique char index %d\n", firstUniqChar("loveleetcode0"))
}
