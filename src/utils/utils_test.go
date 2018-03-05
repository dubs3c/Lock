package utils

import (
    "testing"
    "strings"
    "bytes"
)

func TestPadKey(t *testing.T) {
    key := PadKey("secret_key")
    if len(key) != 16 {
        t.Errorf("Length was incorrect, got: %d, expected: %d", len(key), 16)
    }

    if strings.Compare(key, "secret_key000000") != 0 {
        t.Errorf("key does not match, got: %s, expected: secret_key000000", key)
    }

}

func TestToHex(t *testing.T) {
    b := []byte("hello")
    hex := ToHex(b)
    if bytes.Compare(hex, []byte("68656c6c6f")) != 0 {
        t.Errorf("Hex value was incorrect, got: %s, expected: 68656c6c6f", hex)
    }
}