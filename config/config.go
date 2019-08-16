package config

import (
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

// Service is the main config for this service.
type Service struct {
	Host  string    `yaml:"host"`
	Paths []Mapping `yaml:"paths"`
}

// Mapping will hold the map from a map to repo.
type Mapping struct {
	Path string `yaml:"path"`
	Repo string `yaml:"repo"`
	VCS  string `yaml:"vcs"`
}

// Parse will parse content into service.
func Parse(content []byte) (s *Service, err error) {
	s = &Service{}
	err = yaml.Unmarshal(content, s)
	if err != nil {
		return
	}

	for k, v := range s.Paths {
		if v.VCS == "" {
			s.Paths[k].VCS = "git"
		}
	}

	return
}

// Find will find a path from Service.
func (s Service) Find(path string) (*Mapping, string) {
	for _, v := range s.Paths {
		if !strings.HasPrefix(v.Path, path) {
			log.Printf("Path %s is not found in Service", path)
			continue
		}

		m := v
		if path == v.Path {
			return &m, ""
		}
		return &m, path[len(v.Path)+1:]
	}
	return nil, ""
}
