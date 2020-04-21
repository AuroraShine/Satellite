package decomp

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// DeCompressZip function
// input src file list, output dest file path and algorithm which used in unpack, return error info
// this function will base on algorithm to call correspond function
// src file support both absolute and relative paths, like 'C:\\file.tar.zip' or '../test/data/file.tar.zip'
// dest file also support both absolute and relative paths, like 'C:\\' or '../test/data/'
// return err indicate the success or failure function execute
func DeCompressZip(src string, dest string) (err error) {
	// open the zip reader...
	reader, err := zip.OpenReader(src)
	if err != nil {
		log.Println("Error open zip reader:", err)
		return err
	}
	defer reader.Close()
	// loop decompress src list files
	for _, file := range reader.File {
		path := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Println("Error make dir all:", err)
				return err
			}
		} else {
			// make dir all path...
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				log.Println("Error make dir all:", err)
				return err
			}
			// open the in file
			in, err := file.Open()
			if err != nil {
				log.Println("Error open the in file:", err)
				return err
			}
			defer in.Close()
			// open the out file
			out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				log.Println("Error open the out file:", err)
				return err
			}
			defer out.Close()
			// write decompress data into file
			_, err = io.Copy(out, in)
			if err != nil {
				log.Println("Error write decompress date:", err)
				return err
			}
		}
	}
	return err
}
