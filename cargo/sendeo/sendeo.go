package sendeo

import (
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

func (s sendeoCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch s.GetServiceType() {
	case cargo.SOAP:
		if s.IsDomestic() {
			result, err = soap.CreateDomestic(s.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(s.GetCurrent(), data)
		}
	case cargo.REST:
		if s.IsDomestic() {
			result, err = rest.CreateDomestic(s.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(s.GetCurrent(), data)
		}
	}
	return result, err
}

func (s sendeoCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch s.GetServiceType() {
	case cargo.SOAP:
		if s.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(s.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(s.GetCurrent(), data)
		}
	case cargo.REST:
		if s.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(s.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(s.GetCurrent(), data)
		}
	}
	return result, err
}

func (s sendeoCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch s.GetServiceType() {
	case cargo.SOAP:
		if s.IsDomestic() {
			result, err = soap.CancelCargoDomestic(s.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(s.GetCurrent(), data)
		}
	case cargo.REST:
		if s.IsDomestic() {
			result, err = rest.CancelCargoDomestic(s.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(s.GetCurrent(), data)
		}
	}
	return result, err
}
