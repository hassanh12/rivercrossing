package state

var state = "[kylling rev korn hs ---\\ \\_______/ _________________/---]"

func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return state
}
func PutInBoat() string {
	state = "[rev korn ---\\ \\__kylling___hs__/ _________________/---]"
	return state
}

func CrossRiver() string {
	state = "[rev korn ---\\ _________________\\__kylling___hs__//---]"
	return state
}
