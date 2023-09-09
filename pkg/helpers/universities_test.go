package helpers

import (
	"testing"
)

func TestCheckForSpecialSymbols(t *testing.T) {
	var res bool
	var str string
	var expect bool

	str = "p4ssword"
	res = checkForSpecialSymbols(str, false)
	expect = true
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = "SELECT * FROM universities"
	res = checkForSpecialSymbols(str, false)
	expect = false
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = ";"
	res = checkForSpecialSymbols(str, false)
	expect = false
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = "!!!! !!!!!"
	res = checkForSpecialSymbols(str, false)
	expect = false
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = "p4ssword!"
	res = checkForSpecialSymbols(str, false)
	expect = true
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = "Some Name"
	res = checkForSpecialSymbols(str, true)
	expect = true
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}

	str = "Some Name!!!;"
	res = checkForSpecialSymbols(str, true)
	expect = false
	if res != expect {
		t.Errorf("str=\"%s\" expected=%t got=%t", str, expect, res)
	}
}
