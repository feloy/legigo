package xml

import (
	//	"fmt"
	"strings"
	"testing"
)

func TestDecodeSectionTaUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeSectionTa(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}

func TestDecodeSectionTaInvalid(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000005627819.xml"
	_, err := DecodeSectionTa(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}

func TestDecodeSectionTa1(t *testing.T) {
	f := "specs/section_ta_LEGISCTA000006083673.xml"
	v, err := DecodeSectionTa(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "LEGISCTA000006083673", v.Id)
	assertStringEquals(t, "Partie législative", v.TitreTa)
	assertStringEquals(t, "", v.Commentaire)

	assertLenEquals(t, 0, len(v.LiensArt))
	assertLenEquals(t, 5, len(v.LiensSectionTa))
}

func TestDecodeSectionTa2(t *testing.T) {
	f := "specs/section_ta_LEGISCTA000029502487.xml"
	v, err := DecodeSectionTa(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "LEGISCTA000029502487", v.Id)
	assertStringEquals(t, "Section 6 : Action de groupe intervenant dans le domaine de la concurrence", v.TitreTa)
	assertStringEquals(t, "Néant", v.Commentaire)

	assertLenEquals(t, 0, len(v.LiensArt))
	assertLenEquals(t, 0, len(v.LiensSectionTa))
}

func TestDecodeSectionTa3(t *testing.T) {
	f := "specs/section_ta_LEGISCTA000029502492.xml"
	v, err := DecodeSectionTa(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "LEGISCTA000029502492", v.Id)
	assertStringEquals(t, "Section 1 : Dispositions préliminaires ", v.TitreTa)
	assertStringEquals(t, "", v.Commentaire)

	assertLenEquals(t, 5, len(v.LiensArt))
	assertLenEquals(t, 0, len(v.LiensSectionTa))

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
	assertStringEquals(t, "Livre IV : Les associations de consommateurs", 
		v.Contexte.Tm.Tm.TitreTm.Contenu)
	assertStringEquals(t, "Titre II : Action en justice des associations.", 
		v.Contexte.Tm.Tm.Tm.TitreTm.Contenu)
	assertStringEquals(t, `Chapitre III : Action de groupe`, 
		strings.TrimSpace(v.Contexte.Tm.Tm.Tm.Tm.TitreTm.Contenu))
}
