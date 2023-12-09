package combinedapplication

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	newCache := initLruCache(4)
	printCache(newCache)
    newCache.put(1, "apple")
	newCache.put(2, "huawei")
	newCache.put(3, "xiaomi")
	fmt.Println(newCache.get(2))
	fmt.Println(newCache.get(3))
	newCache.put(4, "meizu")
	printCache(newCache)
	newCache.put(5, "oppo")
	printCache(newCache)
	newCache.remove(2)
	printCache(newCache)
	newCache.remove(5)
	printCache(newCache)
}

func printCache(cache *lruCache) {
	fmt.Print("(")
	for key, data := range cache.nodeMap {
		fmt.Print(key ,":", data.value + " ")
	}
	fmt.Println(")")
}
