package textmatcher

import (
	"fmt"
	"math"
	"strings"
)

func Matcher(textone string, texttwo string) bool {
	var i = 0
	var diff = 0
	if math.Abs(float64(len([]rune(textone))-len([]rune(texttwo)))) > 1 {
		return false
	} else if len([]rune(textone)) > len([]rune(texttwo)) {
		texttwo = fmt.Sprintf("%s%s", texttwo, textone[(len(textone)-1):])
	} else if len([]rune(textone)) > len([]rune(texttwo)) {
		texttwo = strings.TrimSuffix(texttwo, texttwo[(len(texttwo)-1):])
	}

	if textone == texttwo {
		return true
	}

	for len(textone) > i {
		if string([]rune(textone)[i]) != string([]rune(texttwo)[i]) {
			diff += 1
		}

		i++
	}

	if diff > 1 {
		return false
	}

	return true
}
