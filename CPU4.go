package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu   []string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func initialized() {

	cpu = make([]string, 2)
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU1   -> %s\n", cpu[0])
	fmt.Printf("CPU2   -> %s\n", cpu[1])
	fmt.Printf("Ready -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_new(p string) {
	if cpu[0] == "" && cpu[1] == "" {
		cpu[0] = p
	} else if cpu[0] != "" && cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue(ready, p)
	}

}
func command_terminate(terminate string) {
	switch terminate {
	case "cpu1":
		if cpu[0] != "" {
			cpu[0] = deleteQueue(ready)
		}
	case "cpu2":
		if cpu[1] != "" {
			cpu[1] = deleteQueue(ready)
		}
	default:
		fmt.Printf("terminate Error")
	}
}

func command_expire(io string) {

	switch io {
	case "cpu1":
		p := deleteQueue(ready)
		if p == "" {
			return
		}
		insertQueue(ready, cpu[0])
		cpu[0] = p
	case "cpu2":
		p := deleteQueue(ready)
		if p == "" {
			return
		}
		insertQueue(ready, cpu[1])
		cpu[1] = p
	default:
		fmt.Printf("error ")
	}
}

func command_io1(io string) {
	switch io {
	case "cpu1":
		insertQueue(io1, cpu[0])
		cpu[0] = ""
		command_expire(io)
	case "cpu2":
		insertQueue(io1, cpu[1])
		cpu[1] = ""
		command_expire(io)
	default:
		fmt.Printf("error")
	}

}

func command_io2(io string) {
	switch io {
	case "cpu1":
		insertQueue(io2, cpu[0])
		cpu[0] = ""
		command_expire(io)
	case "cpu2":
		insertQueue(io2, cpu[1])
		cpu[1] = ""
		command_expire(io)
	default:
		fmt.Printf("error")
	}
}

func command_io3(io string) {
	switch io {
	case "cpu1":
		insertQueue(io3, cpu[0])
		cpu[0] = ""
		command_expire(io)
	case "cpu2":
		insertQueue(io3, cpu[1])
		cpu[1] = ""
		command_expire(io)
	default:
		fmt.Printf("error")
	}
}

func command_io4(io string) {
	switch io {
	case "cpu1":
		insertQueue(io4, cpu[0])
		cpu[0] = ""
		command_expire(io)
	case "cpu2":
		insertQueue(io4, cpu[1])
		cpu[1] = ""
		command_expire(io)
	default:
		fmt.Printf("error")
	}
}
func command_io1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu[0] == "" {
		cpu[0] = p
	} else if cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu[0] == "" {
		cpu[0] = p
	} else if cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu[0] == "" {
		cpu[0] = p
	} else if cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue(ready, p)
	}
}

func command_io4x() {

	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu[0] == "" {
		cpu[0] = p
	} else if cpu[1] == "" {
		cpu[1] = p
	} else {
		insertQueue(ready, p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_new(commandx[i])
			}

		case "terminate":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_terminate(commandx[i])
			}
		case "expire":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_expire(commandx[i])
			}
		case "io1":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_io1(commandx[i])
			}
		case "io2":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_io2(commandx[i])
			}
		case "io3":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_io3(commandx[i])
			}
		case "io4":
			for i := range commandx {
				if i == 0 {
					continue
				}
				command_io4(commandx[i])
			}
		case "io1x":
			command_io1x()
		case "io2x":
			command_io2x()
		case "io3x":
			command_io3x()
		case "io4x":
			command_io4x()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}
