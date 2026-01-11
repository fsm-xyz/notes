# Go

## GMP 模型

有栈协程，其他语言的async/await是无栈协程

无栈协程是一个可以暂停和恢复的函数，利用promise和future特性进行唤醒和睡眠

G（Goroutine）：轻量级协程，初始栈 2KB；

M（Machine）：OS 线程，由内核调度；

P（Processor）：调度器，绑定 M 并管理本地 G 队列。

调度流程：G 优先入 P 本地队列 → 全局队列 → 从其他 P 窃取（Work Stealing）

## Channel 与同步

```go
type hchan struct {
	qcount   uint           // 当前在队列中的元素数量
	dataqsiz uint           // 环形队列的总容量
	buf      unsafe.Pointer // 环形队列的指针
	elemsize uint16         // 元素的大小
	closed   uint32         // 标记 channel 是否关闭
	elemtype *_type         // 元素类型

	sendx    uint   // 发送操作的索引
	recvx    uint   // 接收操作的索引

	recvq    waitq  // 等待接收的 goroutine 队列
	sendq    waitq  // 等待发送的 goroutine 队列

	lock mutex // 保护 channel 的互斥锁
}

// 互斥锁：保证并发访问的安全性。

// 环形队列：提供有缓冲 channel 的数据存储能力，实现“先进先出”的特性。

// 双向链表：管理等待在 channel 上的 goroutine，实现了 goroutine 的同步和调度。
```

无缓冲 Channel：同步通信（发送阻塞直到接收）

有缓冲 Channel：异步队列（阻塞仅当缓冲区满）

Select：多路复用，随机执行就绪 Case，支持超时控制

向已关闭 Channel 发送数据触发 panic，接收返回零值

重复关闭也会panic

Goroutine 泄露：未设退出条件的 Channel 阻塞导致

```sh
func main() {
	a := make(chan int)
    go fmt.Println(<-a) // a实时求值导致死锁，此时a的接收方没有就绪, 会deadlock
    a <- 5  
}

func main() {
	a := make(chan int)
	go func() {
		a <- 5
	}()
	fmt.Println(<-a)
}
```

## 同步原语

sync.WaitGroup：协程等待（Add/Done/Wait）

sync.Mutex vs RWMutex：互斥锁 vs 读写锁（读多写少场景）

sync.Map：并发安全 Map（适合读多写少）

sync.Once

sync.Cond

---
在 Go 语言中，`sync` 包里的 `Mutex` 和 `RWMutex` 是两种常用的同步原语，用于保护共享资源。理解它们的底层实现原理，对于编写高效且并发安全的代码至关重要。

### Mutex 的底层实现原理

`Mutex`（互斥锁）是一种排他锁，其核心思想是，在任何给定时刻，只有一个 goroutine 可以持有锁。如果一个 goroutine 已经获得了锁，其他试图获取锁的 goroutine 将被阻塞，直到锁被释放。

`Mutex` 的实现是基于 **futex (Fast Userspace muTEX)**。futex 是一种由操作系统内核提供的同步机制，它在用户空间和内核空间之间提供了一个高效的桥梁。

当一个 goroutine 尝试加锁时，`Mutex` 的工作流程大致如下：

1.  **快速路径（Fast Path）**：goroutine 尝试通过原子操作（atomic operation）直接获取锁。如果锁当前没有被占用，原子操作会成功，goroutine 立即获得锁。这是一种无锁的（lock-free）操作，非常快。
2.  **慢速路径（Slow Path）**：如果快速路径失败，说明锁已经被其他 goroutine 占用。此时，goroutine 会进入慢速路径。它会检查锁的状态，并准备进入休眠。
3.  **休眠**：goroutine 会调用内核的 futex 相关系统调用，告诉内核它希望被阻塞（park），直到锁被释放。这个 goroutine 会被放入一个等待队列，并进入休眠状态，释放 CPU 资源。
4.  **唤醒**：当持有锁的 goroutine 调用 `Unlock()` 时，它会检查是否有其他 goroutine 在等待。如果有，它会通过 futex 系统调用通知内核，内核会从等待队列中唤醒一个或多个 goroutine。
5.  **竞争**：被唤醒的 goroutine 会再次尝试获取锁。

