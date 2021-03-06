// Test for name linting.

// Package pkg_with_underscores ...
package pkg_with_underscores // MATCH /underscore.*package name/

import (
	"io"
	"net"
	net_http "net/http" // renamed deliberately
	"net/url"
)

import "C"

var var_name int // MATCH /underscore.*var.*var_name/

type t_wow struct { // MATCH /underscore.*type.*t_wow/
	x_damn int      // MATCH /underscore.*field.*x_damn/
	URL    *url.URL // MATCH /struct field.*URL.*Url/
}

const fooID = "blah" // MATCH /fooID.*fooId/

func f_it() { // MATCH /underscore.*func.*f_it/
	more_underscore := 4 // MATCH /underscore.*var.*more_underscore/
	_ = more_underscore
	var err error
	if isEof := (err == io.EOF); isEof { // should be okay
		more_underscore = 7 // should be okay
	}

	x := net_http.Request{} // should be okay
	_ = x

	var ips []net.IP
	for _, theIp := range ips { // should be okay
		_ = theIp
	}

	switch myJson := g(); { // should be okay
	default:
		_ = myJson
	}
	var y net_http.ResponseWriter // an interface
	switch tApi := y.(type) {     // should be okay
	default:
		_ = tApi
	}

	var c chan int
	select {
	case qId := <-c: // should be okay
		_ = qId
	}
}

// Common styles in other languages that don't belong in Go.
const (
	CPP_CONST   = 1 // MATCH /ALL_CAPS.*CamelCase/
	kLeadingKay = 2 // MATCH /k.*leadingKay/

	HTML    = 3 // MATCH /HTML.*Html/
	X509B   = 4 // ditto
	V1_10_5 = 5 // okay; fewer than two uppercase letters
)

var kVarsAreSometimesUsedAsConstants = 0 // MATCH /k.*varsAreSometimesUsedAsConstants/
var (
	kVarsAreSometimesUsedAsConstants2 = 0 // MATCH /k.*varsAreSometimesUsedAsConstants2/
)

var kThisIsNotOkay = struct { // MATCH /k.*thisIsNotOkay/
	kThisIsOkay bool
}{}

func kThisIsOkay() { // this is okay because this is a function name
	var kThisIsAlsoOkay = 1 // this is okay because this is a non-top-level variable
	_ = kThisIsAlsoOkay
	const kThisIsNotOkay = 2 // MATCH /k.*thisIsNotOkay/
}

var anotherFunctionScope = func() {
	var kThisIsOkay = 1 // this is okay because this is a non-top-level variable
	_ = kThisIsOkay
	const kThisIsNotOkay = 2 // MATCH /k.*thisIsNotOkay/}
}

func f(bad_name int)                    {}            // MATCH /underscore.*func parameter.*bad_name/
func g() (no_way int)                   { return 0 }  // MATCH /underscore.*func result.*no_way/
func (t *t_wow) f(more_under string)    {}            // MATCH /underscore.*method parameter.*more_under/
func (t *t_wow) g() (still_more string) { return "" } // MATCH /underscore.*method result.*still_more/

type i interface {
	CheckHtml() string // okay; interface method names are often constrained by the concrete types' method names

	F(foo_bar int) // MATCH /foo_bar.*fooBar/
}

// All okay; underscore between digits
const case1_1 = 1

type case2_1 struct {
	case2_2 int
}

func case3_1(case3_2 int) (case3_3 string) {
	case3_4 := 4
	_ = case3_4

	return ""
}

type t struct{}

func (t) LastInsertId() (int64, error) { return 0, nil } // okay because it matches a known style violation

//export exported_to_c
func exported_to_c() {} // okay: https://github.com/golang/lint/issues/144

//export exported_to_c_with_arg
func exported_to_c_with_arg(but_use_go_param_names int) // MATCH /underscore.*func parameter.*but_use_go_param_names/

// This is an exported C function with a leading doc comment.
//
//export exported_to_c_with_comment
func exported_to_c_with_comment() {} // okay: https://github.com/golang/lint/issues/144

//export maybe_exported_to_CPlusPlusWithCamelCase
func maybe_exported_to_CPlusPlusWithCamelCase() {} // okay: https://github.com/golang/lint/issues/144

// WhyAreYouUsingCapitalLetters_InACFunctionName is a Go-exported function that
// is also exported to C as a name with underscores.
//
// Don't do that. If you want to use a C-style name for a C export, make it
// lower-case and leave it out of the Go-exported API.
//
//export WhyAreYouUsingCapitalLetters_InACFunctionName
func WhyAreYouUsingCapitalLetters_InACFunctionName() {} // MATCH /underscore.*func.*Why.*CFunctionName/
