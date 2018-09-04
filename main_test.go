package main

import (
	"reflect"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	const template = "a{{ if .test }}b{{ end }}{{ .test2 }}"
	var data = map[string]string{
		"test":  setFlag,
		"test2": "test42",
	}
	const expected = "abtest42"
	r, err := renderTemplate(template, data)
	if err != nil {
		t.Error(err)
	}
	if r != expected {
		t.Errorf("Got %s expected %s\n", r, expected)
	}
}

type testCase struct {
	name   string
	flags  []string
	result map[string]string
	err    error
}

func TestParseFlags(t *testing.T) {
	cases := []testCase{
		{
			name:   "Empty",
			flags:  []string{},
			result: make(map[string]string),
			err:    nil,
		},
		{
			name:   "One flag",
			flags:  []string{"--test"},
			result: map[string]string{"test": setFlag},
			err:    nil,
		},
		{
			name:  "Multi-flags",
			flags: []string{"--test=testArg", "-test2", "--test3", "testArg3"},
			result: map[string]string{
				"test":  "testArg",
				"test2": setFlag,
				"test3": "testArg3",
			},
			err: nil,
		},
		{
			name:   "Error flags",
			flags:  []string{"test", "-test2", "--test3", "testArg3"},
			result: nil,
			err:    ErrWrongFlags,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			parsedFlags, err := parseFlags(c.flags)
			if err != c.err {
				t.Error("Got error", err)
			}
			if !reflect.DeepEqual(parsedFlags, c.result) {
				t.Errorf("Expected: %+v\nGot: %+v\n", c.result, parsedFlags)
			}
		})
	}
}
