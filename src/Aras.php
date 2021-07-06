<?php

class Aras{

    private $soap, $urlMode;
    private $data = [
        "userName" => "neodyum",
        "password" => "nd2580"
    ];
    private $urls = [
        "test" => "https://customerservicestest.araskargo.com.tr/arascargoservice/arascargoservice.asmx?WSDL",
        "live" => "https://customerws.araskargo.com.tr/arascargoservice.asmx?WSDL"
    ];

    function __construct($username=false, $password=false, $urlMode="test"){
        if($username && $password && $urlMode=="live"){
            $this->data["userName"] = $username;
            $this->data["password"] = $password;
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
    }

    // Kargo Oluştur
    public function SetOrder($query=[]){
        try{
            $args = array(
                "orderInfo"=>array(
                    "Order"=> array(
                        "TradingWaybillNumber"  => $query['sale']['id'],
                        "InvoiceNumber" 		=> "",
                        "IntegrationCode"		=> $query['sale']['id'],
                        "ReceiverName"			=> $query['customer']['firstname'].' '.$query['customer']['lastname'],
                        "ReceiverAddress"		=> $query['customer']['address'],
                        "ReceiverCityName"		=> $query['customer']['city'],
                        "ReceiverTownName"		=> $query['customer']['district'],
                        "ReceiverDistrictName"  => "",
                        "ReceiverQuarterName"	=> "",
                        "ReceiverAvenueName"	=> "",
                        "ReceiverStreetName"	=> "",
                        "ReceiverPhone1"		=> $query['customer']['phone'],
                        "ReceiverPhone2"		=> "",
                        "ReceiverPhone3"		=> "",
                        "VolumetricWeight"		=> "",
                        "Weight"				=> "",
                        "CityCode"				=> "",
                        "TownCode"				=> "",
                        "SpecialField1"		    => "",
                        "SpecialField2"		    => "",
                        "SpecialField3"		    => "",
                        "IsCod"				    => in_array($query['sale']['payment_type'], [3,4]) ? "1" : "0",
                        "CodAmount"			    => number_format($query['sale']['grand_total'],2,',','.'),
                        "CodCollectionType"	    => "1",
                        "CodBillingType"		=> "0",
                        "TaxNumber"			    => "",
                        "TaxOffice"			    => "",
                        "PayorTypeCode"		    => in_array($query['sale']['payment_type'], [3,4]) ? "2" : "1",
                        "IsWorldWide"			=> "0",
                        "Description"			=> "",
                        "PrivilegeOrder"        => "",
                        "PieceCount"            => 1,
                        "PieceDetails"          => array(
                            "PieceDetail"       => array(
                                "VolumetricWeight"  => "",
                                "Weight"            => "",
                                "BarcodeNumber"     => $query['sale']['id'],
                                "ProductNumber"     => "",
                                "Description"       => ""
                            )
                        )
                    )
                )
            );
            if($query['sale']['payment_type'] == 3){
                $args['orderInfo']['Order']['CodCollectionType'] = "0";
            }
            if($query['sale']['payment_type'] == 4){
                $args['orderInfo']['Order']['CodCollectionType'] = "1";
            }
            $this->data = array_merge($this->data, $args);
            $response = $this->soap->SetOrder($this->data);
            return ["request"=>"success", "response"=>$response];
        }catch (Exception $e){
            return ["request"=>"error", "response"=>$e->getMessage()];
        }
    }

    private function randStr($length = 10) {
		$str = "";
		$characters = array_merge(range('A','Z'), range('a','z'), range('0','9'));
		$max = count($characters) - 1;
		for ($i = 0; $i < $length; $i++) {
			$rand = mt_rand(0, $max);
			$str .= $characters[$rand];
		}
		return $str;
	}
    

}

