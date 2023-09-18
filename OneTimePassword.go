package algorithms

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"time"
)

func Generator() {
	secret := "secret"
	hash := sha1.New()
	buf := new(bytes.Buffer)
	buf.WriteString(secret)
	timeDynamic := time.Now().UnixNano()
	s := strconv.FormatInt(timeDynamic, 10)
	buf.WriteString(s)
	_, err := hash.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	hashInBytes := hash.Sum(nil)
	passHex := hex.EncodeToString(hashInBytes)

	offset := getLastByte(passHex)

	fmt.Println("offset", offset)
	trunc := hashInBytes[offset : offset+4]

	fmt.Printf("% x", trunc)
	fmt.Println()
	dec, err := strconv.ParseInt(hex.EncodeToString(trunc), 16, 64)
	if err != nil {
		fmt.Println("parse error", err)
	}
	otp := dec % int64(math.Pow10(6))
	fmt.Println("Ma otp", otp)
}

func getLastByte(passHex string) int {
	a := passHex[len(passHex)-1]
	offset, _ := strconv.Atoi(string(a))
	return offset
}
