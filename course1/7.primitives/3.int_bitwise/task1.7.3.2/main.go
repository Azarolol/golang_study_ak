package main

import "fmt"

const (
	Read    = 4
	Write   = 2
	Execute = 1
)

func getFilePermissions(flag int) string {
	output := ""
	for i := 0; i < 3; i++ {
		curFlag := flag % 10
		curState := ""
		if curFlag&Read == Read {
			curState += "Read,"
		} else {
			curState += "-,"
		}
		if curFlag&Write == Write {
			curState += "Write,"
		} else {
			curState += "-,"
		}
		if curFlag&Execute == Execute {
			curState += "Execute"
		} else {
			curState += "-"
		}
		switch i {
		case 0:
			output = " Other:" + curState + output
		case 1:
			output = " Group:" + curState + output
		case 2:
			output = "Owner:" + curState + output
		}
		flag /= 10
	}
	return output
}

func main() {
	fmt.Println(getFilePermissions(755))
	fmt.Println(getFilePermissions(0))
}
