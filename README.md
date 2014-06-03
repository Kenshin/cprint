Color Print: ANSI coloured text to the standard output on Windows
================================

Documentation
---
Visit the docs on [Godoc](http://godoc.org/github.com/Kenshin/cprint)

Importing
---
`go get github.com/Kenshin/cprint`

Example
---
	import . "github.com/Kenshin/cprint"

	// test ont
	P(WARING, "Remote latest version %v %v latest version %v.\n", "0.10.28", "=", "0.1.0.26")

	// test tow
	cp := CP{Red, false, None, false, "="}
	P(NOTICE, "Remote latest version %v %v latest version %v, don't need to upgrade.\n", "0.10.28", cp, "0.1.0.26")

	// test three
	err := errors.New("Variable is not defined.")
	Error(ERROR, "'gnvm updte latest' an error has occurred. \nError: ", err)

Screenshots
--
![color print](http://i.imgur.com/OjFVyKI.png)

Help
---
* Email <kenshin@ksria.com>
* Github issue

CHANGELOG
---
* **2014-05-28, Version `0.0.1` support:**
    * `CP` *struct*
    * `P` *method*
    * `Error` *method*

LICENSE
---
(The MIT License)

Copyright (c) 2014 Kenshin Wang <kenshin@ksria.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
