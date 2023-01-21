package errors_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/haleyrc/errors"
)

type TestError struct{}

func (te TestError) Error() string         { return "oops" }
func (te TestError) GetCode() int          { return http.StatusTeapot }
func (te TestError) GetKind() errors.K     { return errors.K("custom") }
func (te TestError) GetMessage() string    { return "Hello" }
func (te TestError) GetMetadata() errors.M { return errors.M{"env": "prod", "status": 500} }

func TestCodeReturnsADefaultCodeFromAnUncompliantError(t *testing.T) {
	err := fmt.Errorf("oops")
	got := errors.Code(err)
	want := 500
	if got != want {
		t.Errorf("Expected Code to return %d, but got %d.", want, got)
	}
}

func TestCodeReturnsTheCodeFromACompliantError(t *testing.T) {
	var err TestError
	got := errors.Code(err)
	want := 418
	if got != want {
		t.Errorf("Expected Code to return %d, but got %d.", want, got)
	}
}

func TestKindReturnsADefaultKindFromAnUncompliantError(t *testing.T) {
	err := fmt.Errorf("oops")
	got := errors.Kind(err)
	want := errors.KindUnknown
	if got != want {
		t.Errorf("Expected Kind to return %T, but got %T.", want, got)
	}
}

func TestKindReturnsTheKindFromACompliantError(t *testing.T) {
	var err TestError
	got := errors.Kind(err)
	want := "custom"
	if string(got) != want {
		t.Errorf("Expected Kind to return %s, but got %s.", want, got)
	}
}

func TestMessageReturnsADefaultMessageFromAnUncompliantError(t *testing.T) {
	err := fmt.Errorf("oops")
	got := errors.Message(err)
	want := "An unexpected error has occurred."
	if got != want {
		t.Errorf("Expected Message to return %q, but got %q.", want, got)
	}
}

func TestMessageReturnsTheMessageFromACompliantError(t *testing.T) {
	var err TestError
	got := errors.Message(err)
	want := "Hello"
	if string(got) != want {
		t.Errorf("Expected Message to return %s, but got %s.", want, got)
	}
}

func TestMetadataReturnsNilFromAnUncompliantError(t *testing.T) {
	err := fmt.Errorf("oops")
	got := errors.Metadata(err)
	if got != nil {
		t.Errorf("Expected Code to return nil, but got %v.", got)
	}
}

func TestMetadataReturnsTheMetadataFromACompliantError(t *testing.T) {
	var err TestError
	got := errors.Metadata(err)
	want := errors.M{"env": "prod", "status": 500}
	if !metadataEqual(got, want) {
		t.Errorf("Expected Metadata to return %v, but got %v.", want, got)
	}
}

func metadataEqual(first, second errors.M) bool {
	if len(first) != len(second) {
		return false
	}
	for k, v := range first {
		if second[k] != v {
			return false
		}
	}
	return true
}
