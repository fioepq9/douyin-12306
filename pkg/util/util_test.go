package util

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestInt64ToString(t *testing.T) {
	var (
		num        int64
		realAnswer string
		myAnswer   string
	)
	check := func() {
		realAnswer = fmt.Sprintf("%d", num)
		myAnswer = Int64ToString(num)
		if realAnswer != myAnswer {
			t.Fatalf("num = %d, real answer = %s, my answer = %s", num, realAnswer, myAnswer)
		}
	}
	// 边界测试
	nums := []int64{0, -math.MaxInt64, math.MaxInt64}
	for i := 0; i < len(nums); i++ {
		num = nums[i]
		check()
	}
	// 随机测试
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1e6; i++ {
		num = rand.Int63()
		check()
	}
}
