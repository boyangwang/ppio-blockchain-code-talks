// 请补全下面代码中 TODO 的部分，使得根据私钥可以计算得到公钥，且根据公钥可以计算其ETH地址

package ecc

import (
	"crypto/elliptic"
	"crypto/rand"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
)

// 椭圆曲线的接口定义在golang标准库"crypto/elliptic"中
/*
type Curve interface {
	Params() *CurveParams                        // 获取椭圆曲线的参数
	IsOnCurve(x, y *big.Int) bool                // 判断一个已知的点在不在这条椭圆曲线上
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) // 计算椭圆曲线上的点的加法 x + y
	Double(x1, y1 *big.Int) (x, y *big.Int)      // 计算椭圆曲线上的整数2与点x的点积 2x
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) // 计算椭圆曲线上的整数k与点x的点积
	ScalarBaseMult(k []byte) (x, y *big.Int)     // 计算椭圆曲线上整数k与基点g的点积
}
type CurveParams struct {
	P       *big.Int // y² mod p = (x³ + ax + b) mod p 中的质数 p
	N       *big.Int // 基点G的阶
	B       *big.Int // y² mod p = (x³ + ax + b) mod p 中的 b
	Gx, Gy  *big.Int // 基点G的坐标(gx, gy)
	BitSize int      // 横纵坐标的位数
	Name    string   // 椭圆曲线的名字
}
*/

// 我们使用跟比特币和以太坊一致的曲线 secp256k1
var TheCurve = secp256k1.S256()

// See SEC 2 section 2.7.1
// curve parameters taken from:
// http://www.secg.org/sec2-v2.pdf
// theCurve.P, _ = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 0)
// theCurve.N, _ = new(big.Int).SetString("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 0)
// theCurve.B, _ = new(big.Int).SetString("0x0000000000000000000000000000000000000000000000000000000000000007", 0)
// theCurve.Gx, _ = new(big.Int).SetString("0x79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 0)
// theCurve.Gy, _ = new(big.Int).SetString("0x483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 0)
// theCurve.BitSize = 256

// Keccak256 hash函数
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Sha3是Keccak256的别名
func Sha3(data ...[]byte) []byte {
	return Keccak256(data...)
}

// 核心数据结构的定义及其长度
const HASH_LENGTH = 32
const ADDRESS_LENGTH = 20
const PUBLIC_KEY_LENGTH = 65
const PRIVATE_KEY_LENGTH = 32
const SIGNATURE_LENGTH = 65

type Hash [HASH_LENGTH]byte
type Address [ADDRESS_LENGTH]byte
type PublicKey [PUBLIC_KEY_LENGTH]byte
type PrivateKey [PRIVATE_KEY_LENGTH]byte
type Signature [SIGNATURE_LENGTH]byte

// 生成新的私钥
func (sk *PrivateKey) Generate() error {
	// 获取椭圆曲线和随机数读取器
	var c elliptic.Curve = TheCurve
	var r io.Reader = rand.Reader

	// 调用golang标准库里的GenerateKey()随机生成私钥
	b, _, _, err := elliptic.GenerateKey(c, r)
	if err != nil {
		return err
	}
	// 确保私钥长度正确
	if len(b) != PRIVATE_KEY_LENGTH {
		panic("assert: random field element should be 32 bytes")
	}
	copy(sk[:], b)
	return nil
}

// 根据私钥计算公钥
func (sk *PrivateKey) PublicKey() PublicKey {
	// 获取椭圆曲线
	var c elliptic.Curve = TheCurve
	pk := PublicKey{}

	// TODO: 已知私钥pk，计算其在椭圆曲线上的点(x, y)

	// 将公钥(x, y)序列化成数组放到pk中
	// 注意：此时pk是一个长度为65的byte数组，第一个byte为公钥类型：04
	// 接下来的32byte为横坐标值x，最后的32byte为纵坐标y
	copy(pk[:], elliptic.Marshal(c, x, y))
	return pk
}

// 根据公钥计算地址
func (pk *PublicKey) Address() Address {
	addr := Address{}

	// 由于这里序列化之后的公钥是65字节，第一个字节是表示公钥的格式，后面的2个32字节才分别是
	// 公钥的横坐标x和纵坐标y，因此这里先把第一个字节剔除。
	tmp := pk[1:]

	// 计算公钥的hash
	// TODO: 采用Keccak256()函数，计算公钥的hash。
	//       Keccak256()的参数为64字节，前32字节为公钥的横坐标x，后32字节为公钥的纵坐标y

	// 根据以太坊协议，公钥hash的后20个字节为地址
	copy(addr[:], hash[12:])
	return addr
}

// 根据公钥获取其横坐标和纵坐标
func (pk *PublicKey) Point() (*big.Int, *big.Int) {
	x, y := elliptic.Unmarshal(TheCurve, pk[:])
	return x, y
}
