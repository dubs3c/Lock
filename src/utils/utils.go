package utils

import (
    "os"
    "log"
    "encoding/hex"
    "strings"
)

func GetFiles(dir string) []string {
    x, err := os.Open(dir)
    Check(err)
    files, err := x.Readdir(0)
    Check(err)
    var listWithFiles []string
    for _, file := range files {
        info, err := os.Stat(file.Name())
        Check(err)
        if info.IsDir() {
            continue
        }
        if "./"+file.Name() == os.Args[0] {
            continue
        }
        listWithFiles = append(listWithFiles, file.Name())
    }
    return listWithFiles
}

func Check(e error) {
    if e != nil {
        log.Fatal("ERROR:", e)
    }
}

func PanicErr(e error) {
    if e != nil {
        panic(e)
    }
}

func ToHex(src []byte) []byte {
    dst := make([]byte, hex.EncodedLen(len(src)))
    hex.Encode(dst, src)
    return dst
}

func FromHex(src []byte) []byte {
    dst := make([]byte, hex.DecodedLen(len(src)))
    _, err := hex.Decode(dst, src)
    Check(err)
    return dst
}

func PadKey(key string) string {
    var keyLen int = len(key) % 16
    if keyLen != 0 {
        padding := strings.Repeat("0", 16 - keyLen)
        return key+padding
    } else {
        return key
    }
}