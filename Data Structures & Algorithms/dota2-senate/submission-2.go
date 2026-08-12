func predictPartyVictory(senate string) string {
	senates := make(map[rune][]int)
	
	for i,r := range senate {
		senates[r] = append(senates[r],i)
	}

	for len(senates['R']) > 0 && len(senates['D']) > 0 {
		rPos := senates['R'][0]
		senates['R'] = senates['R'][1:]

		dPos := senates['D'][0]
		senates['D'] = senates['D'][1:]

		if rPos < dPos {
			senates['R'] = append(senates['R'],rPos + len(senate))
		} else {
			senates['D'] = append(senates['D'],dPos + len(senate))
		}
	}

	if len(senates['R']) > 0 {
		return "Radiant"
	}

	return "Dire"
}
