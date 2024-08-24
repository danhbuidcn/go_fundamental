# Kiểu Dữ Liệu Trong Go

## Giới Thiệu

Go (hay Golang) là một ngôn ngữ lập trình mạnh mẽ với hệ thống kiểu dữ liệu tĩnh, giúp bạn kiểm soát chính xác các loại dữ liệu mà bạn đang làm việc. Hiểu rõ về các kiểu dữ liệu cơ bản và nâng cao trong Go, cùng với cách khai báo và chuyển đổi giữa các kiểu dữ liệu, là điều rất quan trọng để viết mã nguồn hiệu quả và chính xác.

## Các Kiểu Dữ Liệu Cơ Bản

### 1. **Kiểu Số Nguyên (`int`, `uint`, `byte`, `rune`)**

- **`int`**:
  - **Mô tả**: Là kiểu số nguyên với kích thước phụ thuộc vào hệ thống (32-bit hoặc 64-bit). Được sử dụng phổ biến nhất cho các phép toán số học.
  - **Ứng dụng**: Thích hợp cho các phép toán cần tính toán với số nguyên, ví dụ như đếm số lượng phần tử hoặc tính toán các giá trị không âm.
  - **Ví dụ**:
    ```go
    var age int = 25
    ```

- **`uint`**:
  - **Mô tả**: Là kiểu số nguyên không dấu, kích thước phụ thuộc vào hệ thống. Chỉ lưu trữ giá trị dương và 0.
  - **Ứng dụng**: Sử dụng khi bạn cần số nguyên không âm, chẳng hạn như chỉ số mảng, kích thước của đối tượng hoặc giá trị không thể âm.
  - **Ví dụ**:
    ```go
    var count uint = 100
    ```

- **`byte`**:
  - **Mô tả**: Là alias của `uint8`, lưu trữ giá trị từ 0 đến 255. Thường dùng để xử lý dữ liệu byte.
  - **Ứng dụng**: Thích hợp cho các thao tác với dữ liệu nhị phân hoặc các ký tự ASCII. Ví dụ, lưu trữ dữ liệu từ các tệp hoặc xử lý thông tin mạng.
  - **Ví dụ**:
    ```go
    var b byte = 255
    ```

- **`rune`**:
  - **Mô tả**: Là alias của `int32`, dùng để đại diện cho một ký tự Unicode.
  - **Ứng dụng**: Thích hợp cho việc xử lý các ký tự Unicode trong chuỗi văn bản, đặc biệt là khi bạn cần xử lý văn bản đa ngôn ngữ.
  - **Ví dụ**:
    ```go
    var r rune = 'A'
    ```

### 2. **Kiểu Số Thực (`float32`, `float64`)**

- **`float32`**:
  - **Mô tả**: Kiểu số thực với độ chính xác đơn. Cung cấp khoảng giá trị và độ chính xác thấp hơn.
  - **Ứng dụng**: Phù hợp khi bạn cần tiết kiệm bộ nhớ và không yêu cầu độ chính xác cao, chẳng hạn như trong các ứng dụng đồ họa hoặc tính toán khoa học cơ bản.
  - **Ví dụ**:
    ```go
    var pi float32 = 3.14
    ```

- **`float64`**:
  - **Mô tả**: Kiểu số thực với độ chính xác gấp đôi. Cung cấp độ chính xác cao hơn so với `float32`.
  - **Ứng dụng**: Phù hợp khi bạn cần tính toán với số thực có độ chính xác cao hơn, như trong các ứng dụng tài chính, phân tích số liệu lớn hoặc mô phỏng khoa học.
  - **Ví dụ**:
    ```go
    var e float64 = 2.718281828459
    ```

### 3. **Kiểu Boolean (`bool`)**

- **`bool`**:
  - **Mô tả**: Chỉ có hai giá trị: `true` hoặc `false`.
  - **Ứng dụng**: Thường dùng trong các phép so sánh, điều kiện, và kiểm tra trạng thái. Ví dụ, kiểm tra điều kiện trong cấu trúc điều khiển (`if`, `for`).
  - **Ví dụ**:
    ```go
    var isValid bool = true
    ```

### 4. **Kiểu Chuỗi (`string`)**

- **`string`**:
  - **Mô tả**: Đại diện cho một chuỗi ký tự Unicode.
  - **Ứng dụng**: Dùng để lưu trữ và thao tác với văn bản, chẳng hạn như tên người dùng, thông điệp hoặc dữ liệu văn bản.
  - **Ví dụ**:
    ```go
    var greeting string = "Hello, Go!"
    ```

## Các Kiểu Dữ Liệu Nâng Cao

Go cũng hỗ trợ các kiểu dữ liệu nâng cao cho phép bạn tổ chức và làm việc với dữ liệu một cách linh hoạt hơn.

### 1. **Mảng (`array`)**

