package bench

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	_lastGC     uint64
	_lastGCLock sync.Mutex
)

func Debug() {
	_lastGCLock.Lock()
	defer _lastGCLock.Unlock()

	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)

	ts := time.Now().UnixNano() - int64(_lastGC)
	fmt.Printf("allocated: %.2f mb, lastGC: %.2f ms", (float64(ms.Alloc)/1024.0)/1024.0, float64(ts)/1.0e6)
	if _lastGC != ms.LastGC {
		_lastGC = ms.LastGC
		println("GC!!!!")
	}
}
