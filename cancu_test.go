package cancu

import "testing"

func TestObserver(t *testing.T)  {
	o := Cancu()
	ob := Cancu()
	if o != ob {
		t.Fail()
	}
}
