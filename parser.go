package elgar

import (
  "fmt"
  "github.com/kylelemons/go-gypsy/yaml"
  "io"
  "strconv"
  "strings"
)

type ElgarState struct {
  hosts     map[string]*Host
  services  map[string]*Service
}

type InvalidDeclaration struct {
  Name          string
  Declaration   yaml.Node
}

func (i InvalidDeclaration) Error() string {
  return fmt.Sprintf("Badly formed declaration for %s: %v", i.Name, i.Declaration)
}

func parseHosts(foo yaml.Map) (map[string]*Host, error) {

  hosts := map[string]*Host{}

  for name, val := range foo {

    declaration, _ := val.(yaml.Map)
    dec_type, _    := declaration["type"].(yaml.Scalar)

    if dec_type == "host" {

      address, err := declaration["address"].(yaml.Scalar)

      if err == false {
        return nil, &InvalidDeclaration{ Name: name, Declaration: val }
      }

      host, _ := NewHost(name, address.String())

      hosts[name] = host
    }
  }

  return hosts, nil
}

func parseServices(foo yaml.Map) (map[string]*Service, error) {

  services := map[string]*Service{}

  for name, val := range foo {

    declaration, _ := val.(yaml.Map)
    dec_type, _    := declaration["type"].(yaml.Scalar)

    if dec_type == "service" {

      rule_list, _ := declaration["rules"].(yaml.List)

      rules := []*Rule{}

      for _, params := range rule_list {

        map_for_rule := params.(yaml.Map)

        for command, params := range map_for_rule {
          prs, _ := params.(yaml.Scalar)

          protocol_and_ports := strings.Split(prs.String(), ",")

          port, err := strconv.ParseInt(strings.TrimSpace(protocol_and_ports[1]),0,0)

          if err != nil {
            return nil, &InvalidDeclaration{ Name: name, Declaration: val }
          }

          rules = append(rules, &Rule{ behaviour: command, protocol: protocol_and_ports[0], port: int(port) })
        }
      }

      services[name] = &Service{ name: name, rules: rules }
    }
  }

  return services, nil
}

func ParseFile(r io.Reader) (*ElgarState, error) {

  out, err := yaml.Parse(r)

  if err != nil {
    return nil, err
  }

  out_map, _ := out.(yaml.Map)

  hosts, err := parseHosts(out_map)

  if err != nil {
    return nil, err
  }

  services, err := parseServices(out_map)

  if err != nil {
    return nil, err
  }

  return &ElgarState{ hosts: hosts, services: services }, nil

}

func (es *ElgarState) Service(name string) (*Service, error) {

  if val, ok := es.services[name]; ok {
    return val, nil
  } else {
    return nil, fmt.Errorf("Could not find host: %s", name)
  }

}

func (es *ElgarState) Host(name string) (*Host, error) {

  if val, ok := es.hosts[name]; ok {
    return val, nil
  } else {
    return nil, fmt.Errorf("Could not find host: %s", name)
  }

}
