# More types: structs, slices, and maps.

[moretypes](https://go.dev/tour/moretypes/1)

[pointers](https://kungfutech.edu.vn/bai-viet/go/con-tro-va-gia-tri-trong-go)

## Pointers

Go có con trỏ, dùng để lưu địa chỉ bộ nhớ của một giá trị.

Kiểu `*T` là con trỏ trỏ tới giá trị `T`, với giá trị mặc định là `nil`.
	```go
	var p *int
	```

Toán tử `&` tạo ra con trỏ tới toán hạng của nó, còn toán tử `*` truy xuất giá trị mà con trỏ trỏ tới.
	```go
	i := 42
	p = &i
	```

Toán tử * biểu thị giá trị cơ bản của con trỏ.
	```go
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	```
Điều này được gọi là "dereferencing"(giải tham chiếu) hoặc "indirecting".

Không giống như C, Go không hỗ trợ con trỏ toán học.

## Structs

- `struct` là tập hợp của fields

```go
type Vertex struct {
	X int
	Y int
}
```

### `Struct fields` 

- được truy cập bằng dấu chấm.

```go
v := Vertex{1, 2}
v.X = 4
```

### Pointers to structs

- Các trường của struct có thể được truy cập thông qua con trỏ struct.
- Thay vì phải viết `(*p).X` để truy cập trường `X` của một struct thông qua con trỏ `p`, Go cho phép bạn đơn giản hóa bằng cách chỉ cần viết `p.X` mà không cần giải tham chiếu rõ ràng.

### Struct Literals

- Một struct literal biểu thị một giá trị struct mới được cấp phát bằng cách liệt kê các giá trị của các fields.
- Bạn có thể liệt kê chỉ một phần của các fields bằng cách sử dụng cú pháp `Name:`. (Và thứ tự của các named fields là không quan trọng.)
- Tiền tố đặc biệt `&` trả về một pointer tới giá trị struct.

```go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // struct literal. has type Vertex
	v2 = Vertex{X: 1}  // struct literal. Y:0 is implicit
	v3 = Vertex{}      // struct literal. X:0 and Y:0
	p  = &Vertex{1, 2} // struct literal. has type *Vertex
)
/*
- v1,v2,v3: biến có kiểu dữ liệu Vertex (object trực tiếp), lưu trữ trực tiếp giá trị của struct.
- p: pointer trỏ đến object kiểu Vertes, lưu trữ địa chỉ của Verter trong bộ nhớ.
*/
```

## Arrays

Kiểu `[n]T` là một mảng có `n` giá trị thuộc kiểu `T`.

Ví dụ, biểu thức:
	```go
	var a [10]int
	```
khai báo biến `a` là một mảng gồm 10 số nguyên.

Chiều dài của mảng là một phần của kiểu dữ liệu, do đó mảng không thể thay đổi kích thước sau khi được tạo. Điều này có vẻ hạn chế, nhưng Go cung cấp một cách tiện lợi hơn để làm việc với mảng thông qua slices.

## Slices

- [Slices: usage and internals](https://go.dev/blog/slices-intro)

- Mảng(Array) có kích thước cố định. 
- Mặt khác, slices là chế độ xem có kích thước động, linh hoạt vào các phần tử của mảng. Trong thực tế, slices phổ biến hơn nhiều so với mảng.

- Khai báo Slice:
    - Kiểu: []T là một slice với các phần tử kiểu T.
    - Tạo Slice: Sử dụng cú pháp a[low:high], với low và high là chỉ số bắt đầu và kết thúc (không bao gồm high).
    - Ví dụ: a[1:4] chọn các phần tử từ chỉ số 1 đến 3 của mảng a.
	```go
	func main() {
		primes := [6]int{2, 3, 5, 7, 11, 13}

		var s []int = primes[1:4]
		fmt.Println(s)
	}
	```

### `Slices as References to Arrays`(slice là tham chiếu đến mảng):

- Đặc điểm: Slices không lưu trữ dữ liệu mà chỉ mô tả một phần của mảng cơ sở (underlying array).

- Thay đổi dữ liệu: Thay đổi các phần tử của slice sẽ thay đổi các phần tử tương ứng trong mảng cơ sở.

- Ảnh hưởng đến các slice khác: Những slice khác chia sẻ cùng một mảng cơ sở sẽ thấy các thay đổi này.

```go
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

### `Slice Literal`

- Tương tự như array literal nhưng không có độ dài.

  - Array Literal: `[3]bool{true, true, false}`

  - Slice Literal: `[]bool{true, true, false}`

  - Slice literal tạo ra mảng tương tự như array literal nhưng tạo một slice tham chiếu đến mảng đó.

### So sánh slice và slice literal:

- `Slice Literals`: Khai báo một slice với các giá trị cụ thể ngay lập tức, ví dụ: `[]int{1, 2, 3}`. Không cần chỉ định kích thước.

- `Slice`: Một loại dữ liệu mô tả một phần của mảng, có thể được tạo từ array literals, slice literals, hoặc từ mảng/slice hiện có, ví dụ: `s := []int{1, 2, 3}`.

### `Slice defaults`

- Khi cắt (slicing) một mảng trong Go, bạn có thể bỏ qua chỉ số thấp hoặc cao để sử dụng các giá trị mặc định. 

	- Chỉ số thấp (low bound): Mặc định là `0` nếu không được chỉ định.

	- Chỉ số cao (high bound): Mặc định là chiều dài của mảng nếu không được chỉ định.

- Với mảng `var a [10]int`, các biểu thức cắt sau là tương đương:

	- `a[0:10]` – Cắt từ chỉ số `0` đến `10` (bao gồm từ `a[0]` đến `a[9]`).
	- `a[:10]` – Cắt từ chỉ số `0` đến `10` (mặc định bắt đầu từ `0`).
	- `a[0:]` – Cắt từ chỉ số `0` đến cuối mảng (mặc định kết thúc ở chiều dài của mảng).
	- `a[:]` – Cắt toàn bộ mảng (mặc định bắt đầu từ `0` đến chiều dài của mảng).

```go
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

### `Slice length and capacity`(Kích thước và dung lượng của Slice)

- Kích thước (length): Số lượng phần tử trong slice - `len(s)`.

- Dung lượng (capacity): Số lượng phần tử của mảng cơ sở từ phần tử đầu tiên trong slice - `cap(s)`.

- Bạn có thể mở rộng kích thước của một slice bằng cách cắt lại nó, miễn là nó có đủ dung lượng.

```go
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	a := s[:0]
	printSlice(a)

	// Extend its length.
	b := s[:4]
	printSlice(b)

	// Drop its first two values.
	c := s[2:]
	printSlice(c)
}
```

### `Slice nil`

- Giá trị mặc định của một slice là `nil`.

- Một slice nil có kích thước và dung lượng bằng 0 và không có mảng cơ sở.

- Tạo slice với `make`:
	- `make([]int, 5)`: Tạo một slice với kích thước 5, giá trị khởi tạo là 0.
	- `make([]int, 0, 5)`: Tạo một slice với kích thước 0 và dung lượng 5.
	- `b = b[:cap(b)]`: Mở rộng slice đến dung lượng tối đa của nó.
	- `b = b[1:]`: Cắt slice, làm giảm kích thước và dung lượng của nó.

```go
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
```

### `Slices of slices`

-Slices có thể chứa bất kỳ kiểu dữ liệu nào, bao gồm cả các slice khác.

```go
func main() {
	// Tạo một slice chứa các slice khác
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:", matrix)
}
```

### `Appending to a slice` (Thêm phần tử vào slice):

- Trong Go, việc thêm các phần tử mới vào một slice là rất phổ biến, và ngôn ngữ cung cấp hàm `append` để thực hiện điều này.

- Cú pháp:
	```go
	func append(s []T, vs ...T) []T
	// s: Slice gốc của kiểu dữ liệu `T`.
	// vs: Các giá trị kiểu `T` muốn thêm vào slice.
	```

- Hàm `append` trả về một slice mới chứa tất cả các phần tử của slice gốc cộng với các giá trị mới được thêm vào.

- Nếu mảng sao lưu (backing array) của slice `s` không đủ lớn để chứa tất cả các giá trị mới, một mảng lớn hơn sẽ được cấp phát. Slice trả về sẽ trỏ tới mảng mới được cấp phát.
	```go
	func main() {
		// Tạo một slice ban đầu
		s := []int{1, 2, 3}

		// Thêm các phần tử vào slice
		s = append(s, 4, 5, 6)

		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // Output: len=6 cap=6 [1 2 3 4 5 6]
	}
	```

### `Range`

- Trong Go, từ khóa `range` được sử dụng để lặp qua một slice hoặc map.
- Khi lặp qua một slice, `range` trả về hai giá trị cho mỗi lần lặp: chỉ số (index) và một bản sao của phần tử tại chỉ số đó.
	```go
	func main() {
		s := []int{10, 20, 30}
		
		// Lặp qua slice
		for i, v := range s {
			fmt.Printf("Index: %d, Value: %d\n", i, v)
		}
	}
	/*
	Index: 0, Value: 10
	Index: 1, Value: 20
	Index: 2, Value: 30 
	*/
	```

- Trong Go, bạn có thể bỏ qua chỉ số (index) hoặc giá trị bằng cách gán chúng cho `_`.
	```go
	func main() {
		pow := []int{2, 4, 8, 16}

		// Chỉ lấy chỉ số
		for i := range pow {
			fmt.Println("Index:", i)
		}

		// Chỉ lấy giá trị
		for _, value := range pow {
			fmt.Println("Value:", value)
		}
	}

	/*
	Index: 0
	Index: 1
	Index: 2
	Index: 3
	Value: 2
	Value: 4
	Value: 8
	Value: 16
	*/
	```

## Maps

[strings.Fields](https://pkg.go.dev/strings#Fields)

- Maps là một tập hợp các cặp key-value, trong đó mỗi key là duy nhất và được liên kết với một value

- Một map `nil` không chứa khóa nào và không thể thêm khóa.
- Hàm `make` tạo ra một map với kiểu dữ liệu được chỉ định, được khởi tạo và sẵn sàng sử dụng.
	```go
	type Vertex struct {
		Lat, Long float64
	}

	func main() {
		m := make(map[string]Vertex)
		m["Bell Labs"] = Vertex{40.68433, -74.39967}

		fmt.Println(m["Bell Labs"])
	}
	```

### `Map literals`

- khai báo và khởi tạo một map với các cặp key-value được định nghĩa.
	```go
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{ 40.68433, -74.39967 },
		"Google": Vertex{ 37.42202, -122.08408 },
	}

	func main() {
		fmt.Println(m)
	}
	```

- Khi kiểu dữ liệu cấp cao nhất đã được xác định rõ, Go cho phép bạn bỏ qua việc chỉ định lại kiểu cho từng phần tử trong map literal.
	```go
	func main() {
		// Tạo một map literal với kiểu dữ liệu cấp cao nhất là map[string]int
		scores := map[string]int{
			"Alice": 10,
			"Bob":   15,
			"Eve":   12,
		}

		// In map ra màn hình
		fmt.Println(scores)
	}
	```

### `Mutating Maps` (biến đổi maps):

- Thêm hoặc cập nhật một phần tử trong map `m`:
	```go
	m[key] = elem
  	```

- Truy xuất một phần tử:
	```go
	elem = m[key]
	```

- Xóa một phần tử:
	```go
	delete(m, key)
	```

- Kiểm tra sự tồn tại của một khóa với phép gán hai giá trị:
	```go
	elem, ok = m[key]
	/*
	- Nếu `key` có trong `m`, `ok` sẽ là `true`.
	- Nếu không, `ok` là `false` và `elem` sẽ nhận giá trị zero cho kiểu dữ liệu của phần tử trong map.
	*/
	```

- Lưu ý: Nếu `elem` hoặc `ok` chưa được khai báo, bạn có thể sử dụng cách khai báo ngắn gọn:
	```go
	elem, ok := m[key]
	```

- VD:
	```go
	func main() {
		// Tạo một map
		scores := make(map[string]int)

		// Thêm phần tử vào map
		scores["Alice"] = 10

		// Cập nhật phần tử
		scores["Alice"] = 15

		// Truy xuất phần tử
		score := scores["Alice"]
		fmt.Println("Alice's score:", score)

		// Xóa phần tử
		delete(scores, "Alice")

		// Kiểm tra sự tồn tại của khóa
		score, exists := scores["Alice"]
		fmt.Printf("Alice's score: %d , exists: %v", score, exists)
	}
	```

## Function values

- Hàm cũng là một loại giá trị trong Go, có thể được truyền như các giá trị khác.

- Giá trị hàm có thể được sử dụng làm đối số của hàm khác hoặc làm giá trị trả về.
	```go
	// Định nghĩa một hàm nhận một hàm khác làm đối số
	func applyOperation(x, y int, operation func(int, int) int) int {
		return operation(x, y)
	}

	// Các hàm đơn giản
	func add(a, b int) int {
		return a + b
	}

	func multiply(a, b int) int {
		return a * b
	}

	func main() {
		// Truyền hàm add và multiply làm đối số
		result1 := applyOperation(5, 3, add)
		fmt.Println("5 + 3 =", result1)

		result2 := applyOperation(5, 3, multiply)
		fmt.Println("5 * 3 =", result2)

		// Lưu một hàm vào biến
		subtract := func(a, b int) int {
			return a - b
		}

		// Sử dụng biến hàm
		result3 := applyOperation(5, 3, subtract)
		fmt.Println("5 - 3 =", result3)
	}
	```

### Function closures

- `closure` là một hàm có thể tham chiếu đến các biến được khai báo bên ngoài phạm vi của chính nó. 

- Điều này có nghĩa là closure có thể truy cập và thay đổi các biến này, ngay cả khi hàm đã thoát khỏi phạm vi mà nó được tạo ra.
	
- VD:
	```go
	func adder() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	func main() {
		pos, neg := adder(), adder()
		fmt.Println(pos(1))  // 1
		fmt.Println(pos(2))  // 3
		fmt.Println(pos(3))  // 6
		fmt.Println(neg(-2)) // -2
		fmt.Println(neg(-4)) // -6
	}
	```

- Phân tích:
	- `adder()`: Hàm này trả về một closure - một hàm ẩn danh mà có thể truy cập biến `sum`.
	- `sum`: Biến này được khai báo trong hàm `adder` và được lưu giữ trong closure. Mỗi lần hàm ẩn danh được gọi, nó có thể cập nhật giá trị của `sum`.
	- `pos` và `neg`: Đây là hai closure khác nhau, cả hai đều được tạo từ `adder()`. Mỗi closure có biến `sum` riêng của nó.

- Hoạt động của Closure:
	- Khi gọi `pos(1)`, `pos(2)`, và `pos(3)`, closure được lưu trong `pos` sử dụng và cập nhật biến `sum` cục bộ của nó.
	- Khi gọi `neg(-2)` và `neg(-4)`, closure trong `neg` cũng hoạt động tương tự nhưng với biến `sum` khác.

- Ý nghĩa:
	- Closure cho phép bạn tạo ra các hàm có trạng thái cục bộ mà không cần sử dụng các biến toàn cục hoặc trả về các giá trị phức tạp. 
	- Điều này rất hữu ích trong việc tạo ra các hàm linh hoạt và có khả năng lưu trữ trạng thái qua các lần gọi.

- VD: viết chương trình tính dãy fibonacci
	```go
	func fibonacci() func() int {
		a, b := 0, 1
		
		return func() int {
			result := a
			a,b = b, b+a
			return result
		}
	}

	func main() {
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Printf("%d, ", f())
		}
	}
	```
