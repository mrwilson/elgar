package elgar

import (
  "testing"
  "os"
)

func TestNewHost_validIP(t *testing.T) {

  hostname := "foo"
  address  := "1.2.3.4"

  host, _ := NewHost(hostname, address)

  if host.name != hostname {
    t.Error("Hostnames do not match")
  }

  if host.ip.String() != address {
    t.Error("IP addresses do not match")
  }

}

func TestNewHost_invalidIP(t *testing.T) {

  hostname := "foo"
  address  := "invalid"

  host, err := NewHost(hostname, address)

  if host != nil || err == nil {
    t.Error("Did not return error on invalid address")
  }

}

func TestParseHost_fromYaml(t *testing.T) {

  file, err := os.Open("./resources/example_host.yml")

  if err != nil {
    t.Error("Could not open file")
  }

  out, err := ParseFile(file)

  if err != nil {
    t.Errorf("Error parsing file: %+v", err)
  }

  host, err := out.Host("example_host")

  if err != nil {
    t.Error("Cannot parse host from file")
  }

  if host.name != "example_host" {
    t.Error("Hostname not 'example_host': %s", host.name)
  }

  if host.ip.String() != "1.2.3.4" {
    t.Error("Address does not match '1.2.3.4': %s", host.ip.String())
  }
}
