package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

// should implement ProviderInterface
type Handler struct {}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "aaa",
	}
}

func (h *Handler) DescribeCard(name string) plug.Card {
	return plug.Card{}
}
func (h *Handler) DescribePanel(name string) plug.Panel {
	return plug.Panel{}
}
func (h *Handler) Register(model string, name string) error {
	return nil
}
func (h *Handler) Get(model string, name string) plug.Record {
	return plug.Record{}
}
func (h *Handler) Set(model string, name string, record plug.Record) error {
	return nil
}
func (h *Handler) Del(model string, name string) error {
	return nil
}
