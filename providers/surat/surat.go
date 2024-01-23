package surat

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type suratCargo struct {
	*cargo.Cargo
}

func NewSuratProvider() cargo.Shipper {
	return &suratCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (s suratCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (s suratCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (s suratCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
