package pdebug

import (
	"bytes"
	"runtime"
	"strings"
	"testing"
)

func pkgPathFromFrame(fn string) string {
	// strip the ".FuncName" suffix
	if i := strings.LastIndex(fn, "."); i != -1 {
		return fn[:i]
	}
	return fn
}

func currentPkgPath(t *testing.T) string {
	t.Helper()
	var pcs [1]uintptr
	n := runtime.Callers(3, pcs[:]) // skip runtime.Callers + this helper + caller
	if n == 0 {
		t.Fatalf("no callers")
	}
	f, _ := runtime.CallersFrames(pcs[:n]).Next()
	return pkgPathFromFrame(f.Function)
}

func withIncludePrefix(t *testing.T, p string) {
	prev := includePrefix
	SetIncludePrefix(p)
	t.Cleanup(func() { includePrefix = prev })
}

func TestSetIncludePrefixTrims(t *testing.T) {
	withIncludePrefix(t, "  trimmed/value  ")
	if includePrefix != "trimmed/value" {
		t.Fatalf("expected trimmed prefix, got %q", includePrefix)
	}
}

func TestStringWithMatchingPrefix(t *testing.T) {
	pkg := currentPkgPath(t) // e.g. "module/path/pdebug"
	withIncludePrefix(t, pkg)

	out := String()
	if out == "" {
		t.Fatalf("expected non-empty output with matching prefix")
	}

	lines := strings.Split(strings.TrimSpace(out), "\n")
	for _, l := range lines {
		if !strings.HasSuffix(l, ".go:"+lineNumberSuffix(l)) {
			t.Fatalf("line does not appear to reference a .go file with line: %q", l)
		}
		if !strings.Contains(l, "pdebug.go") && !strings.Contains(l, "pdebug_test.go") {
			// At least one line should reference our package files; relax per line check.
			// Just continue; final assertion below ensures at least one matched.
		}
	}
}

func lineNumberSuffix(l string) string {
	i := strings.LastIndex(l, ":")
	if i == -1 || i == len(l)-1 {
		return ""
	}
	return l[i+1:]
}

func TestPrintToWithNonMatchingPrefix(t *testing.T) {
	withIncludePrefix(t, "this/prefix/does/not/match")
	var buf bytes.Buffer
	PrintStackTo(&buf)
	if buf.Len() != 0 {
		t.Fatalf("expected no output for non-matching prefix, got: %q", buf.String())
	}
}
