# Hỗ trợ lập trình concurrent (đồng thời) rất dễ dàng với Goroutine

https://200lab.io/blog/golang-la-gi/

## 1. **Lập Trình Đồng Thời (Concurrency) Trong Go**

**Concurrency** trong lập trình đề cập đến khả năng quản lý nhiều tác vụ cùng một lúc. Điều này không có nghĩa là các tác vụ được thực thi đồng thời (song song), mà là chương trình có thể chuyển đổi giữa các tác vụ, làm cho tiến trình của nhiều tác vụ được thực hiện mà không cần chờ đợi từng tác vụ hoàn thành.

Trong Go, concurrency được thực hiện bằng cách sử dụng **Goroutine** và **channels**.

## 2. **Goroutines**

**Goroutines** là các luồng thực thi nhẹ được quản lý bởi runtime của Go. Chúng cho phép bạn chạy các hàm đồng thời, nghĩa là nhiều Goroutine có thể được thực thi cùng lúc, độc lập với nhau.

- **Tạo Một Goroutine**: Để khởi tạo một Goroutine, bạn chỉ cần thêm từ khóa `go` trước một lời gọi hàm. Ví dụ:

    ```go
    package main

    import (
        "fmt"
        "time"
    )

    func say(s string) {
        for i := 0; i < 3; i++ {
            time.Sleep(100 * time.Millisecond)
            fmt.Println(s)
        }
    }

    func main() {
        go say("Hello")
        say("World")
    }
    ```

**Giải Thích:**
- Hàm `say` được gọi như một Goroutine bằng cách thêm từ khóa `go` trước nó. Điều này có nghĩa là `say("Hello")` sẽ chạy đồng thời với `say("World")`.
- Goroutine `say("Hello")` có thể bị tạm dừng và tiếp tục chạy sau khi `say("World")` hoàn thành. Kết quả in ra có thể không theo thứ tự.

## 3. **Lợi Ích Của Goroutines**

- **Nhẹ và Tiết Kiệm Tài Nguyên**: Goroutines rất nhẹ so với các thread truyền thống. Khi một thread thông thường tiêu tốn hàng megabyte bộ nhớ, một Goroutine chỉ tiêu tốn khoảng 2 KB.
  
- **Tự Động Quản Lý**: Go runtime tự động quản lý các Goroutines, bao gồm việc phân phối chúng trên nhiều CPU nếu có.

- **Đơn Giản và Tiện Lợi**: Với Goroutines, lập trình đồng thời trở nên rất dễ dàng. Bạn chỉ cần thêm từ khóa `go` để chạy một hàm đồng thời.

## 4. **Nhược Điểm Của Goroutines và Cách Xử Lý**

### 1. **Deadlock**

- **Deadlock** xảy ra khi hai hoặc nhiều Goroutines chờ đợi nhau để hoàn thành một tác vụ, dẫn đến tình trạng tất cả đều bị khóa và không có tiến trình nào có thể tiếp tục. 

    ```go
    package main

    import (
        "fmt"
    )

    func main() {
        ch1 := make(chan int)
        ch2 := make(chan int)

        go func() {
            ch1 <- 1
            <-ch2
        }()

        <-ch1
        ch2 <- 2
    }
    ```

- **Giải Thích**:
    - Chương trình sẽ bị deadlock vì Goroutine đầu tiên đang chờ dữ liệu từ `ch2`, trong khi Goroutine chính đang chờ dữ liệu từ `ch1`. Cả hai Goroutine đều bị chặn chờ nhau, dẫn đến tình trạng không thể tiếp tục.

- **Cách Xử Lý**:

    - **Sử Dụng `select` Với `default`**: `select` trong Go có thể được sử dụng để tránh deadlock bằng cách cung cấp một trường hợp `default`, giúp chương trình không bị khóa khi không có dữ liệu trên bất kỳ kênh nào.

    ```go
    select {
    case val := <-ch1:
        fmt.Println(val)
    case ch2 <- 2:
        // Thực hiện hành động
    default:
        fmt.Println("Không có dữ liệu")
    }
    ```

    - **Thiết Kế Kênh Cẩn Thận**: Đảm bảo rằng mọi Goroutine đều có cách để thoát khỏi việc chờ đợi dữ liệu nếu dữ liệu đó không đến.

    - **Sử Dụng Kênh Có Đệm**: Kênh có đệm có thể giúp giảm thiểu nguy cơ deadlock bằng cách cho phép các giá trị được gửi mà không cần chờ người nhận.

### 2. **Race Condition (Điều Kiện Cạnh Tranh)**

