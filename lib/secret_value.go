package secretvalue

func CalculateSecretValue( lengthOfString int ) (string, bool) {
	var char string 

	switch lengthOfString {
		case 4: 
		char = "E"
		case 5: 
		char = "F"
		case 6: 
		char = "G"
		case 7: 
		char = "H"
		case 8: 
		char = "I"
		case 9: 
		char = "J"
		case 10: 
		char = "K"
		case 11: 
		char = "L"
		case 12: 
		char = "M"
		case 13: 
		char = "N"
		case 14: 
		char = "O"
		case 15: 
		char = "P"
		case 16: 
		char = "Q"
		case 17: 
		char = "R"
		case 18: 
		char = "S"
		case 19: 
		char = "T"
		case 20: 
		char = "U"
		case 21: 
		char = "V"
		case 22: 
		char = "W"
		case 23: 
		char = "X"
		case 24: 
		char = "Y"
		case 25: 
		char = "Z"
		case 26: 
		char = "a"
		case 27: 
		char = "b"
		case 28: 
		char = "c"
		case 29: 
		char = "d"
		case 30: 
		char = "e"
		case 31: 
		char = "f" 
		 case 32: 
		char = "g" 
		 case 33: 
		char = "h" 
		 case 34: 
		char = "i" 
		 case 35: 
		char = "j" 
		 case 36: 
		char = "k" 
		 case 37: 
		char = "l" 
		 case 38: 
		char = "m" 
		 case 39: 
		char = "n" 
		 case 40: 
		char = "o" 
		 case 41: 
		char = "p" 
		 case 42: 
		char = "q" 
		 case 43: 
		char = "r" 
		 case 44: 
		char = "s" 
		 case 45: 
		char = "t" 
		 case 46: 
		char = "u" 
		 case 47: 
		char = "v" 
		 case 48: 
		char = "w" 
		 case 49: 
		char = "x" 
		 case 50: 
		char = "y" 
		 case 51: 
		char = "z" 
		 case 52: 
		char = "0" 
		 case 53: 
		char = "1" 
		 case 54: 
		char = "2" 
		 case 55: 
		char = "3" 
		 case 56: 
		char = "4" 
		 case 57: 
		char = "5" 
		 case 58: 
		char = "6" 
		 case 59: 
		char = "7" 
		 case 60: 
		char = "8" 
		 case 61: 
		char = "9" 
		 case 62: 
		char = "-" 
		 case 63: 
		char = " " 
		 case 64: 
		char = "A" 
		 case 65: 
		char = "B" 
		 case 66: 
		char = "C" 
		 case 67: 
		char = "D" 
		 case 68: 
		char = "E" 
		 case 69: 
		char = "F" 
		 case 70: 
		char = "G" 
		 case 71: 
		char = "H" 
		 case 72: 
		char = "I" 
		 case 73: 
		char = "J" 
		 case 74: 
		char = "K" 
		 case 75: 
		char = "L" 
		 case 76: 
		char = "M" 
		 case 77: 
		char = "N" 
		 case 78: 
		char = "O" 
		 case 79: 
		char = "P" 
		 case 80: 
		char = "Q" 
		 case 81: 
		char = "R" 
		 case 82: 
		char = "S" 
		 case 83: 
		char = "T" 
		 case 84: 
		char = "U" 
		 case 85: 
		char = "V" 
		default:
			return "not found", false
	}
	return char, true

}