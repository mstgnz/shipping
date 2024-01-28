package mng

import (
	"github.com/mstgnz/shipping/cargo/mng/rest"
	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/config"
)

type mngCargo struct {
	*cargo.Cargo
}

func NewMNGCargo() cargo.Shipper {
	return &mngCargo{&cargo.Cargo{}}
}

func (m mngCargo) CreateCargo(data cargo.ShippingData) (map[string]any, error) {
	var err error
	var result map[string]any
	switch m.GetServiceType() {
	case cargo.SOAP:
		if m.IsDomestic() {
			result, err = soap.CreateDomestic(m.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(m.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if m.IsDomestic() {
			result, err = rest.CreateDomestic(m.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(m.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (m mngCargo) WhereIsTheCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) CancelCargo(data cargo.ShippingData) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
