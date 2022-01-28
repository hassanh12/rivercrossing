package state

import "testing"

func TestViewState(t *testing.T) {
	wanted := "[kylling rev korn hs ---\\ \\_______/ _________________/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
func TestPutInBoat(t *testing.T) {
	wanted := "[rev korn ---\\ \\__kylling___hs__/ _________________/---]"
	state := PutInBoat()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestCrossRiver(t *testing.T) {
	wanted := "[rev korn ---\\ _________________\\__kylling___hs__//---]"
	state := CrossRiver()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
