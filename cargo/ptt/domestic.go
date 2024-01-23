package ptt

type PttData struct {
	Input input `json:"input"`
}

type input struct {
	MusteriId  string `json:"musteriId"`
	Sifre      string `json:"sifre"`
	Kullanici  string `json:"kullanici"`
	GonderiTip string `json:"gonderiTip"`
	GonderiTur string `json:"gonderiTur"`
	DosyaAdi   string `json:"dosyaAdi"`
	Dongu      dongu  `json:"dongu"`
}

type dongu struct {
	AAdres                  string `json:"aAdres"`
	AliciAdi                string `json:"aliciAdi"`
	AliciEmail              string `json:"aliciEmail"`
	AliciIlAdi              string `json:"aliciIlAdi"`
	AliciIlceAdi            string `json:"aliciIlceAdi"`
	AliciSms                string `json:"aliciSms"`
	DegerUcreti             string `json:"deger_ucreti"`
	MusteriReferansNo       string `json:"musteriReferansNo"`
	OdemeSartUcreti         string `json:"odeme_sart_ucreti"`
	OdemeSekli              string `json:"OdemeSekli"`
	AAdres2                 string `json:"aAdres2"`
	AAdres3                 string `json:"aAdres3"`
	AIlKodu                 string `json:"aIlKodu"`
	AIlKodu2                string `json:"aIlKodu2"`
	AIlKodu3                string `json:"aIlKodu3"`
	AIlceKodu               string `json:"aIlceKodu"`
	AIlceKodu2              string `json:"aIlceKodu2"`
	AIlceKodu3              string `json:"aIlceKodu3"`
	Agirlik                 string `json:"agirlik"`
	AliciAdi2               string `json:"aliciAdi2"`
	AliciAdi3               string `json:"aliciAdi3"`
	AliciEmail2             string `json:"aliciEmail2"`
	AliciEmail3             string `json:"aliciEmail3"`
	AliciIlAdi2             string `json:"aliciIlAdi2"`
	AliciIlAdi3             string `json:"aliciIlAdi3"`
	AliciIlceAdi2           string `json:"aliciIlceAdi2"`
	AliciIlceAdi3           string `json:"aliciIlceAdi3"`
	AliciSms2               string `json:"aliciSms2"`
	AliciSms3               string `json:"aliciSms3"`
	AliciTel                string `json:"aliciTel"`
	AliciTel2               string `json:"aliciTel2"`
	AliciTel3               string `json:"aliciTel3"`
	Boy                     string `json:"boy"`
	Desi                    string `json:"desi"`
	Ekhizmet                string `json:"ekhizmet"`
	En                      string `json:"en"`
	IadeAAdres              string `json:"iadeAAdres"`
	IadeAIlKodu             string `json:"iadeAIlKodu"`
	IadeAIlceKodu           string `json:"iadeAIlceKodu"`
	IadeAliciAdi            string `json:"iadeAliciAdi"`
	IadeAliciEmail          string `json:"iadeAliciEmail"`
	IadeAliciIlAdi          string `json:"iadeAliciIlAdi"`
	IadeAliciIlceAdi        string `json:"iadeAliciIlceAdi"`
	IadeAliciTel            string `json:"iadeAliciTel"`
	Rezerve1                string `json:"rezerve1"`
	Rezerve2                string `json:"rezerve2"`
	Rezerve3                string `json:"rezerve3"`
	Rezerve4                string `json:"rezerve4"`
	Rezerve5                string `json:"rezerve5"`
	TeslimTip               string `json:"teslim_tip"`
	Ucret                   string `json:"ucret"`
	Yukseklik               string `json:"yukseklik"`
	BarkodNo                string `json:"barkodNo"`
	ParcaliBarkodReferansNo string `json:"parcali_barkod_referans_no"`
}
