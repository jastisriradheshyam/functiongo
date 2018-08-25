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

package functiongo

import (
	"fmt"
	"runtime"
	"strings"
)

// TraceLog shows :
// - Filename of file in which code was written,
// - Line Number of code in the file from where TraceLog function is called
// - Function Name of the function that enapculated the TraceLog function
// return string format - "Filename: file_name | Line: line_number | Function: function_name"
func TraceLog() string {
	// we get the callers as uintptrs - but we just need 2 or more for Frames to work
	fpcs := make([]uintptr, 2)

	// 2nd level to get to the caller of whoever called TraceLog()
	n := runtime.Callers(2, fpcs)
	// Check if pc have data
	if n == 0 {
		return "-"
	}
	// Get the frames for program counter range
	frames := runtime.CallersFrames(fpcs[:n])
	// Get the frame
	frame, _ := frames.Next()
	// Get the function name
	// For example:
	// main.(*ABC).MyFunc --> MyFunc  (with Structure)
	// main.MyFunc --> MyFunc  (Without Structure)
	fn := strings.Split(frame.Function, ".")
	// Return the filename, line number, and function name in formatted string
	return fmt.Sprintf("Filename :%s | Line :%d | Function :%s", frame.File, frame.Line, fn[len(fn)-1])
}
