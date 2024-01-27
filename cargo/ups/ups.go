package ups

import (
	"github.com/mstgnz/shipping/cargo/ups/rest"
	"github.com/mstgnz/shipping/cargo/ups/soap"
	"github.com/mstgnz/shipping/config"
)

type upsCargo struct {
	*cargo.Cargo
}

func NewUPSCargo() cargo.Shipper {
	return &upsCargo{&cargo.Cargo{}}
}

func (u upsCargo) CreateCargo(data map[string]any) (map[string]any, error) {
	var err error
	var result map[string]any
	if u.GetServiceType() == cargo.SOAP {
		if u.IsDomestic() {
			result, err = soap.CreateDomestic(u.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(u.GetCurrentEndpointAndCredential(), data)
		}
	}
	if u.GetServiceType() == cargo.REST {
		if u.IsDomestic() {
			result, err = rest.CreateDomestic(u.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(u.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (u upsCargo) WhereIsTheCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) CancelCargo(tracking string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
