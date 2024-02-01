package turk

import (
	"github.com/mstgnz/shipping/cargo/turk/rest"
	"github.com/mstgnz/shipping/cargo/turk/soap"
	"github.com/mstgnz/shipping/config"
)

type turkCargo struct {
	*cargo.Cargo
}

func NewTurkCargo() cargo.Shipper {
	return &turkCargo{&cargo.Cargo{}}
}

func (t turkCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch t.GetServiceType() {
	case cargo.SOAP:
		if t.IsDomestic() {
			result, err = soap.CreateDomestic(t.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(t.GetCurrent(), data)
		}
	case cargo.REST:
		if t.IsDomestic() {
			result, err = rest.CreateDomestic(t.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(t.GetCurrent(), data)
		}
	}
	return result, err
}

func (t turkCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch t.GetServiceType() {
	case cargo.SOAP:
		if t.IsDomestic() {
			result, err = soap.WhereIsTheCargoDomestic(t.GetCurrent(), data)
		} else {
			result, err = soap.WhereIsTheCargoAbroad(t.GetCurrent(), data)
		}
	case cargo.REST:
		if t.IsDomestic() {
			result, err = rest.WhereIsTheCargoDomestic(t.GetCurrent(), data)
		} else {
			result, err = rest.WhereIsTheCargoAbroad(t.GetCurrent(), data)
		}
	}
	return result, err
}

func (t turkCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch t.GetServiceType() {
	case cargo.SOAP:
		if t.IsDomestic() {
			result, err = soap.CancelCargoDomestic(t.GetCurrent(), data)
		} else {
			result, err = soap.CancelCargoAbroad(t.GetCurrent(), data)
		}
	case cargo.REST:
		if t.IsDomestic() {
			result, err = rest.CancelCargoDomestic(t.GetCurrent(), data)
		} else {
			result, err = rest.CancelCargoAbroad(t.GetCurrent(), data)
		}
	}
	return result, err
}
