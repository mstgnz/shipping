package mng

import (
	"github.com/mstgnz/shipping/config"
)

type mngCargo struct {
	*cargo.Cargo
}

func NewMNGCargo() cargo.Shipper {
	return &mngCargo{&cargo.Cargo{}}
}

func (m mngCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	//current := m.GetCurrentEndpointAndCredential()
	if m.IsDomestic() {

	} else {

	}
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
