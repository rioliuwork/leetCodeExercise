package exercise

import "bytes"

/**

TinyURL是一种URL简化服务， 比如：当你输入一个URL https://leetcode.com/problems/design-tinyurl 时，它将返回一个简化的URL http://tinyurl.com/4e9iAk.

要求：设计一个 TinyURL 的加密 encode 和解密 decode 的方法。你的加密和解密算法如何设计和运作是没有限制的，你只需要保证一个URL可以被加密成一个TinyURL，并且这个TinyURL可以用解密方法恢复成原本的URL。
*/

type Codec struct {
	m      map[string]string
	number int
}

func Constructor535() Codec {
	return Codec{
		m: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	c.number++
	buffer := bytes.Buffer{}
	for c.number > 0 {
		i := c.number % 62
		buffer.WriteRune(rune(chars[i]))
		c.number /= 62
	}
	shortUrl := buffer.String()
	c.m[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.m[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
