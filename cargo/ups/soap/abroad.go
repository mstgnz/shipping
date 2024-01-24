package soap

type shipAccept struct {
	AccessRequest         accessRequest
	ShipmentAcceptRequest shipmentAcceptRequest
}

type shipConfirm struct {
	AccessRequest          accessRequest
	ShipmentConfirmRequest shipmentConfirmRequest
}

type accessRequest struct {
	AccessLicenseNumber string
	UserId              string
	Password            string
}

type shipmentAcceptRequest struct {
	Request        request
	ShipmentDigest string
}

type shipmentConfirmRequest struct {
	Request  request
	Shipment shipment
}

type request struct {
	RequestAction        string
	RequestOption        string
	TransactionReference transactionReference
}

type transactionReference struct {
	CustomerContext string
}

type shipment struct {
	Description        string
	Shipper            shipper
	ShipFrom           shipFrom
	ShipTo             shipTo
	PaymentInformation paymentInformation
	Service            service
	Package            shipmentPackage
}

type shipper struct {
	Name                    string
	AttentionName           string
	PhoneNumber             string
	TaxIdentificationNumber string
	ShipperNumber           string
	Address                 address
}

type address struct {
	AddressLine1      string
	AddressLine2      string
	AddressLine3      string
	City              string
	PostalCode        string
	CountryCode       string
	StateProvinceCode string
}

type shipFrom struct {
	CompanyName             string
	AttentionName           string
	PhoneNumber             string
	TaxIdentificationNumber string
	Address                 address
}

type shipTo struct {
	CompanyName   string
	AttentionName string
	PhoneNumber   string
	Address       address
}

type paymentInformation struct {
	Prepaid prepaid
}

type prepaid struct {
	BillShipper billShipper
}

type billShipper struct {
	AccountNumber string
}

type service struct {
	Code        string
	Description string
}

type shipmentPackage struct {
	Description   string
	PackagingType packagingType
	PackageWeight packageWeight
	Dimensions    dimensions
}

type packagingType struct {
	Code        string
	Description string
}

type packageWeight struct {
	Weight            string
	UnitOfMeasurement unitOfMeasurement
}

type unitOfMeasurement struct {
	Code string
}

type dimensions struct {
	Length            string
	Width             string
	Height            string
	UnitOfMeasurement unitOfMeasurement
}

type labelRecovery struct {
	AccessRequest        accessRequest
	LabelRecoveryRequest labelRecoveryRequest
}

type labelRecoveryRequest struct {
	Request            request
	LabelSpecification labelSpecification
	LabelDelivery      labelDelivery
	TrackingNumber     string
}

type labelSpecification struct {
	HTTPUserAgent    string
	LabelImageFormat labelImageFormat
}

type labelImageFormat struct {
	Code string
}

type labelDelivery struct {
	LabelLinkIndicator   string
	ResendEMailIndicator string
}

type shipCancel struct {
	AccessRequest       accessRequest
	VoidShipmentRequest voidShipmentRequest
}

type voidShipmentRequest struct {
	Request                      request
	ShipmentIdentificationNumber string
}

type shipTracking struct {
	AccessRequest accessRequest
	TrackRequest  trackRequest
}

type trackRequest struct {
	Request        request
	TrackingNumber string
}
