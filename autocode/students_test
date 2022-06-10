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

func TestMatrixNew(t *testing.T) {
	ass := assert.New(t)

	m1, _ := New("")

	ass.Equal(m1, (*Matrix)(nil))

	m2, _ := New("\n")

	ass.Equal(m2, (*Matrix)(nil))

	m3, _ := New("1 2\n1")

	ass.Equal(m3, (*Matrix)(nil))

	_, err1 := New("1")

	ass.Equal(err1, nil)

	_, err2 := New("1 2 3\n1 2 3")

	ass.Equal(err2, nil)
}

func TestMatrixRowsCols(t *testing.T) {
	ass := assert.New(t)

	m, err := New("1 2")

	ass.Equal(err, nil)

	cols := m.Cols()

	ass.Equal(true, cols[0][0] == 1 && cols[1][0] == 2 && len(cols) == 2 && len(cols[0]) == 1 && len(cols[1]) == 1)

	rows := m.Rows()

	ass.Equal(true, rows[0][0] == 1 && rows[0][1] == 2 && len(rows) == 1 && len(rows[0]) == 2)
}

func TestMatrixSet(t *testing.T) {
	ass := assert.New(t)

	m, err := New("0")

	ass.Equal(err, nil)

	ass.Equal(false, m.Set(-1, -1, 1))

	ass.Equal(false, m.Set(-1, 0, 1))

	ass.Equal(false, m.Set(0, -1, 1))

	ass.Equal(false, m.Set(1, 1, 1))

	ass.Equal(false, m.Set(1, 0, 1))

	ass.Equal(false, m.Set(1, 1, 1))

	ass.Equal(true, m.Set(0, 0, 1))

	ass.Equal(1, m.data[0])
}
