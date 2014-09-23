/* #nb:
A Good presenti Example

The presenti Authors
https://www.github.com/dupoxy/presenti/
*/

// A presenti test file.
package main

import (
	"fmt"
)

/* #nb:
* Title of slide or section (must have asterisk)
Some Text
** Subsection
- bullets
- more bullets
- a bullet with
*** Sub-subsection
Some More text
	Preformatted text
	is indented (however you like)
Further Text, including invocations like:
.code test.go /^func main/,/^}/
.play test.go
# .image image.jpg
# .iframe http://foo
*/

/* #nb: Their IS NO new Line at end of comment
 */
func main() {

	/* #nb: THIS IS NOT a top level comment
	Bla bla */

	fmt.Println("good")
}
/* #nb: EOF comment Test. */
