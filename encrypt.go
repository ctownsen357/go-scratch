// encrypt.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"io"
	"os"
)

//Encrypt is intended to be compatible with some existing Python and C processes
//that had already implemented the AES file encryption this way: by reading the source file
//by chunks and encrypting the individual chunks - the other processes did not or at least don't
//implement any sort of streaming so this one doesn't either.  One added benefit is that in this
//implementation you aren't opening the entire file at one time - my latest projects have dealt with
//some very large files so this is important.
func Encrypt(password string, pathToInFile string, pathToOutFile string) error {
	keySize := 32 //this was a default/optional parameter in the original Python code
	bs := int(aes.BlockSize)
	salt := make([]byte, aes.BlockSize-len("Salted__"))
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	key, iv := deriveKeyAndIV([]byte(password), salt, bs, keySize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)

	fin, err := os.Open(pathToInFile)
	defer fin.Close()
	if err != nil {
		return err
	}

	fout, err := os.Create(pathToOutFile)
	defer fout.Close()
	if err != nil {
		return err
	}

	_, err = fout.Write([]byte("Salted__"))
	if err != nil {
		return err
	}
	_, err = fout.Write(salt)
	if err != nil {
		return err
	}

	for true {
		chunk := make([]byte, 1024*bs)

		n, err := fin.Read(chunk)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if n == 0 || n%bs != 0 { //need to pad up to block size :bs
			paddingLength := (bs - n%bs)
			paddingChr := []byte(string(rune(paddingLength)))
			paddingBytes := make([]byte, paddingLength)

			for i := 0; i < paddingLength; i++ {
				paddingBytes[i] = paddingChr[0]
			}

			chunk = append(chunk[0:n], []byte(paddingBytes)...)
		} else {
			chunk = chunk[0:n] //no padding needed
		}

		ciphertext := make([]byte, len(chunk))
		cbc.CryptBlocks(ciphertext, chunk)
		fout.Write(ciphertext)
	}
	return nil
}

//There are certainly alternative and more standard ways of specifying the salt, deriving a key from a password
//and creating an initialization vector but this was designed to match an existing implementation
func deriveKeyAndIV(password []byte, salt []byte, bs int, keySize int) ([]byte, []byte) {
	var digest []byte
	hash := md5.New()
	out := make([]byte, keySize+bs)
	for i := 0; i < keySize+bs; i += len(digest) {
		hash.Reset()
		hash.Write(digest)
		hash.Write(password)
		hash.Write(salt)
		digest = hash.Sum(digest[:0])
		copy(out[i:], digest)
	}
	return out[:keySize], out[keySize : keySize+bs]
}

func main() {
	pwd := "testing123"
	
	err := Encrypt(pwd, "tst.bin", "enctst.bin")
	
	if err != nil{
		panic(err)
	}
}
