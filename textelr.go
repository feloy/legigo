package xml

import (
	"encoding/xml"
	"io/ioutil"
)

type Textelr struct {
	XMLName xml.Name `xml:"TEXTELR"`
	// META
	// - META_COMMUN
	Id       string `xml:"META>META_COMMUN>ID"`
	AncienId string `xml:"META>META_COMMUN>ANCIEN_ID"`
	Origine  string `xml:"META>META_COMMUN>ORIGINE"`
	Url      string `xml:"META>META_COMMUN>URL"`
	Nature   string `xml:"META>META_COMMUN>NATURE"`
	// - META_SPEC
	//   - META_TEXTE_CHRONICLE
	Cid                  string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>CID"`
	Num                  string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>NUM"`
	NumSequence          string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>NUM_SEQUENCE"`
	Nor                  string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>NOR"`
	DatePubli            string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>DATE_PUBLI"`
	DateTexte            string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>DATE_TEXTE"`
	DerniereModification string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>DERNIERE_MODIFICATION"`
	VersionsAVenir       []string `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>VERSIONS_A_VENIR>VERSION_A_VENIR"`
	OriginePubli         string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>ORIGINE_PUBLI"`
	PageDebPubli         string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>PAGE_DEB_PUBLI"`
	PageFinPubli         string   `xml:"META>META_SPEC>META_TEXTE_CHRONICLE>PAGE_FIN_PUBLI"`
	// VERSIONS
	Versions []Version `xml:"VERSIONS>VERSION"`
	// STRUCT
	LiensArt       []LienArt       `xml:"STRUCT>LIEN_ART"`
	LiensSectionTa []LienSectionTa `xml:"STRUCT>LIEN_SECTION_TA"`
}

func DecodeTextelr(filename string) (*Textelr, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v Textelr
	err = xml.Unmarshal(content, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
