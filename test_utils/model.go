package testutils

import "fmt"

type FakePrint struct {
	Args [2]string
}

func (f *FakePrint) FakePrintln(a ...any) (int, error) {
	f.Args[0] = fmt.Sprint(a[0])
	if len(a) > 1 {
		f.Args[1] = fmt.Sprint(a[1])
	}
	return 1, nil
}

type FakeProcess struct {
	Called int
}

func (f *FakeProcess) ProcessFunTest() {
	f.Called = f.Called + 1
	fmt.Println("Testing.....")
}

type FakeExit struct {
	Code int
}

func (f *FakeExit) Exit(code int) {
	f.Code = code
}
