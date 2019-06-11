package frequency

import "testing"

func TestGetPopular(t *testing.T) {
	str := `
	hello hello hello
	the the the the
	this 
	spend spend spend 
	first 
	blind blind  
	surround 
	preciding 
	reduce reduce
	consider 
	ache ache ache ache 
	mercy mercy 
	bow bow bow bow bow bow bow bow bow
	tear 
	forsake forsake`
	res:=GetPopular(str)
	t.Log(res)
	last:=res[0].count
	for _,v := range res {
		if last < v.count {
			t.Error("The algoritm isn't work")
			return 
		} 
		last = v.count
	} 
}
