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

func (s sendeoCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (s sendeoCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (s sendeoCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