`Mutex` 的设计目标是最大化快速路径的效率，因为在大多数情况下，锁的竞争并不激烈。

### RWMutex 的底层实现原理

`RWMutex`（读写互斥锁）是 `Mutex` 的一个扩展，它允许多个 goroutine 同时读共享资源，但只允许一个 goroutine 写共享资源。这种设计可以显著提高读多写少的应用场景的并发性能。

`RWMutex` 的核心思想是区分“读”和“写”两种操作，并分别对待。它的内部实现要比 `Mutex` 复杂，通常会包含以下几个部分：

* **一个互斥锁（`Mutex`）**：用于保护写操作，确保在任何时候只有一个 goroutine 可以进行写操作。
* **一个计数器**：记录当前正在进行读操作的 goroutine 数量。
* **一个信号量或等待队列**：用于管理等待写入的 goroutine。

`RWMutex` 的工作流程大致如下：

#### 读锁（`RLock`）

1.  **获取读锁**：当一个 goroutine 尝试调用 `RLock()` 时，它会原子地增加读操作的计数器。
2.  **写锁检查**：在增加计数器之前，它会检查是否有 goroutine 在等待写锁。如果有，为了避免读操作持续阻塞写操作，它可能会被阻塞，直到写锁被释放。
3.  **同时读**：如果当前没有 goroutine 正在进行写操作，或者没有写操作在等待，多个 goroutine 可以同时成功获取读锁，并继续执行。

#### 读锁释放（`RUnlock`）

1.  **释放读锁**：当一个 goroutine 调用 `RUnlock()` 时，它会原子地减少读操作的计数器。
2.  **唤醒写锁**：如果减少后计数器变为 0，并且有 goroutine 在等待写锁，它会唤醒一个正在等待的写 goroutine。

#### 写锁（`Lock`）

1.  **获取写锁**：当一个 goroutine 尝试调用 `Lock()` 时，它会尝试获取一个内部的互斥锁。
2.  **等待读锁**：如果内部的互斥锁被其他 goroutine 占用（可能正在进行写操作），它会被阻塞。如果内部互斥锁获取成功，它会等待所有读操作完成，也就是等待读计数器变为 0。
3.  **独占**：一旦读计数器变为 0，写 goroutine 就获得了独占的写权限，开始执行写操作。

#### 写锁释放（`Unlock`）

1.  **释放写锁**：当一个 goroutine 调用 `Unlock()` 时，它会释放内部的互斥锁，并通知所有等待读锁或写锁的 goroutine，让它们继续竞争。

总而言之，**`Mutex` 适用于所有需要独占访问资源的场景**，而 **`RWMutex` 适用于读多写少的场景**。因为 `RWMutex` 引入了额外的逻辑和开销，如果读写操作的比例不倾向于读，使用 `RWMutex` 可能并不会带来性能提升，甚至可能因为其复杂性而导致性能下降。

在选择使用哪种锁时，应根据实际的业务场景和对共享资源的访问模式来决定。

在 Go 语言中，waitq（或称 waitqueue）是一个底层的同步机制，用于协程（goroutine）之间的通信和等待。它本质上是一个队列，存储着因某些条件不满足而需要等待的协程。当条件满足时，这些协程会被唤醒并重新调度执行。

Go 的 waitq 数据结构并没有直接暴露给用户，它被封装在许多同步原语的底层实现中，例如 sync.Mutex、sync.WaitGroup、sync.Cond 等。理解 waitq 的工作原理有助于你更好地理解这些同步原语是如何工作的。

waitq 的核心思想
可以把 waitq 理解为一个简单的“等待队列”：

