<?php

class YurticiSelf{

    private $soap, $urlMode;
    private $data = [
        "wsUserName" => "YKTEST",
        "wsPassword" => "YK",
        "wsUserLanguage" => "TR",
        "userLanguage" => "TR"
    ];
    private $urls = [
        "test" => "http://testwebservices.yurticikargo.com:9090/KOPSWebServices/ShippingOrderDispatcherServices?wsdl",
        "live" => "http://webservices.yurticikargo.com:8080/KOPSWebServices/ShippingOrderDispatcherServices?wsdl"
    ];

    function __construct($username=false, $password=false, $urlMode="test"){
        if($username && $password){
            $this->data["wsUserName"] = $username;
            $this->data["wsPassword"] = $password;
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
    }

    // Kargo oluştur
    public function createShipment($query=[]){
        try{
            $args['ShippingOrderVO'] = [
                "cargoKey"              => $query['sale']['id'],
                "invoiceKey"            => $query['sale']['id'],
                "receiverCustName"      => $query['customer']['firstname'].' '.$query['customer']['lastname'],
                "receiverAddress"       => $query['customer']['address'],
                "receiverPhone1"        => $query['customer']['phone'],
                "cityName"              => $query['customer']['city'],
                "townName"              => $query['customer']['district'],
                "waybillNo"             => $query['sale']['id'],
                "taxOfficeId"           => "",
                "cargoCount"            => 1,
                "ttDocumentId"          => $query['sale']['id'],
                "dcSelectedCredit"      => "1",
                "dcCreditRule"          => "1",
                "ttInvoiceAmount"       => $query['sale']['grand_total'],
                "ttDocumentSaveType"    => "0"
            ];
            if($query['sale']['payment_type'] == 3){
                $args['ShippingOrderVO']['ttCollectionType'] = "0";
            }
            if($query['sale']['payment_type'] == 4){
                $args['ShippingOrderVO']['ttCollectionType'] = "1";
            }
            $this->data = array_merge($this->data, $args);
            $response = $this->soap->createShipment($this->data);
            return ["request"=>"success", "response"=>$response];
        }catch (Exception $e){
            return ["request"=>"error", "response"=>$e->getMessage()];
        }
    }

}