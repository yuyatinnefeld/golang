package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"firstname"`
	Age       int      `json:"age"`
	Nicknames []string `json:"your nicknames"`
}

// func (p Person) MarshalJSON() ([]byte, error) {
// 	v, err := json.Marshal(&struct {
// 		Name string
// 	}{
// 		Name: "Mr./Ms." + p.Name,
// 	})
// 	return v, err
// }

func main() {

	fmt.Println("RECEIVE JSON DATA")
	b := []byte(`{"firstname":"mike", "age":20, "your nicknames":["a", "b", "c"]}`)

	fmt.Println(b)

	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	fmt.Println("SEND JSON DATA")
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

}
