# Methods and interfaces

https://go.dev/tour/methods/1

## Type parameters

- Go cho phép viết các hàm có thể hoạt động trên nhiều kiểu dữ liệu khác nhau bằng cách sử dụng tham số kiểu. Tính năng này giúp tạo ra mã linh hoạt và tái sử dụng hơn.

- Cú pháp:
    ```go
    func Index[T comparable](s []T, x T) int
    ```

    - **Tham số kiểu `T`**: Hàm `Index` có một tham số kiểu `T`, được chỉ định trong dấu ngoặc vuông `[T]` trước các tham số của hàm.

    - **Ràng buộc `comparable`**: Tham số kiểu `T` bị ràng buộc bởi `comparable`, nghĩa là kiểu `T` có thể sử dụng các toán tử so sánh như `==` và `!=`.

    - **Tham số của hàm**:
    - `s []T`: `s` là một slice chứa các phần tử có kiểu `T`.
    - `x T`: `x` là một giá trị có kiểu `T`.
- VD:
    ```go
    package main

    import "fmt"

    func Index[T comparable](s []T, x T) int {
        for i, v := range s {
            if v == x {
                return i
            }
        }
        return -1
    }

    func main() {
        ints := []int{10, 20, 30, 40}
        strings := []string{"apple", "banana", "cherry"}

        fmt.Println(Index(ints, 30))    // Kết quả: 2
        fmt.Println(Index(strings, "banana")) // Kết quả: 1
    }
    ```

## Generic Types

- Go hỗ trợ kiểu dữ liệu tổng quát, cho phép bạn định nghĩa các cấu trúc dữ liệu có thể hoạt động với nhiều kiểu khác nhau. Ví dụ:
    - **Danh sách liên kết đơn giản**: Bạn có thể tạo một danh sách liên kết mà mỗi phần tử có thể là bất kỳ kiểu dữ liệu nào (như `int`, `string`,...).
    - **Tham số kiểu**: Được sử dụng để xác định kiểu dữ liệu mà danh sách sẽ chứa.
- Ví dụ:
    ```go
    type Node[T any] struct {
        value T
        next  *Node[T]
    }

    type List[T any] struct {
        head *Node[T]
    }
    ```
    - **Node[T]**: Đại diện cho một nút trong danh sách, chứa giá trị `value` và con trỏ `next`.
    - **List[T]**: Đại diện cho danh sách liên kết, giữ nút đầu tiên (`head`). 
