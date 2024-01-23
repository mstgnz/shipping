package turk

import (
	"context"

	"github.com/mstgnz/cargo/config"
)

type turkCargo struct {
	*cargo.Cargo
}

func NewTurkProvider() cargo.Shipper {
	return &turkCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (t turkCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
