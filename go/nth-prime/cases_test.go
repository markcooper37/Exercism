package prime

import "errors"

// Source: exercism/problem-specifications
// Commit: 4a3ba76 nth-prime: Apply new "input" policy
// Problem Specifications Version: 2.1.0

var tests = []struct {
	description string
	n           int
	p           int
	err         error
}{
	{
		"first prime",
		1,
		2,
		nil,
	},
	{
		"second prime",
		2,
		3,
		nil,
	},
	{
		"sixth prime",
		6,
		13,
		nil,
	},
	{
		"big prime",
		10001,
		104743,
		nil,
	},
	{
		"there is no zeroth prime",
		0,
		0,
		errors.New("input must be greater than 1"),
	},
}
