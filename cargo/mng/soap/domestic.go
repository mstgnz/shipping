package soap

import (
	"encoding/xml"
)

type login struct {
	PKullaniciAdi string `xml:"pKullaniciAdi" json:"pKullaniciAdi"`
	PSifre        string `xml:"pSifre" json:"pSifre"`
	WsUserName    string `xml:"WsUserName" json:"WsUserName"`
	WsPassword    string `xml:"WsPassword" json:"WsPassword"`
	PMusteriNo    string `xml:"pMusteriNo" json:"pMusteriNo"`
}

// SiparisGirisiDetayliV3 http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=SiparisGirisiDetayliV3
type SiparisGirisiDetayliV3 struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SiparisGirisiDetayliV3"`
	login
	PChIrsaliyeNo        string `xml:"pChIrsaliyeNo" json:"pChIrsaliyeNo"`
	PPrKiymet            string `xml:"pPrKiymet" json:"pPrKiymet"`
	PChBarkod            string `xml:"pChBarkod" json:"pChBarkod"`
	PGonderiHizmetSekli  string `xml:"pGonderiHizmetSekli" json:"pGonderiHizmetSekli"`
	PTeslimSekli         string `xml:"pTeslimSekli" json:"pTeslimSekli"`
	PFlAlSms             string `xml:"pFlAlSms" json:"pFlAlSms"`
	PFlGnSms             string `xml:"pFlGnSms" json:"pFlGnSms"`
	PKargoParcaList      string `xml:"pKargoParcaList" json:"pKargoParcaList"`
	PAliciMusteriAdi     string `xml:"pAliciMusteriAdi" json:"pAliciMusteriAdi"`
	PChSiparisNo         string `xml:"pChSiparisNo" json:"pChSiparisNo"`
	PLuOdemeSekli        string `xml:"pLuOdemeSekli" json:"pLuOdemeSekli"`
	PFlAdresFarkli       string `xml:"pFlAdresFarkli" json:"pFlAdresFarkli"`
	PChIl                string `xml:"pChIl" json:"pChIl"`
	PChIlce              string `xml:"pChIlce" json:"pChIlce"`
	PChAdres             string `xml:"pChAdres" json:"pChAdres"`
	PChTelCep            string `xml:"pChTelCep" json:"pChTelCep"`
	PChEmail             string `xml:"pChEmail" json:"pChEmail"`
	PMalBedeliOdemeSekli string `xml:"pMalBedeliOdemeSekli" json:"pMalBedeliOdemeSekli"`
	PFlKapidaOdeme       string `xml:"pFlKapidaOdeme" json:"pFlKapidaOdeme"`
	PChIcerik            string `xml:"pChIcerik" json:"pChIcerik"`
	PAliciMusteriMngNo   string `xml:"pAliciMusteriMngNo" json:"pAliciMusteriMngNo"`
	PAliciMusteriBayiNo  string `xml:"pAliciMusteriBayiNo" json:"pAliciMusteriBayiNo"`
	PChSemt              string `xml:"pChSemt" json:"pChSemt"`
	PChMahalle           string `xml:"pChMahalle" json:"pChMahalle"`
	PChMeydanBulvar      string `xml:"pChMeydanBulvar" json:"pChMeydanBulvar"`
	PChCadde             string `xml:"pChCadde" json:"pChCadde"`
	PChSokak             string `xml:"pChSokak" json:"pChSokak"`
	PChFax               string `xml:"pChFax" json:"pChFax"`
	PChVergiDairesi      string `xml:"pChVergiDairesi" json:"pChVergiDairesi"`
	PChVergiNumarasi     string `xml:"pChVergiNumarasi" json:"pChVergiNumarasi"`
	PPlatformKisaAdi     string `xml:"pPlatformKisaAdi" json:"pPlatformKisaAdi"`
	PPlatformSatisKodu   string `xml:"pPlatformSatisKodu" json:"pPlatformSatisKodu"`
	PChTelEv             string `xml:"pChTelEv" json:"pChTelEv"`
	PChTelIs             string `xml:"pChTelIs" json:"pChTelIs"`
}

