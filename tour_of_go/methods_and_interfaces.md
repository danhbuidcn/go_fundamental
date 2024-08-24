# Methods and interfaces

https://go.dev/tour/methods/1

# Methods

## Overview

- Go không có khái niệm lớp (class) như các ngôn ngữ hướng đối tượng khác. Tuy nhiên, bạn có thể định nghĩa các phương thức (method) cho các kiểu (type) trong Go.

- Phương thức là một hàm có một đối số đặc biệt gọi là `receiver`

- `Receiver` xuất hiện trong danh sách đối số riêng của nó, nằm giữa từ khóa `func` và tên của phương thức.
    ```go
    package main

    import (
        "fmt"
        "math"
    )

    // Định nghĩa một kiểu Vertex với các tọa độ X và Y
    type Vertex struct {
        X, Y float64
    }

    // Định nghĩa phương thức Abs với receiver là kiểu Vertex
    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func main() {
        v := Vertex{3, 4}
        fmt.Println(v.Abs()) // Kết quả: 5
    }
    ```

- Điểm quan trọng:
    - **Không có lớp (Class)**: Go không có lớp, nhưng bạn có thể gắn các phương thức vào các kiểu (type), cung cấp cách tổ chức mã tương tự xung quanh các kiểu dữ liệu.
    - **Receiver Argument**: Receiver giống như một instance trong lập trình hướng đối tượng, liên kết phương thức với một kiểu cụ thể.
    - **Tính linh hoạt**: Phương thức có thể được định nghĩa trên bất kỳ kiểu nào, bao gồm các kiểu dữ liệu có sẵn, struct, và thậm chí là con trỏ.
        ```go
        // Định nghĩa kiểu MyInt dựa trên kiểu int
        type MyInt int

        // Định nghĩa một phương thức Double trên kiểu MyInt
        // Đây là phương thức trên kiểu dữ liệu có sẵn
        func (m MyInt) Double() MyInt {
            return m * 2
        }

        func main() {
            var num MyInt = 10
            fmt.Println(num.Double()) // Kết quả: 20
        }
        ```
    - **Không thể khai báo phương thức cho kiểu từ package khác**:
        ```go
        // không hợp lệ và sẽ gây lỗi
        // func (f float64) Abs() float64 {
        //     if f < 0 {
        //         return -f
        //     }
        //     return f
        // }

        func main() {
            f := -math.Sqrt2
            fmt.Println(f.Abs()) // Sẽ gây lỗi biên dịch nếu phương thức trên được khai báo
        }
        ```
    - **Trong một package, bạn có thể khai báo bao nhiêu kiểu tùy thích, miễn là chúng có tên khác nhau.**
    - **Một hàm chỉ có thể có một receiver duy nhất**

## Methods are functions

- Phương thức thực chất là một hàm với một đối số nhận giá trị (receiver argument).

- So sánh func và method:
    - `Function`: Độc lập, không liên kết với một kiểu/ object cụ thể.
    - `Method`: Gắn với một kiểu/ object, thường dùng để thao tác với dữ liệu của kiểu đó.

- Ví dụ, giả sử bạn có một phương thức `Abs` trong Go, được định nghĩa trên một struct `Vertex`. Bạn có thể viết phương thức này dưới dạng một hàm thông thường mà không làm thay đổi chức năng.

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// Phương thức Abs với receiver là Vertex
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs()) // Kết quả: 5
}
```

## Pointer receivers

- Bạn có thể khai báo phương thức (func) với receiver là con trỏ (pointer receivers).

- Điều này có nghĩa là kiểu của receiver có cú pháp `*T` cho một kiểu T nào đó. (Ngoài ra, T không thể tự nó là một con trỏ như `*int`.)

- Giải thích thêm:
    - **receiver là con trỏ (`*T`)** cho phép bạn thay đổi giá trị gốc mà phương thức được gọi trên.
    - **receiver là giá trị (`T`)** chỉ hoạt động trên một bản sao của giá trị, do đó, bất kỳ thay đổi nào được thực hiện bên trong phương thức sẽ không ảnh hưởng đến giá trị gốc.

- Ví dụ:
```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) NoScale(f float64) {
	v.X *= f
	v.Y *= f
	fmt.Println("v in NoSacle function: ", v)
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
	fmt.Println("v in Scale function: ", v)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println("after scale: ", v)

	v.NoScale(5)
	fmt.Println("after no scale: ", v)
}
```

## Pointers and functions

- Khi bạn sử dụng con trỏ làm receiver trong phương thức(method), bạn có thể thay đổi giá trị của đối tượng mà con trỏ trỏ đến.

- Khi chuyển các phương thức(method) này thành các hàm(func), bạn cũng có thể sử dụng con trỏ để đạt được kết quả tương tự.

```go
type Vertex struct {
    X, Y float64
}

