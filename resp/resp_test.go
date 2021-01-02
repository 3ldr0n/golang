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
