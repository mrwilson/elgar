package elgar

import (
  "net"
)

type Host struct {
  name      string
  ip        *net.IPAddr
}

func NewHost(name, address string) (*Host) {

  ip, _ := net.ResolveIPAddr("ip", address)

  return &Host{ name: name, ip: ip }

}
