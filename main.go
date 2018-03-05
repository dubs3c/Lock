package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "lock/src/crypto"
    "lock/src/utils"
    "lock/src/zip"
)

func main() {
    validateArgs()
    if os.Args[1] == "-e" || os.Args[1] == "--encrypt" {
        performEncryption(os.Args[2])
    } else if os.Args[1] == "-d" || os.Args[1] == "--decrypt" {
        performDecryption(os.Args[2])
    }
}

func validateArgs() {
    if len(os.Args) > 3 {
        fmt.Println("Too many arguments, see usage for help.")
        os.Exit(1)
    } else if len(os.Args) == 2 {
        fmt.Println("Missing an argument, see usage for help")
        os.Exit(1)
    } else if len(os.Args) == 1 {
        usage()
        os.Exit(1)
    } 
}

func usage() {
    fmt.Println(`Place this program in a folder and run it. It will zip all files and encrypt the zip file.

Usage:
    ` + os.Args[0] + ` [-d|--decrypt] KEY
    ` + os.Args[0] + ` [-e|--encrypt] KEY`)
}

func performDecryption(key string) {
    plaintext := crypto.Decrypt("seal.enc", key)
    fmt.Println("[+] Decrypted file...")
    err := ioutil.WriteFile("out.zip", []byte(plaintext), 0644)
    utils.Check(err)
    zip.Unzip("out.zip", ".")
    os.Remove("seal.enc")
    os.Remove("out.zip")
    fmt.Println("[+] Unziped file...")
}

func performEncryption(key string) {
    filesToArchive := utils.GetFiles(".")
    if len(filesToArchive) == 0 {
        fmt.Println("[-] No files to archive...")
        os.Exit(1)
    }
    err := zip.ZipFiles(filesToArchive)
    utils.Check(err)

    fmt.Println("[+] Zip archive has been created.")
    ciphertext := crypto.Encrypt("out.zip", key)
    err = ioutil.WriteFile("seal.enc", []byte(ciphertext), 0644)
    utils.Check(err)
    fmt.Println("[+] Encrypted zip archive has been created...")
    err = os.Remove("out.zip")
    utils.Check(err)

    for _, file := range filesToArchive {
        os.Remove(file)
    }
}









