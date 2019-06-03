# 椭圆曲线练习题

已知以太坊的私钥sk，计算其公钥pk和地址address

1. **下载**

```
$ git clone https://github.com/PPIO/ppio-blockchain-code-talks
$ cd ppio-blockchain-code-talks/ecc
```

2. **修改**

修改 `ecc.go` 中的 `TODO` 部分，使代码逻辑完整。

```
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
```

```
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
```

3. **测试**

```
$ go test
```

验证结果是否正确，可将测试私钥导入以太坊钱包，然后查看地址与代码计算出来的是否一致。