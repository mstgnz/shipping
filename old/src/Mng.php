<?php

class Mng{

    private $soap, $urlMode;
    private $data = [];
    private $urls = [
        "test" => "http://service.mngkargo.com.tr/tservis/musterikargosiparis.asmx?WSDL",
        "live" => "http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?WSDL"
    ];

    function __construct($username=false, $password=false, $urlMode="test"){
        if($username && $password){
            $this->data["pKullaniciAdi"] = $username;
            $this->data["pSifre"] = $password;
            $this->data["WsUserName"] = $username;
            $this->data["WsPassword"] = $password;
            $this->data["pMusteriNo"] = $username;
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
    }

    // Kargo Oluştur
    public function SiparisGirisiDetayliV3($query=[]){
        try{
            $args = [
                "pChIrsaliyeNo"         => $query['sale']['id'],
                "pPrKiymet"             => number_format($query['sale']['grand_total'],2,',','.'),
                "pChBarkod"             => $query['sale']['id'],
                "pGonderiHizmetSekli"   => "NORMAL",
                "pTeslimSekli"          => 1,
                "pFlAlSms"              => 0,
                "pFlGnSms"              => 0,
                "pKargoParcaList"       => "1:1:1:koli:1:;",
                "pAliciMusteriAdi"      => $query['customer']['firstname'].' '.$query['customer']['lastname'],
                "pChSiparisNo"          => $query['sale']['id'],
                "pLuOdemeSekli"         => in_array($query['sale']['payment_type'], [3,4]) ? "U" : "P",
                "pFlAdresFarkli"        => "0",
                "pChIl"                 => strtoupper($query['customer']['city']),
                "pChIlce"               => strtoupper($query['customer']['district']),
                "pChAdres"              => $query['customer']['address'],
                "pChTelCep"             => $query['customer']['phone'],
                "pChEmail"              => $query['customer']['email'],
                "pMalBedeliOdemeSekli"  => "KREDI_KARTI",
                "pFlKapidaOdeme"        => in_array($query['sale']['payment_type'], [3,4]) ? 1 : 0,
                "pChIcerik"             => "",
                "pAliciMusteriMngNo"    => "",
                "pAliciMusteriBayiNo"   => "",
                "pChSemt"               => "",
                "pChMahalle"            => "",
                "pChMeydanBulvar"       => "",
                "pChCadde"              => "",
                "pChSokak"              => "",
                "pChFax"                => "",
                "pChVergiDairesi"       => "",
                "pChVergiNumarasi"      => "",
                "pPlatformKisaAdi"      => "",
                "pPlatformSatisKodu"    => "",
                "pChTelEv"              => $query['customer']['phone'],
                "pChTelIs"              => $query['customer']['phone']
            ];
            $this->data = array_merge($this->data, $args);
            $response = $this->soap->SiparisGirisiDetayliV3($this->data);
            return ["request"=>"success", "response"=>$response];
        }catch (Exception $e){
            return ["request"=>"error", "response"=>$e->getMessage()];
        }
    }

}