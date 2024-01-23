<?php

class UpsTr{

    private $username;
    private $password;
    private $customerNumber;
    private $sessionID;
    private $urlMode;
    private $soap;

    private $urls = [
        "test" => "http://ws.ups.com.tr/wsCreateShipment/wsCreateShipment.asmx?WSDL",
        "live" => "http://ws.ups.com.tr/wsCreateShipment/wsCreateShipment.asmx?WSDL"
    ];

    function __construct($username = false, $password = false, $customerNumber = false, $urlMode="test"){
        if($username && $password && $customerNumber){
            $this->username = $username;
            $this->password = $password;
            $this->customerNumber = $customerNumber;
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
    }

    // Session
    public function Login_Type1(){
        try{
            $data = [
                "CustomerNumber" => $this->customerNumber,
                "UserName" => $this->username,
                "Password" => $this->password
            ];
            $res = $this->soap->Login_Type1($data);
            if(!empty($res->Login_Type1Result->SessionID)){
                $this->sessionID = $res->Login_Type1Result->SessionID;
                return $this->sessionID;
            }
            return false;
        }catch (Exception $e){
            return false;
        }
    }

    // Session
    public function Login_V1(){
        try{
            $data = [
                "CustomerNumber" => $this->customerNumber,
                "UserName" => $this->username,
                "Password" => $this->password
            ];
            $res = $this->soap->Login_V1($data);
            if(!empty($res->Login_V1Result->SessionID)){
                $this->sessionID = $res->Login_V1Result->SessionID;
                return $this->sessionID;
            }
            return false;
        }catch (Exception $e){
            return false;
        }
    }

    // Kargo Oluştur
    public function CreateShipment_Type2($query=[], $areaCode){
        if($this->Login_Type1()){
            try{
                $args = [
                    "ShipperName"                   => $query['sender']['fullname'],
                    "ShipperAddress"                => $query['sender']['address'],
                    "ShipperCityCode"               => $query['sender']['city_id'],
                    "ShipperAreaCode"               => $areaCode[0]['ups_districts']['ups_district_id'],
                    "ShipperMobilePhoneNumber"      => $query['sender']['phone'],
                    "ShipperEMail"                  => $query['sender']['email'],
                    "ConsigneeName"                 => $query['customer']['firstname'].' '.$query['customer']['lastname'],
                    "ConsigneeAddress"              => $query['customer']['address'],
                    "ConsigneeCityCode"             => $query['customer']['city_id'],
                    "ConsigneeAreaCode"             => $areaCode[0]['ups_districts']['ups_district_id'],
                    "ConsigneeMobilePhoneNumber"    => $query['customer']['phone'],
                    "ConsigneeEMail"                => $query['customer']['email'],
                    "PaymentType"                   => in_array($query['sale']['payment_type'], [3,4]) ? 1 : 2,
                    "SmsToConsignee"                => 1,
                    "ShipperAccountNumber"          => $query['cargo']['cust_id'],
                    "ShipperContactName"            => "",
                    "ShipperPostalCode"             => "",
                    "ShipperPhoneExtension"         => "",
                    "ShipperPhoneNumber"            => "",
                    "ShipperExpenseCode"            => "",
                    "ConsigneeAccountNumber"        => "",
                    "ConsigneeContactName"          => "",
                    "ConsigneeExpenseCode"          => "",
                    "ConsigneePostalCode"           => "",
                    "ConsigneePhoneNumber"          => "",
                    "CustomerReferance"             => "",
                    "CustomerInvoiceNumber"         => "",
                    "DescriptionOfGoods"            => "",
                    "DeliveryNotificationEmail"     => "",
                    "NumberOfPackages"              => 1,
                    "PackageType"                   => "K",
                    "ServiceLevel"                  => 3,
                    "IdControlFlag"                 => 0,
                    "PhonePrealertFlag"             => 0,
                    "InsuranceValue"                => 0,
                    "InsuranceValueCurrency"        => "",
                    "ValueOfGoods"                  => 0,
                    "ValueOfGoodsCurrency"          => "",
                    "ValueOfGoodsPaymentType"       => 0,
                    "ThirdPartyAccountNumber"       => "",
                    "ThirdPartyExpenseCode"         => "",
                    "SmsToShipper"                  => 0
                ];
                $data = [
                    "SessionID" => $this->sessionID,
                    "ShipmentInfo" => $args,
                    "ReturnLabelLink" => true,
                    "ReturnLabelImage" => true
                ];
                $res = $this->soap->CreateShipment_Type2($data);
                return ["request"=>"success", "response"=>$res];
            }catch (Exception $e){
                return ["request"=>"error", "response"=>$e->getMessage()];
            }
        }else{
            return ["request"=>"error", "response"=>"Oturum açılamadı!"];
        }
    }

    // Kargo iptal
    public function Cancel_Shipment_V1($customerCode, $waybillNumber){
        if($this->Login_Type1()){
            $data = [
                "sessionId" => $this->sessionID,
                "customerCode" => $customerCode,
                "waybillNumber" => $waybillNumber
            ];
            try{ 
                $res = $this->soap->Cancel_Shipment_V1($data);
                return ["request"=>"success", "response"=>$res];
            }catch (Exception $e){
                return ["request"=>"error", "response"=>$e->getMessage()];
            }
        }else{
            return ["request"=>"error", "response"=>"Oturum açılamadı!"];
        }
    }

    // Kargom Nerede
    public function GetShipmentInfoByTrackingNumber_V1($trackingNumber){
        $url = "http://ws.ups.com.tr/QueryPackageInfo/wsQueryPackagesInfo.asmx?WSDL";
        $this->soap = new \SoapClient($url, array('trace' => true));
        if($this->Login_V1()){
            $data = [
                "SessionID" => $this->sessionID,
                "InformationLevel" => 1,
                "TrackingNumber" => $trackingNumber
            ];
            try{
                $res = $this->soap->GetShipmentInfoByTrackingNumber_V1($data);
                return $res->GetShipmentInfoByTrackingNumber_V1Result->PackageInformation[0];
            }catch (Exception $e){
                return ["request"=>"error", "response"=>$e->getMessage()];
            }
        }else{
            return ["request"=>"error", "response"=>"Oturum açılamadı!"];
        }
    }

    private function xmlOutput($array){
        $xml_data = new SimpleXMLElement('<?xml version="1.0"?><data></data>');
        $this->array_to_xml($array,$xml_data);
        header('Content-Type: application/xml; charset=utf-8');
        echo'<pre>';
        print_r($xml_data->asXML());
        exit;
    }

    private function array_to_xml( $data, &$xml_data ) {
        foreach( $data as $key => $value ) {
            if( is_array($value) ) {
                if( is_numeric($key) ){
                    $key = 'item'.$key;
                }
                $subnode = $xml_data->addChild($key);
                $this->array_to_xml($value, $subnode);
            } else {
                $xml_data->addChild("$key",htmlspecialchars("$value"));
            }
        }
    }
}