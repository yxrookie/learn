package rollingwindow


type Bucket struct {
	Sum float64
	Count int64
}

func (b *Bucket) Add(v float64) {
	b.Sum += v
	b.Count ++
}

func (b *Bucket) Reset() {
	b.Sum = 0
	b.Count = 0
}