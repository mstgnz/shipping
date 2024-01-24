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

func (t tntCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t tntCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
