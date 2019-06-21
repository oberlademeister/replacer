package replacer

import (
	"fmt"
	"strings"
)

// Positional transforms the in argument by replacing all placeholders
// such as {0} by the coresponding array elements of replaceWith
func Positional(in string, debug bool, replaceWith ...string) string {
	var newArgs []string
	for i, a := range replaceWith {
		newArgs = append(newArgs, fmt.Sprintf("{%d}", i))
		if debug {
			newArgs = append(newArgs, fmt.Sprintf("{%d->%s}", i, a))
		} else {
			newArgs = append(newArgs, a)
		}
	}
	r := strings.NewReplacer(newArgs...)
	return r.Replace(in)
}

// Map transforms the in argument by replacing all placeholders
// such as {key} by the corresponding values in the replaceWith Map
func Map(in string, debug bool, replaceWith map[string]string) string {
	var newArgs []string
	for k, v := range replaceWith {
		newArgs = append(newArgs, fmt.Sprintf("{%s}", k))
		if debug {
			newArgs = append(newArgs, fmt.Sprintf("{%s->%s}", k, v))
		} else {
			newArgs = append(newArgs, v)
		}
	}
	r := strings.NewReplacer(newArgs...)
	return r.Replace(in)
}
