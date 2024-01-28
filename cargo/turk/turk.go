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

func (t turkCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch t.GetServiceType() {
	case cargo.SOAP:
		if t.IsDomestic() {
			result, err = soap.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if t.IsDomestic() {
			result, err = rest.CreateDomestic(t.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(t.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (t turkCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (t turkCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	//TODO implement me
	panic("implement me")
}
