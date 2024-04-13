package libtest

import "github.com/deqinganz/lib-test/internal"

type LibTest struct{}

func (lt *LibTest) Hello() string {
	return "hi from lib test " + internal.Queue()
}
