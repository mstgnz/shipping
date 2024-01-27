package iyi

import (
	"github.com/mstgnz/shipping/config"
)

type iyiCargo struct {
	*cargo.Cargo
}

func NewIyiCargo() cargo.Shipper {
	return &iyiCargo{&cargo.Cargo{}}
}

func (i iyiCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (i iyiCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (i iyiCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
