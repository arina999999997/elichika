//go:build embedded

package clientdb

// embedded do not modify client files
func isNotChanged(file string) bool {
	return true
}
