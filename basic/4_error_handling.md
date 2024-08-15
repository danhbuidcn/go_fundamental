# Hướng Dẫn Xử Lý Ngoại Lệ Trong Go

Trong Go, quản lý lỗi và xử lý ngoại lệ là phần quan trọng trong lập trình để đảm bảo rằng chương trình của bạn hoạt động một cách ổn định và có thể xử lý các tình huống không mong muốn. Go không có cơ chế xử lý ngoại lệ (exceptions) như nhiều ngôn ngữ khác (ví dụ: try-catch trong Java hoặc Python). Thay vào đó, Go sử dụng cơ chế trả về lỗi (error handling) đơn giản và hiệu quả. Dưới đây là hướng dẫn chi tiết về cách xử lý lỗi trong Go.

## 1. Khái Niệm Về Lỗi Trong Go

Trong Go, lỗi được đại diện bằng kiểu dữ liệu `error`. Đây là một interface có phương thức duy nhất `Error() string` để trả về mô tả lỗi dưới dạng chuỗi.

### Định Nghĩa Interface `error`:

```go
type error interface {
    Error() string
}
```

Khi một hàm gặp lỗi, nó sẽ trả về một giá trị của kiểu `error`. Nếu không có lỗi, hàm thường trả về `nil` cho lỗi.

## 2. Xử Lý Lỗi Cơ Bản

### a. Hàm Trả Về Lỗi

Khi viết hàm, bạn có thể trả về một giá trị lỗi nếu hàm gặp phải vấn đề nào đó. Dưới đây là ví dụ về cách viết hàm có thể trả về lỗi:

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

- **Giải Thích**:
  - Hàm `divide` nhận hai số nguyên `a` và `b` và trả về kết quả chia cùng với một giá trị lỗi.
  - Nếu `b` bằng 0, hàm trả về một lỗi mới với thông điệp "cannot divide by zero".
  - Nếu không có lỗi, hàm trả về kết quả và `nil` cho lỗi.

### b. Kiểm Tra Lỗi

Trong hàm `main`, bạn kiểm tra giá trị lỗi được trả về từ hàm `divide`:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

- **Giải Thích**:
  - Bạn kiểm tra xem biến `err` có phải là `nil` không. Nếu không, có nghĩa là đã xảy ra lỗi, và bạn in thông báo lỗi ra màn hình.
  - Nếu không có lỗi (`err` là `nil`), bạn tiếp tục sử dụng giá trị kết quả.

## 3. Tạo Lỗi Tùy Chỉnh

Bạn có thể tạo lỗi tùy chỉnh để cung cấp thông tin chi tiết hơn về lỗi. Dưới đây là ví dụ về cách tạo lỗi tùy chỉnh bằng cách định nghĩa một kiểu lỗi mới.

```go
package main

import (
    "fmt"
)

// Define a custom error type
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func performTask(success bool) error {
    if !success {
        return &MyError{Code: 500, Message: "Task failed"}
    }
    return nil
}

func main() {
    err := performTask(false)
    if err != nil {
        fmt.Println("Custom Error:", err)
        return
    }
    fmt.Println("Task completed successfully")
}
```

- **Giải Thích**:
  - Định nghĩa kiểu lỗi tùy chỉnh `MyError` với các trường `Code` và `Message`.
  - Cài đặt phương thức `Error()` để trả về thông điệp lỗi định dạng.
  - Hàm `performTask` trả về lỗi tùy chỉnh khi gặp lỗi.

## 4. Sử Dụng `panic` và `recover`

Go cũng cung cấp cơ chế `panic` và `recover` để xử lý lỗi nghiêm trọng mà bạn không thể xử lý bằng cách thông thường.

### a. `panic`

- **Mô Tả**: Gây ra một tình huống lỗi nghiêm trọng, khiến chương trình dừng lại. Chỉ nên dùng khi bạn gặp phải tình huống không thể phục hồi hoặc lỗi nghiêm trọng.

```go
func mayPanic() {
    panic("something went terribly wrong")
}

func main() {
    mayPanic()
    fmt.Println("This will not be printed")
}
```

