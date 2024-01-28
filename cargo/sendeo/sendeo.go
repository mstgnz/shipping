package sendeo

import (
	"net/http"

	"github.com/mstgnz/shipping/cargo/sendeo/rest"
	"github.com/mstgnz/shipping/cargo/sendeo/soap"
	"github.com/mstgnz/shipping/config"
)

type sendeoCargo struct {
	*cargo.Cargo
}

func NewSendeoCargo() cargo.Shipper {
	return &sendeoCargo{&cargo.Cargo{}}
}

func (s sendeoCargo) CreateCargo(data cargo.ShippingData) (*http.Response, error) {
	var err error
	var result *http.Response
	switch s.GetServiceType() {
	case cargo.SOAP:
		if s.IsDomestic() {
			result, err = soap.CreateDomestic(s.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = soap.CreateAbroad(s.GetCurrentEndpointAndCredential(), data)
		}
	case cargo.REST:
		if s.IsDomestic() {
			result, err = rest.CreateDomestic(s.GetCurrentEndpointAndCredential(), data)
		} else {
			result, err = rest.CreateAbroad(s.GetCurrentEndpointAndCredential(), data)
		}
	}
	return result, err
}

func (s sendeoCargo) WhereIsTheCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s sendeoCargo) CancelCargo(data cargo.ShippingData) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}
