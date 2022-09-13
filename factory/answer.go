package answer


type Type uint


const (
  Text   Type = iota
  Choice
)


type Answer interface {
  Type() Type
}
