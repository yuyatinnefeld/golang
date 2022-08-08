package main

import "fmt"

func by2(input int) string {
	if input%2 == 0 {
		return "yes"
	} else {
		return "no"
	}
}

func main() {
	res := by2(200)
	fmt.Println(res)

	if res_2 := by2(10); res_2 == "yes" {
		fmt.Println("great")

	}

	for i := 1; i < 61; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz, ")
		} else if i%3 == 0 {
			fmt.Print("Fizz, ")
		} else if i%5 == 0 {
			fmt.Print("Buzz, ")
		} else {
			fmt.Print(i, ", ")
		}
	}

	fmt.Println("\n### continue and break ###\n")
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
}
