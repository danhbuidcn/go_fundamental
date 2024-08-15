# Biến trong Go

Go (hay còn gọi là Golang) là một ngôn ngữ lập trình được thiết kế để đơn giản, hiệu quả và dễ đọc. Một phần quan trọng trong việc sử dụng Go là hiểu cách làm việc với biến và kiểu dữ liệu. Bài viết này sẽ cung cấp cái nhìn sâu sắc về cách khai báo, phạm vi, và các loại biến trong Go.

## 1. Phạm Vi Của Biến

### 1.1 **Biến Toàn Cục (Global Variables)**
Biến toàn cục được khai báo ngoài tất cả các hàm và có phạm vi toàn bộ package nơi nó được khai báo. Biến toàn cục có thể được truy cập từ bất kỳ hàm nào trong cùng một package.

**Ví dụ**:
```go
package main

import "fmt"

// Biến toàn cục
var globalVar int = 100

func main() {
    fmt.Println("Global variable:", globalVar) // Sử dụng biến toàn cục
}
```

**Lưu ý**:
- Biến toàn cục nên được sử dụng cẩn thận vì chúng có thể gây ra xung đột tên và khó theo dõi khi dự án mở rộng.

### 1.2 **Biến Cục Bộ (Local Variables)**
Biến cục bộ được khai báo bên trong một hàm và chỉ có phạm vi trong hàm đó. Chúng không thể được truy cập từ bên ngoài hàm.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    // Biến cục bộ
    var localVar string = "Hello, Go!"
    fmt.Println("Local variable:", localVar)

    // Biến cục bộ khác
    a := 10
    b := 20
    fmt.Println("Sum:", a + b)
}
```

**Lưu ý**:
- Biến cục bộ giúp mã nguồn trở nên dễ đọc hơn và giảm khả năng gây lỗi bằng cách giới hạn phạm vi của biến.

### 1.3 **Biến Khối (Block Variables)**
Biến khối được khai báo bên trong các khối code như `if`, `for`, hoặc `switch`. Chúng chỉ tồn tại trong phạm vi của khối code đó.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    x := 5

    if x > 0 {
        // Biến khối
        y := 10
        fmt.Println("x + y:", x + y)
    }

    // fmt.Println(y) // Sẽ gây lỗi vì y không tồn tại ngoài khối if
}
```

**Lưu ý**:
- Biến khối có thể giúp tổ chức mã nguồn tốt hơn và làm cho các biến chỉ tồn tại trong phạm vi cần thiết.

## 2. Các Loại Biến và Kiểu Dữ Liệu

### 2.1 **Biến Hằng (Constants)**
Biến hằng được khai báo với từ khóa `const` và có giá trị không thay đổi sau khi được khởi tạo. Hằng có thể là kiểu số, chuỗi hoặc boolean.

**Ví dụ**:
```go
package main

import "fmt"

// Biến hằng
const Pi = 3.14159

func main() {
    const Greeting = "Hello, World!"
    fmt.Println(Greeting)
    fmt.Println("Value of Pi:", Pi)
}
```

**Lưu ý**:
- Hằng giúp đảm bảo rằng các giá trị không bị thay đổi ngẫu nhiên trong quá trình thực thi chương trình.

### 2.2 **Biến Con Trỏ (Pointer Variables)**
Con trỏ là biến lưu trữ địa chỉ của biến khác. Con trỏ cho phép truy cập và thay đổi giá trị của biến thông qua địa chỉ của nó.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    var a int = 10
    var ptr *int

    ptr = &a  // Lưu trữ địa chỉ của a trong con trỏ ptr
    fmt.Println("Value of a:", a)
    fmt.Println("Address of a:", ptr)
    fmt.Println("Value pointed by ptr:", *ptr)  // Truy cập giá trị thông qua con trỏ
}
```

**Lưu ý**:
- Con trỏ cho phép thay đổi giá trị của biến thông qua địa chỉ của nó, điều này hữu ích trong các tình huống cần sửa đổi giá trị từ nhiều nơi.

### 2.3 **Biến Slice**
Slice là một cấu trúc dữ liệu linh hoạt dùng để lưu trữ các giá trị cùng loại. Độ dài của slice có thể thay đổi trong quá trình thực thi.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    // Tạo slice từ mảng
    var numbers = []int{2, 3, 4, 5}

    fmt.Println("Numbers:", numbers)
    fmt.Println("Length of slice:", len(numbers))
    fmt.Println("Capacity of slice:", cap(numbers))

    // Thêm phần tử vào slice
    numbers = append(numbers, 6)
    fmt.Println("After append:", numbers)
}
```

**Lưu ý**:
- Slice có thể mở rộng và thu hẹp kích thước của nó, điều này làm cho nó linh hoạt hơn so với mảng.

