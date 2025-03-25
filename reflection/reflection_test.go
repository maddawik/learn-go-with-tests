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
}
