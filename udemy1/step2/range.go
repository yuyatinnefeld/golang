package main

import "fmt"

func main() {
	pl := []string{
		"python",
		"go",
		"javascript",
	}

	for i := 0; i < len(pl); i++ {
		fmt.Println(i, pl[i])
	}

	for i, v := range pl {
		fmt.Println(i, v)
	}

	m := map[string]int{
		"apple":  299,
		"banana": 9238,
		"orange": 39,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
