package main

import (
	"testing"
)

func TestConverter(t *testing.T) {
	result := converter("参拾四億九千参百四拾八万五千九百弐拾")
	expect := 3493485920
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case1 pass")
	
	result = converter("九千九百九拾九兆九千九百九拾九億九千九百九拾九万九千九百九拾九")
	expect = 9999999999999999
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case2 pass")
	
	result = converter("壱千弐百参拾四兆五千六百七拾八万")
	expect = 1234000056780000
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case3 pass")
	
	result = converter("壱千兆壱")
	expect = 1000000000000001
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case4 pass")

	result = converter("零")
	expect = 0
	if result != expect {
		t.Error("expect: ", expect, " result: ", result)
	}
	t.Log("case4 pass")

	t.Log("complete")
} 