package iyi

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type iyiCargo struct {
	*cargo.Cargo
}

func NewIyiProvider() cargo.Shipper {
	return &iyiCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (i iyiCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (i iyiCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (i iyiCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
