package elgar

import (
  "testing"
)

func TestCreateHost_validIP(t *testing.T) {

  hostname := "foo"
  address  := "1.2.3.4"

  host := NewHost(hostname, address)

  if host.name != hostname {
    t.Error("Hostnames do not match")
  }

  if host.ip.IP.String() != address {
    t.Error("IP addresses do not match")
  }

}
