package aras

import (
	"github.com/mstgnz/shipping/config"
)

type arasCargo struct {
	*cargo.Cargo
}

func NewArasCargo() cargo.Shipper {
	return &arasCargo{&cargo.Cargo{}}
}

func (a arasCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
