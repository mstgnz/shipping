<?php

class UpsYd{

    private $userId;
    private $password;
    private $accessLicenseNumber;
    private $shipperNumber;
    private $endpoint;
    private $urls = [
        "test" => "https://wwwcie.ups.com/ups.app/xml",
        "live" => "https://onlinetools.ups.com/ups.app/xml"
    ];

    function __construct($userId, $password, $accessLicenseNumber, $shipperNumber, $urlMode="test"){
        if($userId && $password && $accessLicenseNumber && $shipperNumber){
            $this->userId = $userId;
            $this->password = $password;
            $this->accessLicenseNumber = $accessLicenseNumber;
            $this->shipperNumber = $shipperNumber;
        }
        $this->endpoint = $this->urls[$urlMode];
    }

    public function createShipment($query=[]){
        $args = [
            "shipToFullName"            => $query['customer']['firstname'].' '.$query['customer']['lastname'],
            "shipToCountryCode"         => $query['customer']['country_code'].' - '.$query['customer']['country_name'],
            "StateProvinceCode"         => $query['customer']['province_code'],
            "shipToCityName"            => $query['customer']['city'],
            "shipToDistrictName"        => $query['customer']['district'],
            "shipToPostalCode"          => $query['customer']['postal_code'],
            "shipToAddress"             => $query['customer']['address'],
            "shipToPhone"               => $query['customer']['phone'],
            "shipToEMail"               => $query['customer']['email'],
            "shipperFullName"           => $query['sender']['fullname'],
            "shipperAddress"            => $query['sender']['address'],
            "shipperCityName"           => $query['sender']['city_name'],
            "shipperDistrictName"       => $query['sender']['district_name'],
            "shipperPhone"              => $query['sender']['phone'],
            "shipperEmail"              => $query['sender']['email']
        ];
        $response = $this->shipConfirm($args);
        // shipConfirm success ise shipAccept istek at
        if($response['Response']['ResponseStatusCode']){
            $shipId = $response['ShipmentIdentificationNumber'];
            $digest = $response['ShipmentDigest'];
            $response = $this->shipAccept($digest);
            return $response;
        }else{
            return $response;
        }
    }

    public function shipAccept($digest){
        try {
            // Create AccessRequest XMl
            $accessRequestXML = new SimpleXMLElement ("<AccessRequest></AccessRequest>");
            $accessRequestXML->addChild ("AccessLicenseNumber", $this->accessLicenseNumber);
            $accessRequestXML->addChild ("UserId", $this->userId);
            $accessRequestXML->addChild ("Password", $this->password);
            // Create ShipmentAcceptRequest XMl
            $shipmentAcceptRequestXML = new SimpleXMLElement ("<ShipmentAcceptRequest ></ShipmentAcceptRequest>");
            $request = $shipmentAcceptRequestXML->addChild ('Request');
            $request->addChild ("RequestAction", "01");
            $shipmentAcceptRequestXML->addChild ("ShipmentDigest", $digest);
            $requestXML = $accessRequestXML->asXML () . $shipmentAcceptRequestXML->asXML ();
            $ch = curl_init();
            curl_setopt( $ch, CURLOPT_URL, $this->endpoint.'/ShipAccept');
            curl_setopt( $ch, CURLOPT_POST, true);
            curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded'));
            curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt( $ch, CURLOPT_POSTFIELDS, $requestXML);
            $response = curl_exec($ch);
            curl_close($ch);
            if ($response == false) {
                throw new Exception ("Bad data.");
            } else {
                $response = simplexml_load_string($response);
                $response = json_encode($response);
                $response = json_decode($response, true); 
                return $response;
            }
        } catch ( Exception $ex ) {
            return $ex;
        }
    }

