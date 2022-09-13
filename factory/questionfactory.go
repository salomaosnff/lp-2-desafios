package questionfactory


import (
  "question"
  "answer"

  "errors"
)


func Create(answerType, text string, tags *map[string]string, level question.Level, answer answer.Answer) (*question.Question, error) {
  if answerType == "Choice" {
    return question.NewQuestion(text, tags, level, answer), nil
  } else if answerType == "Text" {
    return question.NewQuestion(text, tags, level, answer), nil
  }

  return nil, errors.New("")
}
