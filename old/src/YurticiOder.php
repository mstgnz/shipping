<?php

class YurticiOder{

    private $soap, $urlMode;
    private $data = [
        "wsUserName" => "YKTEST",
        "wsPassword" => "YK",
        "wsUserLanguage" => "TR",
        "userLanguage" => "TR",
        "payerCustData" => [
            "invCustId" => "",
            "invAddressId" => ""
        ]
    ];
    private $urls = [
        "shipment" => [ // kargo işlmeleri için
            "test" => "http://testwebservices.yurticikargo.com:9090/KOPSWebServices/NgiShipmentInterfaceServices?wsdl",
            "live" => "http://webservices.yurticikargo.com:8080/KOPSWebServices/NgiShipmentInterfaceServices?wsdl"
        ],
        "reference" => [ // kargo sorgulama işlemleri için
            "test" => "http://testwebservices.yurticikargo.com:9090/KOPSWebServices/WsReportWithReferenceServices?wsdl",
            "live" => "http://webservices.yurticikargo.com:8080/KOPSWebServices/WsReportWithReferenceServices?wsdl"
        ]
    ];
    
    function __construct($username=false, $password=false, $custId=false, $urlMode="test", $self="shipment"){
        if($username && $password && $custId){
            $this->data["wsUserName"] = $username;
            $this->data["wsPassword"] = $password;
            $this->data["payerCustData"]['invCustId'] = $custId;
            $this->unneccessary['custParamsVO']['invCustIdArray'] = $custId;
            $this->unneccessary['shipmentParamsData']['projectCustIdArray'] = $custId;
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$self][$this->urlMode], array('trace' => true));
    }

    // Kargo Oluştur
    public function createNgiShipmentWithAddress($query=[]){
        try{
            $args = [
                "shipmentData" => array(
                    "ngiDocumentKey"                => $query['sale']['id'],
                    "cargoType"                     => "2",
                    "totalCargoCount"               => "1",
                    "totalDesi"                     => "1",
                    "totalWeight"                   => "1",
                    "personGiver"                   => $query['company']['name'],
                    "productCode"                   => "STA",
                    "description"                   => "",
                    "selectedArrivalUnitId"         => "",
                    "selectedArrivalTransferUnitId" => "",
                    "docCargoDataArray"             => array(
                        "ngiCargoKey"               => $query['sale']['id'],
                        "cargoType"                 => "2",
                        "cargoDesi"                 => "1",
                        "cargoWeight"               => "1",
                        "weightUnit"                => "KGM",
                        "cargoCount"                => "1",
                        "width"                     => "",
                        "height"                    => "",
                        "length"                    => "",
                        "dimensionsUnit"            => ""
                    ),
                    "specialFieldDataArray"         => array(
                        "specialFieldName"          => "",
                        "specialFieldValue"         => ""
                    ),
                    "complementaryProductDataArray" => array(
                        "complementaryProductCode"  => ""
                    ),
                    "codData"                       => array(
                        "ttInvoiceAmount"           => number_format($query['sale']['grand_total'],2,',','.'),
                        "ttDocumentId"              => $query['sale']['id'],
                        "ttCollectionType"          => "",
                        "ttDocumentSaveType"        => "0",
                        "dcSelectedCredit"          => "1",
                        "dcCreditRule"              => "1"
                    )
                ),
                "XSenderCustAddress"                => array(
                    "senderCustName"                => $query['sender']['fullname'],
                    "senderAddress"                 => $query['sender']['address'],
                    "cityId"                        => $query['sender']['city_id'],
                    "townName"                      => $query['sender']['district_name'],
                    "senderMobilePhone"             => $query['sender']['phone'],
                    "senderEmailAddress"            => $query['sender']['email'],
                    "senderPhone"                   => $query['sender']['phone'],
                    "senderCustReferenceId"         => "",
                    "senderAddressReferenceId"      => "",
                    "senderAdditionalInfo"          => "",
                    "latitude"                      => "",
                    "longitude"                     => ""
                ),
                "XConsigneeCustAddress"             => array(
                    "consigneeCustName"             => $query['customer']['firstname'].' '.$query['customer']['lastname'],
                    "consigneeAddress"              => $query['customer']['address'],
                    "cityId"                        => $query['customer']['city_id'],
                    "townName"                      => $query['customer']['district'],
                    "consigneeMobilePhone"          => $query['customer']['phone'],
                    "consigneeEmailAddress"         => $query['customer']['email'],
                    "consigneePhone"                => $query['customer']['phone'],
                    "consigneeCustReferenceId"      => "",
                    "consigneeAddressReferenceId"   => "",
                    "consigneeAdditionalInfo"       => "",
                    "latitude"                      => "",
                    "longitude"                     => ""
                )
            ];
            if($query['sale']['payment_type'] == 3){
                $args['shipmentData']['codData']['ttCollectionType'] = "0";
            }
            if($query['sale']['payment_type'] == 4){
                $args['shipmentData']['codData']['ttCollectionType'] = "1";
            }
            $this->data = array_merge($this->data, $args);
            $response = $this->soap->createNgiShipmentWithAddress($this->data);
            return ["request"=>"success", "response"=>$response];
        }catch (Exception $e){
            return ["request"=>"error", "response"=>$e->getMessage()];
        }
    }


}