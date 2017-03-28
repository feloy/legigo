package xml

import (
	//	"fmt"
	"testing"
)

func TestDecodeTexteVersionUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeTexteVersion(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}

func TestDecodeTexteVersionInvalid(t *testing.T) {
	f := "specs/textelr_LEGITEXT000005627819.xml"
	_, err := DecodeTexteVersion(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}

func TestDecodeTexteVersion1(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000005627819.xml"
	v, err := DecodeTexteVersion(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "CMARPCN0", v.AncienId)
	assertStringEquals(t, "LEGITEXT000005627819", v.Id)
	assertStringEquals(t, "CODE", v.Nature)
	assertStringEquals(t, "LEGI", v.Origine)
	assertStringEquals(t, "texte/version/LEGI/TEXT/00/00/05/62/78/LEGITEXT000005627819.xml", v.Url)
	assertStringEquals(t, "LEGITEXT000005627819", v.Cid)
	assertStringEquals(t, "2999-01-01", v.DatePubli)
	assertStringEquals(t, "2999-01-01", v.DateTexte)
	assertStringEquals(t, "2016-01-01", v.DerniereModification)
	assertStringEquals(t, "", v.Nor)
	assertStringEquals(t, "", v.Num)
	assertStringEquals(t, "0", v.NumSequence)
	assertStringEquals(t, "", v.OriginePubli)
	assertStringEquals(t, "0", v.PageDebPubli)
	assertStringEquals(t, "0", v.PageFinPubli)

	assertLenEquals(t, 1, len(v.VersionsAVenir))
	assertStringEquals(t, "2016-04-01", v.VersionsAVenir[0])

	assertStringEquals(t, "", v.Autorite)
	assertStringEquals(t, "2006-09-01", v.DateDebut)
	assertStringEquals(t, "2016-04-01", v.DateFin)
	assertStringEquals(t, "ABROGE_DIFF", v.Etat)
	assertStringEquals(t, "", v.Ministere)
	assertStringEquals(t, "Code des marchés publics", v.Titre)
	assertStringEquals(t, "Code des marchés publics (édition 2006)", v.TitreFull)
	assertStringEquals(t, "", v.Visas.Contenu)
	assertStringEquals(t, "", v.Signataires.Contenu)
	assertStringEquals(t, "", v.Tp.Contenu)
	assertStringEquals(t, "", v.Nota.Contenu)
	assertStringEquals(t, "", v.Abro.Contenu)
	assertStringEquals(t, "", v.Rect.Contenu)

	assertLenEquals(t, 0, len(v.Liens))

}

func TestDecodeTexteVersion2(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000006069747.xml"
	v, err := DecodeTexteVersion(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "FAIFX", v.AncienId)
	assertStringEquals(t, "LEGITEXT000006069747", v.Id)
	assertStringEquals(t, "ARRETE", v.Nature)
	assertStringEquals(t, "LEGI", v.Origine)
	assertStringEquals(t, "texte/version/LEGI/TEXT/00/00/06/06/97/LEGITEXT000006069747.xml", v.Url)
	assertStringEquals(t, "LEGITEXT000006069747", v.Cid)
	assertStringEquals(t, "2999-01-01", v.DatePubli)
	assertStringEquals(t, "1982-05-28", v.DateTexte)
	assertStringEquals(t, "1982-05-30", v.DerniereModification)
	assertStringEquals(t, "", v.Nor)
	assertStringEquals(t, "", v.Num)
	assertStringEquals(t, "0", v.NumSequence)
	assertStringEquals(t, "", v.OriginePubli)
	assertStringEquals(t, "0", v.PageDebPubli)
	assertStringEquals(t, "0", v.PageFinPubli)

	assertLenEquals(t, 0, len(v.VersionsAVenir))

	assertStringEquals(t, "", v.Autorite)
	assertStringEquals(t, "1982-05-30", v.DateDebut)
	assertStringEquals(t, "2999-01-01", v.DateFin)
	assertStringEquals(t, "VIGUEUR", v.Etat)
	assertStringEquals(t, "", v.Ministere)
	assertStringEquals(t, "Arrêté du 28 mai 1982", v.Titre)
	assertStringEquals(t, "Arrêté du 28 mai 1982 Approbation de la méthode de calcul du complément de rémunération des comptes sur livret d'épargne populaire.", v.TitreFull)
	assertStringEquals(t, `   Le ministre de l'économie et des finances,    Vu le décret n° 82-454 du 28 mai 1982 pris pour l'application de la loi n° 82-357 du 27 avril 1982 portant création d'un régime d'épargne populaire, et notamment son article 16,<br/>
    <br/>
    `, v.Visas.Contenu)
	assertStringEquals(t, `JACQUES DELORS<br/>
    <br/>
    `, v.Signataires.Contenu)
	assertStringEquals(t, "", v.Tp.Contenu)
	assertStringEquals(t, "", v.Nota.Contenu)
	assertStringEquals(t, "", v.Abro.Contenu)
	assertStringEquals(t, "", v.Rect.Contenu)

	assertLenEquals(t, 1, len(v.Liens))
	lien := v.Liens[0]
	assertStringEquals(t, "JORFTEXT000000879675", lien.CidTexte)
	assertStringEquals(t, "1982-05-28", lien.DateSignaTexte)
	assertStringEquals(t, "LEGIARTI000006730900", lien.Id)
	assertStringEquals(t, "DECRET", lien.NatureTexte)
	assertStringEquals(t, "", lien.NorTexte)
	assertStringEquals(t, "16", lien.Num)
	assertStringEquals(t, "82-454", lien.NumTexte)
	assertStringEquals(t, "source", lien.Sens)
	assertStringEquals(t, "TXT_SOURCE", lien.TypeLien)
}
