package ptt

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type pttCargo struct {
	*cargo.Cargo
}

func NewPTTCargo() cargo.Shipper {
	return &pttCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (p pttCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
