package frequency

import (

	"regexp"
	"strings"
	"unicode"
)

func filter(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

type word struct {
	name  string
	count int
}

// GetPopular return 10 popular word
func GetPopular(str string) []word {
	words := strings.FieldsFunc(str, filter)
	strBytes := []byte(str)
	result := make(map[string]int)
	var popular []word
	for _, v := range words {
		if _, ok := result[v]; !ok {
			re := regexp.MustCompile(v)
			result[v] = len(re.FindAll(strBytes, -1))
			popular=append(popular,word{name:v,count:result[v]})
		}
	}
	// Checking if words < 10
	var n int
	if len(result) < 10 {
		n = len(result)
	} else {
		n =10
	}
	return qwicksort(popular)[:n]
}

func qwicksort(arr []word) []word {
	if len(arr) < 2 {
		return arr
	}
	prine := arr[0]
	var left, rigt, result []word
	for _,v := range arr[1:] {
		w:=word{name: v.name, count:v.count}
		if v.count < prine.count {
			rigt = append(rigt,w)
		} else {
			left = append(left, w)
		}
	}

	left = qwicksort(append(left))
	rigt = qwicksort(append(rigt))
	result = append(left, prine)
	result = append(result, rigt...)

	return result
}