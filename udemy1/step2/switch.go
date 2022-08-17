package main

import (
	"fmt"
	"time"
)

func main() {
	os := map[string]string{
		"pc1": "mac",
		"pc2": "windows",
		"pc3": "linux",
		"pc4": "linux",
		"pc5": "mac",
		"pc6": "android",
	}

	for k, v := range os {
		fmt.Println(k, v)
		switch v {
		case "mac":
			fmt.Println("desinger?")
		case "windows":
			fmt.Println("office worker?")
		case "linux":
			fmt.Println("nerd?")
		default:
			fmt.Println("????")
		}
	}

	t := time.Now()
	h := t.Weekday()
	fmt.Println(h)

}
