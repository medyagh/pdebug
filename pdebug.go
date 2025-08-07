package pdebug

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

// Print writes a filtered stack trace to stderr, showing only frames that belong
// to the main module (e.g., "k8s.io/minikube/..."). Each line is "<abs file>:<line>".
// Example:
//
//	/Users/you/workspace/minikube/pkg/minikube/out/out.go:507
//	/Users/you/workspace/minikube/pkg/minikube/out/out_reason.go:99
//	...
func Print() { PrintTo(os.Stderr) }

// PrintTo is like Print, but writes to the given io.Writer.
func PrintTo(w io.Writer) {
	printFiltered(w, 1, detectedMainModulePath())
}

// SetIncludePrefix allows overriding the module prefix used to filter frames.
// Leave empty to revert to auto-detected main module. This is process-global.
func SetIncludePrefix(p string) { includePrefix = strings.TrimSpace(p) }

// String returns the filtered stack as a string.
func String() string {
	var b strings.Builder
	printFiltered(&b, 1, detectedMainModulePath())
	return b.String()
}

// ----- internals -----

var includePrefix string // optional override (e.g., "k8s.io/minikube")

func detectedMainModulePath() string {
	if includePrefix != "" {
		return includePrefix
	}
	if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Path != "" {
		return bi.Main.Path
	}
	// No modules? Fallback: nothing filtered (prints nothing) unless user sets SetIncludePrefix.
	return ""
}

func printFiltered(w io.Writer, callerSkip int, modulePrefix string) {
	// Grab PCs; +3 to skip runtime.Callers + this function + caller.
	const maxDepth = 128
	var pcs [maxDepth]uintptr
	n := runtime.Callers(callerSkip+2, pcs[:])
	if n == 0 {
		return
	}

	frames := runtime.CallersFrames(pcs[:n])
	for {
		f, more := frames.Next()
		// Skip empty frames and our own package frames.
		if f.File != "" && f.Line > 0 && isUserFrame(f, modulePrefix) {
			// Print just "    /abs/path/file.go:LINE"
			fmt.Fprintf(w, "        %s:%d\n", f.File, f.Line)
		}
		if !more {
			break
		}
	}
}

func isUserFrame(f runtime.Frame, modulePrefix string) bool {
	fn := f.Function // fully-qualified, e.g. "k8s.io/minikube/pkg/minikube/out.displayGitHubIssueMessage"
	if modulePrefix == "" {
		// If we couldn't detect a module, default to "print nothing" unless user set a prefix.
		return false
	}
	// Match the module (and subpackages). Guard against weird function names.
	if strings.HasPrefix(fn, modulePrefix) {
		// Exclude tests if someone calls us from tests and you want only product code:
		// optionally add: if strings.HasSuffix(f.File, "_test.go") { return false }
		return true
	}
	return false
}
