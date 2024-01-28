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
	PTeslimSekli         int    `xml:"pTeslimSekli"`
	PFlAlSms             int    `xml:"pFlAlSms"`
	PFlGnSms             int    `xml:"pFlGnSms"`
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
	PFlKapidaOdeme       int    `xml:"pFlKapidaOdeme"`
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
	PSiparisTarihi    string `xml:"pSiparisTarihi"`
}

// MusteriTeslimatIptalIstegi Cancel shipping - if the product has been shipped
// http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MusteriTeslimatIptalIstegi
type MusteriTeslimatIptalIstegi struct {
	login
	PMusteriSiparisNo string `xml:"pMusteriSiparisNo"`
	PIslemAciklama    string `xml:"pIslemAciklama"`
}

// MNGGonderiBarkod http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MNGGonderiBarkod
type MNGGonderiBarkod struct {
	login
	IrsaliyeNo               string     `xml:"IrsaliyeNo"`
	OutBarkodType            string     `xml:"OutBarkodType"`
	FlKapidaTahsilat         string     `xml:"FlKapidaTahsilat"`
	UrunBedeli               string     `xml:"UrunBedeli"`
	ChMesaj                  string     `xml:"ChMesaj"`
	PChSiparisNo             string     `xml:"pChSiparisNo"`
	PChIrsaliyeNo            string     `xml:"pChIrsaliyeNo"`
	ReferansNo               string     `xml:"ReferansNo"`
	PFlKapidaTahsilat        string     `xml:"pFlKapidaTahsilat"`
	EkString1                string     `xml:"EkString1"`
	EkString2                string     `xml:"EkString2"`
	EkString3                string     `xml:"EkString3"`
	EkString4                string     `xml:"EkString4"`
	HatadaReferansBarkoduBas string     `xml:"HatadaReferansBarkoduBas"`
	ParcaBilgi               ParcaBilgi `xml:"ParcaBilgi"`
}

type ParcaBilgi struct {
	GonderiParca []GonderiParca `xml:"GonderiParca"`
}

type GonderiParca struct {
	Kg          int    `xml:"Kg"`
	Desi        int    `xml:"Desi"`
	Adet        int    `xml:"Adet"`
	Icerik      string `xml:"Icerik"`
	ParcaBarkod string `xml:"ParcaBarkod"`
}
