package main

func getLevelChar(x, y, l int32) rune {
	if l == 0 {
		l++
	}
	return rune(LEVEL_DATA[l-1][y][x])
}

func levplan() int32 {
	var l = int32(0) //todo: make stateful
	//if l > 8 {
	//	l = (l & 3) + 5 /* Level plan: 12345678, 678, (5678) 247 times, 5 forever */
	//}
	return l
}