    public function shipConfirm($args){
        $args['shipToCountryCode'] = explode('-',$args['shipToCountryCode']);
        $args['shipToCountryCode'] = trim($args['shipToCountryCode'][0]);
        list($shipperAddress1, $shipperAddress2, $shipperAddress3) = $this->smashAddress($args['shipperAddress']);
        list($shipToAddress1, $shipToAddress2, $shipToAddress3) = $this->smashAddress($args['shipToAddress']);
        try {
            // Create AccessRequest XMl
            $accessRequestXML = new SimpleXMLElement ("<AccessRequest></AccessRequest>");
            $accessRequestXML->addChild ("AccessLicenseNumber", $this->accessLicenseNumber);
            $accessRequestXML->addChild ("UserId", $this->userId);
            $accessRequestXML->addChild ("Password", $this->password);
            
            // Create ShipmentConfirmRequest XMl
            $shipmentConfirmRequestXML = new SimpleXMLElement ("<ShipmentConfirmRequest ></ShipmentConfirmRequest>");
            $request = $shipmentConfirmRequestXML->addChild ('Request');
            $request->addChild ("RequestAction", "ShipConfirm");
            $request->addChild ("RequestOption", "nonvalidate");
            
            /*$labelSpecification = $shipmentConfirmRequestXML->addChild ('LabelSpecification');
            $labelSpecification->addChild ("HTTPUserAgent", "");
            $labelPrintMethod = $labelSpecification->addChild ('LabelPrintMethod');
            $labelPrintMethod->addChild ("Code", "GIF");
            $labelPrintMethod->addChild ("Description", "");
            $labelImageFormat = $labelSpecification->addChild ('LabelImageFormat');
            $labelImageFormat->addChild ("Code", "GIF");
            $labelImageFormat->addChild ("Description", "");*/
            
            $shipment = $shipmentConfirmRequestXML->addChild ('Shipment');
            $shipment->addChild ("Description", "Kobisi.com");
            //$rateInformation = $shipment->addChild ('RateInformation');
            //$rateInformation->addChild ("NegotiatedRatesIndicator", "");
            
            $shipper = $shipment->addChild ('Shipper');
            $shipper->addChild ("Name", $args['shipperFullName']);
            $shipper->addChild ("AttentionName", $args['shipperFullName']);
            $shipper->addChild ("PhoneNumber", $args['shipperPhone']);
            $shipper->addChild ("TaxIdentificationNumber", $args['shipperPhone']);
            $shipper->addChild ("ShipperNumber", $this->shipperNumber);
            $shipperAddress = $shipper->addChild ('Address');
            $shipperAddress->addChild ("AddressLine1", $shipperAddress1);
            $shipperAddress->addChild ("AddressLine2", $shipperAddress2);
            $shipperAddress->addChild ("AddressLine3", $shipperAddress3);
            $shipperAddress->addChild ("City", $args['shipperCityName']);
            //$shipperAddress->addChild ("StateProvinceCode", "TR");
            $shipperAddress->addChild ("PostalCode", "34000");
            $shipperAddress->addChild ("CountryCode", "TR");

            $shipFrom = $shipment->addChild ('ShipFrom');
            $shipFrom->addChild ("CompanyName", "Kobisi.com");
            $shipFrom->addChild ("AttentionName", "Mesut GENEZ");
            $shipFrom->addChild ("PhoneNumber", $args['shipperPhone']);
            $shipFrom->addChild ("TaxIdentificationNumber", $args['shipperPhone']);
            $shipFromAddress = $shipFrom->addChild ('Address');
            $shipFromAddress->addChild ("AddressLine1", $shipperAddress1);
            $shipFromAddress->addChild ("AddressLine2", $shipperAddress2);
            $shipFromAddress->addChild ("AddressLine3", $shipperAddress3);
            $shipFromAddress->addChild ("City", $args['shipperCityName']);
            //$shipFromAddress->addChild ("StateProvinceCode", "TR");
            $shipFromAddress->addChild ("PostalCode", "34000");
            $shipFromAddress->addChild ("CountryCode", "TR");

            $shipTo = $shipment->addChild ('ShipTo');
            $shipTo->addChild ("CompanyName", $args['shipToFullName']);
            $shipTo->addChild ("AttentionName", $args['shipToFullName']);
            $shipTo->addChild ("PhoneNumber", $args['shipToPhone']);
            $shipToAddress = $shipTo->addChild ('Address');
            $shipToAddress->addChild ("AddressLine1", $shipToAddress1);
            $shipToAddress->addChild ("AddressLine2", $shipToAddress2);
            $shipToAddress->addChild ("AddressLine3", $shipToAddress3);
            $shipToAddress->addChild ("City", $args['shipToCityName']);
            $shipToAddress->addChild ("StateProvinceCode", $args['StateProvinceCode']);
            $shipToAddress->addChild ("PostalCode", $args['shipToPostalCode']);
            $shipToAddress->addChild ("CountryCode", $args['shipToCountryCode']);
            
            $paymentInformation = $shipment->addChild ('PaymentInformation');
            $prepaid = $paymentInformation->addChild ('Prepaid');
            $billShipper = $prepaid->addChild ('BillShipper');
            $billShipper->addChild ("AccountNumber", $this->shipperNumber);
            
            $service = $shipment->addChild ('Service');
            $service->addChild ("Code", "08"); // 07 = Express 08 = Expedited 54 = Express Plus 65 = UPS Saver
            $service->addChild ("Description", "");
            
            $package = $shipment->addChild ('Package');
            $package->addChild ("Description", "");

            $packagingType = $package->addChild ('PackagingType');
            $packagingType->addChild ("Code", "04");
            $packagingType->addChild ("Description", "");

            $packageWeight = $package->addChild ('PackageWeight');
            $packageWeight->addChild ("Weight", "1");
            $packageWeightUnitOfMeasurement = $packageWeight->addChild ('UnitOfMeasurement');
            $packageWeightUnitOfMeasurement->addChild("Code","KGS");

            $dimensions = $package->addChild ('Dimensions');
            $dimensions->addChild('Length','1');
            $dimensions->addChild('Width','1');
            $dimensions->addChild('Height','1');
            $dimensionsUnitOfMeasurement = $dimensions->addChild ('UnitOfMeasurement');
            $dimensionsUnitOfMeasurement->addChild('Code','CM');

            $requestXML = $accessRequestXML->asXML () . $shipmentConfirmRequestXML->asXML ();
            
            $ch = curl_init();
            curl_setopt( $ch, CURLOPT_URL, $this->endpoint.'/ShipConfirm');
            curl_setopt( $ch, CURLOPT_POST, true);
            curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded'));
            curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt( $ch, CURLOPT_POSTFIELDS, $requestXML);
            $response = curl_exec($ch);
            curl_close($ch);

            if ($response == false) {
                throw new Exception ("Bad data.");
            } else {
                $response = simplexml_load_string($response);
                $response = json_encode($response);
                $response = json_decode($response, true); 
                return $response;
            } 
        } catch ( Exception $ex ) {
            return $ex;
        }
    }
    
