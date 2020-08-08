package main

import (
	"dict"
	"fmt"
)

func main() {

	s := dict.Dictionary{}
	result := s.GetMeanings([]string{"pretend", "holly", "holly", "holly"})

	for _, ornage := range result {
		fmt.Println(ornage.Word)
	}
}
