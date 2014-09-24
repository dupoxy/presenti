/* #nb:
A presenti test file

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

** Subsection (slide does not support Subsection)

- bullets
- more bullets
- a bullet with

*** Sub-subsection

Some More text

  Preformatted text
  is indented (however you like)

Further Text, including invocations like:

.code test.go /^func main/,/^}/

The above code should have `func main` line highlighted.

.code test.go /^func main/,/^}/ HL001

The above code should have `fmt.Println` line highlighted.

.code -edit test.go /^func main/,/^}/

The above code should have `-edit` activated.

.code -numbers test.go /^func main/,/^}/

The above code should have `-numbers` activated.

.play test.go /^func main/,/^}/

The above code should have `func main` line highlighted and be playable.

.play test.go /^func main/,/^}/ HL001

The above code should have `fmt.Println` line highlighted and be playable.

.play -edit test.go /^func main/,/^}/

The above code should have `-edit` activated and be playable.

.play -numbers test.go /^func main/,/^}/

The above code should have `-numbers` activated and be playable.

# This should return an error ? .html, .code and .play does.
.image image.jpg

The above image should return an error ? .html, .code and .play does.

.iframe http://foo
.link http://foo label
.caption _Gopher_ by [[http://www.reneefrench.com][Renée French]]

Again, more text
*/

/* #nb: Their IS NO new Line at end of comment
 */
func main() { // HL

	/* #nb: THIS IS NOT a top level comment */

	fmt.Println("good") // HL001
}

/* #nb: EOF comment Test. */
