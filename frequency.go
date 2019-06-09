package frequency

import (
	"strings"
	"unicode"
	"regexp"
	"fmt"
)

func filter(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

// GetPopular return 10 popular word
func GetPopular(str string) []string {
	words:=strings.FieldsFunc(str, filter)
	strBytes :=[]byte(str)
	result:=make(map[string]int)
	for _,v := range words {
		if _,ok := result[v]; !ok {
			re := regexp.MustCompile(v)
			result[v]=len(re.FindAll(strBytes, -1))
		}	
	}
	fmt.Println(result)
	var popular = make([]string,10)
	return popular
}


func rex(res []int, i int, s int) int{
	s=s+res[i]
    if i ==0 {
		return  s
	}  
	return rex(res, i-1,s)
} 

func recur() {
	res:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Println("===>",rex(res,len(res)-1,0))
}



func qwicksort(arr []int) []int {
	if len(arr) < 2 {
		return  arr
	}
	prine:=arr[0]
	var left, rigt, result []int
	for _,v := range arr[1:] {	
		if v > prine {
			rigt = append(rigt,v)
		}  else {
			left = append(left,v)
		}
	}

	left = qwicksort(append(left))
	rigt = qwicksort(append(rigt))
	result = append(left,prine)
	result = append(result,rigt...)

	return result
}

func qwicksortstart() {
	res:=[]int{9,16,7,4,45,6,33,8,37,100,79,1}
	fmt.Println(qwicksort(res))
}