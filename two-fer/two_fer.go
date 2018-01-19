
package twofer

// Check input 
func ShareWith(s string) string {
	if len(s) > 0 {
		return "One for " + s + ", one for me."
	} else {
		return "One for you, one for me."
	}

}
