# functiongo
Function related package in go lang

### Function Name
For now this package contains only function name functionality.
Function name will return function name of the function from which the function is called.

## Using the package
### Getting the package
`go get github.com/jastisriradheshyam/functiongo/`

### Importing the package
`import fn "github.com/jastisriradheshyam/functiongo/"`

### Functions

#### CallerName(skiplevel)
Retunrs caller name _(function name from which CallerName is called)_
#### TraceLog()
Returns filename, line number in file in which function invoked,  function name of function in which "TraceLog" function is encapculated   
Format - __"Filename: file_name | Line: line_number | Function: function_name"__
