package errors_test

import (
	"fmt"
	"testing"

	"github.com/haleyrc/errors"
)

func TestDefaultErrorsReturnCorrectCodes(t *testing.T) {
	testcases := []struct {
		err  error
		want int
	}{
		{errors.NotAuthorized{}, 403},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("%T", tc.err)
		t.Run(name, func(t *testing.T) {
			got := errors.Code(tc.err)
			if got != tc.want {
				t.Errorf("Expected Code to return %d, but got %d.", tc.want, got)
			}
		})
	}
}

func TestDefaultErrorsReturnCorrectKinds(t *testing.T) {
	testcases := []struct {
		err  error
		want string
	}{
		{errors.NotAuthorized{}, "not_authorized"},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("%T", tc.err)
		t.Run(name, func(t *testing.T) {
			got := errors.Kind(tc.err)
			if string(got) != tc.want {
				t.Errorf("Expected Kind to return %s, but got %s.", tc.want, got)
			}
		})
	}
}

func TestDefaultErrorsReturnCorrectMessages(t *testing.T) {
	testcases := []struct {
		err  error
		want string
	}{
		{errors.NotAuthorized{Message: "hello"}, "hello"},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("%T", tc.err)
		t.Run(name, func(t *testing.T) {
			got := errors.Message(tc.err)
			if got != tc.want {
				t.Errorf("Expected Message to return %q, but got %q.", tc.want, got)
			}
		})
	}
}

func TestDefaultErrorsReturnCorrectMetadata(t *testing.T) {
	testcases := []struct {
		err  error
		want errors.M
	}{
		{
			errors.NotAuthorized{
				Metadata: errors.M{"env": "prod", "status": 500},
			},
			errors.M{"env": "prod", "status": 500},
		},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("%T", tc.err)
		t.Run(name, func(t *testing.T) {
			got := errors.Metadata(tc.err)
			if !metadataEqual(got, tc.want) {
				t.Errorf("Expected Metadata to return %v, but got %v.", tc.want, got)
			}
		})
	}
}
