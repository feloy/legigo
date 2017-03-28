package xml

import (
	"encoding/xml"
	"io/ioutil"
)

type Article struct {
	XMLName        xml.Name        `xml:"ARTICLE"`
	// META
	// - META_COMMUN
	Id       string `xml:"META>META_COMMUN>ID"`
	AncienId string `xml:"META>META_COMMUN>ANCIEN_ID"`
	Origine  string `xml:"META>META_COMMUN>ORIGINE"`
	Url      string `xml:"META>META_COMMUN>URL"`
	Nature   string `xml:"META>META_COMMUN>NATURE"`
	// - META_SPEC
	//   - META_ARTICLE
	DateDebut string `xml:"META>META_SPEC>META_ARTICLE>DATE_DEBUT"`
	DateFin string `xml:"META>META_SPEC>META_ARTICLE>DATE_FIN"`
	Etat string `xml:"META>META_SPEC>META_ARTICLE>ETAT"`
	Num string `xml:"META>META_SPEC>META_ARTICLE>NUM"`
	Type string `xml:"META>META_SPEC>META_ARTICLE>TYPE"`

	Contexte Contexte `xml:"CONTEXTE>TEXTE"`

	Versions []Version `xml:"VERSIONS>VERSION"`
	
	BlocTextuel struct {
		Contenu string `xml:",innerxml"`
	} `xml:"BLOC_TEXTUEL>CONTENU"`
	Nota struct {
		Contenu string `xml:",innerxml"`
	} `xml:"NOTA>CONTENU"`

	Liens []Lien `xml:"LIENS>LIEN"`
}

func DecodeArticle(filename string) (*Article, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v Article
	err = xml.Unmarshal(content, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
