package seventeenth

import (
	"adventofcode/2024-go/util"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// 3 bit opcodes are 0-7
// registers are ints
// opcode is succeeded by operand
// instruciton pointer -> opcode, after it is done, pointer +2 -> next opcode (except jnz 3)
// after instruction, pointer goes 2 right
// out of bounds - halt
type Computer struct {
	a       int64
	b       int64
	c       int64
	program []int64
	output  []int64
}

func (c *Computer) outputStr() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.output)), ","), "[]")
}

func parseInput(lines []string) *Computer {
	instructions := strings.Split(strings.Split(lines[4], "Program: ")[1], ",")
	var program []int64 = make([]int64, 0)
	for _, i := range instructions {
		val, _ := strconv.ParseInt(i, 10, 64)
		program = append(program, val)
	}

	a, _ := strconv.ParseInt(strings.Split(lines[0], "A: ")[1], 10, 64)
	b, _ := strconv.ParseInt(strings.Split(lines[1], "B: ")[1], 10, 64)
	c, _ := strconv.ParseInt(strings.Split(lines[2], "C: ")[1], 10, 64)
	computer := Computer{
		a: a, b: b, c: c,
		program: program,
	}
	return &computer
}

// literal operand 0-7
// combo operand 0-3 literal, 4-6 ABC, 7 NOT POSSIBLE (yet)
func (c *Computer) comboOperatorValue(combo int64) int64 {
	if combo <= 3 {
		return combo
	}
	switch combo {
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	}
	panic("combo operand should not be 7 or more")
}

func (c *Computer) dv(combo int64) int64 {
	pow := math.Pow(2, float64(c.comboOperatorValue(combo)))
	return c.a / int64(pow)
}

// 0 adv / div 			-> A = A div 2^combo operand //// truncate to integer
// 1 bxl / bitwise xor 	-> B = B xor literal
// 2 bst / modulo 8		-> B = combo operand % 8
// 3 jnz / nthn or jump -> A != 0 ? jmp to instruction pointer +1 INSTEAD of +2 the pointer)
// 4 bxc / bitwise xor  -> B = B xor C (ignore operand, still +2)
// 5 out / output combo op mod 8 -> out << combo operator
// 6 bdv / div 			-> B !!! = A div 2^combo operand //// truncate to integer
// 7 cdv / div			-> C !!! = A div 2^combo operand //// truncate to integer
func (c *Computer) run() {
	for i := 0; i < len(c.program)-1; {
		opcode := c.program[i]
		switch int64(opcode) {
		case int64(0):
			c.a = c.dv(c.program[i+1])
		case 1:
			c.b = c.b ^ c.program[i+1]
		case 2:
			c.b = c.comboOperatorValue(c.program[i+1]) % 8
		case 3:
			if c.a != 0 {
				i = int(c.program[i+1])
			} else {
				i += 2
			}
		case 4:
			c.b = c.b ^ c.c
		case 5:
			c.output = append(c.output, c.comboOperatorValue(c.program[i+1])%8)
		case 6:
			c.b = c.dv(c.program[i+1])
		case 7:
			c.c = c.dv(c.program[i+1])
		}
		if opcode != 3 {
			i += 2
		}
	}
	log.Println("HALT")
}

func Run() {
	file := util.OpenFileOrPanicPlz("./17/input.txt")
	defer util.CloseFileOrPanicPlz(file)
	scanner := util.GiveMeAScannerPlz(file)
	input := util.ReadLines(scanner)

	computer := parseInput(input)
	computer.run()

	fmt.Println("17.1 out", computer.outputStr())
}
