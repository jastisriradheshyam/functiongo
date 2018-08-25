# functiongo

## function_name.go

### CallerName
CallerName is the function which shows the name of the function from which function "CallerName" is called.  
In other words : To get the function name   

#### Function Name in the function stack
It can also usefull to know the name of the function in the function stack.  
This can be done by providing the number as argument to the CallerName function

##### For Example :  
Code is like this  

```
func main() {
testfn1()
}

func testfn1() {
testfn2()
}

func testfn2() {
// some code
}
```
Then the stack is  
from testfn2() in go lang

testfn2 -> testfn1 -> main -> main -> goexit  

#### Using the CallerName
Code :
```
import(
  "fmt"
  fn "github.com/jastisriradheshyam/functiongo"
)
func main() {
testfn1()
}

func testfn1() {
testfn2()
}

func testfn2() {
fmt.Println(fn.CallerName(arg))
}
```
> If *__arg__ is __blank__ or __2__* then it will print __"testfn2"__  
> If *__arg__ is __3__* then it will print __"testfn1"__  
> If *__arg__ is __4__* then it will print __"main"__  
> If *__arg__ is __5__* then it will print __"main"__  
> If *__arg__ is __6__* then it will print __"goexit"__  

### References :
[1] : https://stackoverflow.com/questions/35212985/is-it-possible-get-information-about-caller-function-in-golang
