package soap

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/mstgnz/shipping/config"
)

// CreateAbroad global service
func CreateAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CreateDomestic local service
// TEST REQUEST: https://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL
// LIVE REQUEST: https://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL
func CreateDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {

	orderData := SiparisGirisiDetayliV3{}
	switch content := data.(type) {
	case SiparisGirisiDetayliV3:
		orderData = content
	case []byte:
		if json.Valid(content) {
			if err := json.Unmarshal(content, &orderData); err != nil {
				return nil, fmt.Errorf("JSON Unmarshal Error: %s", err)
			}
		} else {
			if err := xml.Unmarshal(content, &orderData); err != nil {
				return nil, fmt.Errorf("XML Unmarshal Error: %s", err)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported data type: %T", data)
	}

	orderData.PKullaniciAdi = current.Credential.GetUsername()
	orderData.WsUserName = current.Credential.GetUsername()
	orderData.PMusteriNo = current.Credential.GetUsername()
	orderData.WsPassword = current.Credential.GetPassword()
	orderData.PSifre = current.Credential.GetPassword()

	xmlHeader := `<?xml version="1.0" encoding="utf-8"?><soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body>`

	xmlFooter := `</soap:Body></soap:Envelope>`

	xmlMarshal, err := xml.MarshalIndent(orderData, "", "    ")
	if err != nil {
		return nil, err
	}
	xmlContent := append(append([]byte(xmlHeader), xmlMarshal...), []byte(xmlFooter)...)

	resp, err := http.Post(current.Endpoint.GetDevelopment(), "text/xml; charset=utf-8", bytes.NewBuffer(xmlContent))
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &cargo.Response{
		Status: resp.Status,
		Code:   resp.StatusCode,
		Data:   body,
		Length: resp.ContentLength,
	}

	return res, err
}

// WhereIsTheCargoAbroad global service
func WhereIsTheCargoAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// WhereIsTheCargoDomestic local service
func WhereIsTheCargoDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CancelCargoAbroad global service
func CancelCargoAbroad(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}

// CancelCargoDomestic local service
func CancelCargoDomestic(current cargo.Current, data cargo.ShippingData) (*cargo.Response, error) {
	return nil, nil
}
