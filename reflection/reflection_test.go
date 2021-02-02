package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{{
		"Struct with one string",
		struct {
			Name string
		}{"Matt"},
		[]string{"Matt"},
	}, {
		"Struct with two strings",
		struct {
			Name string
			City string
		}{
			"Matt",
			"DC",
		},
		[]string{"Matt", "DC"},
	}, {
		"Struct with non string",
		struct {
			Name string
			Age  int
		}{
			"Matt",
			69,
		},
		[]string{"Matt"},
	}, {
		"Nested fields",
		Person{
			"Chris",
			Profile{33, "London"},
		},
		[]string{"Chris", "London"},
	}, {
		"Pointers",
		&Person{
			"Matt",
			Profile{12, "DC"},
		},
		[]string{"Matt", "DC"},
	},
		{
			"Slice slice",
			[]Profile{
				{33, "London"},
				{21, "DC"},
			},
			[]string{"London", "DC"},
		}, {
			"Arrays",
			[2]Profile{
				{69, "Memphis"},
				{20, "Needham"},
			},
			[]string{"Memphis", "Needham"},
		}, {
			"Map",
			map[string]string{
				"Foo": "Bar",
				"Boo": "Peek",
			},
			[]string{"Bar", "Peek"},
		}}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Boo": "Peek",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Peek")
	})

	t.Run("Channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "SF"}
			aChannel <- Profile{21, "LA"}
			close(aChannel)
		}()

		var got []string

		want := []string{"SF", "LA"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{12, "Seattle"}, Profile{59, "Austin"}
		}

		var got []string

		want := []string{"Seattle", "Austin"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
