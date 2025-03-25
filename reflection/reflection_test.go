package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name     string
	Children Children
}

type Children struct {
	First  string
	Second string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Kayla"},
			[]string{"Kayla"},
		},
		{
			"struct with two string fields",
			struct {
				Name    string
				Partner string
			}{"Kayla", "EJ"},
			[]string{"Kayla", "EJ"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Kayla", 33},
			[]string{"Kayla"},
		},
		{
			"nested fields",
			Person{
				"Kayla",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			},
			[]string{"Kayla", "Calvin", "Emilia"},
		},
		{
			"pointers to things",
			&Person{
				"EJ",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			},
			[]string{"EJ", "Calvin", "Emilia"},
		},
		{
			"slices",
			[]Person{{
				"EJ",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			}, {
				"Kayla",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			}},
			[]string{"EJ", "Calvin", "Emilia", "Kayla", "Calvin", "Emilia"},
		},
		{
			"arrays",
			[2]Person{{
				"EJ",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			}, {
				"Kayla",
				Children{
					First:  "Calvin",
					Second: "Emilia",
				},
			}},
			[]string{"EJ", "Calvin", "Emilia", "Kayla", "Calvin", "Emilia"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moooo",
				"Sheep": "Baaaa",
			},
			[]string{"Moooo", "Baaaa"},
		},
	}

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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moooo",
			"Sheep": "Baaaa",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moooo")
		assertContains(t, got, "Baaaa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Person)

		go func() {
			aChannel <- Person{"Kayla", Children{"Calvin", "Emilia"}}
			aChannel <- Person{"EJ", Children{"Calvin", "Emilia"}}
			close(aChannel)
		}()

		var got []string
		want := []string{"Kayla", "Calvin", "Emilia", "EJ", "Calvin", "Emilia"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
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
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
