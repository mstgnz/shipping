package vatan

import (
	"github.com/mstgnz/shipping/config"
)

type vatanCargo struct {
	*cargo.Cargo
}

func NewVatanCargo() cargo.Shipper {
	return &vatanCargo{&cargo.Cargo{}}
}

func (v vatanCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (v vatanCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (v vatanCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
