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
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with two string fields and one int field",
			struct {
				Name string
				City string
				Age  int
			}{"Chris", "London", 1},
			[]string{"Chris", "London"},
		},
		{
			"nested fields",
			Person{
				"Cris",
				Profile{
					35,
					"London",
				},
			},
			[]string{"Cris", "London"},
		},
		{
			"value passed in a pointer",
			&Person{
				"Cris",
				Profile{35, "London"},
			},
			[]string{"Cris", "London"},
		},
		{
			"slices",
			[]Profile{
				{35, "London"},
				{42, "Dublin"},
			},
			[]string{"London", "Dublin"},
		},
		{
			"pure string",
			"Teste",
			[]string{"Teste"},
		},
		{
			"string slice",
			[]string{"Teste1", "Teste2"},
			[]string{"Teste1", "Teste2"},
		},
		{
			"arrays",
			[2]Profile{
				{32, "London"},
				{42, "Dublin"},
			},
			[]string{"London", "Dublin"},
		},
	}

	for _, value := range cases {
		t.Run(value.Name, func(t *testing.T) {
			var got []string

			walk(value.Input, func(str string) {
				got = append(got, str)
			})

			if !reflect.DeepEqual(got, value.ExpectedCalls) {
				t.Errorf("got %v want %v", got, value.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		var got []string
		aMap := map[string]string{
			"Foo": "Bar",
			"Bar": "Foo",
		}

		walk(aMap, func(str string) {
			got = append(got, str)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Foo")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{40, "Krakow"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Krakow"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{35, "Berlin"}, Profile{42, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, neddle string) {
	t.Helper()
	contains := false
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == neddle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, neddle)
	}
}
