# errors

[![Build Status](https://github.com/haleyrc/errors/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/haleyrc/errors/actions?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/haleyrc/errors?status.svg)](https://pkg.go.dev/github.com/haleyrc/errors?tab=doc)

[TODO: description]

## Install

```
$ go get -u github.com/haleyrc/errors
```

## Motivation

The methods in Interface and their corresponding concepts are born out of experience attempting to group and work with errors in a meaningful way. The combination of the "Big Four" values covers the most common needs of both frontend clients who receive these errors in response to failed requests as well as backend engineers attempting to diagnose root cause among other things.

### Code

The simplest of the Big Four values is the code. This allows any part of your request pipeline to return an error that can self-identify which HTTP status code corresponds to the error condition. This allows the same kind of mapping behavior that you might otherwise implement with a type switch, but in a more natural way than maintaining a big, literal map from type to code.

As a quick example, assume you get an `sql.ErrNoRows` error back when attempting to get a record from your database. You could pass this back up through your stack and, at the point you are creating the response, have something like the following:

```go
var code int

cause := errors.Cause(err)

switch cause.(type) {
    case sql.ErrNoRows:
        code = http.StatusNotfound
    // A bunch more cases...
    default:
        code = http.StatusInternalServerError
}
```

This approach has a number of problems, however. First, you end up needing a potentially massive switch statement to handle all of the various errors that may have gotten you here. On top of that, however, you may even need to do additional processing to get an indicator that lets you know what code to return if the type alone isn't meaningful. Finally, whatever code is doing this mapping now has dependencies on any packages that can return errors that we need to check again.

By using the `errors` package, you can centralize this processing where it makes the most sense and only have a single dependency at the end of your stack. Assuming your data layer now handles the translation of errors from the `sql` package to errors that implement `errors.Interface`, your response code can be obtained as simply as:

```go
code := errors.Code(err)
```

The `Code` function will even do the work of obtaining the root error for you, so you're free to wrap your errors to your heart's content.

### Kind

While many applications can function just fine using only HTTP status codes, it's often the case that HTTP codes alone don't adequately describe the space of error conditions that can occur or are too vague to be useful. For that reason, it's often helpful to surface a separate "code" that is still machine-readable, but can be customized to whatever your API contract requires. That's the purpose of an error's "kind" in this package.

While the specific values of the kind string aren't important (as long as they're meaningful in your context), they should be considered "static" in much the same way as the HTTP codes. This allows clients to specificy behavior based on the "class" of the error without worrying about that behavior changing
unexpectedly.

This package exposes a number of kinds for use in the default error types, but you can define your own kinds easily enough to ensure that you're accurately modeling your own domain.

### Message

Errors that are returned to clients as the result of a failed request can be the result of a fairly complicated error state. As a result, attempting to compile a meaningful error message to present to users client-side can be difficult or impossible. The result of this is a spate of "Something went wrong" style messages that severely degrade the user experience.

This package takes an alternate approach of encouraging the server to return a
meaningful, user-friendly error message. This is a much simpler task since the server has all of the context surrounding the error.

There aren't a lot of rules around writing a good message here, but as some general guidance I find that it's best to imagine this message in a flash message within your (or a client's) UI. If you think about what a user's next step would need to be to recover, your message should guide them to that move. If they can change a value and resubmit a form, for instance, then you might be able to use something like `The date you selected for your appointment is unavailable. Select a different date and try again.`. If there's no good next step as is often the case in a server error, then it can be helpful to default to something like `Something unexpected went wrong. Try again in a few minutes and if the issue persists, contact the help desk at help@example.com or by calling 1-800-555-1234.`.

**N.B.** Since these messages are intended to be display to users, it is critical to avoid any kind of sensitive information. This includes sensitive user information such as PHI as well as potentially dangerous server information such as stack traces.

### Metadata

## Usage

```json
{
  "error": {
    "kind": "not_authorized",
    "message": "You do not have permission to invite users to the organization. Please contact your organization administrator."
  }
}
```

### Interface

The heart of the errors package is the Interface. This is the contract that all of the default errors implement, meaning that all of the top-level "extraction" functions (Code, Kind, Message, Metadata) work natively with them. You are not restricted to using the default errors however. Any error that implements the same interface can be used in exactly the same way. To keep from blowing up your application when an "uncompliant" error (that is, one that does not implement Interface) makes it through the stack. In those cases, the extraction methods simply return default values that should be "safe" to use (if not a bit generic).
