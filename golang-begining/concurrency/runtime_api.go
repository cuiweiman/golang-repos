package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/* runtime 包中定义的 协程相关的 API */
func RuntimePk() {
	// fmt.Println("------------------ Gosched 让出CPU时间片------------------")
	// goschedDemo()
	// fmt.Println("------------------ Goexit 当前协程直接退出------------------")
	// goExitDemo()
	// fmt.Println("------------------ GOMAXPROCS 设置CPU最大核心数:默认使用当前资源的最大核心数 ------------------")
	// goMaxprocsDemo()
	// fmt.Println("------------------ mutex 协程加互斥锁 ------------------")
	// mutexDemo()
	fmt.Println("------------------ 并发 原子变量 的引入 ------------------")
	atomicDemo()
}

/* 让出CPU执行时间片 */
func goschedDemo() {
	go goschedDemoShowMsg("java")
	for i := 0; i < 2; i++ {
		// 通过 Gosched 可以观察 协程方法的执行和 主协程的执行
		// 让出 CPU 时间片，使其可以执行 协程(子线程)
		runtime.Gosched()
		fmt.Println(runtime.NumCPU(), " go sched")
	}
}

func goschedDemoShowMsg(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Println(msg)
	}
}

/* 直接退出当前协程 */
func goExitDemo() {
	go doGoExitDemo()
	time.Sleep(time.Second * 1)
}

func doGoExitDemo() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			runtime.Goexit()
		}
		fmt.Println("i = ", i)
	}
}

/* 限制CPU核心数，默认使用当前资源的最大核心数， */
func goMaxprocsDemo() {
	fmt.Println("当前CPU核心数: ", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go doGoMaxprocsDemo("CWM")
	go doGoMaxprocsDemo("SYY")
	go doGoMaxprocsDemo("CS")
	time.Sleep(time.Second * 3)
}

func doGoMaxprocsDemo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ": ", i)
	}
}

/* 互斥锁 */
var mutexWp sync.WaitGroup
var mutexLock sync.Mutex
var mutexSum int = 0

func mutexDemo() {
	for i := 0; i < 100; i++ {
		go doMutexDemoA()
		mutexWp.Add(1)
		go doMutexDemoB()
		mutexWp.Add(1)
	}
	mutexWp.Wait()
	fmt.Println("mutexSum = ", mutexSum)
}

func doMutexDemoA() {
	defer mutexWp.Done()
	mutexLock.Lock()
	time.Sleep(time.Millisecond * 2)
	mutexSum++
	mutexLock.Unlock()
}

func doMutexDemoB() {
	defer mutexWp.Done()
	mutexLock.Lock()
	time.Sleep(time.Millisecond * 2)
	mutexSum--
	mutexLock.Unlock()
}

var atomicI int32 = 0

/* 原子操作: CAS 原理 */
func atomicDemo() {
	for i := 0; i < 100; i++ {
		go addAtomic()
		go subAtomic()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("atomicI: %v\n", atomicI)
	loadAtomic()
	casLoad()
}

/* 原子增减操作 */
func addAtomic() {
	// 原子操作: CAS 原理
	atomic.AddInt32(&atomicI, 1)
}

func subAtomic() {
	atomic.AddInt32(&atomicI, -1)
}

/* 原子载入操作: 即读取出 原子的value值 */
func loadAtomic() {
	/* 原子操作 CAS 赋值 */
	atomic.StoreInt32(&atomicI, 89)
	fmt.Printf("atomic.LoadInt32(&atomicI): %v\n", atomic.LoadInt32(&atomicI))
}

/* 原子CAS 操作: 是否允许 CAS赋值。旧变量,期望值,新值 */
func casLoad() {
	var atomicK int32 = 59
	var cas bool = atomic.CompareAndSwapInt32(&atomicK, 59, 0)
	fmt.Printf("cas: %v\n", cas)
}
