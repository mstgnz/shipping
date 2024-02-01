package dhl

import (
	"github.com/mstgnz/shipping/cargo/dhl/rest"
	"github.com/mstgnz/shipping/cargo/dhl/soap"
	"github.com/mstgnz/shipping/config"
)

type dhlCargo struct {
	*cargo.Cargo
}

func NewDHLCargo() cargo.Shipper {
	return &dhlCargo{&cargo.Cargo{}}
}

func (d dhlCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch d.GetServiceType() {
	case cargo.SOAP:
		if d.IsDomestic() {
			result, err = soap.CreateDomestic(d.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(d.GetCurrent(), data)
		}
	case cargo.REST:
		if d.IsDomestic() {
			result, err = rest.CreateDomestic(d.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(d.GetCurrent(), data)
		}
	}
	return result, err
}

func (d dhlCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch d.GetServiceType() {
	case cargo.SOAP:
		if d.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(d.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(d.GetCurrent(), data)
		}
	case cargo.REST:
		if d.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(d.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(d.GetCurrent(), data)
		}
	}
	return result, err
}

func (d dhlCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch d.GetServiceType() {
	case cargo.SOAP:
		if d.IsDomestic() {
			result, err = soap.CancelCargoDomestic(d.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(d.GetCurrent(), data)
		}
	case cargo.REST:
		if d.IsDomestic() {
			result, err = rest.CancelCargoDomestic(d.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(d.GetCurrent(), data)
		}
	}
	return result, err
}
