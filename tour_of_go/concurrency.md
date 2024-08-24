# Concurrency

https://go.dev/tour/concurrency/1

## Goroutines

- Định nghĩa:
    - **Goroutine** là một luồng nhẹ (lightweight thread) được quản lý bởi Go runtime.
    - Khi bạn gọi một hàm với từ khóa `go` trước nó, ví dụ: `go f(x, y, z)`, điều này sẽ khởi chạy một **goroutine** mới để thực thi hàm `f(x, y, z)`.
    - Việc tính toán các giá trị `f`, `x`, `y`, và `z` diễn ra trong goroutine hiện tại, trong khi việc thực thi hàm `f` diễn ra trong goroutine mới.
    - **Goroutines** chạy trong cùng một không gian địa chỉ (address space), vì vậy việc truy cập bộ nhớ chia sẻ phải được đồng bộ hóa để tránh xung đột. 

- **Lưu ý**: Go cung cấp các công cụ đồng bộ hóa trong gói `sync`, nhưng thường không cần sử dụng nhiều vì Go có các công cụ khác để quản lý đồng bộ hóa.
- VD:
    ```go
    package main

    import (
        "fmt"
        "time"
    )

    // Hàm printNumbers in ra các số từ 1 đến 5 với một khoảng thời gian chờ giữa các số.
    func printNumbers() {
        for i := 1; i <= 5; i++ {
            fmt.Println(i) // In ra số i
            time.Sleep(100 * time.Millisecond) // Giả lập một công việc mất thời gian 100ms
        }
    }

    // Hàm printLetters in ra các chữ cái từ A đến E với một khoảng thời gian chờ giữa các chữ cái.
    func printLetters() {
        letters := []string{"A", "B", "C", "D", "E"}
        for _, letter := range letters {
            fmt.Println(letter) // In ra chữ cái
            time.Sleep(150 * time.Millisecond) // Giả lập một công việc mất thời gian 150ms
        }
    }

    func main() {
        // Khởi chạy một goroutine mới để thực thi hàm printNumbers.
        go printNumbers()

        // Khởi chạy một goroutine mới để thực thi hàm printLetters.
        go printLetters()

        // Đợi một thời gian (1 giây) để các goroutine có cơ hội hoàn thành công việc của chúng.
        // Thời gian chờ này đảm bảo rằng chương trình chính không kết thúc trước khi các goroutine hoàn thành.
        time.Sleep(1 * time.Second)
    }
    ```

## Channels

- **Channels** giúp đồng bộ hóa và giao tiếp giữa các goroutines mà không cần sử dụng khóa hoặc biến điều kiện.
- **Channels** chặn (block) khi gửi và nhận giá trị, giúp các goroutines đồng bộ hóa với nhau một cách dễ dàng.

- VD: 
    ```go
    package main

    import (
        "fmt"
    )

    // Hàm sum nhận một slice các số nguyên và một channel.
    // Nó tính tổng của các số trong slice và gửi kết quả qua channel.
    func sum(numbers []int, ch chan int) {
        total := 0
        for _, number := range numbers {
            total += number
        }
        ch <- total // Gửi tổng vào channel
    }

    func main() {
        // Slice chứa các số nguyên để tính tổng
        numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        
        // Tạo hai channel để nhận kết quả từ hai goroutines
        ch1 := make(chan int)
        ch2 := make(chan int)
        
        // Chia slice thành hai phần và khởi chạy hai goroutines để tính tổng
        go sum(numbers[:len(numbers)/2], ch1) // Tính tổng phần đầu tiên của slice
        go sum(numbers[len(numbers)/2:], ch2) // Tính tổng phần còn lại của slice
        
        // Nhận kết quả từ cả hai channel
        total1 := <-ch1 // Nhận tổng từ channel đầu tiên
        total2 := <-ch2 // Nhận tổng từ channel thứ hai
        
        // Tính tổng cuối cùng bằng cách cộng tổng từ hai channel
        total := total1 + total2
        
        // In ra kết quả
        fmt.Println("Tổng của các số là:", total)
    }
    ```

