package dhl

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type dhlCargo struct {
	*cargo.Cargo
}

func NewDHLProvider() cargo.Shipper {
	return &dhlCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (d dhlCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (d dhlCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (d dhlCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
