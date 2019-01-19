package norepeatingsubstr

func findMaxSubStr(str string) int {
	charIndex := map[byte]int{}

	bytes := []byte(str)
	if len(bytes) == 0 {
		return 0
	}
	var s int = 0
	var ss, ee int = 0, 0
	for i, ch := range bytes {
		if index, ok := charIndex[ch]; ok && index >= s {
			s = index + 1
		} else if i-s+1 > ee-ss+1 {
			ee = i
			ss = s
		}
		charIndex[ch] = i
	}
	return ee - ss + 1
}
