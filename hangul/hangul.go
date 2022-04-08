package hangul

var (
	start = rune(44032)
	end   = rune(55204)
)

func HanConsonantSuffix(s string) bool {
	var result bool = false
	numEnds := 28
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}