    public function labelRecovery(){
        try {
            // Create AccessRequest XMl
            $accessRequesttXML = new SimpleXMLElement ("<AccessRequest></AccessRequest>");
            $accessRequesttXML->addChild ("AccessLicenseNumber", $this->accessLicenseNumber);
            $accessRequesttXML->addChild ("UserId", $this->userId);
            $accessRequesttXML->addChild ("Password", $this->password);
            
            // Create LabelRecoveryRequest XMl
            $labelRecoveryRequestXML = new SimpleXMLElement ("<LabelRecoveryRequest ></LabelRecoveryRequest >");
            $request = $labelRecoveryRequestXML->addChild ('Request');
            $request->addChild ("RequestAction", "LabelRecovery");
            
            $labelSpecification = $labelRecoveryRequestXML->addChild ('LabelSpecification');
            $labelSpecification->addChild ("HTTPUserAgent");
            $labelImageFormat = $labelSpecification->addChild ('LabelImageFormat');
            $labelImageFormat->addChild ("Code", "GIF");
            
            $labelDelivery = $labelRecoveryRequestXML->addChild ('LabelDelivery');
            $labelDelivery->addChild ("LabelLinkIndicator");
            $labelDelivery->addChild ("ResendEMailIndicator");
            
            $labelRecoveryRequestXML->addChild ("TrackingNumber", "Your Tracking Number");
            
            $requestXML = $accessRequesttXML->asXML () . $labelRecoveryRequestXML->asXML ();
            
            $ch = curl_init();
            curl_setopt( $ch, CURLOPT_URL, $this->endpoint.'/LabelRecovery');
            curl_setopt( $ch, CURLOPT_POST, true);
            curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded'));
            curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt( $ch, CURLOPT_POSTFIELDS, $requestXML);
            $response = curl_exec($ch);
            curl_close($ch);
            
            if ($response == false) {
                throw new Exception ("Bad data.");
            } else {
                $response = simplexml_load_string($response);
                $response = json_encode($response);
                $response = json_decode($response, true); 
                return $response;
            }
            
            Header ('Content-type: text/xml');
        } catch ( Exception $ex ) {
            return $ex;
        }
    }

