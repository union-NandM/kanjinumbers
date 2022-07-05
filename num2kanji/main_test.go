package main

import (
	"testing"
)

func TestConverter(t *testing.T) {
	result := converter(3493485920)
	expect := "参拾四億九千参百四拾八万五千九百弐拾"
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case1 pass")
	
	result = converter(9999999999999999)
	expect = "九千九百九拾九兆九千九百九拾九億九千九百九拾九万九千九百九拾九"
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case2 pass")
	
	result = converter(1234000056780000)
	expect = "壱千弐百参拾四兆五千六百七拾八万"
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case3 pass")
	
	result = converter(1000000000000001)
	expect = "壱千兆壱"
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case4 pass")

	result = converter(0)
	expect = "零"
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case5 pass")

	t.Log("complete")
} 