- **Giải Thích**: Khi `panic` được gọi, chương trình dừng lại và không thực hiện các dòng mã sau đó.

### b. `recover`

- **Mô Tả**: Được sử dụng để phục hồi từ tình huống `panic`. Thường dùng trong các hàm `defer` để xử lý lỗi nghiêm trọng.

```go
func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("something went terribly wrong")
}

func main() {
    mayPanic()
    fmt.Println("Program continues after panic")
}
```

- **Giải Thích**: 
  - Sử dụng `defer` để gọi một hàm phục hồi khi có `panic`.
  - `recover()` sẽ lấy giá trị của `panic` và tiếp tục thực hiện chương trình sau khi lỗi được phục hồi.

## Tóm Tắt

- **Xử lý lỗi cơ bản**: Sử dụng giá trị lỗi trả về để kiểm tra và xử lý lỗi.
- **Tạo lỗi tùy chỉnh**: Định nghĩa kiểu lỗi mới để cung cấp thông tin chi tiết hơn.
- **`panic` và `recover`**: Dùng cho các tình huống lỗi nghiêm trọng mà không thể xử lý bằng cách thông thường.

# Giải thích block code

```go
package main

import (
    "fmt"
)

// Define a custom error type
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func performTask(success bool) error {
    if !success {
        return &MyError{Code: 500, Message: "Task failed"}
    }
    return nil
}

func main() {
    err := performTask(false)
    if err != nil {
        fmt.Println("Custom Error:", err)
        return
    }
    fmt.Println("Task completed successfully")
}
```

## Giải Thích Cách Khai Báo Hàm

### 1. **Hàm `Error`**

```go
func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
```

- **Khai Báo**:
  - `func` là từ khóa dùng để khai báo hàm.
  - `(e *MyError)` là receiver (người nhận). Đây là kiểu dữ liệu mà phương thức này sẽ hoạt động trên. `e` là tên của receiver, và `*MyError` cho biết phương thức thuộc về kiểu `MyError`.
  - `Error() string` chỉ định rằng đây là một phương thức không nhận tham số đầu vào và trả về một giá trị kiểu `string`.

- **Chức Năng**:
  - Phương thức `Error` là một phương thức của kiểu `MyError`. Nó trả về một chuỗi mô tả lỗi theo định dạng "Error Code: Message" sử dụng `fmt.Sprintf`.

### 2. **Hàm `performTask`**

```go
func performTask(success bool) error {
    if !success {
        return &MyError{Code: 500, Message: "Task failed"}
    }
    return nil
}
```

- **Khai Báo**:
  - `func` là từ khóa khai báo hàm.
  - `performTask(success bool)` là tên hàm và danh sách tham số. Hàm `performTask` nhận một tham số `success` kiểu `bool`.
  - `error` là kiểu dữ liệu trả về của hàm. Đây là kiểu lỗi chuẩn trong Go.

- **Chức Năng**:
  - Hàm `performTask` kiểm tra giá trị của tham số `success`.
  - Nếu `success` là `false`, hàm tạo và trả về một lỗi tùy chỉnh của kiểu `MyError` với mã lỗi 500 và thông điệp "Task failed".
  - Nếu `success` là `true`, hàm trả về `nil`, biểu thị không có lỗi xảy ra.

### 3. **Hàm `main`**

```go
func main() {
    err := performTask(false)
    if err != nil {
        fmt.Println("Custom Error:", err)
        return
    }
    fmt.Println("Task completed successfully")
}
```

- **Khai Báo**:
  - `func` là từ khóa khai báo hàm.
  - `main()` là tên hàm. Hàm `main` là điểm vào của chương trình Go. Đây là nơi chương trình bắt đầu thực thi.
  - Hàm `main` không nhận tham số và không trả về giá trị.

- **Chức Năng**:
  - Hàm `main` gọi hàm `performTask` với giá trị `false` và lưu kết quả trả về vào biến `err`.
  - Nếu `err` không phải là `nil`, điều này có nghĩa là đã xảy ra lỗi, và thông điệp lỗi sẽ được in ra màn hình.
  - Nếu không có lỗi (`err` là `nil`), thông điệp "Task completed successfully" sẽ được in ra.

