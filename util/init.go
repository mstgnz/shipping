package util

import (
	"context"
	"errors"

	"github.com/mstgnz/shipping/cargo/aras"
	"github.com/mstgnz/shipping/cargo/dhl"
	"github.com/mstgnz/shipping/cargo/fedex"
	"github.com/mstgnz/shipping/cargo/iyi"
	"github.com/mstgnz/shipping/cargo/mng"
	"github.com/mstgnz/shipping/cargo/ptt"
	"github.com/mstgnz/shipping/cargo/sendeo"
	"github.com/mstgnz/shipping/cargo/surat"
	"github.com/mstgnz/shipping/cargo/tnt"
	"github.com/mstgnz/shipping/cargo/turk"
	"github.com/mstgnz/shipping/cargo/ups"
	"github.com/mstgnz/shipping/cargo/vatan"
	"github.com/mstgnz/shipping/cargo/yurtici"
	"github.com/mstgnz/shipping/config"
)

// providerMap is a map that associates Provider names with their corresponding constructor functions.
var providerMap = make(map[string]interface{})

// NewProviderByName returns a new preconfigured provider instance by its name identifier.
func NewProviderByName(name string) (cargo.Shipper, error) {
	constructorFunc, ok := providerMap[name]
	if ok {
		construct := constructorFunc.(func() cargo.Shipper)()
		construct.SetContext(context.Background())
		return construct, nil
	}
	return nil, errors.New("Missing provider " + name)
}

// registerProvider registers a Provider constructor function with a given name.
func registerProvider[T cargo.Shipper](name string, constructor func() T) {
	providerMap[name] = constructor
}

func init() {
	registerProvider("aras", aras.NewArasCargo)
	registerProvider("dhl", dhl.NewDHLCargo)
	registerProvider("fedex", fedex.NewFedexCargo)
	registerProvider("iyi", iyi.NewIyiCargo)
	registerProvider("mng", mng.NewMNGCargo)
	registerProvider("ptt", ptt.NewPTTCargo)
	registerProvider("sendeo", sendeo.NewSendeoCargo)
	registerProvider("surat", surat.NewSuratCargo)
	registerProvider("tnt", tnt.NewTNTCargo)
	registerProvider("turk", turk.NewTurkCargo)
	registerProvider("ups", ups.NewUPSCargo)
	registerProvider("vatan", vatan.NewVatanCargo)
	registerProvider("yurtici", yurtici.NewYurticiCargo)
}
