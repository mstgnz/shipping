package yurtici

import (
	"github.com/mstgnz/shipping/config"
)

type yurticiCargo struct {
	*cargo.Cargo
}

func NewYurticiCargo() cargo.Shipper {
	return &yurticiCargo{&cargo.Cargo{}}
}

func (y yurticiCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