## Tóm Tắt

- **`func`**: Từ khóa khai báo hàm.
- **Receiver**: Được sử dụng trong phương thức kiểu để xác định kiểu dữ liệu mà phương thức hoạt động trên.
- **Danh Sách Tham Số**: Các tham số hàm được khai báo trong dấu ngoặc đơn sau tên hàm.
- **Kiểu Trả Về**: Xác định kiểu dữ liệu mà hàm sẽ trả về.

# VD Lưu Dữ Liệu vào Cơ Sở Dữ Liệu

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // Sử dụng driver PostgreSQL
)

// Hàm kết nối tới cơ sở dữ liệu
func connectToDB() (*sql.DB, error) {
    connStr := "user=username dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Hàm lưu dữ liệu vào cơ sở dữ liệu
func saveData(db *sql.DB, data string) error {
    query := "INSERT INTO mytable (data) VALUES ($1)"
    _, err := db.Exec(query, data)
    if err != nil {
        return err // Trả về lỗi nếu lưu dữ liệu thất bại
    }
    return nil // Không có lỗi, lưu dữ liệu thành công
}

func main() {
    // Kết nối đến cơ sở dữ liệu
    db, err := connectToDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Dữ liệu cần lưu
    data := "Some data"

    // Thực hiện lưu dữ liệu vào cơ sở dữ liệu
    err = saveData(db, data)
    if err != nil {
        fmt.Println("Error saving data:", err)
        // Xử lý lỗi khi lưu dữ liệu thất bại
        // Có thể là ghi log, thông báo cho người dùng, hoặc thực hiện một thao tác khác
    } else {
        fmt.Println("Data saved successfully!")
        // Thực hiện các thao tác tiếp theo sau khi lưu dữ liệu thành công
    }
}
```

# So sánh cách bắt lỗi trong ruby và go

Trong Ruby, khi gặp lỗi, bạn thường thấy một *exception* được ném ra và có thể bắt nó bằng cách sử dụng `rescue`. Điều này giúp bạn xử lý lỗi một cách rõ ràng trong một khối riêng biệt.

Ngược lại, trong Go, không có cơ chế *exception* như trong Ruby. Thay vào đó, Go sử dụng một cách tiếp cận khác để xử lý lỗi: hầu hết các hàm trong Go sẽ trả về một giá trị lỗi (`error`) nếu có điều gì đó không ổn. Lập trình viên sẽ phải kiểm tra giá trị này và xử lý lỗi ngay tại chỗ.

## Điểm chính:

- **Ruby (Exception Handling)**: Khi có lỗi xảy ra, một exception được ném ra và bạn có thể bắt nó bằng `begin...rescue...end`. Nếu không có `rescue`, chương trình sẽ dừng và hiển thị stack trace của lỗi.
- **Go (Error Handling)**: Go không có khái niệm exception như Ruby. Thay vào đó, các hàm thường trả về một giá trị lỗi cùng với giá trị chính. Bạn phải kiểm tra và xử lý lỗi này sau mỗi lần gọi hàm.

## Ví dụ so sánh:

### Ruby:

```ruby
begin
  # Code có thể gây ra lỗi
  user.save!
rescue ActiveRecord::RecordInvalid => e
  puts "Validation failed: #{e.message}"
end
```

### Go:

```go
err := saveUser(user)
if err != nil {
    fmt.Println("Error saving user:", err)
}
```

## Khác biệt:

- Trong Ruby, lỗi được "quăng" ra khỏi dòng chảy của chương trình và cần phải được bắt lại.
- Trong Go, lỗi là một giá trị mà bạn cần kiểm tra ngay lập tức sau khi thực hiện một thao tác nào đó.

Sự khác biệt này phản ánh triết lý thiết kế của Go, nơi mọi thứ cần rõ ràng và không có bất ngờ (no magic). Điều này giúp mã Go dễ đọc và dễ hiểu hơn, nhưng đòi hỏi lập trình viên phải chủ động hơn trong việc xử lý lỗi.
