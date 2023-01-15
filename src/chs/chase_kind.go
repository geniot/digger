package chs

// Kind* constants refer to tile kinds for input and output.
const (
	KindTunnel = iota
	KindField
	KindBag
	KindFrom
	KindTo
	KindPath
)

// KindRunes map tile kinds to output runes.
var KindRunes = map[int]rune{
	KindTunnel: '.',
	KindField:  'X',
	KindBag:    'B',
	KindFrom:   'F',
	KindTo:     'T',
	KindPath:   '●',
}

// RuneKinds map input runes to tile kinds.
var RuneKinds = map[rune]int{
	'.': KindTunnel,
	'X': KindField,
	'B': KindBag,
	'F': KindFrom,
	'T': KindTo,
}

// KindCosts map tile kinds to movement costs.
var KindCosts = map[int]float64{
	KindTunnel: 1.0,
	KindFrom:   1.0,
	KindBag:    10.0,
	KindTo:     1.0,
}

var AllField = [9]int{
	KindField, KindField, KindField,
	KindField, KindField, KindField,
	KindField, KindField, KindField}

var SField = [9]int{
	KindField, KindTunnel, KindField,
	KindTunnel, KindTunnel, KindTunnel,
	KindField, KindTunnel, KindField}

var VField = [9]int{
	KindField, KindTunnel, KindField,
	KindField, KindTunnel, KindField,
	KindField, KindTunnel, KindField}

var HField = [9]int{
	KindField, KindField, KindField,
	KindTunnel, KindTunnel, KindTunnel,
	KindField, KindField, KindField}
