package ups

import (
	"net/http"

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

func (u upsCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch u.GetServiceType() {
	case cargo.SOAP:
		if u.IsDomestic() {
			result, err = soap.CreateDomestic(u.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(u.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if u.IsDomestic() {
			result, err = rest.CreateDomestic(u.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(u.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (u upsCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u upsCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
