package main

func PlayerWins(playerHand string, dealerHand string) bool {
	plrScr := countScore(playerHand)
	if plrScr == 0 || plrScr > 21 {
		return false
	}
	dlrScr := countScore(dealerHand)
	if dlrScr > 21 {
		return false
	}
	return plrScr > dlrScr
}

func countScore(dt string) int8 {
	var rst int8
	for _, rn := range dt {
		rst += convertToInt(rn)
		if rst >= 22 {
			return rst
		}
	}
	return rst
}

func convertToInt(dt int32) int8 {
	if dt >= 50 && dt <= 57 {
		return int8(dt - 48)
	}

	switch dt {
	case 88, 74, 81, 75:
		return 10
	case 65:
		return 11
	}
	return 0
}
