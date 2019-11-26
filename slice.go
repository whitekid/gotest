package main

import "fmt"

func printSlice(s []int) string {
	return fmt.Sprintf("len = %d, cap = %d, %+v", len(s), cap(s), s)
}

// removeAt remove slice element index, panic if index is out of range
func removeAt(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
