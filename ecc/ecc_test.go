package ecc

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// 「随机」产生一个私钥，然后计算其公钥和地址
func Test_CalculateAddressByRandomPrivateKey(t *testing.T) {
	fmt.Println("Calculating the address of a random private key")
	sk := PrivateKey{}
	sk.Generate()
	fmt.Println("private", hex.EncodeToString(sk[:]))
	pk := sk.PublicKey()
	fmt.Println("public ", hex.EncodeToString(pk[:]))
	x, y := pk.Point()
	fmt.Println("point  ", fmt.Sprintf("(0x%s,0x%s)", hex.EncodeToString(x.Bytes()), hex.EncodeToString(y.Bytes())))
	addr := pk.Address()
	fmt.Println("address", "0x"+hex.EncodeToString(addr[:]))
	fmt.Println()
}

// 「给定」已知的私钥，然后计算其公钥和地址
func Test_CalculateAddressByFixedPrivateKey(t *testing.T) {
	key := "44a11d14bc8a27714ce1ab5238bf9ab08e72f5b463b6a4f0157463c8993e34ea"
	fmt.Println("Calculating the address of", key)
	sk := PrivateKey{}
	b, _ := hex.DecodeString(key)
	copy(sk[:], b)
	fmt.Println("private", hex.EncodeToString(sk[:]))
	pk := sk.PublicKey()
	fmt.Println("public ", hex.EncodeToString(pk[:]))
	x, y := pk.Point()
	fmt.Println("point  ", fmt.Sprintf("(0x%s,0x%s)", hex.EncodeToString(x.Bytes()), hex.EncodeToString(y.Bytes())))
	addr := pk.Address()
	fmt.Println("address", "0x"+hex.EncodeToString(addr[:]))
}
