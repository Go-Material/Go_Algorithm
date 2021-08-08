/*
	Diffie-Hellman密钥交换协议/算法，确保共享KEY安全穿越不安全网络的方法。
	可以用这个方法确定对称密钥. 然后可以用这个密钥进行加密和解密，但是不能用于数据的加密和解密
 */
package diffiehelkeyexchange

const (
	generator         = 3 //生成器
	primeNumber int64 = 6700417 // 素数模
)

/*
	通过私钥、生成器、素数模生成公钥
	shareKey = (generator ^ key) % primeNumber
 */

func GenerateShareKey(prvKey int64) int64 {
	return modularExponentiation(generator, prvKey, primeNumber)
}

/*
	加密密钥生成 : 只有协商双方才知道密钥是什么，窃听者无法得知密钥。双方生成的密钥一定是一致的
	假设Bob、Slice的私钥是p1，p2, 他们生成的公钥分别为ret1 = (3^p1) % primeNumber, ret2 = (3^p2) % primeNumber
	Bob生成的密钥为m1 = (ret2^p1)%primeNumber = (((3^p2) % primeNumber)^p1)%primeNumber = (3^p2^p1)%primeNumber
    Slice生成的密钥为m2 = (ret1^p2)%primeNumber = (((3^p1) % primeNumber)^p2)%primeNumber = (3^p1^p2)%primeNumber
	所以他们两个人生成的密钥是一致的
	mutualKey = (shareKey ^ prvKey ) % primeNumber
 */
func GenerateMutualKey(prvKey, shareKey int64) int64 {
	return modularExponentiation(shareKey, prvKey, primeNumber)
}

/*
	计算思路
	快速幂算法
	ab mod c = ((a mod c) * (b mod c)) mod c
 */
func modularExponentiation(b, e, mod int64) int64 {
	if mod == 1 {
		return 0
	}
	var r int64 = 1
	b = b % mod
	for e > 0 {
		if e&1 == 1 {
			r = (r * b) % mod
		}
		e = e >> 1 // 等价于e /= 2
		b = (b * b) % mod
	}
	return r
}
