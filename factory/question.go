package questions


import (
  "answer"
)


type Level uint


const (
  Easy     Level = 25
  Medium   Level = 50
  Hard     Level = 75
  VeryHard Level = 100
)


type Question struct {
  text string
  tags *map[string]string
  level Level
  answer answer.Answer
}


func (this *Question) Text() string {
  return this.text
}

func (this *Question) Tags() *map[string]string {
  return this.tags
}

func (this *Question) Level() Level {
  return this.level
}

func (this *Question) Answer() answer.Answer {
  return this.answer
}


func NewQuestion(text string, tags *map[string]string, level Level, answer answer.Answer) *Question {
  return &Question{text, tags, level, answer}
}
