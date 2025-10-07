package lintchecker

import "testing"

type testingT struct{ *testing.T }

func (t testingT) Fatalf(format string, args ...any) {
	t.T.Fatalf("[lesiw.io/lintchecker] "+format, args...)
}

func (t testingT) Errorf(format string, args ...any) {
	t.T.Errorf("[lesiw.io/lintchecker] "+format, args...)
}

func (t testingT) Logf(format string, args ...any) {
	t.T.Logf("[lesiw.io/lintchecker] "+format, args...)
}