## Buffered Channels

- **Channel đệm (`Buffered Channels`)** cho phép lưu trữ một số lượng giá trị nhất định trong channel trước khi việc gửi giá trị mới bị chặn. 

    - Gửi dữ liệu: Các giá trị được gửi vào channel. Nếu buffer đầy, việc gửi dữ liệu sẽ chặn cho đến khi có giá trị được nhận từ channel, giải phóng chỗ trong buffer.

    - Nhận dữ liệu: Các giá trị được nhận từ channel. Việc nhận dữ liệu sẽ chặn nếu buffer trống, cho đến khi có giá trị mới được gửi vào channel.

- VD:
    ```go
    package main

    import "fmt"

    func sendData(ch chan int, data []int) {
        // Gửi các giá trị từ slice data vào channel
        for _, value := range data {
            ch <- value // Gửi giá trị vào channel
            fmt.Println("Sent:", value) // In giá trị đã gửi
        }
        close(ch) // Đóng channel sau khi gửi xong tất cả dữ liệu
    }

    func main() {
        // Tạo một buffered channel với kích thước buffer là 3
        ch := make(chan int, 3)

        // Dữ liệu để gửi vào channel
        data := []int{1, 2, 3, 4, 5, 6}

        // Khởi chạy goroutine để gửi dữ liệu vào channel
        go sendData(ch, data)

        // Nhận dữ liệu từ channel cho đến khi channel bị đóng
        for value := range ch {
            fmt.Println("Received:", value) // In giá trị đã nhận
        }
    }
    ```

## Range and Close

- **Đóng Channel**: Sender có thể đóng channel để thông báo không còn giá trị nào sẽ được gửi thêm.
- **Kiểm tra Channel Đã Đóng**: Receiver kiểm tra bằng cách sử dụng `v, ok := <-ch`. Nếu `ok` là `false`, channel đã đóng.
- **Vòng Lặp `range`**: Dùng để nhận giá trị từ channel liên tục cho đến khi channel đóng.
- **Lưu ý**: Chỉ sender mới nên đóng channel. Đóng channel không phải lúc nào cũng cần, chỉ cần thiết khi receiver cần biết không còn giá trị nào sẽ được gửi.

- Ví dụ:
    ```go
    package main

    import "fmt"

    func sendData(ch chan int) {
        for i := 1; i <= 5; i++ {
            ch <- i // Gửi giá trị vào channel
        }
        close(ch) // Đóng channel sau khi gửi xong tất cả dữ liệu
    }

    func main() {
        ch := make(chan int) // Tạo một channel kiểu int

        go sendData(ch) // Khởi chạy goroutine để gửi dữ liệu vào channel

        // Nhận dữ liệu từ channel cho đến khi channel bị đóng
        for value := range ch {
            fmt.Println("Received:", value) // In giá trị đã nhận
        }
    }
    ```

## Select

- **`select` Statement**: Cho phép một goroutine chờ đợi trên nhiều hoạt động giao tiếp khác nhau.
- **Hoạt Động**:
  - `select` chặn cho đến khi một trong các trường hợp (cases) có thể thực thi.
  - Nếu nhiều trường hợp đều sẵn sàng, `select` sẽ chọn ngẫu nhiên một trường hợp để thực thi.

- Ví dụ
    ```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        // Tạo hai kênh để sử dụng trong ví dụ
        ch1 := make(chan string)
        ch2 := make(chan string)

        // Goroutine gửi giá trị vào ch1 sau 2 giây
        go func() {
            time.Sleep(2 * time.Second)
            ch1 <- "Message from ch1"
        }()

        // Goroutine gửi giá trị vào ch2 sau 1 giây
        go func() {
            time.Sleep(1 * time.Second)
            ch2 <- "Message from ch2"
        }()

        // Sử dụng select để chờ nhận giá trị từ một trong các kênh
        select {
        case msg1 := <-ch1:
            // Nếu ch1 có giá trị sẵn, in giá trị đó ra
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            // Nếu ch2 có giá trị sẵn, in giá trị đó ra
            fmt.Println("Received:", msg2)
        case <-time.After(3 * time.Second):
            // Nếu không có kênh nào gửi giá trị trong 3 giây, in thông báo timeout
            fmt.Println("Timeout")
        }
    }
    ```

