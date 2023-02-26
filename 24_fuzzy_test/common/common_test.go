package commmon

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	//arrange
	testCases := []string{"Luiz Moitinho", "   ", ":3123123"}

	for _, test := range testCases {
		//act
		result, _ := Reverse(test)
		resultCheck, _ := Reverse(result)

		//assert
		if test != resultCheck {
			t.Errorf("was expected %q and got %q", test, resultCheck)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"Luiz Moitinho", "   ", ":3123123"}

	for _, test := range testCases {
		f.Add(test)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		//act
		result, _ := Reverse(orig)
		resultCheck, _ := Reverse(result)

		//assert
		if orig != resultCheck {
			t.Errorf("was expected %q and got %q", orig, resultCheck)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(result) {
			t.Errorf("the reverse string not is UTF-8: %q", result)
		}
	})
}
