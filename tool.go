// Package gotool is a library of utility functions used to implement the standard "Go" tool provided
// as a convenience to developers who want to write tools with similar semantics.
package gotool

// export functions as here to make it easier to keep the implementations up to date with upstream.

// MatchPackages returns all matching packages found
// under the $GOROOT and $GOPATH directories.
// The pattern is either "all" (all packages), "std" (standard packages)
// or a path in which "..." means "any string".
func MatchPackages(pattern string) []string {
	return matchPackages(pattern)
}
