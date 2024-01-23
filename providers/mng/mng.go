package mng

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type mngCargo struct {
	*cargo.Cargo
}

func NewMNGProvider() cargo.Shipper {
	return &mngCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (m mngCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
