package ex191

import "strings"

func wordPattern(pattern string, str string) bool {
	strArray := strings.Split(str, " ")
	patternArray := strings.Split(pattern, "")

	if len(strArray) != len(patternArray) {
		return false
	}

	// TODO: 优化
	hMap1 := make(map[string]string, 0)
	hMap2 := make(map[string]string, 0)

	for i := 0; i < len(patternArray); i++ {
		if _, exist := hMap1[patternArray[i]]; !exist {
			hMap1[patternArray[i]] = strArray[i]
		}

		if _, exist := hMap2[strArray[i]]; !exist {
			hMap2[strArray[i]] = patternArray[i]
		}

		if hMap1[patternArray[i]] != strArray[i] || hMap2[strArray[i]] != patternArray[i] {
			return false
		}
	}

	return true
}
