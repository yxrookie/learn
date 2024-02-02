package rollingwindow


type window struct {
	buckets []*Bucket
	size int
}

func Newwindow(size int) *window {
	bucksets := make([]*Bucket, size)
	// 自己做时，没有对 buckets 中的每个 Bucket 进行内存分配
	for i := 0; i < size; i++ {
		bucksets[i] = new(Bucket)
	}
	return &window{
		buckets: bucksets,
		size: size,
	}
} 

func (w *window) Add(v float64, offset int) {
	w.buckets[offset%w.size].Add(v)
}

func (w *window) Reset(offset int) {
	w.buckets[offset%w.size].Reset()
}

// 统计时间窗口内的有效桶中的值
// 因为取余操作形成了一个循环数组，所以只要不是当前窗口失效的桶就都进行统计
func (w *window) Reduce(start, count int, fn func(b *Bucket)) {
	for i := 0; i < count; i ++ {
		fn(w.buckets[(start+i) % w.size])
	}
}

