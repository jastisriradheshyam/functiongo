# functiongo

## trace.go

### TraceLog
TraceLog is the function which shows the filename, line number and function name of the function from which function "TraceLog" is called.  
In other words : To get the function name, filename, filename from where "TraceLog" called in source code   
__Format__ : "Filename: file_name | Line: line_number | Function: function_name"    

#### __Example__

__Code__:
```
package main

import (
	"fmt"
	fn "github.com/jastisriradheshyam/functiongo"
)
func main() {
	fmt.Println("hello, world\n")
	fmt.Println(fn.TraceLog())
}
```

__Output__:
```
hello, world
Filename: pathtothefile/hello.go | Line: 9 | Function: main
```
### References :
[1] : https://stackoverflow.com/questions/25927660/golang-get-current-scope-of-function-name   
[2] : https://play.golang.org/p/z2kHQJoeUo 

### Credits :
___ksrb___ -> https://stackoverflow.com/users/1490379/ksrb
