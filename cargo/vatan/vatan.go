package vatan

import (
	"github.com/mstgnz/shipping/cargo/vatan/rest"
	"github.com/mstgnz/shipping/cargo/vatan/soap"
	"github.com/mstgnz/shipping/config"
)

type vatanCargo struct {
	*cargo.Cargo
}

func NewVatanCargo() cargo.Shipper {
	return &vatanCargo{&cargo.Cargo{}}
}

func (v vatanCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch v.GetServiceType() {
	case cargo.SOAP:
		if v.IsDomestic() {
			result, err = soap.CreateDomestic(v.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(v.GetCurrent(), data)
		}
	case cargo.REST:
		if v.IsDomestic() {
			result, err = rest.CreateDomestic(v.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(v.GetCurrent(), data)
		}
	}
	return result, err
}

func (v vatanCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch v.GetServiceType() {
	case cargo.SOAP:
		if v.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(v.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(v.GetCurrent(), data)
		}
	case cargo.REST:
		if v.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(v.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(v.GetCurrent(), data)
		}
	}
	return result, err
}

func (v vatanCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch v.GetServiceType() {
	case cargo.SOAP:
		if v.IsDomestic() {
			result, err = soap.CancelCargoDomestic(v.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(v.GetCurrent(), data)
		}
	case cargo.REST:
		if v.IsDomestic() {
			result, err = rest.CancelCargoDomestic(v.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(v.GetCurrent(), data)
		}
	}
	return result, err
}