// Abs phương thức
func (v *Vertex) Abs() float64 {
    return v.X * v.X + v.Y * v.Y
}

// Scale phương thức
func (v *Vertex) Scale(f float64) {
    v.X *= f
    v.Y *= f
}

// Hàm Abs
func Abs(v Vertex) float64 {
    return v.X * v.X + v.Y * v.Y
}

// Hàm Scale
func Scale(v *Vertex, f float64) {
    v.X *= f
    v.Y *= f
}

func main() {
    v := Vertex{3, 4}
	// hàm
    Scale(&v, 10)  // Cần truyền con trỏ để thay đổi giá trị gốc
    fmt.Println(v) // Output: {30 40}
	fmt.Println("abs: ", Abs(v)) 
	
	v = Vertex{3, 4}
	// phương thức
	v.Scale(10)  // gọi đến method của receiver để thay đổi giá trị gốc
    fmt.Println(v) // Output: {30 40}
	fmt.Println("abs: ", v.Abs())
}
```

## Methods and pointer indirection

### a.So Sánh Giữa Phương Thức và Hàm Với Con Trỏ

```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

#### 1.Phương Thức với Con Trỏ

- **Phương Thức với Con Trỏ Nhận Giá Trị:**
  - Khi bạn định nghĩa một phương thức với con trỏ (`*T`) làm receiver, Go cho phép bạn gọi phương thức đó bằng cả giá trị và con trỏ.
  - Ví dụ:
    ```go
    var v Vertex
    v.Scale(5)  // OK - gọi bằng giá trị
    (&v).Scale(10) // OK - gọi bằng con trỏ
    ```

- **Tự Động Chuyển Đổi:**
  - Khi gọi phương thức trên một giá trị (`v`), Go tự động chuyển đổi thành con trỏ (`&v`) nếu phương thức yêu cầu con trỏ. Điều này giúp việc gọi phương thức trở nên linh hoạt và thuận tiện hơn.

#### 2.Hàm Với Con Trỏ

- **Hàm với Con Trỏ Nhận Giá Trị:**
  - Đối với hàm có tham số là con trỏ, bạn cần phải truyền một con trỏ vào hàm. Hàm không tự động chuyển đổi giá trị thành con trỏ như phương thức.
  - Ví dụ:
    ```go
    var v Vertex
    ScaleFunc(v, 5)  // Lỗi biên dịch!
    ScaleFunc(&v, 5) // Đúng
    ```

#### 3.Tóm Tắt

- **Phương Thức:** Có thể gọi bằng giá trị hoặc con trỏ, Go tự động chuyển giá trị thành con trỏ nếu cần thiết.
- **Hàm:** Yêu cầu truyền tham số chính xác theo kiểu đã định nghĩa, không tự động chuyển đổi giữa giá trị và con trỏ.

### b.So Sánh Giữa Phương Thức và Hàm Với Tham Số Nhận Giá Trị

