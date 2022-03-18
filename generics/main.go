package main

import "fmt"

func main() {
	fmt.Println(includesInt([]int{1, 2, 3}, 3))
	fmt.Println(includesInt([]int{1, 2, 3}, 5))
	fmt.Println(includesString([]string{"ab", "cd", "ef"}, "ab"))
	fmt.Println(includesString([]string{"ab", "cd", "ef"}, "gh"))
	// generic !
	fmt.Println(includes([]int{1, 2, 3}, 3))
	fmt.Println(includes([]int{1, 2, 3}, 5))
	fmt.Println(includes([]string{"ab", "cd", "ef"}, "ab"))
	fmt.Println(includes([]string{"ab", "cd", "ef"}, "gh"))
}

func includes[comparable int | string](values []comparable, value comparable) bool {
	for _, i := range values {
		if i == value {
			return true
		}
	}
	return false
}

func includesInt(values []int, value int) bool {
	for _, i := range values {
		if i == value {
			return true
		}
	}
	return false
}

func includesString(values []string, value string) bool {
	for _, i := range values {
		if i == value {
			return true
		}
	}
	return false
}
