package elgar

import (
  "fmt"
)

type Rule struct {
  behaviour string
  protocol  string
  port      int
}

type Service struct {
  name  string
  rules []*Rule
}

func NewService(name string, rules... *Rule) (*Service, error) {

  if len(rules) == 0 {
    return nil, fmt.Errorf("Empty ruleset for Service: %s",name)
  }

  return &Service{ name: name, rules: rules }, nil

}
