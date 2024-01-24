package soap

type self struct {
	wsUserName      string
	wsPassword      string
	ShippingOrderVO shippingOrderVO
}

type shippingOrderVO struct {
	cargoKey           string
	invoiceKey         string
	receiverCustName   string
	receiverAddress    string
	receiverPhone1     string
	cityName           string
	townName           string
	waybillNo          string
	taxOfficeId        string
	cargoCount         string
	ttDocumentId       string
	dcSelectedCredit   string
	dcCreditRule       string
	ttInvoiceAmount    string
	ttDocumentSaveType string
	ttCollectionType   string
}

type order struct {
	wsUserName            string
	wsPassword            string
	payerCustData         payerCustData
	shipmentData          shipmentData
	XSenderCustAddress    xSenderCustAddress
	XConsigneeCustAddress xConsigneeCustAddress
}

type payerCustData struct {
	invCustId string
}

type shipmentData struct {
	ngiDocumentKey                string
	cargoType                     string
	totalCargoCount               string
	totalDesi                     string
	totalWeight                   string
	personGiver                   string
	productCode                   string
	description                   string
	selectedArrivalUnitId         string
	selectedArrivalTransferUnitId string
	docCargoDataArray             docCargoDataArray
	specialFieldDataArray         specialFieldDataArray
	complementaryProductDataArray complementaryProductDataArray
	codData                       codData
}

type docCargoDataArray struct {
	ngiCargoKey    string
	cargoType      string
	cargoDesi      string
	cargoWeight    string
	weightUnit     string
	cargoCount     string
	width          string
	height         string
	length         string
	dimensionsUnit string
}

type specialFieldDataArray struct {
	specialFieldName  string
	specialFieldValue string
}

type complementaryProductDataArray struct {
	complementaryProductCode string
}

type codData struct {
	ttInvoiceAmount    string
	ttDocumentId       string
	ttCollectionType   string
	ttDocumentSaveType string
	dcSelectedCredit   string
	dcCreditRule       string
}

type xSenderCustAddress struct {
	senderCustName           string
	senderAddress            string
	cityId                   string
	townName                 string
	senderMobilePhone        string
	senderEmailAddress       string
	senderPhone              string
	senderCustReferenceId    string
	senderAddressReferenceId string
	senderAdditionalInfo     string
	latitude                 string
	longitude                string
}

type xConsigneeCustAddress struct {
	consigneeCustName           string
	consigneeAddress            string
	cityId                      string
	townName                    string
	consigneeMobilePhone        string
	consigneeEmailAddress       string
	consigneePhone              string
	consigneeCustReferenceId    string
	consigneeAddressReferenceId string
	consigneeAdditionalInfo     string
	latitude                    string
	longitude                   string
}
