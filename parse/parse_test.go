/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import ()

var (
	src = `
/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse
import ()

// @ginger
type SlackUser struct {
	Name string
}
// don't scan
type SlackChannel struct {
	Name string
}
// @ginger
type SlackMessage struct {
	Name string
}

func main() {
}
`
)

func ExampleParse() {
	Scan(src, "")
	//Output:
	//SlackUser
	//SlackMessage
}
