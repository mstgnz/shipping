package ptt

import (
	"github.com/mstgnz/shipping/cargo/ptt/rest"
	"github.com/mstgnz/shipping/cargo/ptt/soap"
	"github.com/mstgnz/shipping/config"
)

type pttCargo struct {
	*cargo.Cargo
}

func NewPTTCargo() cargo.Shipper {
	return &pttCargo{&cargo.Cargo{}}
}

func (p pttCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch p.GetServiceType() {
	case cargo.SOAP:
		if p.IsDomestic() {
			result, err = soap.CreateDomestic(p.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(p.GetCurrent(), data)
		}
	case cargo.REST:
		if p.IsDomestic() {
			result, err = rest.CreateDomestic(p.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(p.GetCurrent(), data)
		}
	}
	return result, err
}

func (p pttCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch p.GetServiceType() {
	case cargo.SOAP:
		if p.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(p.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(p.GetCurrent(), data)
		}
	case cargo.REST:
		if p.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(p.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(p.GetCurrent(), data)
		}
	}
	return result, err
}

func (p pttCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch p.GetServiceType() {
	case cargo.SOAP:
		if p.IsDomestic() {
			result, err = soap.CancelCargoDomestic(p.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(p.GetCurrent(), data)
		}
	case cargo.REST:
		if p.IsDomestic() {
			result, err = rest.CancelCargoDomestic(p.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(p.GetCurrent(), data)
		}
	}
	return result, err
}
