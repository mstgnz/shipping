package fedex

import (
	"github.com/mstgnz/shipping/cargo/fedex/rest"
	"github.com/mstgnz/shipping/cargo/fedex/soap"
	"github.com/mstgnz/shipping/config"
)

type fedexCargo struct {
	*cargo.Cargo
}

func NewFedexCargo() cargo.Shipper {
	return &fedexCargo{&cargo.Cargo{}}
}

func (f fedexCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch f.GetServiceType() {
	case cargo.SOAP:
		if f.IsDomestic() {
			result, err = soap.CreateDomestic(f.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(f.GetCurrent(), data)
		}
	case cargo.REST:
		if f.IsDomestic() {
			result, err = rest.CreateDomestic(f.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(f.GetCurrent(), data)
		}
	}
	return result, err
}

func (f fedexCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch f.GetServiceType() {
	case cargo.SOAP:
		if f.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(f.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(f.GetCurrent(), data)
		}
	case cargo.REST:
		if f.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(f.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(f.GetCurrent(), data)
		}
	}
	return result, err
}

func (f fedexCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch f.GetServiceType() {
	case cargo.SOAP:
		if f.IsDomestic() {
			result, err = soap.CancelCargoDomestic(f.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(f.GetCurrent(), data)
		}
	case cargo.REST:
		if f.IsDomestic() {
			result, err = rest.CancelCargoDomestic(f.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(f.GetCurrent(), data)
		}
	}
	return result, err
}
