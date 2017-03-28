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
