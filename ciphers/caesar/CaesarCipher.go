package caesar

/*
	凯撒加密算法：针对因为字母'a'-'z'，'A'-'Z',将字母向由迁移key个位置
	假设字母为a，key=1，则加密转换后为字母'b'，对于其他英文字母就保持原来的值就好
 */
func Encrypt(input string, key int) string {
	// 为了避免key是负数或者key大于26或者小于-26，首先将key对26取模后加26再对26取模
	key = (key%26 + 26) % 26

	var outputBuffer []byte
	for _, r := range input {
		// r就是字母的数字编码形式
		if 'A' <= r && r <= 'Z' {
			outputBuffer = append(outputBuffer, byte('A' + (r - 'A' + int32(key))%26))
		} else if 'a' <= r && r <= 'z' {
			outputBuffer = append(outputBuffer, byte('a' + (r -'a' + int32(key))%26))
		} else {
			outputBuffer = append(outputBuffer, byte(r))
		}
	}
	return string(outputBuffer)
}

// Decrypt 解码就是向左迁移key，等价于向右迁移26-key个位置
func Decrypt(input string, key int) string {
	return Encrypt(input, 26-key)
}
