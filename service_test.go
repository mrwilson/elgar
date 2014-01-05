package elgar

import (
  "testing"
)

func TestNewService_nonEmptyRuleSet(t *testing.T) {

  _, err := NewService("foo")

  if err == nil {
    t.Error("Service with empty ruleset should error")
  }

}
