# map

-----

### `map`

`map` 是 Go 语言内置的数据类型，用于存储键值对。它不是并发安全的，如果在多个 goroutine 中同时对一个 `map` 进行读写操作，会引发 **竞态条件（race condition）**，导致程序崩溃。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	// 模拟并发写入
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i
		}(i)
	}

	wg.Wait()
	fmt.Println("Map size:", len(m)) // 在没有锁的情况下，这个大小可能小于1000，并且程序可能会崩溃
}
```

为了解决并发问题，通常需要手动添加锁（例如 `sync.Mutex` 或 `sync.RWMutex`）。

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]int)
	var wg sync.WaitGroup

	// 模拟并发写入
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[i] = i
		}(i)
	}

	wg.Wait()
	fmt.Println("Map size:", len(m))
}
```

-----

### `sync.Map`

`sync.Map` 是 Go 语言标准库 `sync` 包中提供的一种特殊的并发安全的映射（Map）类型。它专门为 **“读多写少”** 的场景进行了优化，旨在提供比使用 `sync.Mutex` 保护的 `map` 更高的并发性能。

`sync.Map` 的主要特点：

  * **无需初始化**：`sync.Map` 是一个结构体，可以直接使用 `var m sync.Map` 或 `m := sync.Map{}` 声明，不需要像 `map` 那样用 `make` 来初始化。
  * **并发安全**：内置了锁机制，可以在多个 goroutine 中安全地读写，而不会导致竞态条件。
  * **优化读操作**：它内部维护了两个 `map`：一个用于存储只读数据，另一个用于存储需要写入的数据。读取操作会首先尝试从只读 `map` 中获取，如果获取失败再从写 `map` 中获取，这样可以减少锁的使用，提高读性能。
  * **API 差异**：`sync.Map` 的操作方法与内置 `map` 不同，它没有 `len()` 方法，也不能使用 `for range` 循环。

`sync.Map` 的常用方法：

  * **`Store(key, value interface{})`**：存储键值对。
  * **`Load(key interface{}) (value interface{}, ok bool)`**：加载键值，返回一个值和一个布尔值表示键是否存在。
  * **`LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)`**：如果键存在则返回现有值，否则存储新值并返回。
  * **`Delete(key interface{})`**：删除键。
  * **`Range(f func(key, value interface{}) bool)`**：遍历 `sync.Map`。

<!-- end list -->

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	// 模拟并发写入
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}

	wg.Wait()

	// 遍历并打印
	count := 0
	m.Range(func(key, value interface{}) bool {
		count++
		return true // 返回 true 继续遍历
	})
	fmt.Println("Map size:", count)
}
```

-----

### 总结与选择建议

| 特性 | `map` | `sync.Map` |
| :--- | :--- | :--- |
| **并发安全** | 否，需要手动加锁 | 是，内置并发控制 |
| **初始化** | 必须使用 `make()` | 无需初始化 |
| **性能** | **单线程**下性能极佳 | **读多写少**的并发场景下性能更优 |
| **API** | 内置语法，如 `m[key]` | 方法调用，如 `m.Load(key)` |
| **遍历/大小**| 可使用 `for range` 和 `len()` | 只能使用 `Range()` 方法遍历，没有 `len()` 方法 |
| **适用场景** | 绝大多数场景，特别是在单线程或手动控制并发的场景 | 高并发、读多写少的场景，例如缓存 |

**何时选择哪个？**

  * **大多数情况下，使用 `map` 加锁（`sync.Mutex`）**。这种方式简单、直观，并且在 **写操作频繁** 或 **读写比例接近** 的场景下，性能通常不比 `sync.Map` 差，甚至更好。
  * **当你的程序面临** **高并发** **且** **读操作远多于写操作** **时，请使用 `sync.Map`**。它的设计正是为了优化这类场景的性能。

简单来说，如果无法确定读写比例，或者写操作很频繁，就用 `map` 加锁。如果确定是读多写少，`sync.Map` 是更好的选择。