### 2.4 **Biến Map**
Map là một cấu trúc dữ liệu lưu trữ các cặp khóa - giá trị. Mỗi khóa là duy nhất và ánh xạ đến một giá trị cụ thể.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    // Tạo một map
    var capitals = map[string]string{
        "France": "Paris",
        "Japan": "Tokyo",
    }

    fmt.Println("Map:", capitals)
    fmt.Println("Capital of Japan:", capitals["Japan"])

    // Thêm một cặp khóa - giá trị mới
    capitals["Vietnam"] = "Hanoi"
    fmt.Println("Updated Map:", capitals)
}
```

**Lưu ý**:
- Map cho phép tìm kiếm, thêm và xóa các cặp khóa - giá trị nhanh chóng, nhưng không duy trì thứ tự của các phần tử.

### 2.5 **Biến Struct**
Struct là một kiểu dữ liệu do người dùng định nghĩa, bao gồm nhiều trường dữ liệu khác nhau. Struct cho phép nhóm các biến lại với nhau thành một kiểu dữ liệu duy nhất.

**Ví dụ**:
```go
package main

import "fmt"

// Định nghĩa một struct
type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    // Tạo một instance của struct
    var person = Person{Name: "Alice", Age: 30, City: "New York"}

    fmt.Println("Person:", person)
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("City:", person.City)
}
```

**Lưu ý**:
- Struct giúp nhóm dữ liệu có liên quan lại với nhau, làm cho mã nguồn dễ hiểu và quản lý hơn.

### 2.6 **Biến Interface**
Interface định nghĩa một tập hợp các phương thức mà các kiểu dữ liệu phải thực thi. Interface cho phép một biến chứa bất kỳ giá trị nào thỏa mãn các phương thức mà nó định nghĩa.

**Ví dụ**:
```go
package main

import "fmt"

// Định nghĩa interface
type Shape interface {
    Area() float64
}

// Struct thực thi interface
type Circle struct {
    Radius float64
}

// Phương thức Area của Circle thực thi interface Shape
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape = Circle{Radius: 5}
    fmt.Println("Area of circle:", s.Area())
}
```

**Lưu ý**:
- Interface là công cụ mạnh mẽ để đạt được tính trừu tượng và đa hình trong Go, giúp mã nguồn trở nên linh hoạt và dễ bảo trì.

## 3. Cách Khai Báo Biến

### 3.1 **Khai Báo Biến với `var`**
Từ khóa `var` được dùng để khai báo biến và có thể khởi tạo giá trị cho biến.

**Ví dụ**:
```go
package main

import "fmt"

// Khai báo biến với từ khóa var
var x int = 10
var y = 20 // Từ khóa var có thể suy diễn kiểu dữ liệu

func main() {
    fmt.Println("x:", x)
    fmt.Println("y:", y)
}
```

**Lưu ý**:
- Bạn có thể khai báo nhiều biến cùng một lúc với từ khóa `var`.

### 3.2 **Khai Báo Biến với `:=`**
Từ khóa `:=` được dùng để khai báo và khởi tạo biến mới trong phạm vi của một hàm. Đây là cách khai báo ngắn gọn và không yêu cầu bạn chỉ định kiểu dữ liệu.

**Ví dụ**:
```go
package main

import "fmt"

func main() {
    // Khai báo biến với := 
    x := 10
   

 y := "Hello, Go!"
    fmt.Println("x:", x)
    fmt.Println("y:", y)
}
```

**Lưu ý**:
- `:=` chỉ có thể được sử dụng trong phạm vi hàm và không thể dùng để khai báo biến toàn cục.

## 4. Các Lưu Ý Khi Sử Dụng Biến

1. **Phạm Vi Biến**:
   - Hiểu rõ phạm vi của biến để tránh xung đột và lỗi không mong muốn.
   - Sử dụng biến cục bộ và biến khối để quản lý phạm vi và bảo vệ dữ liệu.

2. **Kiểu Dữ Liệu**:
   - Lựa chọn kiểu dữ liệu phù hợp với nhu cầu của ứng dụng.
   - Kiểm tra và xử lý lỗi khi làm việc với con trỏ để tránh lỗi thời gian chạy.

3. **Hiệu Suất**:
   - Sử dụng struct và interface một cách hiệu quả để tổ chức dữ liệu và hành vi.
   - Cân nhắc khi sử dụng map và slice, đặc biệt là khi làm việc với dữ liệu lớn hoặc yêu cầu hiệu suất cao.

4. **Tính Khả Đọc**:
   - Đặt tên biến rõ ràng và có ý nghĩa để mã nguồn dễ đọc và bảo trì.
   - Tạo cấu trúc dữ liệu hợp lý để tăng cường khả năng hiểu biết về mã nguồn.

# Tìm hiểu thêm

https://kiendinh.space/golang-slice-tat-tan-tat-va-nhung-dieu-co-the-ban-chua-biet/
