//A hash function takes a set of data and reduces it to a smaller fixed
//size. Hashes are frequently used in programming for everything from
//looking up data to easily detecting changes. Hash functions in Go are
//broken into two categories: cryptographic and non-cryptographic.

package main

//The non-cryptographic hash functions can be found underneath the hash package and include adler32, crc32, crc64 and fnv.
import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("sonerdm"))
	v := h.Sum32()
	fmt.Println(v)

	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
