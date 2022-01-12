package main

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

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string // used to name test
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 25},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with nested struct",
			Input: Person{"Chris",
				Profile{Age: 35, City: "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
		{
			Name: "Arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Babaz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Babaz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{228, "Berlin"}
			aChannel <- Profile{1337, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{35, "Moscow"}
		}

		var got []string
		want := []string{"Berlin", "Moscow"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}

func assertContains(t *testing.T, got []string, s string) {
	t.Helper()
	contains := false
	for _, x := range got {
		if x == s {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but id didn't", got, s)
	}
}