```go
type Vertex struct {
	X, Y float64
}

// v Vertex là receiver kiểu giá trị.
// Sao chép giá trị của receiver, không thay đổi giá trị gốc.
// Go tự động chuyển đổi con trỏ thành giá trị để gọi phương thức
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
- v *Vertex là receiver kiểu con trỏ.
- Sử dụng con trỏ đến giá trị, cho phép thay đổi giá trị gốc
- Go sẽ tự động chuyển đổi giá trị thành con trỏ để gọi phương thức
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/

// v Vertex là receiver kiểu giá trị.
// Abs func nhận bản sao của giá trị gốc -> các thay đổi không ảnh hưởng đến giá trị gốc
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

#### 1.Hàm Với Tham Số Nhận Giá Trị

- **Hàm Nhận Giá Trị:**
  - Hàm yêu cầu tham số phải khớp chính xác với kiểu đã định nghĩa. Nếu hàm yêu cầu tham số là giá trị, bạn không thể truyền con trỏ cho hàm đó.
  - Ví dụ:
    ```go
    var v Vertex
    fmt.Println(AbsFunc(v))  // OK
    fmt.Println(AbsFunc(&v)) // Lỗi biên dịch!
    ```

#### 2.Phương Thức Với Tham Số Nhận Giá Trị

- **Phương Thức Nhận Giá Trị:**
  - Phương thức với tham số nhận giá trị có thể được gọi bằng cả giá trị hoặc con trỏ. Go tự động chuyển đổi con trỏ thành giá trị nếu phương thức yêu cầu giá trị.
  - Ví dụ:
    ```go
    var v Vertex
    fmt.Println(v.Abs()) // OK
    p := &v
    fmt.Println(p.Abs()) // OK
    ```
  - Trong trường hợp này, `p.Abs()` được hiểu là `(*p).Abs()`, giúp phương thức có thể hoạt động với cả giá trị và con trỏ.

#### 3.Tóm Tắt

- **Hàm:** Yêu cầu tham số phải khớp chính xác với kiểu định nghĩa, không có tự động chuyển đổi giữa giá trị và con trỏ.
- **Phương Thức:** Phương thức với tham số nhận giá trị có thể được gọi bằng cả giá trị hoặc con trỏ, Go tự động chuyển đổi nếu cần thiết.

## Choosing a value or pointer receiver (Receiver Giá Trị và Receiver Con Trỏ)

### 1. Sửa Đổi Giá Trị

- **Receiver Con Trỏ**: Cho phép phương thức sửa đổi giá trị mà receiver trỏ đến.
- **Receiver Giá Trị**: Làm việc với một bản sao của giá trị. Các sửa đổi được thực hiện trên bản sao, do đó giá trị gốc không bị thay đổi.

### 2. Hiệu Suất

- **Receiver Con Trỏ**: Tránh việc sao chép các struct lớn, điều này có thể hiệu quả hơn.
- **Receiver Giá Trị**: Sao chép toàn bộ giá trị, có thể kém hiệu quả hơn đối với các struct lớn.

**Ví Dụ**:

```go
type LargeStruct struct {
    // Nhiều trường
}

func (s *LargeStruct) Process() {
    // Sửa đổi struct
}

func (s LargeStruct) Calculate() int {
    // Sử dụng struct mà không sửa đổi
    return 0
}
```

### 3.Tóm tắt

- Sử dụng **receiver con trỏ** nếu bạn cần sửa đổi receiver hoặc nếu receiver là một struct lớn. 

- Sử dụng **receiver giá trị** nếu phương thức không cần sửa đổi receiver và giá trị là nhỏ hoặc không tốn kém khi sao chép. 

=> Đảm bảo tính đồng nhất trong loại receiver giữa các phương thức cho một kiểu.

# Interfaces

## Định nghĩa

- **Giao Diện (Interfaces)**: Định nghĩa các phương thức mà một giá trị phải thực thi để phù hợp với giao diện đó.
- **Kiểu Con Trỏ và Giá Trị**: Đảm bảo rằng phương thức của giao diện có thể được thực thi bởi kiểu mà bạn sử dụng. Nếu phương thức được định nghĩa trên kiểu con trỏ, bạn cần phải sử dụng kiểu con trỏ để thực thi giao diện.

**Ví Dụ**:

```go
type Abser interface {
    Abs() float64
}

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    var a Abser
    a = Vertex{4, 3}
    fmt.Println(a.Abs())
}
```

Hoặc
```go
type Abser interface {
    Abs() float64
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    var a Abser
    a = &Vertex{4, 3}
    fmt.Println(a.Abs())
}
```


## Interfaces are implemented implicitly (Interfaces Được Thực Hiện Ngầm Định)

- **Thực Hiện Ngầm Định**: Một kiểu thực hiện giao diện nếu nó có các phương thức cần thiết, mà không cần khai báo chính thức.
- **Linh Hoạt**: Điều này cho phép giao diện và các thực hiện của chúng được tách biệt và sử dụng linh hoạt trong các gói khác nhau mà không cần cấu hình trước.
```go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

