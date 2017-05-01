package legigo

// Used in Textelr and SectionTa and Article
type LienArt struct {
	Debut   string `xml:"debut,attr"`
	Etat    string `xml:"etat,attr"`
	Fin     string `xml:"fin,attr"`
	Id      string `xml:"id,attr"`
	Num     string `xml:"num,attr"`
	Origine string `xml:"origine,attr"`
}

// Used in Textelr and SectionTa
type LienSectionTa struct {
	Cid     string `xml:"cid,attr"`
	Debut   string `xml:"debut,attr"`
	Etat    string `xml:"etat,attr"`
	Fin     string `xml:"fin,attr"`
	Id      string `xml:"id,attr"`
	Niv     string `xml:"niv,attr"`
	Url     string `xml:"url,attr"`
	Contenu string `xml:",chardata"`
}

// Used in SectionTa and Article
type Contexte struct {
	Autorite      string   `xml:"autorite,attr"`
	Cid           string   `xml:"cid,attr"`
	DatePubli     string   `xml:"date_publi,attr"`
	DateSignature string   `xml:"date_signature,attr"`
	Ministere     string   `xml:"ministere,attr"`
	Nature        string   `xml:"nature,attr"`
	Nor           string   `xml:"nor,attr"`
	Num           string   `xml:"num,attr"`
	TitreTxt      TitreTxt `xml:"TITRE_TXT"`
	Tm            Tm       `xml:"TM"`
}

type TitreTxt struct {
	CTitreCourt string `xml:"c_titre_court,attr"`
	Debut       string `xml:"debut,attr"`
	Fin         string `xml:"fin,attr"`
	IdTxt       string `xml:"id_txt,attr"`
	Contenu     string `xml:",chardata"`
}

type Tm struct {
	TitreTm TitreTm `xml:"TITRE_TM"`
	Tm      *Tm     `xml:"TM"`
}

type TitreTm struct {
	Debut   string `xml:"debut,attr"`
	Fin     string `xml:"fin,attr"`
	Id      string `xml:"id,attr"`
	Contenu string `xml:",chardata"`
}

type Version struct {
	Etat    string  `xml:"etat,attr"`
	LienTxt LienTxt `xml:"LIEN_TXT"`
	LienArt LienArt `xml:"LIEN_ART"`
}

type LienTxt struct {
	Debut string `xml:"debut,attr"`
	Fin   string `xml:"fin,attr"`
	Id    string `xml:"id,attr"`
	Num   string `xml:"num,attr"`
}

type Lien struct {
	CidTexte       string `xml:"cidtexte,attr"`
	DateSignaTexte string `xml:"datesignatexte,attr"`
	Id             string `xml:"id,attr"`
	NatureTexte    string `xml:"naturetexte,attr"`
	NorTexte       string `xml:"nortexte,attr"`
	Num            string `xml:"num,attr"`
	NumTexte       string `xml:"numtexte,attr"`
	Sens           string `xml:"sens,attr"`
	TypeLien       string `xml:"typelien,attr"`
	Contenu        string `xml:",chardata"`
}