- **Mô tả**: Là một tập hợp các phần tử có cùng kiểu dữ liệu với kích thước cố định.
- **Ứng dụng**: Dùng khi bạn biết trước số lượng phần tử và không cần thay đổi kích thước. Thích hợp cho các tác vụ với dữ liệu có kích thước cố định.
- **Khai báo và sử dụng**:
  ```go
  var numbers [5]int
  numbers[0] = 1
  numbers[1] = 2
  fmt.Println("array:", numbers)
  ```
- **Lưu ý**: Kích thước của mảng không thể thay đổi sau khi khai báo.

### 2. **Slice (`slice`)**

- **Mô tả**: 
  - Là một kiểu dữ liệu động, có thể thay đổi kích thước. Slices là một cách tiện lợi hơn để làm việc với tập hợp dữ liệu so với mảng.
  - slice chỉ chứa thông tin về mảng mà nó tham chiếu, bao gồm con trỏ đến phần tử đầu tiên của mảng, độ dài và sức chứa.
- **Ứng dụng**: Thích hợp cho các tác vụ cần thay đổi kích thước động, như xử lý danh sách các mục dữ liệu hoặc tập hợp các đối tượng mà kích thước không được xác định trước.
- **Khai báo và sử dụng**:
  ```go
  var numbers []int = []int{1, 2, 3, 4, 5} // slice number
  fmt.Println("slice:", numbers)
  ```
- **Lưu ý**: Slices có thể thay đổi kích thước và có thể trỏ đến một mảng.

```go
// không cần phải trả về dữ liệu vì các thay đổi được thực hiện trên mảng mà slice tham chiếu sẽ phản ánh ngay lập tức bên ngoài hàm.
func updateSlice(s []int) {
    // thay đổi các giá trị trong mảng mà slice tham chiếu
    s[0] = 10
    s[1] = 20
    s[2] = 30
}

func main() {
    s := []int{1, 2, 3} // slice tham chiếu đến mảng
    fmt.Println("Before update:", s) // [1 2 3]
    updateSlice(s)
    fmt.Println("After update:", s) // [10 20 30]
}
```

### 3. **Struct (`struct`)**

- **Mô tả**: Là một kiểu dữ liệu tùy chỉnh cho phép nhóm nhiều giá trị có các kiểu dữ liệu khác nhau.
- **Ứng dụng**: Thích hợp khi bạn cần nhóm dữ liệu có cấu trúc phức tạp. Ví dụ, nhóm thông tin về một người dùng như tên, tuổi, địa chỉ, v.v.
- **Khai báo và sử dụng**:
  ```go
  type Person struct {
      Name string
      Age  int
  }

  person := Person{Name: "John", Age: 30}
  fmt.Println("struct:", person)
  ```
- **Lưu ý**: Struct rất hữu ích để tổ chức dữ liệu có cấu trúc, và bạn có thể định nghĩa các phương thức cho struct.

### 4. **Map (`map`)**

- **Mô tả**: 
  - Là một kiểu dữ liệu cấu trúc dạng bảng băm (hash table) lưu trữ các cặp khóa-giá trị.
  - Là cấu trúc dữ liệu tham chiếu, nghĩa là khi bạn truyền một map vào hàm, bạn đang truyền một tham chiếu đến map đó.
  - Do đó, bạn không cần sử dụng con trỏ với maps.
- **Ứng dụng**: Thích hợp cho việc tra cứu dữ liệu nhanh chóng dựa trên khóa, chẳng hạn như lưu trữ cấu hình, lập chỉ mục, và xử lý dữ liệu trong các ứng dụng web.
- **Khai báo và sử dụng**:
  ```go
  myMap := make(map[string]int)
  myMap["a"] = 1
  myMap["b"] = 2
  fmt.Println("map:", myMap)
  ```
- **Lưu ý**: Maps rất linh hoạt và nhanh chóng cho việc tra cứu dữ liệu. Tuy nhiên, không có thứ tự cho các phần tử trong map.

```go
package main

import "fmt"

func updateMap(m map[string]int) {
    m["Alice"] = 30
}

func main() {
    m := map[string]int{"Bob": 25}
    fmt.Println("Before update:", m)
    updateMap(m)
    fmt.Println("After update:", m)
}
```

### 5. **Interface (`interface`)**

- **Mô tả**: Là một kiểu dữ liệu trừu tượng cho phép định nghĩa các phương thức mà các kiểu dữ liệu khác phải thực hiện.
- **Ứng dụng**: Thích hợp cho các tình huống cần sử dụng nhiều kiểu dữ liệu khác nhau nhưng cùng thực hiện các phương thức giống nhau. Ví dụ, sử dụng các phương thức của interface trong các thư viện hoặc API.
- **Khai báo và sử dụng**:
  ```go
  type Speaker interface {
      Speak() string
  }

  type Person struct {
      Name string
  }

  func (p Person) Speak() string {
      return "Hello, my name is " + p.Name
  }

  var s Speaker = Person{Name: "John"}
  fmt.Println("interface:", s.Speak())
  ```
- **Lưu ý**: Interfaces cho phép viết mã linh hoạt và mở rộng, với khả năng sử dụng các kiểu dữ liệu khác nhau mà không cần biết rõ kiểu cụ thể.

