package soap

type loginData struct {
	CustomerNumber string
	UserName       string
	Password       string
}

type createData struct {
	SessionID        string
	ShipmentInfo     shipmentInfo
	ReturnLabelLink  bool
	ReturnLabelImage bool
}

type shipmentInfo struct {
	ShipperName                string
	ShipperAddress             string
	ShipperCityCode            string
	ShipperAreaCode            string
	ShipperMobilePhoneNumber   string
	ShipperEMail               string
	ConsigneeName              string
	ConsigneeAddress           string
	ConsigneeCityCode          string
	ConsigneeAreaCode          string
	ConsigneeMobilePhoneNumber string
	ConsigneeEMail             string
	PaymentType                string
	SmsToConsignee             string
	ShipperAccountNumber       string
	ShipperContactName         string
	ShipperPostalCode          string
	ShipperPhoneExtension      string
	ShipperPhoneNumber         string
	ShipperExpenseCode         string
	ConsigneeAccountNumber     string
	ConsigneeContactName       string
	ConsigneeExpenseCode       string
	ConsigneePostalCode        string
	ConsigneePhoneNumber       string
	CustomerReferance          string
	CustomerInvoiceNumber      string
	DescriptionOfGoods         string
	DeliveryNotificationEmail  string
	NumberOfPackages           string
	PackageType                string
	ServiceLevel               string
	IdControlFlag              string
	PhonePrealertFlag          string
	InsuranceValue             string
	InsuranceValueCurrency     string
	ValueOfGoods               string
	ValueOfGoodsCurrency       string
	ValueOfGoodsPaymentType    string
	ThirdPartyAccountNumber    string
	ThirdPartyExpenseCode      string
	SmsToShipper               string
}

type cancelData struct {
	sessionId     string
	customerCode  string
	waybillNumber string
}

type trackingData struct {
	SessionID        string
	InformationLevel string
	TrackingNumber   string
}
