package methods

import (
	"fmt"
	"reflect"
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

var minorBytes = makeBytes(1, 0)

type Data struct {
	Host string

	bytes      []byte
	minorBytes []byte
}

func (data *Data) supplyBytes(bytes []byte) {
	data.bytes = bytes
}
