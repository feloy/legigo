package xml

// Used in Textelr and SectionTa
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
