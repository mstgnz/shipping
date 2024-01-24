package turk

import (
	"context"

	"github.com/mstgnz/shipping/config"
)

type turkCargo struct {
	*cargo.Cargo
}

func NewTurkCargo() cargo.Shipper {
	return &turkCargo{&cargo.Cargo{
		Ctx: context.Background(),
	}}
}

func (t turkCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
