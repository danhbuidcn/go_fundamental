# Cách Viết Mã Go

https://go.dev/doc/code

## Mục Lục
- [Giới thiệu](#giới-thiệu)
- [Tổ chức mã](#tổ-chức-mã)
- [Chương trình đầu tiên của bạn](#chương-trình-đầu-tiên-của-bạn)
  - [Nhập các package từ module của bạn](#nhập-các-package-từ-module-của-bạn)
  - [Nhập các package từ module từ xa](#nhập-các-package-từ-module-từ-xa)
- [Kiểm thử](#kiểm-thử)
- [Tiếp theo là gì](#tiếp-theo-là-gì)

## Giới thiệu

Tài liệu này trình bày cách phát triển một package Go đơn giản trong một module và giới thiệu công cụ `go`, cách tiêu chuẩn để lấy, xây dựng và cài đặt các module, package, và lệnh Go.

## Tổ chức mã

- **Package**:
  - Các chương trình Go được tổ chức thành các **package**. Một package là tập hợp các tệp mã nguồn (file `.go`) nằm trong cùng một thư mục. Những tệp này được biên dịch cùng nhau để tạo thành một phần của chương trình.
  - Các hàm, kiểu dữ liệu, biến và hằng số được định nghĩa trong một tệp nguồn của package có thể được sử dụng từ tất cả các tệp khác trong cùng package đó. Ví dụ, nếu bạn có một hàm trong tệp `a.go`, bạn có thể gọi hàm đó từ tệp `b.go` trong cùng thư mục.

- **Module**:
  - Một **module** là tập hợp các package Go có liên quan và được phát hành cùng nhau. Một kho lưu trữ (repository) thường chứa một module. 
  - Trong kho lưu trữ, module thường nằm ở thư mục gốc và được khai báo trong tệp `go.mod`. Tệp `go.mod` định nghĩa đường dẫn của module, tức là tên của module và được dùng làm tiền tố cho các đường dẫn import trong các package.
  - Module chứa các package không chỉ trong thư mục chứa tệp `go.mod`, mà còn trong các thư mục con của nó. Đến khi gặp một thư mục con có tệp `go.mod` khác, module sẽ dừng lại.

- **Kho lưu trữ**:
  - Bạn không cần phải xuất bản mã của mình lên kho lưu trữ từ xa (như GitHub) để có thể biên dịch nó. Một module có thể tồn tại cục bộ trên máy tính của bạn mà không cần phải thuộc về một kho lưu trữ từ xa. Tuy nhiên, việc tổ chức mã như thể bạn sẽ xuất bản nó trong tương lai là một thói quen tốt.

- **Đường dẫn import**:
  - Đường dẫn import là một chuỗi ký tự dùng để chỉ định package khi bạn sử dụng lệnh import trong mã nguồn Go.
  - Đường dẫn import của một package được cấu thành từ đường dẫn của module và tên thư mục chứa package đó. Ví dụ, nếu module của bạn là `github.com/google/go-cmp` và package nằm trong thư mục `cmp`, thì đường dẫn import của package đó là `github.com/google/go-cmp/cmp`.
  - Các package trong thư viện chuẩn của Go không có đường dẫn module; chúng chỉ sử dụng tên gốc của package, chẳng hạn như `fmt` hoặc `net/http`.

## Chương trình đầu tiên của bạn

Để biên dịch và chạy một chương trình đơn giản, trước tiên hãy chọn một đường dẫn module (chúng ta sẽ sử dụng `example/user/hello`) và tạo một tệp `go.mod` để khai báo nó:

```bash
$ mkdir hello # Hoặc, clone nếu nó đã tồn tại trong hệ thống kiểm soát phiên bản.
$ cd hello
$ go mod init example/user/hello # khởi tạo một module Go mới với tên được chỉ định
go: creating new go.mod: module example/user/hello
$ cat go.mod
module example/user/hello

go 1.16
$
```

Câu lệnh đầu tiên trong tệp nguồn Go phải là tên package. Các lệnh thực thi phải luôn sử dụng package `main`.

Tiếp theo, tạo một tệp có tên `hello.go` bên trong thư mục đó với nội dung mã Go sau:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world !")
}
```

Bây giờ bạn có thể xây dựng và cài đặt chương trình đó với công cụ `go`:

```bash
$ go install example/user/hello # biên dịch và cài đặt tệp thực thi vào thư mục được chỉ định bởi biến môi trường GOBIN hoặc GOPATH/bin
```

Lệnh này sẽ xây dựng lệnh `hello`, tạo ra một tệp nhị phân thực thi. Sau đó, nó sẽ cài đặt tệp nhị phân đó vào `/go/bin/hello` (hoặc, trên Windows, `%USERPROFILE%\go\bin\hello.exe`).

Thư mục cài đặt được điều khiển bởi các biến môi trường `GOPATH` và `GOBIN`. Nếu `GOBIN` được thiết lập, các tệp nhị phân sẽ được cài đặt vào thư mục đó. Nếu `GOPATH` được thiết lập, các tệp nhị phân sẽ được cài đặt vào thư mục con `bin` của thư mục đầu tiên trong danh sách `GOPATH`. Nếu không, các tệp nhị phân sẽ được cài đặt vào thư mục con `bin` của `GOPATH` mặc định (`$HOME/go` hoặc `%USERPROFILE%\go`).

Bạn có thể sử dụng lệnh `go env` để đặt giá trị mặc định cho một biến môi trường cho các lệnh `go` trong tương lai:

```bash
$ go env -w GOBIN=/somewhere/else/bin
$
```

Để xóa một biến đã được thiết lập trước đó bằng `go env -w`, sử dụng `go env -u`:

```bash
$ go env -u GOBIN
$
```

Các lệnh như `go install` hoạt động trong ngữ cảnh của module chứa thư mục làm việc hiện tại. Nếu thư mục làm việc không nằm trong module `example/user/hello`, lệnh `go install` có thể thất bại.

Để tiện lợi, các lệnh `go` chấp nhận các đường dẫn tương đối với thư mục làm việc, và mặc định đến package trong thư mục làm việc hiện tại nếu không có đường dẫn nào khác được cung cấp. Vì vậy, trong thư mục làm việc của chúng ta, các lệnh sau đây đều tương đương:

```bash
$ go install example/user/hello
$ go install .
$ go install
```

Tiếp theo, hãy chạy chương trình để đảm bảo nó hoạt động. Để tiện lợi hơn, chúng ta sẽ thêm thư mục cài đặt vào `PATH` để dễ dàng chạy các tệp nhị phân:

```bash
# Người dùng Windows nên tham khảo /wiki/SettingGOPATH
# để thiết lập %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world !
$
```

Nếu bạn đang sử dụng hệ thống kiểm soát nguồn, bây giờ là thời điểm thích hợp để khởi tạo một kho lưu trữ, thêm các tệp và commit thay đổi đầu tiên của bạn. Một lần nữa, bước này là tùy chọn: bạn không cần sử dụng kiểm soát nguồn để viết mã Go.

```bash
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```

Lệnh `go` xác định kho lưu trữ chứa một đường dẫn module cụ thể bằng cách yêu cầu một URL HTTPS tương ứng và đọc metadata nhúng trong phản hồi HTML (xem `go help importpath`). Nhiều dịch vụ lưu trữ đã cung cấp metadata đó cho các kho chứa mã Go, vì vậy cách dễ nhất để làm cho module của bạn sẵn có cho người khác sử dụng là làm cho đường dẫn module của nó khớp với URL của kho lưu trữ.

### Nhập các package từ module của bạn

Hãy viết một package `morestrings` và sử dụng nó từ chương trình `hello`. Trước tiên, tạo một thư mục cho package có tên `/hello/morestrings`, và sau đó một tệp có tên `reverse.go` trong thư mục đó với nội dung sau:

```go
// Package morestrings triển khai các hàm bổ sung để xử lý các chuỗi UTF-8
// được mã hóa, vượt xa những gì được cung cấp trong package chuẩn "strings".
package morestrings

// ReverseRunes trả về chuỗi đầu vào đã được đảo ngược ký tự từ trái sang phải.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

Vì hàm `ReverseRunes` của chúng ta bắt đầu bằng một chữ cái viết hoa, nó được xuất khẩu (exported) và có thể được sử dụng trong các package khác mà import package `morestrings` của chúng ta.

Hãy kiểm tra xem package đó có biên dịch được không bằng `go build`:

```bash
$ cd morestrings
$ go build
```

Điều này sẽ không tạo ra một tệp đầu ra. Thay vào đó, nó lưu trữ package đã biên dịch trong bộ nhớ cache xây dựng cục bộ.

Sau khi xác nhận rằng package `morestrings` đã được xây dựng, hãy sử dụng nó từ chương trình `hello`. Để thực hiện việc này, sửa đổi tệp gốc `$HOME/hello/hello.go` để sử dụng package `morestrings`:

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
```

Cài đặt lại chương trình `hello`:

```bash
$ go install example/user/hello
```

Chạy phiên bản mới của chương trình, bạn sẽ thấy một thông báo đảo ngược mới:

```bash
$ hello
Hello, Go!
```

### Nhập các package từ module từ xa

Một đường dẫn import có thể mô tả cách lấy mã nguồn của package bằng cách sử dụng hệ thống kiểm soát phiên bản như Git hoặc Mercurial. Công cụ `go` sử dụng thuộc tính này để tự động tải các package từ các kho lưu trữ từ xa. Ví dụ, để sử dụng `github.com/google/go-cmp/cmp` trong chương trình của bạn:

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

Bây giờ bạn đã có một phụ thuộc vào một module bên ngoài, bạn cần tải module đó và ghi lại phiên bản của nó trong tệp `go.mod` của bạn. Lệnh `go mod tidy` thêm các yêu cầu module bị thiếu cho các package đã import và xóa các yêu cầu về các module không còn được sử dụng nữa.

```bash
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

Các phụ thuộc module tự động được tải xuống vào thư mục `pkg/mod` của thư mục được chỉ định bởi biến môi trường `GOPATH`. Các nội dung đã tải xuống cho một phiên bản cụ thể của một module được chia sẻ giữa tất cả các module khác yêu cầu phiên bản đó, vì vậy lệnh `go` đánh dấu các tệp và thư mục đó là chỉ đọc. Để xóa tất cả các module đã tải xuống, bạn có thể sử dụng cờ `-modcache` với lệnh `go clean`:

```bash
$ go clean -modcache
$
```

## Kiểm thử

Go có một khung kiểm thử nhẹ bao gồm lệnh `go test` và package `testing`.

Bạn viết một bài kiểm thử bằng cách tạo một tệp có tên kết thúc bằng `_test.go` chứa các hàm có tên `TestXXX` với chữ ký `func (t *testing.T)`. Khung kiểm thử sẽ chạy từng hàm như vậy; nếu hàm gọi một hàm thất bại như `t.Error` hoặc `t.Fail`, bài kiểm thử được coi là đã thất bại.

Thêm một bài kiểm thử vào package `morestrings` bằng cách tạo tệp `$HOME/hello/morestrings/reverse_test.go` chứa mã Go sau:

```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

Sau đó chạy bài kiểm thử với `go test`:

```bash
$ cd $HOME/hello/morestrings
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

Chạy lệnh `go help test` và xem tài liệu của package `testing` để biết thêm chi tiết.

## Tiếp theo là gì

Đăng ký vào danh sách gửi thư golang-announce để được thông báo khi có phiên bản ổn định mới của Go.

Xem [Effective Go](https://go.dev/doc/effective_go) để nhận các mẹo viết mã Go rõ ràng và idiomatic.

Thực hiện [Khám Phá Go](https://tour.golang.org/) để học ngôn ngữ này một cách đúng đắn.

Truy cập trang tài liệu để xem các bài viết chi tiết về ngôn ngữ Go và các thư viện, công cụ của nó.

# Giải thích các đoạn code:

- **`go mod init`**:
  - **Mục đích**: Khởi tạo một module Go mới.
  - **Chi tiết**: Lệnh này tạo ra một tệp `go.mod` trong thư mục hiện tại. Tệp `go.mod` khai báo đường dẫn module của dự án (tên module) và giúp Go nhận diện và quản lý các phụ thuộc của dự án. Module là đơn vị cơ bản để tổ chức mã nguồn và phụ thuộc trong Go. Khi bạn sử dụng `go mod init`, bạn đặt tên cho module của mình và Go sẽ sử dụng tên này để quản lý các package và phiên bản của chúng.

- **`go install`**:
  - **Mục đích**: Biên dịch và cài đặt một chương trình Go.
  - **Chi tiết**: Lệnh này biên dịch chương trình Go thành một tệp nhị phân (executable) và cài đặt nó vào thư mục `bin` của Go. Đối với hệ điều hành Unix (như Linux và macOS), tệp nhị phân sẽ được đặt trong thư mục `$HOME/go/bin`, và trên Windows, trong thư mục `%USERPROFILE%\go\bin`. Khi bạn sử dụng `go install`, bạn có thể chạy chương trình từ dòng lệnh mà không cần chỉ định đường dẫn đầy đủ đến tệp thực thi.

- **`go build`**:
  - **Mục đích**: Biên dịch một package Go thành tệp nhị phân hoặc package nhị phân.
  - **Chi tiết**: Lệnh này biên dịch mã nguồn Go trong thư mục hiện tại hoặc các package con mà không tạo ra một tệp thực thi. Thay vào đó, các file đã biên dịch được lưu trữ trong bộ nhớ cache của Go. Điều này hữu ích khi bạn đang phát triển và cần biên dịch mã nguồn mà không muốn tạo tệp thực thi, hoặc khi bạn chỉ muốn kiểm tra xem mã nguồn có biên dịch thành công hay không.

- **`go run`**:
  - **Mục đích**: Biên dịch và chạy chương trình Go ngay lập tức mà không tạo tệp nhị phân.
  - **Chi tiết**: Lệnh này là cách nhanh chóng để biên dịch và chạy một hoặc nhiều tệp mã nguồn Go mà không cần phải tạo và quản lý tệp nhị phân. Khi bạn sử dụng `go run`, Go sẽ biên dịch mã nguồn và thực thi chương trình ngay lập tức. Đây là công cụ lý tưởng để kiểm tra mã nguồn nhỏ hoặc thử nghiệm nhanh mà không cần phải lưu trữ hoặc phân phối mã.

- **`go env`**:
  - **Mục đích**: Xem và thiết lập các biến môi trường của Go.
  - **Chi tiết**: Lệnh này cho phép bạn kiểm tra và thiết lập các biến môi trường liên quan đến Go, như `GOPATH`, `GOBIN`, và `GOROOT`. Những biến môi trường này kiểm soát nơi Go tìm kiếm các package, lưu trữ các tệp nhị phân đã cài đặt, và tìm các thư viện cơ bản của Go. Sử dụng `go env`, bạn có thể thay đổi cấu hình để phù hợp với môi trường phát triển của bạn.

- **`go mod tidy`**:
  - **Mục đích**: Làm sạch tệp `go.mod` bằng cách thêm các yêu cầu module bị thiếu và loại bỏ các yêu cầu không còn sử dụng.
  - **Chi tiết**: Lệnh này giúp duy trì tệp `go.mod` gọn gàng và chính xác. Khi bạn thêm hoặc xóa các phụ thuộc trong mã nguồn, tệp `go.mod` có thể trở nên lỗi thời hoặc không chính xác. `go mod tidy` sẽ tự động thêm các phụ thuộc còn thiếu mà mã nguồn yêu cầu và loại bỏ các phụ thuộc không còn được sử dụng nữa, giúp đảm bảo rằng dự án của bạn chỉ chứa các phụ thuộc cần thiết.
