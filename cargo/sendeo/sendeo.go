package sendeo

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type sendeoCargo struct {
	*cargo.Cargo
}

func NewSendeoCargo() cargo.Shipper {
	return &sendeoCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (s sendeoCargo) CreateCargo() error {
	//TODO implement me
	panic("implement me")
}

func (s sendeoCargo) WhereIsTheCargo() error {
	//TODO implement me
	panic("implement me")
}

func (s sendeoCargo) CancelCargo() error {
	//TODO implement me
	panic("implement me")
}
