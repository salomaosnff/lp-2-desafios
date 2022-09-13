package answerchoice


type AnswerChoice struct {
  options *map[string]string
}


func (this *AnswerChoice) Type() Type {
  return Choice
}

func (this *AnswerChoice) Options() *map[string]string {
  return this.options
}

func NewAnswerChoice(options *map[string]string) *AnswerChoice {
  return &AnswerChoice{options}
}