当一个协程尝试获取一个锁或等待一个条件，但发现条件不满足（例如，锁已经被占用），它不会一直忙等，而是会把自己“挂”在这个等待队列上。

这个“挂”的过程意味着协程进入一个休眠状态，不再占用 CPU 资源。

当持有锁的协程释放锁或者某个条件被满足时，它会从 waitq 中唤醒一个或多个等待的协程。

被唤醒的协程会被重新放入调度器的可运行队列中，等待下一次被调度执行。

waitq 在 sync.Mutex 中的应用
以 sync.Mutex 为例，它的内部实现就依赖于 waitq。一个 Mutex 结构体中包含一个 state 字段，以及一个用于等待的 waitq。

加锁 (Lock): 当一个协程调用 mu.Lock() 时，它会尝试原子性地修改 state 字段。

如果修改成功（state 从 0 变为 1），表示协程成功获取了锁，继续执行。

如果修改失败（state 已经被其他协程持有），那么该协程会进入一个 自旋（spinning） 阶段，在短时间内尝试再次获取锁。

如果自旋后仍然无法获取锁，协程会把自己挂在 Mutex 内部的 waitq 上，然后进入休眠状态，让出 CPU。

解锁 (Unlock): 当持有锁的协程调用 mu.Unlock() 时，它会释放锁。

释放锁后，它会检查 waitq 中是否有等待的协程。

如果 waitq 非空，它会唤醒队列中的一个或多个协程，让它们重新竞争锁。

## Context 控制链

传递超时、取消信号（WithTimeout/Cancel）

典型场景：取消下游 HTTP 请求、终止耗时计算


## Slice 深度考点

a := []int{2:1}  指定下标2的值为1

底层结构：(ptr *array, len, cap)

扩容规则：

容量 <1024：双倍扩容；

容量 ≥1024：1.25 倍扩容

共享问题：切片截取共享底层数组，append 触发扩容后分离

## Map 实现与陷阱

哈希表结构：Bucket + Overflow 链表（解决哈希冲突）

非并发安全：并发写触发 fatal error，需用 sync.Map 或 Mutex

无序性：需借助 Slice 排序 Key 后遍历

Go 的哈希表是一个由 buckets 组成的数组。每个 bucket 实际上是一个包含键值对的数组。当一个新的键值对被插入时，Go 会计算键的哈希值，并根据这个哈希值来决定它应该被存放到哪一个 bucket 里。

hmap bmap 

渐进式hash

hash冲突的其他方案: 


开放定址法 (Open Addressing)
开放定址法是一种将所有元素都存放在哈希表本身的方案。当发生哈希冲突时，它会按一定的规则在哈希表中寻找下一个空的槽位（slot）来存放冲突的元素。

多重哈希
它使用多个哈希函数来解决冲突。当第一个哈希函数发生冲突时，就使用第二个哈希函数进行哈希，如果还冲突，就用第三个

桶哈希法 (Bucket Hashing)
桶哈希法是把哈希表中的每个位置看作一个桶，每个桶可以存放多个元素。当发生哈希冲突时，冲突的元素都被放入同一个桶中。这个桶可以用一个数组或者链表来实现，所以链表法实际上也可以被看作是桶哈希的一种特殊形式，即每个桶都用链表来存储元素。

## 字符串优化

高效拼接：strings.Builder（零拷贝） > bytes.Buffer 

不可变性：修改需转 []byte，重新分配内存

## 其他

GC 三色标记法：白→灰→黑，与 STW（Stop-The-World）优化，写屏障

并发标记

defer recover
错误逐层添加上下文（Wrap），避免吞没原始错误
recover 需在 defer 中调用，且仅能捕获当前 Goroutine 的 panic

## 内存优化

对象池 sync.Pool
减少内存频繁分配，切片预分配
小对象尽量使用栈上存储

## 协程泄露

没有合理的退出

+ 空的无限制for 循环
+ channel未释放
+ 锁导致协程不释放

work-poll解决
sync.Poll