// KargoBilgileriByReferans http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=KargoBilgileriByReferans
type KargoBilgileriByReferans struct {
	login
	PSiparisNo  string `xml:"pSiparisNo" xml:"pSiparisNo"`
	PGonderiNo  string `xml:"pGonderiNo" json:"pGonderiNo"`
	PFaturaSeri string `xml:"pFaturaSeri" json:"pFaturaSeri"`
	PFaturaNo   string `xml:"pFaturaNo" json:"pFaturaNo"`
	PIrsaliyeNo string `xml:"pIrsaliyeNo" json:"pIrsaliyeNo"`
	EFaturaNo   string `xml:"eFaturaNo" json:"eFaturaNo"`
	PRaporType  string `xml:"pRaporType" json:"pRaporType"`
}

// MusteriSiparisIptal Cancel shipping - if the product has not shipped
// http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MusteriSiparisIptal
type MusteriSiparisIptal struct {
	login
	PMusteriSiparisNo string `xml:"pMusteriSiparisNo" json:"pMusteriSiparisNo"`
	PSiparisTarihi    string `xml:"pSiparisTarihi" json:"pSiparisTarihi"`
}

// MusteriTeslimatIptalIstegi Cancel shipping - if the product has been shipped
// http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MusteriTeslimatIptalIstegi
type MusteriTeslimatIptalIstegi struct {
	login
	PMusteriSiparisNo string `xml:"pMusteriSiparisNo" json:"pMusteriSiparisNo"`
	PIslemAciklama    string `xml:"pIslemAciklama" json:"pIslemAciklama"`
}

// MNGGonderiBarkod http://service.mngkargo.com.tr/musterikargosiparis/musterikargosiparis.asmx?op=MNGGonderiBarkod
type MNGGonderiBarkod struct {
	login
	IrsaliyeNo               string     `xml:"IrsaliyeNo" json:"IrsaliyeNo"`
	OutBarkodType            string     `xml:"OutBarkodType" json:"OutBarkodType"`
	FlKapidaTahsilat         string     `xml:"FlKapidaTahsilat" json:"FlKapidaTahsilat"`
	UrunBedeli               string     `xml:"UrunBedeli" json:"UrunBedeli"`
	ChMesaj                  string     `xml:"ChMesaj" json:"ChMesaj"`
	PChSiparisNo             string     `xml:"pChSiparisNo" json:"pChSiparisNo"`
	PChIrsaliyeNo            string     `xml:"pChIrsaliyeNo" json:"pChIrsaliyeNo"`
	ReferansNo               string     `xml:"ReferansNo" json:"ReferansNo"`
	PFlKapidaTahsilat        string     `xml:"pFlKapidaTahsilat" json:"pFlKapidaTahsilat"`
	EkString1                string     `xml:"EkString1" json:"EkString1"`
	EkString2                string     `xml:"EkString2" json:"EkString2"`
	EkString3                string     `xml:"EkString3" json:"EkString3"`
	EkString4                string     `xml:"EkString4" json:"EkString4"`
	HatadaReferansBarkoduBas string     `xml:"HatadaReferansBarkoduBas" json:"HatadaReferansBarkoduBas"`
	ParcaBilgi               ParcaBilgi `xml:"ParcaBilgi" json:"ParcaBilgi"`
}

type ParcaBilgi struct {
	GonderiParca []GonderiParca `xml:"GonderiParca" json:"GonderiParca"`
}

type GonderiParca struct {
	Kg          int    `xml:"Kg" json:"Kg"`
	Desi        int    `xml:"Desi" json:"Desi"`
	Adet        int    `xml:"Adet" json:"Adet"`
	Icerik      string `xml:"Icerik" json:"Icerik"`
	ParcaBarkod string `xml:"ParcaBarkod" json:"ParcaBarkod"`
}
