package xml

import (
	"encoding/xml"
	"io/ioutil"
)

type SectionTa struct {
	XMLName        xml.Name        `xml:"SECTION_TA"`
	Id             string          `xml:"ID"`
	TitreTa        string          `xml:"TITRE_TA"`
	Commentaire    string          `xml:"COMMENTAIRE"`
	LiensArt       []LienArt       `xml:"STRUCTURE_TA>LIEN_ART"`
	LiensSectionTa []LienSectionTa `xml:"STRUCTURE_TA>LIEN_SECTION_TA"`
	Contexte       Contexte        `xml:"CONTEXTE>TEXTE"`
}

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

func DecodeSectionTa(filename string) (*SectionTa, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v SectionTa
	err = xml.Unmarshal(content, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
