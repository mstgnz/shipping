package turk

import (
	"github.com/mstgnz/shipping/cargo/turk/rest"
	"github.com/mstgnz/shipping/cargo/turk/soap"
	"github.com/mstgnz/shipping/config"
)

type turkCargo struct {
	*cargo.Cargo
}

func NewTurkCargo() cargo.Shipper {
	return &turkCargo{&cargo.Cargo{}}
}

func (t turkCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if t.GetServiceType() == cargo.SOAP {
		if t.IsDomestic() {
			result, err = soap.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	if t.GetServiceType() == cargo.REST {
		if t.IsDomestic() {
			result, err = rest.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (t turkCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