## Default Selection

- **Trường hợp default** cho phép bạn thực hiện một hành động mà không bị chặn, ngay cả khi không có giá trị nào sẵn sàng trên các kênh khác.
- VD:
    ```go
    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        // Tạo một kênh để sử dụng trong ví dụ
        c := make(chan int)

        // Goroutine gửi giá trị vào kênh sau 1 giây
        go func() {
            time.Sleep(1 * time.Second)
            c <- 42
        }()

        // Sử dụng select với trường hợp default
        select {
        case i := <-c:
            // Nếu kênh c có giá trị sẵn, nhận giá trị và in ra
            fmt.Println("Received:", i)
        default:
            // Nếu không có giá trị sẵn trong kênh c, thực hiện hành động này
            fmt.Println("No value received, executing default case")
        }

        // Chờ một chút trước khi kết thúc chương trình để đảm bảo tất cả các goroutine đã hoàn thành
        time.Sleep(2 * time.Second)
    }
    ```

## sync.Mutex

- Khi làm việc với goroutines, nếu bạn cần đảm bảo rằng chỉ một goroutine có thể truy cập một biến tại một thời điểm để tránh xung đột, bạn sẽ sử dụng cơ chế khóa (mutex). 

- **Mutex (Mutual Exclusion)** giúp bạn quản lý quyền truy cập đồng thời vào tài nguyên chia sẻ. Go cung cấp một loại mutex thông qua gói `sync` với kiểu `sync.Mutex`. 

- `sync.Mutex` cung cấp hai phương thức chính:
    - **`Lock`**: Được sử dụng để chiếm quyền truy cập vào tài nguyên.
    - **`Unlock`**: Được sử dụng để giải phóng quyền truy cập vào tài nguyên.

- Khi một goroutine gọi `Lock`, nó sẽ chiếm quyền truy cập vào mutex và các goroutine khác sẽ phải chờ cho đến khi mutex được giải phóng bằng cách gọi `Unlock`.

- VD:
    ```go
    package main

    import (
        "fmt"
        "sync"
    )

    // Định nghĩa một cấu trúc với một mutex
    type Counter struct {
        mu    sync.Mutex // Mutex để bảo đảm loại trừ lẫn nhau
        count int        // Biến đếm
    }

    // Phương thức Inc để tăng giá trị biến đếm một cách an toàn
    func (c *Counter) Inc() {
        c.mu.Lock()         // Khóa mutex để đảm bảo chỉ một goroutine có thể truy cập vào biến count
        c.count++           // Phần quan trọng: thay đổi biến chia sẻ
        c.mu.Unlock()       // Mở khóa mutex để các goroutine khác có thể truy cập
    }

    // Phương thức Value để đọc giá trị biến đếm một cách an toàn
    func (c *Counter) Value() int {
        c.mu.Lock()         // Khóa mutex để đảm bảo chỉ một goroutine có thể đọc biến count
        defer c.mu.Unlock() // Mở khóa mutex khi phương thức kết thúc, kể cả khi có lỗi xảy ra
        return c.count      // Phần quan trọng: đọc giá trị biến chia sẻ
    }

    func main() {
        var wg sync.WaitGroup
        counter := Counter{} // Khởi tạo biến đếm

        // Tạo 100 goroutine để tăng giá trị biến đếm
        for i := 0; i < 100; i++ {
            wg.Add(1) // Tăng số lượng goroutine cần chờ
            go func() {
                defer wg.Done() // Giảm số lượng goroutine cần chờ khi kết thúc
                counter.In()    // Gọi phương thức Inc để tăng giá trị biến đếm
            }()
        }

        // Chờ tất cả các goroutine kết thúc
        wg.Wait()

        // In ra giá trị cuối cùng của biến đếm
        fmt.Println("Final count:", counter.Value())
    }
    ```
 
