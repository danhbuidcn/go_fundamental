# Giải Thích Bộ Nhớ Trong Go

Để hiểu cách Go cấp phát bộ nhớ trên heap và stack, và làm thế nào để tối ưu chương trình Go, chúng ta cần tìm hiểu chi tiết về hai loại bộ nhớ này và cách chúng hoạt động cùng nhau trong Go. Dưới đây là một giải thích dễ hiểu về các khái niệm và cách tối ưu hóa chúng.

## 1. **Heap và Stack Là Gì?**

### **1.1. Stack**

- **Khái Niệm**: Stack là một vùng bộ nhớ nhỏ, có kích thước cố định, được sử dụng để lưu trữ các biến cục bộ và thông tin về các hàm khi chúng được gọi. Stack hoạt động theo nguyên tắc "LIFO" (Last In, First Out), nghĩa là phần tử được đưa vào cuối cùng sẽ được lấy ra trước.
  
- **Tính Chất**:
  - **Tự Động**: Bộ nhớ trên stack được cấp phát và giải phóng tự động khi các hàm được gọi và kết thúc.
  - **Tốc Độ**: Stack rất nhanh vì nó chỉ cần di chuyển con trỏ stack lên và xuống.
  - **Kích Thước**: Thường có kích thước nhỏ hơn và hạn chế, vì vậy chỉ dùng cho các biến cục bộ nhỏ và tạm thời.

### **1.2. Heap**

- **Khái Niệm**: Heap là một vùng bộ nhớ lớn hơn, không có kích thước cố định, dùng để lưu trữ các đối tượng và dữ liệu có thời gian sống lâu hơn. Khi một biến hoặc đối tượng cần sống lâu hơn vòng đời của hàm, nó được cấp phát trên heap.
  
- **Tính Chất**:
  - **Quản Lý Bằng Garbage Collection**: Bộ nhớ trên heap cần được quản lý bằng garbage collection (GC). Khi dữ liệu không còn được sử dụng, GC sẽ tự động giải phóng bộ nhớ.
  - **Tốc Độ**: Cấp phát và giải phóng bộ nhớ trên heap chậm hơn so với stack vì cần quản lý nhiều hơn.

## 2. **Cấp Phát Bộ Nhớ Trong Go**

### **2.1. Cấp Phát Trên Stack**

- **Biến Cục Bộ**: Khi bạn khai báo biến trong hàm, Go sẽ cấp phát bộ nhớ cho biến đó trên stack. Ví dụ:

  ```go
  func add(a int, b int) int {
      result := a + b
      return result
  }
  ```

  Ở đây, `result`, `a`, và `b` được lưu trên stack.

- **Lưu Trữ Thông Tin Hàm**: Mỗi lần bạn gọi một hàm, Go sẽ tạo một frame mới trên stack để lưu trữ các tham số và biến cục bộ của hàm đó.

### **2.2. Cấp Phát Trên Heap**

- **Biến Toàn Cục và Đối Tượng**: Khi bạn tạo một đối tượng lớn hoặc một biến toàn cục, hoặc khi bạn sử dụng từ khóa `new` hoặc `make`, Go sẽ cấp phát bộ nhớ trên heap. Ví dụ:

  ```go
  func createSlice() []int {
      slice := make([]int, 1000)
      return slice
  }
  ```

  Ở đây, `slice` được cấp phát trên heap vì nó có thể cần không gian lớn và thời gian sống lâu hơn.

- **Dữ Liệu Được Trả Về**: Nếu hàm trả về một giá trị hoặc đối tượng lớn, nó có thể được cấp phát trên heap.

## 3. **Cách Xác Định Bộ Nhớ Được Cấp Phát Ở Đâu**

### **3.1. Đối Với Biến Cục Bộ Trong Hàm**

- **Cấp Phát Trên Stack**: Nếu biến (ví dụ, một slice) được khai báo và sử dụng hoàn toàn trong một hàm, thì nó thường được cấp phát trên stack. Stack rất nhanh và tự động giải phóng bộ nhớ khi hàm kết thúc.

  **Ví dụ**:
  ```go
  func example() {
      slice := []int{1, 2, 3, 4, 5} // Cấp phát trên stack
      fmt.Println(slice)
  }
  ```

  Ở đây, `slice` chỉ tồn tại trong hàm `example`. Bộ nhớ cho `slice` được cấp phát trên stack và sẽ tự động giải phóng khi hàm `example` kết thúc.

### **3.2. Đối Với Biến Toàn Cục Hoặc Đối Tượng Trả Về**

