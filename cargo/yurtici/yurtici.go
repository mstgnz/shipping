package yurtici

import (
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

func (y yurticiCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch y.GetServiceType() {
	case cargo.SOAP:
		if y.IsDomestic() {
			result, err = soap.CreateDomestic(y.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(y.GetCurrent(), data)
		}
	case cargo.REST:
		if y.IsDomestic() {
			result, err = rest.CreateDomestic(y.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(y.GetCurrent(), data)
		}
	}
	return result, err
}

func (y yurticiCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch y.GetServiceType() {
	case cargo.SOAP:
		if y.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(y.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(y.GetCurrent(), data)
		}
	case cargo.REST:
		if y.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(y.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(y.GetCurrent(), data)
		}
	}
	return result, err
}

func (y yurticiCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch y.GetServiceType() {
	case cargo.SOAP:
		if y.IsDomestic() {
			result, err = soap.CancelCargoDomestic(y.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(y.GetCurrent(), data)
		}
	case cargo.REST:
		if y.IsDomestic() {
			result, err = rest.CancelCargoDomestic(y.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(y.GetCurrent(), data)
		}
	}
	return result, err
}
