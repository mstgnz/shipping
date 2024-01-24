package ups

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type upsCargo struct {
	*cargo.Cargo
}

func NewUPSCargo() cargo.Shipper {
	return &upsCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (u upsCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
