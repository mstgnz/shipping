package surat

import (
	"github.com/mstgnz/shipping/config"
)

type suratCargo struct {
	*cargo.Cargo
}

func NewSuratCargo() cargo.Shipper {
	return &suratCargo{&cargo.Cargo{}}
}

func (s suratCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (s suratCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (s suratCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
