package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	secretFlag := flag.String("secret", "", "secret message to hide (or set SECRET env)")
	keyFlag := flag.String("key", "", "key for XOR cipher (or set KEY env)")
	rawFlag := flag.Bool("raw", false, "print raw bytes instead of base64")
	flag.Parse()

	secret := *secretFlag
	if secret == "" {
		secret = os.Getenv("SECRET")
	}
	key := *keyFlag
	if key == "" {
		key = os.Getenv("KEY")
	}
	if secret == "" || key == "" {
		fmt.Println("Usage: go run . -secret=\"message\" -key=\"mykey\"")
		fmt.Println("Or set env vars: SECRET and KEY")
		return
	}

	out := cipher(secret, []byte(key))

	fmt.Println("--- GARBAGE STREAM ---")
	if *rawFlag {
		fmt.Println(string(out))
	} else {
		fmt.Println(base64.StdEncoding.EncodeToString(out))
	}
}
