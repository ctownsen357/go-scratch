# go-scratch
Scratch utilities and how to items, examples, in Go
### binaryReadWrite.go
Very simple example of read & write of packed binary data.  I implemented something like this when replacing a Python process that generated large binary files to feed some storm simulation routines.  
 
### encrypt.go 
An AES file encryption routine in Go designed to encrypt in chunks so it can handle large files without exceeding memory and to be compatible with existing Python and C implementation of the Encrypt/Decrypt routines as coded in enc.py
