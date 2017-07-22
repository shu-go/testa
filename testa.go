package testa

import (
	"flag"
)

var testing bool

func init() {
	testing = (flag.Lookup("test.v") != nil)
}

func IsTesting() bool {
	return testing
}
