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
	//TODO implement me
	panic("implement me")
}

func (p pttCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	//TODO implement me
	panic("implement me")
}
