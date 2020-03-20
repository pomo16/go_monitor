package base

import (
	"math/rand"
	"sync"
)

type SafeRand struct {
	mu  *sync.Mutex
	rng *rand.Rand
}

func NewSafeRand(seed int64) *SafeRand {
	return &SafeRand{
		mu:  &sync.Mutex{},
		rng: rand.New(rand.NewSource(seed)),
	}
}

func (r *SafeRand) Float32() float32 {
	r.mu.Lock()
	i := r.rng.Float32()
	r.mu.Unlock()
	return i
}

func RandK(k int32) int32 {
	return int32(rand.New(rand.NewSource(2147483647)).Float32() * float32(k))
}

//FixedRandK 生成0~K之间的随机数，固定随机数发生器，使用封装的线程安全的随机发生器
func FixedRandK(k int32, rng *SafeRand) int32 {
	return int32(rng.Float32() * float32(k))
}
