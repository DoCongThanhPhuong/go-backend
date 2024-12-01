package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
  // var (
  //   input = 1
  //   output = 2
  // )

  // actual := AddOne(1)
  // if(actual != output) {
  //   t.Errorf("AddOne(%d), input=%d, autual=%d", input, output, actual)
  // }

  assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
  assert.NotEqual(t, AddOne(2), 4)
  assert.Nil(t, nil, nil)
}

// func TestRequire(t *testing.T) {
//   require.Equal(t, 2, 3)
//   fmt.Println("Not execute")
// }

// func TestAssert(t *testing.T) {
//   assert.Equal(t, 2, 3)
//   fmt.Println("Execute")
// }

// CMD
// go test --v --coverprofile=coverage.out
// go tool cover --html=coverage.out -o coverage.html