    public function shipCancel($shipId){
        try {
            // Create AccessRequest XMl
            $accessRequestXML = new SimpleXMLElement ("<AccessRequest></AccessRequest>");
            $accessRequestXML->addChild ("AccessLicenseNumber", $this->accessLicenseNumber);
            $accessRequestXML->addChild ("UserId", $this->userId);
            $accessRequestXML->addChild ("Password", $this->password);
            
            // Create VoidShipmentRequest XMl
            $voidShipmentRequestXML = new SimpleXMLElement ("<VoidShipmentRequest ></VoidShipmentRequest >");
            $request = $voidShipmentRequestXML->addChild ('Request');
            $request->addChild ("RequestAction", "1");
            $transaction = $request->addChild ("TransactionReference");
            $transaction->addChild('CustomerContext');
            
            $voidShipmentRequestXML->addChild ("ShipmentIdentificationNumber", $shipId);
            
            $requestXML = $accessRequestXML->asXML () . $voidShipmentRequestXML->asXML ();
            
            $ch = curl_init();
            curl_setopt( $ch, CURLOPT_URL, $this->endpoint.'/Void');
            curl_setopt( $ch, CURLOPT_POST, true);
            curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded'));
            curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt( $ch, CURLOPT_POSTFIELDS, $requestXML);
            $response = curl_exec($ch);
            curl_close($ch);
            if ($response == false) {
                throw new Exception ("Bad data.");
            } else {
                $response = simplexml_load_string($response);
                $response = json_encode($response);
                $response = json_decode($response, true); 
                return $response;
            }       
        } catch ( Exception $ex ) {
            echo $ex;
        }
    }

    public function shipTracking($shipId){
        try {
	
            // Create AccessRequest XMl
            $accessRequestXML = new SimpleXMLElement ("<AccessRequest></AccessRequest>");
            $accessRequestXML->addChild ("AccessLicenseNumber", $this->accessLicenseNumber);
            $accessRequestXML->addChild ("UserId", $this->userId);
            $accessRequestXML->addChild ("Password", $this->password);
            
            // Create TrackRequest XMl
            $trackRequestXML = new SimpleXMLElement ("<TrackRequest></TrackRequest  >");
            $request = $trackRequestXML->addChild ('Request');
            $request->addChild ("RequestAction", "Track");
            $request->addChild ("RequestOption", "activity");
            $transaction = $request->addChild ("TransactionReference");
            $transaction->addChild('CustomerContext');
            
            $trackRequestXML->addChild ("TrackingNumber", $shipId);
            
            $requestXML = $accessRequestXML->asXML () . $trackRequestXML->asXML ();
            
            $ch = curl_init();
            curl_setopt( $ch, CURLOPT_URL, $this->endpoint.'/Track');
            curl_setopt( $ch, CURLOPT_POST, true);
            curl_setopt( $ch, CURLOPT_HTTPHEADER, array('Content-Type: application/x-www-form-urlencoded'));
            curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);
            curl_setopt( $ch, CURLOPT_POSTFIELDS, $requestXML);
            $response = curl_exec($ch);
            curl_close($ch);
            if ($response == false) {
                throw new Exception ("Bad data.");
            } else {
                $response = simplexml_load_string($response);
                $response = json_encode($response);
                $response = json_decode($response, true); 
                return $response;
            }       
        } catch ( Exception $ex ) {
            echo $ex;
        }
    }

    private function smashAddress($address){
        $address1 = $address2 = $address3 = "";
        $smash = explode(' ',$address);
        foreach ($smash as $value) {
            if((strlen($address1) + strlen($value)) < 30){
                $address1 .= $value." ";
            }elseif((strlen($address2) + strlen($value)) < 30){
                $address2 .= $value." ";
            }elseif((strlen($address3) + strlen($value)) < 30){
                $address3 .= $value." ";
            }
        }
        return [$address1,$address2,$address3];
    }
}