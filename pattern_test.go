package gox_test

import (
	"testing"

	"github.com/gopub/gox"
	"github.com/stretchr/testify/assert"
)

func TestIsBirthDate(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "2008-01-03",
			Expected: true,
		},
		{
			Input:    "2008-01-3",
			Expected: true,
		},
		{
			Input:    "2008-1-03",
			Expected: true,
		},
		{
			Input:    "2008-1-3",
			Expected: true,
		},
		{
			Input:    "2008-11-3",
			Expected: true,
		},
		{
			Input:    "2008-11-30",
			Expected: true,
		},
		{
			Input:    "200801-03",
			Expected: false,
		},
		{
			Input:    "2008-001-3",
			Expected: false,
		},
		{
			Input:    "2008-1-003",
			Expected: false,
		},
		{
			Input:    "2008-21-3",
			Expected: false,
		},
		{
			Input:    "2008-9-31",
			Expected: false,
		},
		{
			Input:    "1900-02-29",
			Expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			assert.Equal(t, tc.Expected, gox.IsBirthDate(tc.Input))
		})
	}
}

func TestIsEmail(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "jim01@gmail.com",
			Expected: true,
		},
		{
			Input:    "jim.green@gmail.com",
			Expected: true,
		},
		{
			Input:    "jim.green.haha@gmail.com",
			Expected: true,
		},
		{
			Input:    "2008@gmail.com",
			Expected: true,
		},
		{
			Input:    "2008jim@gmail.com",
			Expected: true,
		},
		{
			Input:    "2008-jim@gmail.com",
			Expected: true,
		},
		{
			Input:    "2008-jim-001@gmail.com",
			Expected: true,
		},
		{
			Input:    "2008_jim@gmail.com",
			Expected: true,
		},
		{
			Input:    "abc@xyz",
			Expected: true,
		},
		{
			Input:    "abc",
			Expected: false,
		},
		{
			Input:    "abc@xyz..com",
			Expected: false,
		},
		{
			Input:    "abc@xyz@com",
			Expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			assert.Equal(t, tc.Expected, gox.IsEmail(tc.Input))
		})
	}
}

func TestIsLink(t *testing.T) {
	ok := gox.IsLink("")
	assert.Empty(t, ok)

	ok = gox.IsLink("abc")
	assert.Empty(t, ok)
}
