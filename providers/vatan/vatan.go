package vatan

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type vatanCargo struct {
	*cargo.Cargo
}

func NewVatanProvider() cargo.Shipper {
	return &vatanCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (v vatanCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (v vatanCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (v vatanCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
