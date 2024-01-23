<?php

class Mng
{

    private $soap, $urlMode;
    private $data = [];
    private $urls = [
        "test" => "https://pttws.ptt.gov.tr/PttVeriYuklemeTest/services/Sorgu?wsdl",
        "live" => "https://pttws.ptt.gov.tr/PttVeriYukleme/services/Sorgu?wsdl"
    ];

    function __construct($username = false, $password = false, $urlMode = "test")
    {
        if ($username && $password) {
            $this->data["input"]["musteriId"] = $username;
            $this->data["input"]["sifre"] = $password;
            $this->data["input"]["kullanici"] = "PttWs";
            $this->data["input"]["gonderiTip"] = "NORMAL";
            $this->data["input"]["gonderiTur"] = "KARGO";
        }
        $this->urlMode = $urlMode;
        $this->soap = new \SoapClient($this->urls[$this->urlMode], array('trace' => true));
    }

    // Kargo Oluştur
    public function kabulEkle2($query = [])
    {
        try {
            $this->data["input"]["dongu"] = [
                "aAdres"                => $query['customer']['address'],
                "aliciAdi"              => $query['customer']['firstname'] . ' ' . $query['customer']['lastname'],
                "aliciEmail"            => $query['customer']['email'],
                "aliciIlAdi"            => strtoupper($query['customer']['city']),
                "aliciIlceAdi"          => strtoupper($query['customer']['district']),
                "aliciSms"              => $query['customer']['phone'],
                "deger_ucreti"          => number_format($query['sale']['grand_total'], 2, ',', '.'),
                "musteriReferansNo"     => $query['sale']['id'],
                "odeme_sart_ucreti"     => number_format($query['sale']['grand_total'], 2, ',', '.'),
                "OdemeSekli"            => in_array($query['sale']['payment_type'], [3, 4]) ? "UA" : "MH",
                "aAdres2"               => "",
                "aAdres3"               => "",
                "aIlKodu"               => "",
                "aIlKodu2"              => "",
                "aIlKodu3"              => "",
                "aIlceKodu"             => "",
                "aIlceKodu2"            => "",
                "aIlceKodu3"            => "",
                "agirlik"               => "1000",
                "aliciAdi2"             => "",
                "aliciAdi3"             => "",
                "aliciEmail2"           => "",
                "aliciEmail3"           => "",
                "aliciIlAdi2"           => "",
                "aliciIlAdi3"           => "",
                "aliciIlceAdi2"         => "",
                "aliciIlceAdi3"         => "",
                "aliciSms2"             => "",
                "aliciSms3"             => "",
                "aliciTel"              => "",
                "aliciTel2"             => "",
                "aliciTel3"             => "",
                "boy"                   => "",
                "desi"                  => "1",
                "ekhizmet"              => "",
                "en"                    => "",
                "iadeAAdres"            => "",
                "iadeAIlKodu"           => "",
                "iadeAIlceKodu"         => "",
                "iadeAliciAdi"          => "",
                "iadeAliciEmail"        => "",
                "iadeAliciIlAdi"        => "",
                "iadeAliciIlceAdi"      => "",
                "iadeAliciTel"          => "",
                "parcali_barkod_referans_no"    => "",
                "rezerve1"              => "",
                "rezerve2"              => "",
                "rezerve3"              => "",
                "rezerve4"              => "",
                "rezerve5"              => "",
                "teslim_tip"            => "",
                "ucret"                 => "",
                "yukseklik"             => "",
                "barkodNo"              => ""
            ];
            $this->data["input"]["dosyaAdi"] = $this->data["input"]["musteriId"] . "-" . $query['sale']['id'];
            $response = $this->soap->kabulEkle2($this->data);
            return ["request" => "success", "response" => $response];
        } catch (Exception $e) {
            return ["request" => "error", "response" => $e->getMessage()];
        }
    }
}
