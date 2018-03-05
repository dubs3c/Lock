package crypto

import (
    "io"
    "io/ioutil"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "strings"
    "lock/src/utils"
)

func Encrypt(filename string, secret string) string {
    paddedKey := utils.PadKey(secret)
    key := []byte(paddedKey)
    plaintext, err := ioutil.ReadFile(filename)
    utils.Check(err)

    block, err := aes.NewCipher(key)
    utils.Check(err)

    // Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err.Error())
    }

    aesgcm, err := cipher.NewGCM(block)
    utils.Check(err)

    ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
    
    nonceHex := utils.ToHex(nonce)
    ciphertextHex := utils.ToHex(ciphertext)

    ciphertextString := string(nonceHex[:])+"::"+string(ciphertextHex[:])
    return string(ciphertextString)
}

func Decrypt(filename string, secret string) string {
    paddedKey := utils.PadKey(secret)
    key := []byte(paddedKey)
    ciphertextHex, err := ioutil.ReadFile(filename)
    utils.Check(err)

    ciphertextString := string(ciphertextHex[:])
    split := strings.Split(ciphertextString,"::")
    nonce, _ := hex.DecodeString(split[0])
    ciphertext, _ := hex.DecodeString(split[1])

    block, err := aes.NewCipher(key)
    utils.Check(err)

    aesgcm, err := cipher.NewGCM(block)
    utils.Check(err)

    plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
    utils.Check(err)

    return string(plaintext[:])
}
