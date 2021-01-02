package resp

import (
	"fmt"
)

// https://redis.io/topics/protocol

const (
	SimpleString = '+'
	Error        = '-'
	Integer      = ':'
	BulkString   = '$'
	Array        = '*'
	CRLF         = "\r\n"
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
	}

	return len(data)
}
