# Packages, variables, and functions.

https://go.dev/tour/basics/1

## Packages

Mỗi chương trình Go bao gồm các gói (packages). 

Chương trình bắt đầu từ gói `main`. 

Chương trình này sử dụng các gói "fmt" và "math/rand". 

Theo quy ước, tên gói giống như phần cuối của đường dẫn nhập. Ví dụ, gói "math/rand" có tên gói là `rand`.

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Int())
}
```

## Imports

Code này nhóm các import vào một câu lệnh import "factored". Bạn cũng có thể viết nhiều câu lệnh import, nhưng nhóm lại là cách tốt hơn.

```
import (
	"fmt"
	"math"
)
```

## Exported names

Trong Go, tên bắt đầu bằng chữ cái hoa được export và có thể được sử dụng từ các gói khác. Tên bắt đầu bằng chữ cái thường không được export và không thể truy cập từ bên ngoài gói.

Khi import một gói, bạn chỉ có thể sử dụng các tên được export.

## Functions

Trong Go, một hàm có thể không nhận hoặc nhận nhiều đối số.

Lưu ý rằng kiểu dữ liệu được đặt sau tên biến.

[Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)

```go
func add(x, y int) int {
	return x + y
}
// OR
func add(X int, Y int) int {
	return x + y
}
```

## Multiple results

Một hàm trong Go có thể trả về nhiều kết quả.

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

## Named return values

Trong Go, giá trị trả về của hàm có thể được đặt tên và chúng sẽ được coi như các biến định nghĩa ở đầu hàm. 

Những tên này giúp tài liệu hóa ý nghĩa của các giá trị trả về. 

Một câu lệnh `return` không có đối số sẽ trả về các giá trị trả về đã đặt tên. 

Câu lệnh `return` không có đối số (naked return) chỉ nên dùng trong các hàm ngắn để tránh làm giảm tính dễ đọc trong các hàm dài.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

## Variables

Câu lệnh `var` khai báo một danh sách biến, với kiểu dữ liệu xuất hiện sau tên biến. Câu lệnh `var` có thể được đặt ở cấp độ gói hoặc cấp độ hàm.

## Variables with initializers

Một khai báo var có thể bao gồm các bộ khởi tạo, mỗi bộ khởi tạo một biến.

Nếu có giá trị khởi tạo, kiểu dữ liệu có thể bị bỏ qua; biến sẽ nhận kiểu dữ liệu từ giá trị khởi tạo.

## Short variable declarations

Trong hàm, có thể sử dụng câu lệnh gán ngắn `:=` thay cho khai báo biến với kiểu dữ liệu ngầm định (implicit type). 

Ngoài hàm, câu lệnh phải bắt đầu bằng từ khóa (như `var`, `func`), vì vậy không thể sử dụng cú pháp `:=` bên ngoài hàm.

## Basic types

Các loại cơ bản của Go là

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Zero values

Các biến được khai báo mà không có giá trị khởi tạo rõ ràng sẽ được gán giá trị bằng không.

Giá trị bằng không là:

	- 0 đối với kiểu số,
	- false đối với kiểu boolean và
	- "" (chuỗi rỗng) đối với chuỗi.

## Type conversions

Biểu thức T(v) chuyển đổi giá trị v thành loại T.

Một số chuyển đổi số:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
// OR
i := 42
f := float64(i)
u := uint(f)
```

## Type inference

Khi khai báo một biến mà không chỉ định loại rõ ràng (bằng cách sử dụng cú pháp := hoặc cú pháp biểu thức var =), loại của biến được suy ra từ giá trị ở phía bên phải. 

Khi gõ vào phía bên phải của khai báo, biến mới có cùng kiểu:
```go
var i int
j := i // j is an int
```

Nhưng khi phía bên phải chứa hằng số chưa được gõ, biến mới có thể là int, float64 hoặc complex128 tùy thuộc vào độ chính xác của hằng số:
```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## Constants

Các hằng số được khai báo giống như các biến nhưng với từ khóa const.

Các hằng số có thể là giá trị ký tự, chuỗi, boolean hoặc số. 

Các hằng số không thể được khai báo bằng cú pháp :=.

## Numeric Constants

Hằng số là giá trị có độ chính xác cao. 

Một hằng số chưa được gõ sẽ nhận loại mà ngữ cảnh của nó cần. 

Hãy thử in NeedInt(Big). 
(Một int có thể lưu trữ tối đa số nguyên 64 bit và đôi khi ít hơn.)

```go
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```
