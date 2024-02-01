package ups

import (
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

func (u upsCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch u.GetServiceType() {
	case cargo.SOAP:
		if u.IsDomestic() {
			result, err = soap.CreateDomestic(u.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(u.GetCurrent(), data)
		}
	case cargo.REST:
		if u.IsDomestic() {
			result, err = rest.CreateDomestic(u.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(u.GetCurrent(), data)
		}
	}
	return result, err
}

func (u upsCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch u.GetServiceType() {
	case cargo.SOAP:
		if u.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(u.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(u.GetCurrent(), data)
		}
	case cargo.REST:
		if u.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(u.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(u.GetCurrent(), data)
		}
	}
	return result, err
}

func (u upsCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch u.GetServiceType() {
	case cargo.SOAP:
		if u.IsDomestic() {
			result, err = soap.CancelCargoDomestic(u.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(u.GetCurrent(), data)
		}
	case cargo.REST:
		if u.IsDomestic() {
			result, err = rest.CancelCargoDomestic(u.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(u.GetCurrent(), data)
		}
	}
	return result, err
}
