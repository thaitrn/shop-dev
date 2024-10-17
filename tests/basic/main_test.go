package basic

import (
	"testing"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), output = %d, actual = %d", input, output, actual)
	// }

	// assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 3
	)

	actual := AddOne2(1)
	if actual != output {
		t.Errorf("AddOne2(%d), output = %d, actual = %d", input, output, actual)
	}

	// assert.Equal(t, AddOne2(2), 3, "AddOne(2) should be 3")
}

// func TestRequire(t *testing.T) {
// 	require.Equal(t, 2, 3)
// 	fmt.Println("not execute")
// }

// func TestAssert(t *testing.T) {
// 	assert.Equal(t, 2, 3)
// 	fmt.Println("execute...")
// }
