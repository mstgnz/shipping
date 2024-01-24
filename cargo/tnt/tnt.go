package tnt

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type tntCargo struct {
	*cargo.Cargo
}

func NewTNTCargo() cargo.Shipper {
	return &tntCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (t tntCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}