package main

import (
	"fmt"
	"strings"
)

func main() {
  // var name string = "Dent, Arthur"
  // name := "Dent, Arthur"

  // var score = 87
  // score := 87

  // name, score := "Dent, Arthur", 87

  // students := []string{"Dent, Arthur",
  //   "MacMillan, Tricia",
  //   "Ford, Prefect",
  // }
  // scores := []int{87, 96, 64}
  // scores := map[string]int{
  //   students[0]: 87,
  //   students[1]: 96,
  //   students[2]: 64,
  // }

  type score struct {
    name  string
    score int
  }

  scores := []score{
    {name: "Dent, Arthur", score: 87},
    {name: "MacMillan, Tricia", score: 96},
    {name: "Ford, Prefect", score: 64},
  }

  fmt.Println("Select score to print (1 - 3):")
  var option string
  fmt.Scanln(&option)

  fmt.Println("Student scores")
  fmt.Println(strings.Repeat("-", 14))
  var index int
  // if option == "1" { // should use strconv package in production!
  //   index = 0
  // } else if option == "2" {
  //   index = 1
  // } else if option == "3" {
  //   index = 2
  // } else {
  //   fmt.Println("Unknown option, defaulting to 1")
  //   index = 0
  // }
  switch option {
  case "1":
    index = 0
  case "2":
    index = 1
  case "3":
    index = 2
  default:
    fmt.Println("Unknown option, defaulting to 1")
    index = 0
  }
  fmt.Println(scores[index].name, scores[index].score)
}