```go
package main

import "fmt"

type Describer interface {
    Describe()
}

type Person struct {
    Name string
    Age  int
}

func (p *Person) Describe() {
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    var d Describer = &p
    d.Describe()
}
```

## Chuyển Đổi Kiểu Dữ Liệu

Trong Go, **type conversion** (chuyển đổi kiểu dữ liệu) là quá trình chuyển đổi một giá trị từ kiểu dữ liệu này sang kiểu dữ liệu khác

## 1. Chuyển Đổi Giữa Các Kiểu Số

Go hỗ trợ chuyển đổi giữa các kiểu số, chẳng hạn như từ `int` sang `float64`, hoặc từ `float32` sang `int`. Để chuyển đổi kiểu số, bạn có thể sử dụng cú pháp đơn giản là `Type(value)`.

**Ví dụ**:
  ```go
  package main

  import "fmt"

  func main() {
      var intVal int = 42
      var floatVal float64

      // Chuyển đổi int thành float64
      floatVal = float64(intVal)

      fmt.Println("Integer value:", intVal)
      fmt.Println("Float value:", floatVal)

      var anotherFloat float32 = 3.14
      var intFromFloat int

      // Chuyển đổi float32 thành int
      intFromFloat = int(anotherFloat)

      fmt.Println("Float32 value:", anotherFloat)
      fmt.Println("Integer from Float32:", intFromFloat)
  }
  ```

**Lưu ý**:
  - Khi chuyển từ `float` sang `int`, phần thập phân sẽ bị cắt bỏ.
  - Cần cẩn thận khi chuyển đổi giữa các kiểu số để tránh mất dữ liệu hoặc lỗi trong phép toán.

## 2. Chuyển Đổi Giữa Các Kiểu Dữ Liệu Cơ Bản Và `string`

Để chuyển đổi giữa các kiểu dữ liệu cơ bản và kiểu `string`, bạn có thể sử dụng các hàm chuyển đổi tích hợp sẵn như `strconv.Itoa()` để chuyển đổi số nguyên thành chuỗi, hoặc `strconv.ParseInt()` để chuyển đổi chuỗi thành số nguyên.

**Ví dụ**:
  ```go
  package main

  import (
      "fmt"
      "strconv"
  )

  func main() {
      var num int = 123
      var str string

      // Chuyển đổi int thành string
      str = strconv.Itoa(num)

      fmt.Println("Integer:", num)
      fmt.Println("String:", str)

      var strNum string = "456"
      var parsedNum int
      var err error

      // Chuyển đổi string thành int
      parsedNum, err = strconv.Atoi(strNum)
      if err != nil {
          fmt.Println("Error converting string to int:", err)
      } else {
          fmt.Println("Parsed integer:", parsedNum)
      }
  }
  ```

**Lưu ý**:
  - Luôn kiểm tra lỗi khi chuyển đổi từ chuỗi sang số, vì chuỗi có thể không phải là định dạng số hợp lệ.

## 3. Chuyển Đổi Giữa Các Kiểu Dữ Liệu Tuỳ Chỉnh (Custom Types)

Nếu bạn định nghĩa các kiểu dữ liệu của riêng mình thông qua struct hoặc type alias, bạn có thể chuyển đổi giữa các kiểu này nếu các kiểu dữ liệu có thể tương thích.

**Ví dụ**:
  ```go
  package main

  import "fmt"

  type Celsius float64
  type Fahrenheit float64

  func main() {
      var c Celsius = 37.0
      var f Fahrenheit

      // Chuyển đổi Celsius thành Fahrenheit
      f = Fahrenheit(c*9/5 + 32)

      fmt.Printf("Celsius: %.2f\n", c)
      fmt.Printf("Fahrenheit: %.2f\n", f)
  }
  ```

**Lưu ý**:
  - Trong trường hợp chuyển đổi giữa các kiểu dữ liệu tuỳ chỉnh, đảm bảo rằng các phép toán và chuyển đổi đều có ý nghĩa và chính xác.

## Một Số Lưu Ý Khi Chuyển Đổi Kiểu Dữ Liệu

1. **Sự Mất Dữ Liệu**:
  - Chuyển đổi từ kiểu số có độ chính xác cao sang kiểu số có độ chính xác thấp hơn (như từ `float64` sang `float32`) có thể dẫn đến mất dữ liệu.
  - Chuyển đổi từ `float` sang `int` sẽ cắt bỏ phần thập phân.

2. **Kiểm Tra Lỗi**:
  - Khi chuyển đổi giữa kiểu `string` và kiểu số, luôn kiểm tra lỗi để xử lý các trường hợp chuỗi không hợp lệ.

3. **Kiểu Dữ Liệu Tuỳ Chỉnh**:
  - Đối với các kiểu dữ liệu tuỳ chỉnh, đảm bảo rằng các phép toán và chuyển đổi kiểu có nghĩa và chính xác về mặt logic.

