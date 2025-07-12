package errors

import (
	"errors"
	"testing"
)

func TestCompareErrors(t *testing.T) {
	tests := []struct {
		name     string
		input1   error
		input2   error
		expected bool
	}{
		{"CompareErrorsBothNil",
			nil,
			nil,
			true},
		{"CompareErrorsFirstNil",
			nil,
			ErrInvalidEmpty,
			false},
		{"CompareErrorsSecondNil",
			ErrInvalidEmpty,
			nil,
			false},
		{"CompareErrorsSame",
			ErrInvalidEmpty,
			ErrInvalidEmpty,
			true},
		{"CompareErrorsDiff",
			ErrInvalidEmpty,
			ErrInvalidEmail,
			false},
		{"CompareErrorsJoinedSame",
			errors.Join(ErrInvalidEmpty, ErrInvalidEmail),
			errors.Join(ErrInvalidEmpty, ErrInvalidEmail),
			true},
		{"CompareErrorsJoinedSameDiffOrder",
			errors.Join(ErrInvalidEmpty, ErrInvalidEmail),
			errors.Join(ErrInvalidEmail, ErrInvalidEmpty),
			false},
		{"CompareErrorsJoinedDiff",
			errors.Join(ErrInvalidEmpty, ErrInvalidEmail),
			errors.Join(ErrInvalidEmpty, ErrInvalidLength),
			false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if CompareErrors(test.input1, test.input2) != test.expected {
				t.Errorf("CompareErrors() = %v, want %v", CompareErrors(test.input1, test.input2), test.expected)
			}
		})
	}
}
