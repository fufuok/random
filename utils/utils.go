package utils

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"text/tabwriter"
	"unsafe"
)

// S2B StringToBytes
func S2B(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}

// B2S BytesToString
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// AssertEqual checks if values are equal
// Ref: gofiber/utils
func AssertEqual(t testing.TB, expected, actual interface{}, description ...string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	var aType = "<nil>"
	var bType = "<nil>"

	if expected != nil {
		aType = reflect.TypeOf(expected).String()
	}
	if actual != nil {
		bType = reflect.TypeOf(actual).String()
	}

	testName := "AssertEqual"
	if t != nil {
		testName = t.Name()
	}

	_, file, line, _ := runtime.Caller(1)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 0, 5, ' ', 0)
	_, _ = fmt.Fprintf(w, "\nTest:\t%s", testName)
	_, _ = fmt.Fprintf(w, "\nTrace:\t%s:%d", filepath.Base(file), line)
	if len(description) > 0 {
		_, _ = fmt.Fprintf(w, "\nDescription:\t%s", description[0])
	}
	_, _ = fmt.Fprintf(w, "\nExpect:\t%v\t(%s)", expected, aType)
	_, _ = fmt.Fprintf(w, "\nResult:\t%v\t(%s)", actual, bType)

	result := ""
	if err := w.Flush(); err != nil {
		result = err.Error()
	} else {
		result = buf.String()
	}

	if t != nil {
		t.Fatal(result)
	} else {
		log.Fatal(result)
	}
}
