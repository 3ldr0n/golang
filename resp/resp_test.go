package resp

import (
	"fmt"
	"testing"
)

func TestRESPString(t *testing.T) {
	command := []byte("+OK\r\n")

	size := ReadRESP(command)

	fmt.Println(size)
}

func TestRESPError(t *testing.T) {
	command := []byte("-ERROR an error ocurred\r\n")

	size := ReadRESP(command)

	fmt.Println(size)
}

func TestRESPEmptyArray(t *testing.T) {
	command := []byte("*0\r\n")

	size := ReadRESP(command)

	fmt.Println(size)
}

func TestRESPNonEmptyArray(t *testing.T) {
	command := []byte("*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n")

	size := ReadRESP(command)

	fmt.Println(size)
}
