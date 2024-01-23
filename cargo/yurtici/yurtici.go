package yurtici

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type yurticiCargo struct {
	*cargo.Cargo
}

func NewYurticiProvider() cargo.Shipper {
	return &yurticiCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (y yurticiCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
