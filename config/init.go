package cargo

import (
	"errors"

	"github.com/mstgnz/shipping/cargo/aras"
	"github.com/mstgnz/shipping/cargo/dhl"
	"github.com/mstgnz/shipping/cargo/fedex"
	"github.com/mstgnz/shipping/cargo/mng"
	"github.com/mstgnz/shipping/cargo/ptt"
	"github.com/mstgnz/shipping/cargo/sendeo"
	"github.com/mstgnz/shipping/cargo/surat"
	"github.com/mstgnz/shipping/cargo/tnt"
	"github.com/mstgnz/shipping/cargo/turk"
	"github.com/mstgnz/shipping/cargo/ups"
	"github.com/mstgnz/shipping/cargo/yurtici"
)

// providerMap is a map that associates Provider names with their corresponding constructor functions.
var providerMap = make(map[string]interface{})

// NewProviderByName returns a new preconfigured provider instance by its name identifier.
func NewProviderByName(name string) (Shipper, error) {
	constructorFunc, ok := providerMap[name]
	if ok {
		return constructorFunc.(func() Shipper)(), nil
	}
	return constructorFunc.(func() Shipper)(), errors.New("Missing provider " + name)
}

// registerProvider registers a Provider constructor function with a given name.
func registerProvider[T Shipper](name string, constructor func() T) {
	providerMap[name] = constructor
}

func init() {
	registerProvider("aras", aras.NewArasProvider)
	registerProvider("dhl", dhl.NewDHLProvider)
	registerProvider("fedex", fedex.NewFedexProvider)
	registerProvider("mng", mng.NewMNGProvider)
	registerProvider("ptt", ptt.NewPTTProvider)
	registerProvider("sendeo", sendeo.NewSendeoProvider)
	registerProvider("sürat", surat.NewSuratProvider)
	registerProvider("tnt", tnt.NewTNTProvider)
	registerProvider("kargo-turk", turk.NewTurkProvider)
	registerProvider("ups", ups.NewUPSProvider)
	registerProvider("yurtiçi", yurtici.NewYurticiProvider)
}
