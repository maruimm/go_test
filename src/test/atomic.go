package main



import (
	"fmt"
	"runtime"
	"sync"
	//"sync/atomic"
	"math"
	"time"
)


func initCounter(lTime, TestTime int64) (map[int64] *uint64, []int64) {
	Counter := make(map[int64] *uint64)
	var timeList []int64
	for  i := lTime; i <= lTime + TestTime; i++ {
		var val = new(uint64)
		Counter[i] = val
		timeList = append(timeList, i)
	}

	return Counter,timeList
}


func main() {

	TestTime := 10

	originMaxProcs := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("originMaxProcs:%d ,MaxUint64 :%d\n", originMaxProcs,uint64(math.MaxUint64))
	Counter,timeList := initCounter(time.Now().Unix(), int64(TestTime))

	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			lTime := time.Now().Unix()
			runtime.LockOSThread()
			for {
				if _, ok := Counter[lTime]; !ok {
					wg.Done()
					break
				}
				//atomic.AddUint64(Counter[lTime], 1)
				*Counter[lTime] += 1
				nTime := time.Now().Unix()
				if nTime > lTime {
					lTime = nTime
					runtime.Gosched()
				}
			}
		}()
	}
	wg.Wait()


	for _, key := range timeList {
		if val, ok := Counter[key]; ok {
			fmt.Printf("key:%d, val:%d\n", key, val)
		}
	}

}