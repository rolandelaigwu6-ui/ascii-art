package fontbuilder

func BuildFont() map[byte][5]string {
	return map[byte][5]string{

		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'2': {
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
	}

}
