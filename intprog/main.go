//Package intprog implements integer programs for advent of code 19
package intprog

import "fmt"

type Error struct {
	cause string
	ip    int
}

func (err *Error) Error() string {
	return fmt.Sprintf("IntProg: Terminated with error %s with ip at %d", err.cause, err.ip)
}

func Run(nums []int, noun, verb int) ([]int, error) {
	ip := 0
	mem := make([]int, len(nums))
	copy(mem, nums)
	mem[1] = noun
	mem[2] = verb
	for {
		if ip < 0 || ip >= len(mem) {
			return mem, &Error{"Invalid IP", ip}
		}
		switch mem[ip] {
		case 99:
			//fmt.Println("99 exit")
			return mem, nil
		case 1:
			if ip+3 >= len(mem) {
				return mem, &Error{"add: operand out of range", ip}
			}
			op1 := mem[mem[ip+1]]
			op2 := mem[mem[ip+2]]
			dst := mem[ip+3]
			if dst >= len(mem) {
				return mem, &Error{"add: dst out of range", ip}
			}
			mem[dst] = op1 + op2
		case 2:
			if ip+3 >= len(mem) {
				return mem, &Error{"mul: operand out of range", ip}
			}
			op1 := mem[mem[ip+1]]
			op2 := mem[mem[ip+2]]
			dst := mem[ip+3]
			if dst >= len(mem) {
				return mem, &Error{"mul: dst out of range", ip}
			}
			mem[dst] = op1 * op2
		}
		ip += 4
	}

}
