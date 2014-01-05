package elgar

import (
  "testing"
)

func TestNewHost_validIP(t *testing.T) {

  hostname := "foo"
  address  := "1.2.3.4"

  host, _ := NewHost(hostname, address)

  if host.name != hostname {
    t.Error("Hostnames do not match")
  }

  if host.ip.IP.String() != address {
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
