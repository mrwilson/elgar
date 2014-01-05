package elgar

import (
  "net"
  "fmt"
)

type Host struct {
  name      string
  ip        *net.IP
}

func NewHost(name, address string) (*Host, error) {

  ip := net.ParseIP(address)

  if ip == nil {
    return nil, fmt.Errorf("Badly formed address: %s",ip)
  }

  return &Host{ name: name, ip: &ip }, nil

}
