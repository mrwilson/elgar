package elgar

import (
  "net"
)

type Host struct {
  name      string
  ip        *net.IPAddr
}

func NewHost(name, address string) (*Host, error) {

  ip, err := net.ResolveIPAddr("ip", address)

  if err != nil {
    return nil, err
  }

  return &Host{ name: name, ip: ip }, nil

}
