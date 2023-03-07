package main

import (
	"fmt"

	"github.com/microcosm-cc/bluemonday"
)

func main4() {
	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.UGCPolicy()
	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)
	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
	htmlInject := `Hello <STYLE>.XSS{background-image:url("javascript:alert('XSS')");}</STYLE><A CLASS=XSS></A>World`
	htmlInjectSan := p.Sanitize(htmlInject)
	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(htmlInjectSan)
}
