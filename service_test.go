package elgar

import (
  "os"
  "testing"
)

func TestNewService_nonEmptyRuleSet(t *testing.T) {

  _, err := NewService("foo")

  if err == nil {
    t.Error("Service with empty ruleset should error")
  }

}

func TestParseService_fromYaml(t *testing.T) {

  file, _      := os.Open("./resources/example_service.yml")
  out,  _      := ParseFile(file)
  service, err := out.Service("example_service")

  if err != nil { t.Error("Cannot parse host from file") }
  if service.name != "example_service" { t.Error("Hostname not 'example_host': %s", service.name) }

  if len(service.rules) != 1 {
    t.Error("Service contains no rules")
  }

  rule := service.rules[0]

  t.Log(rule)

  if rule.protocol != "tcp" {
    t.Errorf("Rule parsed protocol incorrectly: %s", rule.protocol)
  }

  if rule.port != 9000 {
    t.Errorf("Rule parsed port incorrectly: %d", rule.port)
  }
}
