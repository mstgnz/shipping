package fedex

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type fedexCargo struct {
	*cargo.Cargo
}

func NewFedexProvider() cargo.Shipper {
	return &fedexCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (f fedexCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (f fedexCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (f fedexCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
