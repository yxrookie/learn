package go_consistenthash

import (
	"fmt"
	"testing"
)

func TestNewConsistentHashBalance(t *testing.T) {
	rb := NewHashBanlance(10, nil)
	rb.Add("127.0.0.1:2003")  // 0
	rb.countNode()
	rb.Add("127.0.0.1:2004")
	rb.countNode() 
	rb.Add("127.0.0.1:2005")
	rb.countNode()
	rb.Add("127.0.0.1:2006")
	rb.countNode()
	rb.Add("127.0.0.1:2007")
	rb.countNode()

	// url hash
	fmt.Println(rb.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(rb.Get("http://127.0.0.1:2002/base/error"))
	fmt.Println(rb.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(rb.Get("http://127.0.0.1:2002/base/changepwd"))
    
	// ip hash
	fmt.Println(rb.Get("127.0.0.1"))
	fmt.Println(rb.Get("192.168.0.1"))
	fmt.Println(rb.Get("127.0.0.1"))
    
	// 
	rb.Remove("127.0.0.1:2003")
	rb.countNode()
	rb.Remove("127.0.0.1:2006")
	rb.countNode()
}