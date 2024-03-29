package aras

import (
	"github.com/mstgnz/shipping/cargo/aras/rest"
	"github.com/mstgnz/shipping/cargo/aras/soap"
	"github.com/mstgnz/shipping/config"
)

type arasCargo struct {
	*cargo.Cargo
}

func NewArasCargo() cargo.Shipper {
	return &arasCargo{&cargo.Cargo{}}
}

func (a arasCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch a.GetServiceType() {
	case cargo.SOAP:
		if a.IsDomestic() {
			result, err = soap.CreateDomestic(a.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(a.GetCurrent(), data)
		}
	case cargo.REST:
		if a.IsDomestic() {
			result, err = rest.CreateDomestic(a.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(a.GetCurrent(), data)
		}
	}
	return result, err
}

func (a arasCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch a.GetServiceType() {
	case cargo.SOAP:
		if a.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(a.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(a.GetCurrent(), data)
		}
	case cargo.REST:
		if a.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(a.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(a.GetCurrent(), data)
		}
	}
	return result, err
}

func (a arasCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch a.GetServiceType() {
	case cargo.SOAP:
		if a.IsDomestic() {
			result, err = soap.CancelCargoDomestic(a.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(a.GetCurrent(), data)
		}
	case cargo.REST:
		if a.IsDomestic() {
			result, err = rest.CancelCargoDomestic(a.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(a.GetCurrent(), data)
		}
	}
	return result, err
}
