package libtest

import "lib-test/internal"

type LibTest struct{}

func (lt *LibTest) hello() string {
	return "hi from lib test " + internal.Queue()
}
