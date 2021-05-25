package primitive

import (
	"os"
	"testing"
)

func TestMode(t *testing.T) {

	// exp := func() []string {
	// 	return []string{"-m", "0"}
	// }
	_ = WithMode(ModeCircle)

}

func TestTransform(t *testing.T) {
	opts := WithMode(ModeCircle)
	file, _ := os.Open("input.jpeg")
	_, err := Transform(file, ".jpeg", 10, opts)
	if err != nil {
		panic(err)
	}
}

func TestTransformMock(t *testing.T) {
	opts := WithMode(ModeCircle)
	f1, _ := os.Open("input.jpeg")
	f2, _ := os.Open("input.jpeg")
	f3, _ := os.Open("input.jpeg")
	f4, _ := os.Open("input.jpeg")
	f5, _ := os.Open("input.jpeg")

	Mock1 = true
	_, err := Transform(f1, ".jpeg", 10, opts)
	CheckError(err)
	Mock1 = false

	Mock2 = true
	_, err = Transform(f2, ".jpeg", 10, opts)
	CheckError(err)
	Mock2 = false

	Mock3 = true
	_, err = Transform(f3, ".jpeg", 10, opts)
	CheckError(err)
	Mock3 = false

	Mock4 = true
	_, err = Transform(f4, ".jpeg", 10, opts)
	CheckError(err)
	Mock4 = false

	Mock5 = true
	_, err = Transform(f5, ".jpeg", 10, opts)
	CheckError(err)
	Mock5 = false

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func TestPrimitive(t *testing.T) {

	var args []string
	opts := WithMode(ModeCircle)

	args = append(args, opts()...)
	_, err := Primitive("input.jpeg", "out.jpeg", 10, args...)
	if err != nil {
		t.Fatalf("%v", args)
	}
}

func TestTempFile(t *testing.T) {
	Mock6 = true
	_, err := tempFile("", "")
	CheckError(err)
	Mock6 = false
	_, err = tempFile("", "")
	if err != nil {
		panic(err)
	}
}
