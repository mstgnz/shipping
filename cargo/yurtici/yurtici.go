package yurtici

import (
	"net/http"

	"github.com/mstgnz/shipping/cargo/yurtici/rest"
	"github.com/mstgnz/shipping/cargo/yurtici/soap"
	"github.com/mstgnz/shipping/config"
)

type yurticiCargo struct {
	*cargo.Cargo
}

func NewYurticiCargo() cargo.Shipper {
	return &yurticiCargo{&cargo.Cargo{}}
}

func (y yurticiCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch y.GetServiceType() {
	case cargo.SOAP:
		if y.IsDomestic() {
			result, err = soap.CreateDomestic(y.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(y.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if y.IsDomestic() {
			result, err = rest.CreateDomestic(y.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(y.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (y yurticiCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (y yurticiCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
