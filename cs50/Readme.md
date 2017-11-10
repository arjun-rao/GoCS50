## CS50 Library for Go
---
### To install:

```
$go get github.com/arjun-rao/cs50/cs50
```

### Import the package in your code like so:
<pre>
package main

import (
    "fmt"
    <b>"github.com/arjun-rao/cs50/cs50"</b>
)

func main() {
    name := cs50.GetString("Enter your name:")
    fmt.Printf("Hello %s", name)
}
</pre>

