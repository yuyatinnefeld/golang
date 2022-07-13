package main

import (
    "fmt"
    "github.com/go-gota/gota/series"
    "github.com/go-gota/gota/dataframe"
	  "strings"
)

func main() {

    a := map[string]series.Type{
		"A": series.String,
		"D": series.Bool,
	}

    fmt.Println(a)


    df := dataframe.New(
        series.New([]string{"b", "a"}, series.String, "COL.1"),
        series.New([]int{1, 2}, series.Int, "COL.2"),
        series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
    )

    fmt.Println(df)


    type Dog struct {
        Name     string
        Colour      string
        Height  int
      Vaccinated  bool
    }
    
    dogs := []Dog{
        {"buster", "black", 56, false},
        {"jake", "white", 61, false},
        {"bingo", "brown", 50, true},
        {"gray", "cream", 68, false},
    }
    
    dogsDf := dataframe.LoadStructs(dogs) 
    fmt.Println(dogsDf)

    jsonString := `[
        {
          "Name": "John",
          "Age": 44,
          "Colour": "Red",
          "Height(ft)": 6.7
        },
        {
          "Name": "Mary",
          "Age": 40,
          "Colour": "Blue",
          "Height(ft)": 5.7
        }
      ]`
      
    dataRead := dataframe.ReadJSON(strings.NewReader(jsonString))
    fmt.Println(dataRead)
}