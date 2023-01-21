package errors

import "net/http"

// The following defaults are returned by the top-level extraction functions
// when given a non-compliant error (one that doesn't implement Interface).
const (
	DefaultCode    = http.StatusInternalServerError
	DefaultKind    = KindUnknown
	DefaultMessage = "An unexpected error has occurred."
)

// Code returns the code associated with the provided error if any. If the error
// doesn't implement Interface, a default code is returned.
func Code(err error) int {
	if c, ok := err.(Interface); ok {
		return c.GetCode()
	}
	return DefaultCode
}

// Message returns the message associated with the provided error if any. If the
// error doesn't implement Interface, a default message is returned.
func Message(err error) string {
	if m, ok := err.(Interface); ok {
		return m.GetMessage()
	}
	return DefaultMessage
}

// Interface is the contract that errors can implement in order to be used with
// the top-level methods such as Code, Kind, Message, and Metadata. This allows
// you to define your own error types and as long as you use the errors package
// to process them, you can enjoy all of the consistency benefits it provides.
type Interface interface {
	// GetCode returns the HTTP status code which may be returned to clients.
	GetCode() int

	// GetKind returns a static, machine-readable label for the class of errors
	// this error belongs to e.g. not authorized, etc.
	GetKind() K

	// GetMessage returns a human-readable message suitable for display to an
	// end user. Unlike the standard Go error message format, these can be
	// freely capitalized and end in a period (or whatever scheme makes sense in
	// your application). Care should be taken, however, to ensure that no
	// sensitive information is included in this message as it may be
	// transmitted over the wire to clients and even shown to users.
	GetMessage() string

	// GetMetadata returns arbitrary metadata attached to this instance of the
	// error. This is primarily useful for sending additional data to error
	// platforms like Sentry, where it can be used to troubleshoot the error or
	// allow for more granular error grouping to detect trends, etc.
	GetMetadata() M
}

// K represents an error kind which is a static, machine-readable label for a
// class of errors. You can think of kinds as an extensible alternative to HTTP
// status codes, and like status codes should be identical for errors resulting
// from the same error condition e.g. not authorized, not found, etc.
type K string

// Kind returns the kind associated with the provided error if any. If the error
// doesn't implement Interface, a default kind is returned.
func Kind(err error) K {
	if k, ok := err.(Interface); ok {
		return k.GetKind()
	}
	return DefaultKind
}

// The following kinds are exposed by the errors package directly, but you can
// enumerate your own kinds if you have an alternative labeling scheme that you
// prefer.
const (
	KindUnknown       K = "unknown"
	KindNotAuthorized K = "not_authorized"
)

// M represents arbitrary metadata attached to a specific instance of an error.
// This may be the parameters to the erroring function or method or other data
// that may aid in troubleshooting, grouping, etc.
type M map[string]interface{}

// Metadata returns the metadata associated with the provided error if any. If the
// error doesn't implement Interface, Metadata returns nil.
func Metadata(err error) M {
	if m, ok := err.(Interface); ok {
		return m.GetMetadata()
	}
	return nil
}