- **Race Condition** xảy ra khi nhiều Goroutines truy cập và thay đổi cùng một biến hoặc vùng nhớ mà không có sự đồng bộ, dẫn đến kết quả không xác định.

    ```go
    package main

    import (
        "fmt"
        "sync"
    )

    var counter int

    func increment(wg *sync.WaitGroup) {
        defer wg.Done()
        for i := 0; i < 1000; i++ {
            counter++
        }
    }

    func main() {
        var wg sync.WaitGroup
        wg.Add(2)

        go increment(&wg)
        go increment(&wg)

        wg.Wait()
        fmt.Println("Final Counter:", counter)
    }
    ```

- **Giải Thích**:
    - Kết quả của biến `counter` có thể không chính xác do race condition, vì hai Goroutines có thể cùng lúc truy cập và thay đổi biến này mà không có sự đồng bộ hóa.

- **Cách Xử Lý**:

    - **Sử Dụng `sync.Mutex`**: `sync.Mutex` có thể được sử dụng để khóa các vùng nhớ mà chỉ một Goroutine có thể truy cập tại một thời điểm.

        ```go
        var mu sync.Mutex

        func increment(wg *sync.WaitGroup) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                mu.Lock()
                counter++
                mu.Unlock()
            }
        }
        ```

    - **Sử Dụng `sync.RWMutex`**: `sync.RWMutex` cung cấp các khóa cho cả đọc và ghi, cho phép nhiều Goroutines đọc đồng thời nhưng chỉ một Goroutine có thể ghi.

    - **Sử Dụng Kênh (Channels)**: Channels có thể được sử dụng để điều phối việc truy cập vào dữ liệu, đảm bảo rằng chỉ một Goroutine có thể thay đổi dữ liệu tại một thời điểm.

    ```go
    counter := make(chan int)

    go func() {
        for c := range counter {
            // Xử lý dữ liệu
        }
    }()

    counter <- 1
    ```

### 3. **Quản Lý Số Lượng Goroutines**

Việc tạo quá nhiều Goroutines mà không kiểm soát có thể dẫn đến việc tiêu tốn nhiều tài nguyên hệ thống, thậm chí gây ra lỗi do hết tài nguyên.

- **Cách Xử Lý**:

    - **Hạn Chế Số Lượng Goroutines**: Sử dụng một kênh đệm để giới hạn số lượng Goroutines hoạt động cùng lúc.

    ```go
    var wg sync.WaitGroup
    sem := make(chan struct{}, 10) // Giới hạn 10 Goroutines

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sem <- struct{}{}
            // Thực hiện công việc
            <-sem
        }()
    }

    wg.Wait()
    ```

    - **Sử Dụng Worker Pools**: Thiết kế chương trình với một tập hợp các Goroutines cố định (worker pool) để xử lý các tác vụ từ một hàng đợi công việc.

    ```go
    func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
        defer wg.Done()
        for j := range jobs {
            results <- j * 2
        }
    }

    func main() {
        jobs := make(chan int, 100)
        results := make(chan int, 100)
        var wg sync.WaitGroup

        for w := 1; w <= 3; w++ {
            wg.Add(1)
            go worker(jobs, results, &wg)
        }

        for j := 1; j <= 5; j++ {
            jobs <- j
        }
        close(jobs)

        wg.Wait()
        close(results)

        for result := range results {
            fmt.Println("Result:", result)
        }
    }
    ```

## 5. **Ví Dụ Về Sử Dụng Goroutines**

Dưới đây là một ví dụ đơn giản về việc sử dụng Goroutines và channels để xử lý đồng thời nhiều tác vụ:

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("Worker", id, "processing job", j)
        time.Sleep(time.Second) // Giả lập việc xử lý tác vụ
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results) // Tạo 3 Goroutine worker
    }

    for

 j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 5; a++ {
        fmt.Println("Result:", <-results)
    }
}
```

- **Giải Thích**:
    - **Worker**: Mỗi worker là một Goroutine xử lý công việc từ kênh `jobs`.
    - **Channels**: Channels (`jobs` và `results`) được sử dụng để truyền dữ liệu giữa các Goroutines.

## 6. **Kết Luận**

Goroutines trong Go là một công cụ mạnh mẽ cho lập trình đồng thời, giúp cho việc xử lý nhiều tác vụ cùng lúc trở nên dễ dàng và hiệu quả. Tuy nhiên, chúng cũng đi kèm với một số nhược điểm như deadlock và race condition, đòi hỏi lập trình viên cần có kiến thức và cẩn thận khi sử dụng. Bằng cách áp dụng các kỹ thuật đã đề cập, bạn có thể tận dụng tối đa sức mạnh của Goroutines trong khi hạn chế các rủi ro tiềm ẩn.

# Channels Trong Go

**Channels** là một cơ chế mạnh mẽ trong Go để giao tiếp giữa các Goroutines và đồng bộ hóa chúng. Chúng cho phép các Goroutines gửi và nhận dữ liệu với nhau một cách an toàn và hiệu quả.

## 1. **Khái Niệm Cơ Bản Về Channels**

- **Tạo Channel**: Bạn tạo một channel bằng cách sử dụng hàm `make`. Ví dụ: `ch := make(chan int)` tạo một channel truyền dữ liệu kiểu `int`.

- **Gửi Dữ Liệu**: Để gửi dữ liệu vào channel, bạn sử dụng cú pháp `<-`. Ví dụ: `ch <- 1` gửi giá trị `1` vào channel `ch`.

- **Nhận Dữ Liệu**: Để nhận dữ liệu từ channel, bạn cũng sử dụng cú pháp `<-`. Ví dụ: `val := <-ch` nhận giá trị từ channel `ch` và gán nó cho biến `val`.

- **Đóng Channel**: Khi bạn hoàn tất việc gửi dữ liệu, bạn nên đóng channel bằng cách sử dụng `close(ch)`. Điều này giúp các Goroutines khác biết rằng không còn dữ liệu nào để nhận.

## 2. **Ví Dụ Cơ Bản Về Channels**

Dưới đây là một ví dụ đơn giản về cách sử dụng channels để gửi và nhận dữ liệu giữa các Goroutines:

```go
package main

