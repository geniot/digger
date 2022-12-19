package glb

// Kind* constants refer to tile kinds for input and output.
const (
	KindTunnel = iota
	KindField
	KindFrom
	KindTo
	KindPath
)

// KindRunes map tile kinds to output runes.
var KindRunes = map[int]rune{
	KindTunnel: '.',
	KindField:  'X',
	KindFrom:   'F',
	KindTo:     'T',
	KindPath:   '●',
}

// RuneKinds map input runes to tile kinds.
var RuneKinds = map[rune]int{
	'.': KindTunnel,
	'X': KindField,
	'F': KindFrom,
	'T': KindTo,
}

// KindCosts map tile kinds to movement costs.
var KindCosts = map[int]float64{
	KindTunnel: 1.0,
	KindFrom:   1.0,
	KindTo:     1.0,
}
