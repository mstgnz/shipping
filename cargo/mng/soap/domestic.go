package soap

type login struct {
	PKullaniciAdi string `xml:"pKullaniciAdi"`
	PSifre        string `xml:"pSifre"`
	WsUserName    string `xml:"WsUserName"`
	WsPassword    string `xml:"WsPassword"`
	PMusteriNo    string `xml:"pMusteriNo"`
}

// SiparisGirisiDetayliV3 http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=SiparisGirisiDetayliV3
type SiparisGirisiDetayliV3 struct {
	login
	PChIrsaliyeNo        string `xml:"pChIrsaliyeNo"`
	PPrKiymet            string `xml:"pPrKiymet"`
	PChBarkod            string `xml:"pChBarkod"`
	PGonderiHizmetSekli  string `xml:"pGonderiHizmetSekli"`
	PTeslimSekli         string `xml:"pTeslimSekli"`
	PFlAlSms             string `xml:"pFlAlSms"`
	PFlGnSms             string `xml:"pFlGnSms"`
	PKargoParcaList      string `xml:"pKargoParcaList"`
	PAliciMusteriAdi     string `xml:"pAliciMusteriAdi"`
	PChSiparisNo         string `xml:"pChSiparisNo"`
	PLuOdemeSekli        string `xml:"pLuOdemeSekli"`
	PFlAdresFarkli       string `xml:"pFlAdresFarkli"`
	PChIl                string `xml:"pChIl"`
	PChIlce              string `xml:"pChIlce"`
	PChAdres             string `xml:"pChAdres"`
	PChTelCep            string `xml:"pChTelCep"`
	PChEmail             string `xml:"pChEmail"`
	PMalBedeliOdemeSekli string `xml:"pMalBedeliOdemeSekli"`
	PFlKapidaOdeme       string `xml:"pFlKapidaOdeme"`
	PChIcerik            string `xml:"pChIcerik"`
	PAliciMusteriMngNo   string `xml:"pAliciMusteriMngNo"`
	PAliciMusteriBayiNo  string `xml:"pAliciMusteriBayiNo"`
	PChSemt              string `xml:"pChSemt"`
	PChMahalle           string `xml:"pChMahalle"`
	PChMeydanBulvar      string `xml:"pChMeydanBulvar"`
	PChCadde             string `xml:"pChCadde"`
	PChSokak             string `xml:"pChSokak"`
	PChFax               string `xml:"pChFax"`
	PChVergiDairesi      string `xml:"pChVergiDairesi"`
	PChVergiNumarasi     string `xml:"pChVergiNumarasi"`
	PPlatformKisaAdi     string `xml:"pPlatformKisaAdi"`
	PPlatformSatisKodu   string `xml:"pPlatformSatisKodu"`
	PChTelEv             string `xml:"pChTelEv"`
	PChTelIs             string `xml:"pChTelIs"`
}

// KargoBilgileriByReferans http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=KargoBilgileriByReferans
type KargoBilgileriByReferans struct {
	login
	PSiparisNo  string `xml:"pSiparisNo"`
	PGonderiNo  string `xml:"pGonderiNo"`
	PFaturaSeri string `xml:"pFaturaSeri"`
	PFaturaNo   string `xml:"pFaturaNo"`
	PIrsaliyeNo string `xml:"pIrsaliyeNo"`
	EFaturaNo   string `xml:"eFaturaNo"`
	PRaporType  string `xml:"pRaporType"`
}

// MusteriSiparisIptal Cancel shipping - if the product has not shipped
// http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MusteriSiparisIptal
type MusteriSiparisIptal struct {
	login
	PMusteriSiparisNo string `xml:"pMusteriSiparisNo"`
}

// MusteriTeslimatIptalIstegi Cancel shipping - if the product has been shipped
// http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MusteriTeslimatIptalIstegi
type MusteriTeslimatIptalIstegi struct {
	login
	PMusteriSiparisNo string `xml:"pMusteriSiparisNo"`
}

// MNGGonderiBarkod http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MNGGonderiBarkod
type MNGGonderiBarkod struct {
	login
	PChSiparisNo             string `xml:"pChSiparisNo"`
	PChIrsaliyeNo            string `xml:"pChIrsaliyeNo"`
	ReferansNo               string `xml:"ReferansNo"`
	IrsaliyeNo               string `xml:"IrsaliyeNo"`
	OutBarkodType            string `xml:"OutBarkodType"`
	PFlKapidaTahsilat        string `xml:"pFlKapidaTahsilat"`
	FlKapidaTahsilat         string `xml:"FlKapidaTahsilat"`
	HatadaReferansBarkoduBas string `xml:"HatadaReferansBarkoduBas"`
	UrunBedeli               string `xml:"UrunBedeli"`
	ChMesaj                  string `xml:"ChMesaj"`
	EkString1                string `xml:"EkString1"`
	EkString2                string `xml:"EkString2"`
	EkString3                string `xml:"EkString3"`
	EkString4                string `xml:"EkString4"`
	ParcaBilgi               string `xml:"ParcaBilgi"`
}
