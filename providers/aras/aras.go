package aras

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type arasCargo struct {
	*cargo.Cargo
}

func NewArasProvider() cargo.Shipper {
	return &arasCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (a arasCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (a arasCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