## Interface values

- Là một cặp `(value, type)`.
- **Thực Thi Phương Thức**: Khi gọi phương thức trên `interface values`, Go thực thi phương thức của kiểu cụ thể chứa trong giao diện.
```go
type I interface {
	M()
}

type T struct {
	S string
}

type F float64

// method of T struct
func (t *T) M() {
	fmt.Println(t.S)
}

// method of F float64
func (f F) M() {
	fmt.Println(f)
}

// function
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i) // print value and type of variable
}

func main() {
	var i I // Initialize interface

	i = &T{"Hello"} // assign value (*T) for interface
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
```

## Interface Values with Nil Underlying Values

- **Giao diện với giá trị cụ thể là `nil`**: Phương thức sẽ được gọi với receiver là `nil`. 
- **Xử lý phương thức với receiver `nil`**: Phương thức nên kiểm tra và xử lý trường hợp receiver `nil` để tránh lỗi và đảm bảo chương trình hoạt động ổn định.

```go
// Định nghĩa giao diện
type I interface {
    M()
}

// Định nghĩa kiểu với phương thức M
type T struct {
    S string
}

// Phương thức M cho kiểu *T
func (t *T) M() {
    if t == nil {
        fmt.Println("Nil receiver")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I // Khai báo giao diện với giá trị nil

    // Gán một con trỏ nil cho giao diện
    i = (*T)(nil)
    i.M() // Gọi phương thức M trên giao diện

    i = &T{"hello"}
	i.M()
}
```

## The empty interface

- **Giao diện Trống**: Là giao diện không có phương thức nào, có thể chứa giá trị của bất kỳ kiểu dữ liệu nào.
- **Ứng Dụng**: Được sử dụng trong các tình huống mà kiểu của giá trị không biết trước, rất hữu ích cho việc xử lý các giá trị có kiểu không xác định và truyền dữ liệu trong các hàm hoặc cấu trúc cần hỗ trợ nhiều loại dữ liệu khác nhau.

```go
package main

import "fmt"

func main() {
    var i interface{}

    fmt.Println(i) // In ra: <nil>

    i = 42          // i có thể chứa kiểu int
    fmt.Println(i) // In ra: 42

    i = []int{1, 2, 3} // i có thể chứa kiểu slice int
    fmt.Println(i) // In ra: [1 2 3]
}
```

## Type assertions

- **Type assertions** là một cơ chế trong Go cho phép bạn truy cập vào giá trị cụ thể mà một giá trị giao diện (interface) đang nắm giữ. Đây là cách để xác định và làm việc với giá trị cụ thể trong một giao diện mà không biết trước kiểu cụ thể của nó.

### Cú Pháp

1. **Type Assertion Đơn Giản**
   ```go
   t := i.(T)
   ```
   - **Ý Nghĩa**: Câu lệnh này kiểm tra xem giá trị giao diện `i` có phải là kiểu `T` không. Nếu đúng, nó gán giá trị kiểu `T` vào biến `t`.
   - **Kết Quả**: Nếu `i` không chứa giá trị kiểu `T`, chương trình sẽ gây ra một panic.

2. **Type Assertion Với Hai Giá Trị**
   ```go
   t, ok := i.(T)
   ```
   - **Ý Nghĩa**: Câu lệnh này kiểm tra xem giá trị giao diện `i` có phải là kiểu `T` không. Nếu đúng, nó gán giá trị kiểu `T` vào biến `t` và `ok` sẽ là `true`. Nếu không, `ok` sẽ là `false` và `t` sẽ là giá trị mặc định của kiểu `T`.
   - **Kết Quả**: Không xảy ra panic nếu `i` không chứa giá trị kiểu `T`.