- **Cấp Phát Trên Heap**: Nếu một biến hoặc đối tượng (ví dụ, một slice) được lưu trữ ngoài hàm (như biến toàn cục) hoặc được trả về từ một hàm, bộ nhớ cho nó sẽ được cấp phát trên heap. Điều này đảm bảo rằng dữ liệu vẫn tồn tại sau khi hàm kết thúc hoặc cho các phần của chương trình khác có thể truy cập.

  **Ví dụ**:
  ```go
  var globalSlice = []int{1, 2, 3, 4, 5} // Cấp phát trên heap

  func getSlice() []int {
      slice := []int{1, 2, 3, 4, 5} // Cấp phát trên heap vì nó được trả về
      return slice
  }
  ```

  Ở đây, `globalSlice` và `slice` trong hàm `getSlice` đều được cấp phát trên heap. `globalSlice` tồn tại suốt thời gian chạy của chương trình, trong khi `slice` được cấp phát trên heap vì nó cần tồn tại sau khi hàm `getSlice` kết thúc và được trả về cho các phần khác của chương trình.

### **3.3. Sử Dụng Công Cụ Để Phân Tích**

- **Sử Dụng pprof**: Công cụ profiling `pprof` trong Go có thể giúp bạn phân tích bộ nhớ và xác định nơi dữ liệu được cấp phát. Bạn có thể sử dụng `pprof` để theo dõi và phân tích bộ nhớ heap và stack trong thời gian thực.

  **Cách sử dụng pprof**:
  1. **Cài Đặt pprof**: `pprof` là một phần của Go, vì vậy bạn không cần cài đặt thêm gì.
  2. **Thêm Phần Mềm Để Thu Thập Profiling**: Thêm đoạn mã sau vào chương trình của bạn để thu thập thông tin profiling:
     ```go
     import (
         "net/http"
         _ "net/http/pprof"
     )

     func main() {
         go func() {
             log.Println(http.ListenAndServe("localhost:6060", nil))
         }()
         // Phần còn lại của chương trình
     }
     ```
  3. **Chạy Chương Trình và Xem Profiling**: Chạy chương trình của bạn và mở trình duyệt hoặc công cụ dòng lệnh để xem các số liệu phân tích từ `pprof`:
     ```
     go tool pprof http://localhost:6060/debug/pprof/heap
     go tool pprof http://localhost:6060/debug/pprof/stack
     ```

  Các công cụ này sẽ cung cấp thông tin chi tiết về việc sử dụng bộ nhớ trên heap và stack, giúp bạn hiểu cách chương trình của bạn phân bổ bộ nhớ.

## 4. **Tối Ưu Hóa Bộ Nhớ Trong Go**

### **4.1. Tối Ưu Hóa Sử Dụng Stack**

- **Sử Dụng Biến Cục Bộ Thay Vì Biến Toàn Cục**: Biến cục bộ thường nhanh hơn và tiết kiệm bộ nhớ hơn vì chúng được cấp phát trên stack.
- **Tránh Gọi Hàm Quá Sâu**: Việc gọi hàm sâu có thể làm stack tràn, dẫn đến lỗi "stack overflow". Nên thiết kế chương trình để tránh gọi hàm quá sâu.

### **4.2. Tối Ưu Hóa Sử Dụng Heap**

- **Giảm Số Lượng Đối Tượng**: Cố gắng giảm số lượng đối tượng lớn và chỉ tạo các đối tượng khi thực sự cần thiết.
- **Sử Dụng Garbage Collection Hiệu Quả**: Go có garbage collector để tự động giải phóng bộ nhớ. Tuy nhiên, bạn nên cố gắng tránh giữ tham chiếu đến các đối tượng không còn sử dụng để GC có thể thu hồi bộ nhớ hiệu quả hơn.
- **Sử Dụng Bộ Nhớ Tinh Vi**: Sử dụng các cấu trúc dữ liệu và thuật toán phù hợp để tiết kiệm bộ nhớ.

## 5. **Ví Dụ Cụ Thể**

Dưới đây là một ví dụ đơn giản để minh họa sự khác biệt giữa việc sử dụng stack và heap trong Go:

```go
package main

import "fmt"

// Hàm tạo một đối tượng lớn và trả về nó
func createLargeObject() []int {
    largeArray := make([]int, 1000000) // Cấp phát trên heap
    return largeArray
}

func main() {
    a := 10 // Cấp phát trên stack
    b := 20 // Cấp phát trên stack

    result := a + b // Cấp phát trên stack
    fmt.Println("Result:", result)

    largeObject := createLargeObject() // Đối tượng lớn được cấp phát trên heap
    fmt.Println("Large object length:", len(largeObject))
}
```

**Giải Thích**:
- `a`, `b`, và `result` được cấp phát trên stack.
- `largeObject` được cấp phát trên heap vì nó là một đối tượng lớn và có thể cần không gian bộ nhớ lớn.

# Tìm hiểu thêm

https://200lab.io/blog/golang-cap-phat-bo-nho-nhu-the-nao/
