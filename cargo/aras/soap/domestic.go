package soap

type createData struct {
	userName  string
	password  string
	orderInfo orderInfo
}

type orderInfo struct {
	Order order
}

type order struct {
	TradingWaybillNumber string
	InvoiceNumber        string
	IntegrationCode      string
	ReceiverName         string
	ReceiverAddress      string
	ReceiverCityName     string
	ReceiverTownName     string
	ReceiverDistrictName string
	ReceiverQuarterName  string
	ReceiverAvenueName   string
	ReceiverStreetName   string
	ReceiverPhone1       string
	ReceiverPhone2       string
	ReceiverPhone3       string
	VolumetricWeight     string
	Weight               string
	CityCode             string
	TownCode             string
	SpecialField1        string
	SpecialField2        string
	SpecialField3        string
	IsCod                string
	CodAmount            string
	CodCollectionType    string
	CodBillingType       string
	TaxNumber            string
	TaxOffice            string
	PayorTypeCode        string
	IsWorldWide          string
	Description          string
	PrivilegeOrder       string
	PieceCount           string
	PieceDetails         pieceDetails
}

type pieceDetails struct {
	PieceDetail pieceDetail
}

type pieceDetail struct {
	VolumetricWeight string
	Weight           string
	BarcodeNumber    string
	ProductNumber    string
	Description      string
}
