package xml

import (
	"encoding/xml"
	"io/ioutil"
)

type TexteVersion struct {
	XMLName xml.Name `xml:"TEXTE_VERSION"`
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
	//   - META_TEXTE_VERSION
	Autorite  string `xml:"META>META_SPEC>META_TEXTE_VERSION>AUTORITE"`
	DateDebut string `xml:"META>META_SPEC>META_TEXTE_VERSION>DATE_DEBUT"`
	DateFin   string `xml:"META>META_SPEC>META_TEXTE_VERSION>DATE_FIN"`
	Etat      string `xml:"META>META_SPEC>META_TEXTE_VERSION>ETAT"`
	Liens     []Lien `xml:"META>META_SPEC>META_TEXTE_VERSION>LIENS>LIEN"`
	Ministere string `xml:"META>META_SPEC>META_TEXTE_VERSION>MINISTERE"`
	Titre     string `xml:"META>META_SPEC>META_TEXTE_VERSION>TITRE"`
	TitreFull string `xml:"META>META_SPEC>META_TEXTE_VERSION>TITREFULL"`
	Visas     struct {
		Contenu string `xml:",innerxml"`
	} `xml:"VISAS>CONTENU"`
	Signataires struct {
		Contenu string `xml:",innerxml"`
	} `xml:"SIGNATAIRES>CONTENU"`
	Tp struct {
		Contenu string `xml:",innerxml"`
	} `xml:"TP>CONTENU"`
	Nota struct {
		Contenu string `xml:",innerxml"`
	} `xml:"NOTA>CONTENU"`
	Abro struct {
		Contenu string `xml:",innerxml"`
	} `xml:"ABRO>CONTENU"`
	Rect struct {
		Contenu string `xml:",innerxml"`
	} `xml:"RECT>CONTENU"`
}

func DecodeTexteVersion(filename string) (*TexteVersion, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v TexteVersion
	err = xml.Unmarshal(content, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
