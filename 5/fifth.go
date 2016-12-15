package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "ffykfhsq"
	var password string

	for i := 0; len(password) < 8; i++ {
		hash := md5.Sum([]byte(input + strconv.Itoa(i)))
		res := hex.EncodeToString([]byte(hash[:]))
		if strings.HasPrefix(string(res), "00000") {
			password += res[5:6]
		}
	}
	fmt.Printf("Password %s", password)
}
