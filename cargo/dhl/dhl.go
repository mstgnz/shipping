package dhl

import (
	"github.com/mstgnz/shipping/config"
)

type dhlCargo struct {
	*cargo.Cargo
}

func NewDHLCargo() cargo.Shipper {
	return &dhlCargo{&cargo.Cargo{}}
}

func (d dhlCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (d dhlCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (d dhlCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
