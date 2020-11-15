package fracdex

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert := assert.New(t)

	test := func(a, b, exp string) {
		act, err := KeyBetween(a, b)
		if err != nil {
			assert.Equal("", act)
			assert.Equal(exp, err.Error())
		} else {
			assert.Nil(err)
			assert.Equal(exp, act)
		}
	}

	test("", "", "a0")
	test("", "a0", "Zz")
	test("", "Zz", "Zy")
	test("a0", "", "a1")
	test("a1", "", "a2")
	test("a0", "a1", "a0V")
	test("a1", "a2", "a1V")
	test("a0V", "a1", "a0l")
	test("Zz", "a0", "ZzV")
	test("Zz", "a1", "a0")
	test("", "Y00", "Xzzz")
	test("bzz", "", "c000")
	test("a0", "a0V", "a0G")
	test("a0", "a0G", "a08")
	test("b125", "b129", "b127")
	test("a0", "a1V", "a1")
	test("Zz", "a01", "a0")
	test("", "a0V", "a0")
	test("", "b999", "b99")
	test(
		"",
		"A00000000000000000000000000",
		"invalid order key: A00000000000000000000000000",
	)
	test("", "A000000000000000000000000001", "A000000000000000000000000000V")
	test("zzzzzzzzzzzzzzzzzzzzzzzzzzy", "", "zzzzzzzzzzzzzzzzzzzzzzzzzzz")
	test("zzzzzzzzzzzzzzzzzzzzzzzzzzz", "", "zzzzzzzzzzzzzzzzzzzzzzzzzzzV")
	test("a00", "", "invalid order key: a00")
	test("a00", "a1", "invalid order key: a00")
	test("0", "1", "invalid order key head: 0")
	test("a1", "a0", "a1 >= a0")
}

func TestNKeys(t *testing.T) {
	assert := assert.New(t)

	test := func(a, b string, n uint, exp string) {
		actSlice, err := NKeysBetween(a, b, n)
		act := strings.Join(actSlice, " ")
		if err != nil {
			assert.Equal("", act)
			assert.Equal(exp, err.Error())
		} else {
			assert.Nil(err)
			assert.Equal(exp, act)
		}
	}
	test("", "", 5, "a0 a1 a2 a3 a4")
	test("a4", "", 10, "a5 a6 a7 a8 a9 aA aB aC aD aE")
	test("", "a0", 5, "Zv Zw Zx Zy Zz")
	test(
		"a0",
		"a2",
		20,
		"a04 a08 a0G a0K a0O a0V a0Z a0d a0l a0t a1 a14 a18 a1G a1O a1V a1Z a1d a1l a1t",
	)
}
