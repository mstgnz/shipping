package soap

type createData struct {
	input input
}

type input struct {
	musteriId  string
	sifre      string
	kullanici  string
	gonderiTip string
	gonderiTur string
	dosyaAdi   string
	dongu      dongu
}

type dongu struct {
	aAdres                     string
	aliciAdi                   string
	aliciEmail                 string
	aliciIlAdi                 string
	aliciIlceAdi               string
	aliciSms                   string
	deger_ucreti               string
	musteriReferansNo          string
	odeme_sart_ucreti          string
	OdemeSekli                 string
	aAdres2                    string
	aAdres3                    string
	aIlKodu                    string
	aIlKodu2                   string
	aIlKodu3                   string
	aIlceKodu                  string
	aIlceKodu2                 string
	aIlceKodu3                 string
	agirlik                    string
	aliciAdi2                  string
	aliciAdi3                  string
	aliciEmail2                string
	aliciEmail3                string
	aliciIlAdi2                string
	aliciIlAdi3                string
	aliciIlceAdi2              string
	aliciIlceAdi3              string
	aliciSms2                  string
	aliciSms3                  string
	aliciTel                   string
	aliciTel2                  string
	aliciTel3                  string
	boy                        string
	desi                       string
	ekhizmet                   string
	en                         string
	iadeAAdres                 string
	iadeAIlKodu                string
	iadeAIlceKodu              string
	iadeAliciAdi               string
	iadeAliciEmail             string
	iadeAliciIlAdi             string
	iadeAliciIlceAdi           string
	iadeAliciTel               string
	rezerve1                   string
	rezerve2                   string
	rezerve3                   string
	rezerve4                   string
	rezerve5                   string
	teslim_tip                 string
	ucret                      string
	yukseklik                  string
	barkodNo                   string
	parcali_barkod_referans_no string
}