### Ví Dụ

```go
func main() {
    var i interface{} = "Hello, world!" // Giao diện chứa giá trị kiểu string

    // Kiểm tra kiểu cụ thể của giao diện
    s, ok := i.(string)
    if ok {
        fmt.Println("Giao diện chứa giá trị kiểu string:", s)
    } else {
        fmt.Println("Giao diện không chứa giá trị kiểu string.")
    }

    // Cố gắng lấy giá trị kiểu int từ giao diện chứa giá trị kiểu string
    n, ok := i.(int)
    if ok {
        fmt.Println("Giao diện chứa giá trị kiểu int:", n)
    } else {
        fmt.Println("Giao diện không chứa giá trị kiểu int.")
    }
}
```

## Type switches

- **Type Switch** cho phép bạn kiểm tra kiểu của một giá trị trong giao diện thông qua các trường hợp (cases) khác nhau.
- **Cú pháp**: `switch v := i.(type)` với các trường hợp kiểu cụ thể và một trường hợp `default` cho các kiểu không khớp.
- **Sử Dụng**: Hữu ích khi bạn cần xử lý nhiều loại kiểu khác nhau trong một giao diện mà không cần phải xác định kiểu cụ thể trước.

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Giá trị là kiểu int với giá trị %d\n", v)
    case string:
        fmt.Printf("Giá trị là kiểu string với nội dung %s\n", v)
    case float64:
        fmt.Printf("Giá trị là kiểu float64 với giá trị %f\n", v)
    default:
        fmt.Printf("Giá trị có kiểu không được biết: %T\n", v)
    }
}
```

## Stringer

- **`Stringer`** là một giao diện trong Go giúp các kiểu dữ liệu tự mô tả chính mình dưới dạng chuỗi.
- **Cú pháp**: Giao diện `Stringer` yêu cầu phương thức `String()` trả về một chuỗi mô tả.
    ```go
    type Stringer interface {
        String() string
    }
    ```
- **Sử Dụng**: Khi một kiểu dữ liệu triển khai `Stringer`, gói `fmt` và các gói khác sẽ tự động sử dụng phương thức `String()` để hiển thị giá trị của kiểu đó.

```go
type Point struct {
    X, Y int
}

