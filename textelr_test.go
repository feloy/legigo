package xml

import (
	//	"fmt"
	"testing"
)

func TestDecodeTexteLrUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeTextelr(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}

func TestDecodeTexteLrInvalid(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000005627819.xml"
	_, err := DecodeTexteVersion(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}

func TestDecodeTextelr(t *testing.T) {
	f := "specs/textelr_LEGITEXT000005627819.xml"
	v, err := DecodeTextelr(f)
	if err != nil {
		t.Errorf("Error decoding file")
	}

	assertStringEquals(t, "CMARPCN0", v.AncienId)
	assertStringEquals(t, "LEGITEXT000005627819", v.Id)
	assertStringEquals(t, "CODE", v.Nature)
	assertStringEquals(t, "LEGI", v.Origine)
	assertStringEquals(t, "texte/struct/LEGI/TEXT/00/00/05/62/78/LEGITEXT000005627819.xml", v.Url)
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

	assertLenEquals(t, 0, len(v.LiensArt))

	assertLenEquals(t, 7, len(v.LiensSectionTa))
	lien := v.LiensSectionTa[0]
	assertStringEquals(t, "LEGISCTA000006083133", lien.Cid)
	assertStringEquals(t, "2006-09-01", lien.Debut)
	assertStringEquals(t, "ABROGE_DIFF", lien.Etat)
	assertStringEquals(t, "2016-04-01", lien.Fin)
	assertStringEquals(t, "LEGISCTA000006083133", lien.Id)
	assertStringEquals(t, "1", lien.Niv)
	assertStringEquals(t, "/LEGI/SCTA/00/00/06/08/31/LEGISCTA000006083133.xml", lien.Url)
	assertStringEquals(t, "PREMIÈRE PARTIE : DISPOSITIONS APPLICABLES AUX POUVOIRS ADJUDICATEURS", lien.Contenu)

	assertLenEquals(t, 1, len(v.Versions))
	version := v.Versions[0]
	assertStringEquals(t, "ABROGE_DIFF", version.Etat)
	assertStringEquals(t, "2006-09-01", version.LienTxt.Debut)
	assertStringEquals(t, "2016-04-01", version.LienTxt.Fin)
	assertStringEquals(t, "LEGITEXT000005627819", version.LienTxt.Id)
	assertStringEquals(t, "", version.LienTxt.Num)

}
