package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello world !")
	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	fmt.Println("x is", x)
	fmt.Println("y is", y)
	sum1, prod := learnMultiple(x, y)
	fmt.Println("Sum is", sum1, "and Prod is", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Hello this is Khanh"
	s2 := `Raw string
	with line break`
	fmt.Println(str, s2)
	f := 3.14195
	c := 3 + 4i
	var u uint = 17
	var pi float32 = 22. / 3
	fmt.Println(f, c, u, pi)

	var a4 [4]int
	a4[1] = 123
	fmt.Println(a4)

	a5 := [...]int{1, 3, 5, 7, 9, 11, 13}
	a5[6] = 123
	fmt.Println(a5)

	s3 := []int{1, 2, 3, 4}
	fmt.Println(s3)
	s4 := make([]int, 4)
	fmt.Println(s4)
	bs := []byte("A Slice")
	fmt.Println(bs)

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)
	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"a": 1, "b": 7}
	m["c"] = 19
	fmt.Println(m)
	fmt.Println(m["c"])
	k := "b"
	fmt.Println(m[k])

	testz := learnNamedReturn(2, 3)
	fmt.Println(testz)

	// file, _ := os.Create("output.txt")
	// fmt.Fprint(file, "This is KhanhIceTea")
	// file.Close()

	learnFlowControl()
}

func learnNamedReturn(x, y int) (z int) {
	z = x + y
	return
}

func learnMemory() (p, q *int) {
	p = new(int)
	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func learnFlowControl() {
	if true {
		fmt.Println("We are universe !")
	}

	if false {

	} else {

	}

	x := 42.3
	switch x {
	case 0:
	case 1:
	case 42:
		fmt.Println("There is it !")
	default:
		fmt.Println("Ok boss !")
	}

	for x := 0; x < 3; x++ {
		fmt.Println("Step", x)
	}

	for k, v := range map[string]int{"a": 1, "b": 3} {
		fmt.Println(k, "=>", v)
	}

	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}

	if y := learnNamedReturn(3, 4); y < 8 {
		fmt.Println("7 < 8")
	}

	xBig := func() bool {
		return x > 10000
	}

	x = 99999
	fmt.Println("xBig:", xBig())
	x = 1.3e3
	fmt.Println("xBig:", xBig())

	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	goto love
love:
	// learnFunctionFactory()
	// learnDefer()
	learnInterfaces()
}

func learnFunctionFactory() {
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() {
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

func (p pair) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	p := pair{3, 4}
	fmt.Println(p.String())

	var i Stringer
	i = p
	fmt.Println(i.String())
	p.x = 7
	fmt.Println(p)
	fmt.Println(i)
	// learnVariadicParams("great", "learning", "here!")
	learnVariadicIntParams(1, 2, 3, 4, 5, 7)
	learnErrorHandling()
}

func learnVariadicParams(myStrings ...interface{}) {
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	fmt.Println("params:", fmt.Sprintln(myStrings...))
}

func learnVariadicIntParams(myInts ...int) {
	for _, param := range myInts {
		fmt.Println("param:", param)
	}
}

func learnErrorHandling() {
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok {
		fmt.Println("no one there")
	} else {
		fmt.Print(x)
	}
	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}

	learnWebProgramming()
}

func learnWebProgramming() {

	// First parameter of ListenAndServe is TCP address to listen to.
	// Second parameter is an interface, specifically http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

// Make pair an http.Handler by implementing its only method, ServeHTTP.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve data with a method of http.ResponseWriter.
	w.Write([]byte("You learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}
