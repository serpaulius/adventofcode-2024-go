package seventeenth

import (
	"fmt"
	"testing"
)

func TestOperand2(t *testing.T) {
	computer := Computer{
		c:       9,
		program: []int64{2, 6},
	}
	var wantB int64 = 1
	computer.run()
	if computer.a == 0 && computer.b == wantB && computer.c == 9 {
		return
	}
	t.Fatal(`Opcode 2 failed, only B register should be changed to 1, got`, computer)
}

func TestOperand5(t *testing.T) {
	computer := Computer{
		a:       10,
		program: []int64{5, 0, 5, 1, 5, 4},
	}
	computer.run()
	if fmt.Sprint(computer.output) == fmt.Sprint([]int64{0, 1, 2}) {
		return
	}
	t.Fatal(`Opcode 5 failed, should output 0,1,2, got`, computer)
}

func TestOperands03(t *testing.T) {
	computer := Computer{
		a:       2024,
		program: []int64{0, 1, 5, 4, 3, 0},
	}
	computer.run()
	if fmt.Sprint(computer.output) == fmt.Sprint([]int64{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}) && computer.a == int64(0) {
		return
	}
	t.Fatal(`Opcode 0, 5 or 3 failed, should output 4,2,5,6,7,7,7,7,3,1,0, got`, computer.output)
}

func TestOperand1(t *testing.T) {
	computer := Computer{
		b:       29,
		program: []int64{1, 7},
	}
	computer.run()
	if computer.a == 0 && computer.b == int64(26) && computer.c == 0 {
		return
	}
	t.Fatal(`Opcode 1 failed, B should be 26, got`, computer.b)
}

func TestOperand4(t *testing.T) {
	computer := Computer{
		b:       2024,
		c:       43690,
		program: []int64{4, 0},
	}
	computer.run()
	if computer.a == 0 && computer.b == int64(44354) && computer.c == 43690 {
		return
	}
	t.Fatal(`Opcode 4 failed, B should be 44354, got`, computer.b)
}

func TestOperandCombo(t *testing.T) {
	computer := Computer{
		a:       729,
		program: []int64{0, 1, 5, 4, 3, 0},
	}
	computer.run()
	if fmt.Sprint(computer.output) == fmt.Sprint([]int64{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}) {
		return
	}
	t.Fatal(`Opcode 0, 5 or 3 failed, should output 4,6,3,5,6,3,5,2,1,0, got`, computer.output)
}
