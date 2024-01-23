package ptt

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type pttCargo struct {
	*cargo.Cargo
}

func NewPTTProvider() cargo.Shipper {
	return &pttCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (p pttCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
