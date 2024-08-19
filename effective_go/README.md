# Effective Go

https://go.dev/doc/effective_go

## Table of Contents

- [Introduction](#introduction)
    - [Examples](#examples)
- [Formatting](#formatting)
- [Commentary](#commentary)
- [Names](#names)
    - [Package names](#package_names)
    - [Getters](#getters)
    - [Interface names](#interface_names)
    - [MixedCaps](#mixed_caps)
- [Semicolons](#semicolons)
- [Control structures](#control_structures)
    - [If](#if)
    - [Redeclaration and reassignment](#redeclaration_and_reassignment)
    - [For](#for)
    - [Switch](#switch)
    - [Type switch](#type_switch)
- [Functions](#functions)
    - [Multiple return values](#multiple_return_values)
    - [Named result parameters](#named_result_parameters)
    - [Defer](#defer)
- [Data](#data)
    - [Allocation with](#allocation_with)
    - [Constructors and composite literals](#constructors_and_composite_literals)
    - [Allocation with](#allocation_with)
    - [Arrays](#arrays)
    - [Slices](#slices)
    - [Two-dimensional slices](#two_dimensional_slices)
    - [Maps](#maps)
    - [Printing](#printing)
    - [Append](#append)
- [Initialization](#initialization)
    - [Constants](#constants)
    - [Variables](#variables)
    - [The init function](#the_init_function)
- [Methods](#methods)
    - [Pointers vs. Values](#pointers_vs_values)
- [Interfaces and other types](#interfaces_and_other_types)
- [Interfaces](#interfaces)
    - [Conversions](#conversions)
    - [Interface conversions and type assertions](#interface_conversions_and_type_assertions)
    - [Generality](#generality)
    - [Interfaces and methods](#interfaces_and_methods)
- [The blank identifier](#the_blank_identifier)
    - [The blank identifier in multiple assignment](#the_blank_identifier_in_multiple_assignment)
    - [Unused imports and variables](#unused_imports_and_variables)
    - [Import for side effect](#import_for_side_effect)
    - [Interface checks](#interface_checks)
- [Embedding](#embedding)
- [Concurrency](#concurrency)
    - [Share by communicating](#share_by_communicating)
    - [Goroutines](#goroutines)
    - [Channels](#channels)
    - [Channels of channels](#channels_of_channels)
    - [Parallelization](#parallelization)
    - [A leaky buffer](#a_leaky_buffer)
- [Errors](#errors)
    - [Panic](#panic)
    - [Recover](#recover)
- [A web server](#a_web_server)

Dưới đây là bản dịch tiếng Việt của đoạn văn bạn cung cấp:

## Introduction¶

Go là một ngôn ngữ mới. Mặc dù nó vay mượn ý tưởng từ các ngôn ngữ hiện có, Go có những đặc điểm độc đáo khiến các chương trình Go hiệu quả khác biệt so với các chương trình viết bằng các ngôn ngữ khác. Một bản dịch trực tiếp từ chương trình C++ hoặc Java sang Go khó có thể cho ra kết quả thỏa đáng — các chương trình Java được viết bằng Java, không phải Go. Ngược lại, khi suy nghĩ về vấn đề từ góc nhìn của Go, bạn có thể tạo ra một chương trình thành công nhưng khá khác biệt. Nói cách khác, để viết Go tốt, điều quan trọng là phải hiểu các đặc tính và cách diễn đạt đặc trưng của nó. Ngoài ra, cần phải nắm rõ các quy ước đã được thiết lập khi lập trình bằng Go, chẳng hạn như cách đặt tên, định dạng, cấu trúc chương trình, và các yếu tố khác, để các chương trình bạn viết sẽ dễ dàng được các lập trình viên Go khác hiểu.

Tài liệu này cung cấp các mẹo để viết mã Go rõ ràng, theo phong cách đặc trưng của Go. Nó bổ sung cho [language specification](https://go.dev/ref/spec), [Tour of Go](https://go.dev/tour/welcome/1) và [How to Write Go Code](https://go.dev/doc/code), tất cả đều là những tài liệu bạn nên đọc trước.

**Lưu ý bổ sung tháng 1, 2022**: Tài liệu này được viết cho phiên bản Go ra mắt năm 2009 và chưa được cập nhật đáng kể kể từ đó. Mặc dù nó là một hướng dẫn tốt để hiểu cách sử dụng ngôn ngữ này, nhưng do sự ổn định (stability) của ngôn ngữ, nó ít đề cập đến các thư viện và không đề cập đến những thay đổi quan trọng của hệ sinh thái Go kể từ khi được viết, chẳng hạn như build system, testing, modules, and polymorphism. Không có kế hoạch cập nhật tài liệu này vì quá nhiều thứ đã xảy ra và một tập hợp lớn và ngày càng tăng của các tài liệu, blog và sách đã làm tốt công việc mô tả cách sử dụng Go hiện đại. Effective Go vẫn có giá trị, nhưng người đọc nên hiểu rằng nó không phải là hướng dẫn đầy đủ. Xem [issue 28782](https://github.com/golang/go/issues/28782) để biết thêm thông tin.

### Examples

[Go package sources](https://go.dev/src/) không chỉ nhằm mục đích làm thư viện cốt lõi mà còn là các ví dụ về cách sử dụng ngôn ngữ này. Hơn nữa, nhiều gói chứa các ví dụ tự hoạt động mà bạn có thể chạy trực tiếp từ trang web [go.dev](https://go.dev/), chẳng hạn như [ví dụ này](https://pkg.go.dev/strings#example-Map) (nếu cần, nhấp vào từ "Example" để mở nó ra). Nếu bạn có câu hỏi về cách tiếp cận một vấn đề hoặc cách triển khai một thứ gì đó, tài liệu, mã và ví dụ trong thư viện có thể cung cấp câu trả lời, ý tưởng và nền tảng.

## Formatting

Các vấn đề về định dạng thường gây tranh cãi nhất nhưng lại ít quan trọng nhất. Mọi người có thể thích nghi với các phong cách định dạng khác nhau, nhưng tốt hơn là họ không cần phải làm như vậy, và sẽ ít thời gian hơn được dành cho chủ đề này nếu mọi người tuân theo cùng một phong cách. Vấn đề là làm thế nào để tiếp cận Utopia này mà không cần một hướng dẫn phong cách quy định dài dòng.

Với Go, chúng ta áp dụng một cách tiếp cận khác thường và để máy móc xử lý hầu hết các vấn đề định dạng. Chương trình gofmt (cũng có sẵn dưới dạng lệnh `go fmt`, hoạt động ở cấp độ gói thay vì cấp độ tệp nguồn) đọc một chương trình Go và xuất mã nguồn theo kiểu chuẩn của việc thụt lề và căn chỉnh dọc, giữ nguyên và nếu cần thiết, định dạng lại các chú thích. Nếu bạn muốn biết cách xử lý một tình huống bố cục mới, hãy chạy gofmt; nếu câu trả lời không có vẻ đúng, hãy sắp xếp lại chương trình của bạn (hoặc gửi lỗi về gofmt), đừng cố gắng khắc phục.

Ví dụ, không cần dành thời gian để căn chỉnh các chú thích trên các trường của một cấu trúc. Gofmt sẽ làm điều đó cho bạn. Với khai báo:

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

gofmt sẽ căn chỉnh các cột:

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

Tất cả mã Go trong các gói tiêu chuẩn đã được định dạng bằng gofmt.

Một số chi tiết định dạng vẫn còn. Rất ngắn gọn:

- **Thụt lề**: Chúng tôi sử dụng tab để thụt lề và gofmt sẽ phát ra tab theo mặc định. Chỉ sử dụng khoảng trắng nếu bạn phải làm vậy.
- **Độ dài dòng**: Go không có giới hạn độ dài dòng. Đừng lo lắng về việc tràn một thẻ đục lỗ. Nếu một dòng cảm thấy quá dài, hãy ngắt nó và thụt lề bằng một tab bổ sung.
- **Dấu ngoặc**: Go cần ít dấu ngoặc hơn C và Java: các cấu trúc điều khiển (`if`, `for`, `switch`) không có dấu ngoặc trong cú pháp. Ngoài ra, hệ thống ưu tiên toán tử ngắn hơn và rõ ràng hơn, vì vậy `x<<8 + y<<16` có nghĩa là những gì khoảng cách ngụ ý, không giống như trong các ngôn ngữ khác.



