
package bob
import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

remark = strings.Trim(remark," ")

if strings.HasSuffix(remark,"?"){
	if strings.ToUpper(remark) == remark && strings.ContainsAny(remark,("ABCDEFGHIJKLMNOPQRSTUVWXYZ")){
		return "Calm down, I know what I'm doing!"
	} else {
		return "Sure."
	}

} else if strings.ToUpper(remark) == remark && strings.ContainsAny(remark,("ABCDEFGHIJKLMNOPQRSTUVWXYZ")) {
	return "Whoa, chill out!"

	} else if remark == "" ||  strings.ContainsAny(remark,("\t")){
	return "Fine. Be that way!"

	} else {
	return "Whatever."
}

}
