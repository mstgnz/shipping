package cargo

import (
	"errors"

	"github.com/mstgnz/cargo/providers/aras"
	"github.com/mstgnz/cargo/providers/dhl"
	"github.com/mstgnz/cargo/providers/fedex"
	"github.com/mstgnz/cargo/providers/mng"
	"github.com/mstgnz/cargo/providers/ptt"
	"github.com/mstgnz/cargo/providers/sendeo"
	"github.com/mstgnz/cargo/providers/surat"
	"github.com/mstgnz/cargo/providers/tnt"
	"github.com/mstgnz/cargo/providers/turk"
	"github.com/mstgnz/cargo/providers/ups"
	"github.com/mstgnz/cargo/providers/yurtici"
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
