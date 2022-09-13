package main


import (
  "questionfactory"
  "fmt"
)



func main() {
  q1, _ := questionfactory.Create("Choice", "", &map[string]string{}, Hard, NewAnswerChoice(&map[string]string{}))
}
