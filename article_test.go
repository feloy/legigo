package xml

import (
//	"fmt"
	"strings"
	"testing"
)

func TestDecodeArticleUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeArticle(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}

func TestDecodeArticleInvalid(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000005627819.xml"
	_, err := DecodeArticle(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}

func TestDecodeArticle(t *testing.T) {
	f := "specs/article_LEGIARTI000006291810.xml"
	v, err := DecodeArticle(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "LEGIARTI000006291810", v.Id)
	assertStringEquals(t, "AZAXXXXXXXX3X215D16AXXAA", v.AncienId)
	assertStringEquals(t, "LEGI", v.Origine)
	assertStringEquals(t, "article/LEGI/ARTI/00/00/06/29/18/LEGIARTI000006291810.xml", v.Url)
	assertStringEquals(t, "Article", v.Nature)

	assertStringEquals(t, "2007-10-18", v.DateDebut)
	assertStringEquals(t, "2012-11-11", v.DateFin)
	assertStringEquals(t, "MODIFIE", v.Etat)
	assertStringEquals(t, "D215-16", v.Num)
	assertStringEquals(t, "AUTONOME", v.Type)

	assertStringEquals(t, "", v.Contexte.Autorite)
	assertStringEquals(t, "LEGITEXT000006069565", v.Contexte.Cid)
	assertStringEquals(t, "2999-01-01", v.Contexte.DatePubli)
	assertStringEquals(t, "2999-01-01", v.Contexte.DateSignature)
	assertStringEquals(t, "", v.Contexte.Ministere)
	assertStringEquals(t, "CODE", v.Contexte.Nature)
	assertStringEquals(t, "", v.Contexte.Nor)
	assertStringEquals(t, "", v.Contexte.Num)

	assertStringEquals(t, "Code de la consommation", v.Contexte.TitreTxt.CTitreCourt)
	assertStringEquals(t, "1993-07-27", v.Contexte.TitreTxt.Debut)
	assertStringEquals(t, "2999-01-01", v.Contexte.TitreTxt.Fin)
	assertStringEquals(t, "LEGITEXT000006069565", v.Contexte.TitreTxt.IdTxt)
	assertStringEquals(t, "Code de la consommation", v.Contexte.TitreTxt.Contenu)

	assertStringEquals(t, "1997-04-03", v.Contexte.Tm.TitreTm.Debut)
	assertStringEquals(t, "2016-10-01", v.Contexte.Tm.TitreTm.Fin)
	assertStringEquals(t, "LEGISCTA000006084139", v.Contexte.Tm.TitreTm.Id)
	assertStringEquals(t, "Partie réglementaire", 
		v.Contexte.Tm.TitreTm.Contenu)
	assertStringEquals(t, "Livre II : Qualité des produits et des services", 
		v.Contexte.Tm.Tm.TitreTm.Contenu)
	assertStringEquals(t, "Titre Ier : Conformité", 
		v.Contexte.Tm.Tm.Tm.TitreTm.Contenu)
	assertStringEquals(t, "Chapitre V : Pouvoirs d'enquête", 
		v.Contexte.Tm.Tm.Tm.Tm.TitreTm.Contenu)
	assertStringEquals(t, "Section 2 : Recherche et constatation.", 
		v.Contexte.Tm.Tm.Tm.Tm.Tm.TitreTm.Contenu)

	assertLenEquals(t, 3, len(v.Versions))
	version := v.Versions[0]
	assertStringEquals(t, "MODIFIE", version.Etat)
	assertStringEquals(t, "2007-10-18", version.LienArt.Debut)
	assertStringEquals(t, "MODIFIE", version.LienArt.Etat)
	assertStringEquals(t, "2012-11-11", version.LienArt.Fin)
	assertStringEquals(t, "LEGIARTI000006291810", version.LienArt.Id)
	assertStringEquals(t, "D215-16", version.LienArt.Num)
	assertStringEquals(t, "LEGI", version.LienArt.Origine)

	assertStringEquals(t, `<br/>   Le remboursement des frais exposés pour la recherche et la constatation des infractions au livre II du présent code et des textes pris pour son application est effectué à l'appui d'un titre de perception unique émis par le préfet et recouvré par le comptable du Trésor public conformément aux dispositions prévues aux articles 80 à 95 du décret n° 62-1587 du 29 décembre 1962 portant règlement général sur la comptabilité publique.<br/>
      <br/>   Ce titre précisera, par poste de dépense, les coûts indiqués par l'agent verbalisateur mentionné à l'article L. 215-1 et faisant l'objet de la demande de remboursement.<br/>
      <br/>   Les postes de dépenses sont :<br/>
      <br/>   a) Les prélèvements et le transport des échantillons, dont le montant est fixé forfaitairement à 220 Euros TTC ;<br/>
      <br/>   b) Les analyses et essais, dont le montant est établi sur la base des coûts de revient supportés par le service auquel appartient l'agent verbalisateur.<br/>`, strings.TrimSpace(v.BlocTextuel.Contenu))

	assertLenEquals(t, 4, len(v.Liens))
	lien := v.Liens[0]
	assertStringEquals(t, "JORFTEXT000000299367", lien.CidTexte)
	assertStringEquals(t, "1962-12-29", lien.DateSignaTexte)
	assertStringEquals(t, "JORFTEXT000000299367", lien.Id)
	assertStringEquals(t, "DECRET", lien.NatureTexte)
	assertStringEquals(t, "", lien.NorTexte)
	assertStringEquals(t, "", lien.Num)
	assertStringEquals(t, "62-1587", lien.NumTexte)
	assertStringEquals(t, "source", lien.Sens)
	assertStringEquals(t, "CITATION", lien.TypeLien)
}

