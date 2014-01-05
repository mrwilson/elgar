package elgar

import (
  "fmt"
  "github.com/kylelemons/go-gypsy/yaml"
  "io"
)

type ElgarState struct {
  hosts     map[string]*Host
  services  map[string]*Service
}

func parseHosts(foo yaml.Map) (map[string]*Host, error) {

  hosts := map[string]*Host{}

  for name, val := range foo {

    declaration, _ := val.(yaml.Map)

    dec_type, _ := declaration["type"].(yaml.Scalar)

    if dec_type == "host" {

      address, _ := declaration["address"].(yaml.Scalar)
      host, _ := NewHost(name, address.String())

      hosts[name] = host
    }
  }

  return hosts, nil

}

func ParseFile(r io.Reader) (*ElgarState, error) {

  out, err := yaml.Parse(r)

  if err != nil {
    return nil, err
  }

  out_map, _ := out.(yaml.Map)

  hosts, _ := parseHosts(out_map)

  return &ElgarState{ hosts: hosts, services: map[string]*Service{} }, nil

}

func (es *ElgarState) Host(name string) (*Host, error) {

  if val, ok := es.hosts[name]; ok {
    return val, nil
  } else {
    return nil, fmt.Errorf("Could not find host: %s", name)
  }

}
