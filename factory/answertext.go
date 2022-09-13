package answertext


type AnswerText struct {}


func (this *AnswerText) Type() Type {
  return Text
}

func NewAnswerText() *AnswerText {
  return new(AnswerText)
}
