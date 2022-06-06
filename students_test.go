package coverage

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleLen(t *testing.T) {
	ass := assert.New(t)
	p := People{}

	ass.GreaterOrEqual(0, p.Len())
}

func TestPeopleLess(t *testing.T) {
	ass := assert.New(t)

	tm := time.Now()

	p1 := People{
		{firstName: "test", lastName: "test", birthDay: tm.Add(time.Second)},
		{firstName: "test", lastName: "test", birthDay: tm},
	}

	ass.Equal(true, p1.Less(0, 1))

	p2 := People{
		{firstName: "test", lastName: "test", birthDay: tm},
		{firstName: "test", lastName: "test2", birthDay: tm},
	}

	ass.Equal(true, p2.Less(0, 1))

	p3 := People{
		{firstName: "test", lastName: "test", birthDay: tm},
		{firstName: "test2", lastName: "test", birthDay: tm},
	}

	ass.Equal(true, p3.Less(0, 1))

	p4 := People{
		{firstName: "test", lastName: "test", birthDay: tm},
		{firstName: "test", lastName: "test", birthDay: tm},
	}

	ass.Equal(false, p4.Less(0, 1))
}

func TestPeopleSwap(t *testing.T) {
	ass := assert.New(t)

	tm := time.Now()

	p := People{
		{firstName: "test1", lastName: "test", birthDay: tm},
		{firstName: "test2", lastName: "test", birthDay: tm},
	}

	p.Swap(0, 1)

	ass.Equal(false, p.Less(0, 1))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

//func TestMatrixNew(t *testing.T) {
//	ass := assert.New(t)
//
//	}
