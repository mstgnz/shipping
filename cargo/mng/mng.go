package mng

import (
	"github.com/mstgnz/shipping/cargo/mng/rest"
	"github.com/mstgnz/shipping/cargo/mng/soap"
	"github.com/mstgnz/shipping/config"
)

// TEST: https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx
// LIVE: https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx
type mngCargo struct {
	*cargo.Cargo
}

func NewMNGCargo() cargo.Shipper {
	return &mngCargo{&cargo.Cargo{}}
}

func (m mngCargo) CreateCargo(data cargo.ShippingData) (*cargo.Response, error) {
	var err error
	var result *cargo.Response
	switch m.GetServiceType() {
	case cargo.SOAP:
		if m.IsDomestic() {
			result, err = soap.CreateDomestic(m.GetCurrent(), data)
		} else {
			result, err = soap.CreateAbroad(m.GetCurrent(), data)
		}
	case cargo.REST:
		if m.IsDomestic() {
			result, err = rest.CreateDomestic(m.GetCurrent(), data)
		} else {
			result, err = rest.CreateAbroad(m.GetCurrent(), data)
		}
	}
	return result, err
}

func (m mngCargo) WhereIsTheCargo(data cargo.ShippingData) (*cargo.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (m mngCargo) CancelCargo(data cargo.ShippingData) (*cargo.Response, error) {
	//TODO implement me
	panic("implement me")
}
