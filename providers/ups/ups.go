package ups

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type upsCargo struct {
	*cargo.Cargo
}

func NewUPSProvider() cargo.Shipper {
	return &upsCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (u upsCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