import (
    "fmt"
    "time"
)

func sendData(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
        time.Sleep(time.Millisecond * 100)
    }
    close(ch) // Đóng channel khi hoàn tất gửi dữ liệu
}

func main() {
    ch := make(chan int)

    go sendData(ch) // Khởi tạo Goroutine để gửi dữ liệu

    for val := range ch { // Nhận dữ liệu từ channel cho đến khi channel bị đóng
        fmt.Println(val)
    }
}
```

- **Giải Thích**:
    - **`sendData` Function**: Goroutine `sendData` gửi các số từ `0` đến `4` vào channel `ch` và đóng channel sau khi hoàn tất.
    - **`main` Function**: Nhận dữ liệu từ channel `ch` và in ra màn hình. Vòng lặp `for val := range ch` sẽ tiếp tục cho đến khi channel bị đóng và không còn dữ liệu để nhận.

## 3. **Ví Dụ Về Channels Có Đệm**

Channels có đệm cho phép bạn gửi dữ liệu vào channel mà không cần phải có một Goroutine khác nhận ngay lập tức. Đây là ví dụ về việc sử dụng channel có đệm:

```go
package main

import (
    "fmt"
)

func main() {
    ch := make(chan int, 3) // Tạo channel có đệm với kích thước 3

    ch <- 1
    ch <- 2
    ch <- 3

    // Không cần Goroutine để nhận dữ liệu ngay lập tức

    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

- **Giải Thích**:
    - **Tạo Channel Có Đệm**: Channel `ch` được tạo với kích thước 3, cho phép gửi 3 giá trị mà không cần phải nhận ngay lập tức.
    - **Gửi Dữ Liệu**: Các giá trị `1`, `2`, và `3` được gửi vào channel mà không gây ra deadlock vì channel có đệm.
    - **Nhận Dữ Liệu**: Các giá trị được nhận từ channel và in ra màn hình.

## 4. **Ví Dụ Về Channel Trong Worker Pool**

Dưới đây là một ví dụ về việc sử dụng channel trong một worker pool để xử lý các tác vụ đồng thời:

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d đang xử lý job %d\n", id, job)
        results <- job * 2 // Xử lý job và gửi kết quả vào channel
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    var wg sync.WaitGroup

    // Tạo 3 worker
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Gửi các công việc vào channel
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // Đóng channel jobs khi hoàn tất gửi công việc

    // Đợi tất cả worker hoàn thành
    wg.Wait()
    close(results) // Đóng channel results khi tất cả công việc đã được xử lý

    // Nhận kết quả từ channel results
    for result := range results {
        fmt.Println("Kết quả:", result)
    }
}
```

- **Giải Thích**:
    - **Worker**: Mỗi worker nhận công việc từ channel `jobs`, xử lý nó và gửi kết quả vào channel `results`.
    - **Channel `jobs`**: Được đóng sau khi tất cả công việc đã được gửi.
    - **Channel `results`**: Được đóng sau khi tất cả công việc đã được xử lý. Vòng lặp `for result := range results` nhận và in các kết quả.

## Kết Luận

Channels là công cụ mạnh mẽ trong Go cho phép giao tiếp và đồng bộ hóa giữa các Goroutines. Chúng giúp đảm bảo rằng các Goroutines có thể phối hợp hiệu quả và an toàn trong việc chia sẻ dữ liệu và xử lý tác vụ đồng thời. Bằng cách sử dụng channels đúng cách, bạn có thể tận dụng tối đa các lợi ích của concurrency trong Go.
