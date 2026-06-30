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

  fmt.Println("Student scores")
  fmt.Println(strings.Repeat("-", 14))
  fmt.Println(scores[0].name, scores[0].score)
  fmt.Println(scores[1].name, scores[1].score)
  fmt.Println(scores[2].name, scores[2].score)
}
