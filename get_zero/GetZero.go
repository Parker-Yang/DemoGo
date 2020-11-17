package get_zero

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
)

func GetZero() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	r := make(chan string)
	d := make(chan []byte)
	for i := 0; i < numCPU; i++ {
		j := i
		go func() {
			for i := j; ; i += numCPU {
				data := []byte(fmt.Sprintf("srv_trust-%v", i))
				hash := sha256.Sum256(data)
				result := hex.EncodeToString(hash[:])
				if result[:6] == strings.Repeat("0", 6) {
					d <- data
					r <- result
					break
				}
			}
		}()
	}
	data := <-d
	result := <-r
	fmt.Println(string(data), result)
}
func GetZeroTwo() {
	var(
		data []byte
		result string
	)
	for i := 0; ; i++ {
		data = []byte(fmt.Sprintf("yangcan-%v", i))
		hash := sha256.Sum256(data)
		result = hex.EncodeToString(hash[:])
		if result[:6] == strings.Repeat("0", 6) {
			break
		}
	}
}