// Triển khai phương thức String() cho kiểu Point
func (p Point) String() string {
    return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

func main() {
    p := Point{3, 4}
    fmt.Println(p) // In ra: Point(3, 4)
}
```

## Errors

- **Lỗi `nil`**: Chỉ ra rằng thao tác đã thành công.
- **Lỗi không phải `nil`**: Chỉ ra rằng thao tác không thành công và lỗi cung cấp thông tin về sự cố.

- **Phương thức `Error()`**: Phương thức này trả về một chuỗi mô tả lỗi.
    ```go
    type error interface {
        Error() string
    }
    ```
- VD:
    ```go
    package main

    import (
        "fmt"
        "time"
    )
    // Định nghĩa kiểu lỗi tùy chỉnh
    type MyError struct {
        When time.Time // Thời điểm xảy ra lỗi
        What string // Mô tả lỗi
    }
    // Phương thức Error() của MyError để thực hiện giao diện error
    func (e *MyError) Error() string {
        return fmt.Sprintf("at %v, %s", e.When, e.What)
    }
    // Hàm run trả về lỗi tùy chỉnh
    func run() error {
        return &MyError{ time.Now(), "it didn't work" }
    }

    func main() {
        if err := run(); err != nil {
            fmt.Println(err) // In lỗi ra màn hình
        }
    }
    ```

- VD 2:
    ```go
    package main

    import (
        "fmt"
        "math"
    )

    // Định nghĩa kiểu lỗi tùy chỉnh
    type ErrNegativeSqrt float64

    // Cài đặt phương thức Error() cho kiểu ErrNegativeSqrt
    func (e ErrNegativeSqrt) Error() string {
        return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
    }

    // Hàm Sqrt tính căn bậc hai và trả về lỗi nếu đầu vào là số âm
    func Sqrt(x float64) (float64, error) {
        if x < 0 {
            return 0, ErrNegativeSqrt(x) // Trả về lỗi khi x âm
        }
        return math.Sqrt(x), nil // Tính căn bậc hai khi x không âm
    }

    func Calculator(num float64) {
        result, err := Sqrt(num)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println("Sqrt(2) =", result)
    }

    func main() {
        // Kiểm tra hàm Sqrt với số dương
        Calculator(2)

        // Kiểm tra hàm Sqrt với số âm
        Calculator(-2)
    }
    ```

## Readers

- **`io.Reader`** là một interface trong Go dùng để đọc dữ liệu từ các nguồn khác nhau như tệp tin, kết nối mạng, hoặc chuỗi. Nó có phương thức `Read` với cú pháp sau:
    ```go
    func (T) Read(b []byte) (n int, err error)
    ```
    - **Tham số `b`**: Slice byte để lưu dữ liệu đọc được.
    - **Trả về**:
    - `n`: Số byte đã đọc.
    - `err`: Lỗi nếu có, hoặc `io.EOF` khi kết thúc dòng dữ liệu.

- Ví Dụ
    ```go
    package main

    import (
        "fmt"
        "io"
        "strings"
    )

    func main() {
        r := strings.NewReader("Hello, world!") // Tạo một strings.Reader
        buf := make([]byte, 8) // Slice byte với kích thước 8

        for {
            n, err := r.Read(buf) // Đọc dữ liệu vào buf
            if err == io.EOF {
                break // Kết thúc khi đọc xong
            }
            if err != nil {
                fmt.Println("Error:", err)
                break
            }
            fmt.Printf("Read %d bytes: %s\n", n, buf[:n]) // In dữ liệu đã đọc
        }
    }
    ```

- Exercise:
    ```go
    package main

    import (
        "io"
        "os"
        "strings"
    )

    // rot13Reader bao bọc một io.Reader và áp dụng mã hóa ROT13 cho đầu ra của nó.
    type rot13Reader struct {
        r io.Reader
    }

    // Phương thức Read áp dụng ROT13 cho dữ liệu đầu vào.
    func (r *rot13Reader) Read(b []byte) (int, error) {
        n, err := r.r.Read(b) // Đọc dữ liệu từ reader gốc
        if err != nil {
            return n, err
        }

        // Áp dụng mã hóa ROT13
        for i := 0; i < n; i++ {
            if (b[i] >= 'A' && b[i] <= 'Z') {
                b[i] = (b[i]-'A'+13)%26 + 'A'
            } else if (b[i] >= 'a' && b[i] <= 'z') {
                b[i] = (b[i]-'a'+13)%26 + 'a'
            }
        }

        return n, nil
    }

    func main() {
        s := strings.NewReader("Lbh penpxrq gur pbqr!") // "You cracked the code!" bằng ROT13
        r := rot13Reader{s}
        io.Copy(os.Stdout, &r) // Giải mã thông điệp và in ra
    }
    ```

## Interface Image

- [Package `image`](https://pkg.go.dev/image#Image) định nghĩa một interface tên là `Image`. Interface này bao gồm ba phương thức chính:

    1. **ColorModel**: Trả về kiểu màu của hình ảnh dưới dạng `color.Model`. Điều này xác định cách mà màu sắc được biểu diễn trong hình ảnh, chẳng hạn như RGB hoặc RGBA.

    2. **Bounds**: Trả về một `Rectangle` (hình chữ nhật) xác định phạm vi của hình ảnh (tức là kích thước chiều rộng và chiều cao).

    3. **At**: Trả về một giá trị `color.Color` đại diện cho màu sắc tại một tọa độ cụ thể `(x, y)` trong hình ảnh.

- Nếu bạn muốn làm việc với hình ảnh trong Go, bạn có thể sử dụng các gói [`image`, `image/color` ](https://pkg.go.dev/image/color) và các interface này để thao tác và xử lý hình ảnh ở mức độ thấp.
