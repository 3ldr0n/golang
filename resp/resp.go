package resp

import (
	"fmt"
	"strconv"
)

// https://redis.io/topics/protocol

const (
	SimpleString = '+'
	Error        = '-'
	Integer      = ':'
	BulkString   = '$'
	Array        = '*'
	CR           = '\r'
	LF           = '\n'
	CRLF         = CR + LF
)

func ReadRESP(chunk []byte) int {
	if len(chunk) == 0 {
		return 0
	}
	// First byte is the data type.
	dataType := chunk[0]
	// From the second byte up is the data, minus the last two ones
	// (which are CRLF).
	data := chunk[1 : len(chunk)-2]

	if dataType == SimpleString {
		fmt.Println("Is string")
		fmt.Println(string(data))
	} else if dataType == Error {
		fmt.Println("Is Error")
		fmt.Println(string(data))
	} else if dataType == Integer {
		fmt.Println("Is Integer")
	} else if dataType == BulkString {
		fmt.Println("Is Bulk String")
		fmt.Println(string(data))
	} else if dataType == Array {
		fmt.Println("Is Array")
		// Read until the first CRLF, until the first CRLF appears the
		// array size should be sent.
		size := readUntilCRLF(data)
		fmt.Println("Array size", size)
	}

	return len(data)
}

func readUntilCRLF(data []byte) int {
	if len(data) <= 2 {
		return 0
	}

	i := 0
	for ; i < len(data)-1; i++ {
		if data[i] == CR {
			if data[i+1] == LF {
				break
			}
		}
	}
	integer, _ := strconv.Atoi(string(data[:i]))
	return integer
}
