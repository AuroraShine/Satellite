package unpack

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
)

func Unpack(src string, dest string) (err error) {
	// first, open the file
	file, err := os.Open(src)
	if err != nil {
		log.Println("Error open file:", err)
		return err
	}
	// second, read file data
	buf := make([]byte, 60)
	rd := bufio.NewReader(file)
	_, err = rd.Read(buf)
	if err != nil {
		log.Println("Error read file:", err)
		return err
	}
	// third, close the file
	err = file.Close()
	if err != nil {
		log.Println("Error close file:", err)
		return err
	}
	// fourth, find the algorithm
	buf = buf[48:56]
	index := bytes.IndexByte(buf, 0)
	tp := string(buf[0:index])
	switch tp {
	case "AES", "aes":
		err = UnpackAES(src, dest)
	case "DES", "des":
		err = UnpackDES(src, dest)
	case "3DES", "3des":
		err = Unpack3DES(src, dest)
	case "RSA", "rsa":
		err = UnpackRSA(src, dest)
	case "BASE64", "base64":
		err = UnpackBase64(src, dest)
	default:
		s := fmt.Sprint("Undefined unpack algorithm.")
		err = errors.New(s)
	}
	return err
}
