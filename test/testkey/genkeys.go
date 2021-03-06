// +build ignore

package main

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"go/format"
	"math/rand"
	"os"
)

const (
	header = `// THIS FILE IS GENERATED. DO NOT EDIT THIS FILE DIRECTLY.
package testkey

var (
`
	footer = `)
`
)

var (
	// This "random" source returns a fixed sequence. This will keep the test
	// keys stable between invocations to minimize diffs when more keys are
	// generated
	fixedRand *rand.Rand
)

func main() {
	rsa1024 := flag.Int("rsa1024", 0, "Number of rsa1024 keys to generate")
	rsa2048 := flag.Int("rsa2048", 0, "Number of rsa2048 keys to generate")
	rsa4096 := flag.Int("rsa4096", 0, "Number of rsa4096 keys to generate")
	ec256 := flag.Int("ec256", 0, "Number of ec256 keys to generate")
	ec384 := flag.Int("ec384", 0, "Number of ec384 keys to generate")
	seed := flag.Int64("seed", 0, "Random number seed for deterministic key generation")
	flag.Parse()

	fixedRand = rand.New(rand.NewSource(*seed))

	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, header)

	fmt.Fprintln(buf, "rsa1024Keys = []string{")
	for i := 0; i < *rsa1024; i++ {
		fmt.Fprintf(buf, "`%s`,\n", toPEM(genRSA(1024)))
	}
	fmt.Fprintln(buf, "}")

	fmt.Fprintln(buf, "rsa2048Keys = []string{")
	for i := 0; i < *rsa2048; i++ {
		fmt.Fprintf(buf, "`%s`,\n", toPEM(genRSA(2048)))
	}
	fmt.Fprintln(buf, "}")

	fmt.Fprintln(buf, "rsa4096Keys = []string{")
	for i := 0; i < *rsa4096; i++ {
		fmt.Fprintf(buf, "`%s`,\n", toPEM(genRSA(4096)))
	}
	fmt.Fprintln(buf, "}")

	fmt.Fprintln(buf, "ec256Keys = []string{")
	for i := 0; i < *ec256; i++ {
		fmt.Fprintf(buf, "`%s`,\n", toPEM(genEC(elliptic.P256())))
	}
	fmt.Fprintln(buf, "}")

	fmt.Fprintln(buf, "ec384Keys = []string{")
	for i := 0; i < *ec384; i++ {
		fmt.Fprintf(buf, "`%s`,\n", toPEM(genEC(elliptic.P384())))
	}
	fmt.Fprintln(buf, "}")

	fmt.Fprintln(buf, footer)

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stderr.Write(buf.Bytes())
		panic(err)
	}
	_, err = os.Stdout.Write(formatted)
	check(err)
}

func genRSA(bits int) *rsa.PrivateKey {
	key, err := rsa.GenerateKey(fixedRand, bits)
	check(err)
	return key
}

func genEC(curve elliptic.Curve) *ecdsa.PrivateKey {
	key, err := ecdsa.GenerateKey(curve, fixedRand)
	check(err)
	return key
}

func toPEM(key crypto.PrivateKey) string {
	data, err := x509.MarshalPKCS8PrivateKey(key)
	check(err)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: data,
	}))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
