package models

type Album struct {
	ID     string
	Name   string
	Artist string
	Songs  []*Song
}
