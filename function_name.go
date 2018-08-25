// Copyright 2018 Jasti Sri Radhe Shyam

// Permission is hereby granted, free of charge, to any person obtaining 
// a copy of this software and associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation 
// the rights to use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, and to permit persons to whom the 
// Software is furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in 
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR 
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN 
// THE SOFTWARE.
//
// MIT License : https://opensource.org/licenses/MIT

package function_name

// Get the fucntion name that called this fucntion
// @args levelArg ...int `to set the level (depth of stack) for that function name`
// levelArg --> 2 ,then caller function name
// levelArg --> 3, then callers caller function name
// levelArg --> n, then n-2th callers caller name
// levelArg --> 1, then function being called that is "CallerName" function it self
func CallerName(levelArg ...int) string {
	
	// level is a variable that sets the level of stack of functions
	var level int
	
	// Check the argument
	// If no arguments then default value of level = 2
	// else first argument is assigened to level
	if len(levelArg) == 0 {
		level = 2
	} else {
		level = levelArg[0]
	}
	
	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip levels to get to the caller of whoever called Caller()
	n := runtime.Callers(level, fpcs)
	if n == 0 {
		return "No function in the stack at the skip level : "+ strconv.Itoa(level)
	}

	// get the info of the actual function that's in the pointer
	function := runtime.FuncForPC(fpcs[0])
	if function == nil {
		return "No function name is available"
	}

	// split the string e.g: main.functionName = main , functionName
	functionSplit := strings.Split(function.Name(), ".")
	// return the function name
	// Gets the last element of the array(Slice) functionSplit and removes inconsistency
	// For example:
	// main.(*ABC).MyFunc --> MyFunc  (with Structure)
	// main.MyFunc --> MyFunc  (Without Structure)
	return functionSplit[len(functionSplit)-1]
}
