package main

import "fmt"

func main() {

	english := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
	}

	turkish := make(map[string]string, len(english)) // Map can be growth than given Size "len(english)" -> Not like Slice
	for k, v := range english {
		turkish[v] = k
	}

	if v, ok := english["good"]; ok {
		fmt.Printf("Good in Turkish %q\n", v)
	}

	if v, ok := turkish["harika"]; ok {
		fmt.Printf("harika in English %q\n", v)
	}
}
