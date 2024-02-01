package iyi

import (
	"github.com/mstgnz/shipping/cargo/iyi/rest"
	"github.com/mstgnz/shipping/cargo/iyi/soap"
	"github.com/mstgnz/shipping/config"
)

type iyiCargo struct {
	*cargo.Cargo
}

func NewIyiCargo() cargo.Shipper {
	return &iyiCargo{&cargo.Cargo{}}
}

func (i iyiCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch i.GetServiceType() {
	case cargo.SOAP:
		if i.IsDomestic() {
			result, err = soap.CreateDomestic(i.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(i.GetCurrent(), data)
		}
	case cargo.REST:
		if i.IsDomestic() {
			result, err = rest.CreateDomestic(i.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(i.GetCurrent(), data)
		}
	}
	return result, err
}

func (i iyiCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch i.GetServiceType() {
	case cargo.SOAP:
		if i.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(i.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(i.GetCurrent(), data)
		}
	case cargo.REST:
		if i.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(i.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(i.GetCurrent(), data)
		}
	}
	return result, err
}

func (i iyiCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch i.GetServiceType() {
	case cargo.SOAP:
		if i.IsDomestic() {
			result, err = soap.CancelCargoDomestic(i.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(i.GetCurrent(), data)
		}
	case cargo.REST:
		if i.IsDomestic() {
			result, err = rest.CancelCargoDomestic(i.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(i.GetCurrent(), data)
		}
	}
	return result, err
}
