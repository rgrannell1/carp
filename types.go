package main

type Dependency map[string]string

// Group defines a logically coherant group of dependencies
type Group struct {
	Requires []Dependency `json:"requires"`
}

type DependencyResult struct {
	Met    bool
	Reason []string
}
