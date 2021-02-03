package prop

import "testing"

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Title string
		Num   int
		Want  string
	}{
		{"1 to I", 1, "I"},
		{"2 to II", 2, "II"},
		{"3 to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Title, func(t *testing.T) {
			got := ConvertToRoman(test.Num)

			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)

			}
		})
	}
}
