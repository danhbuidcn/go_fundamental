# Flow control statements: for, if, else, switch and defer

https://go.dev/tour/flowcontrol/1

## For 

Go chỉ có một cấu trúc lặp, vòng lặp for.

Vòng lặp for cơ bản có ba thành phần được phân tách bằng dấu chấm phẩy:
    - câu lệnh init: được thực hiện trước lần lặp đầu tiên
    - biểu thức điều kiện: được đánh giá trước mỗi lần lặp
    - câu lệnh post: được thực hiện vào cuối mỗi lần lặp

Câu lệnh init thường là một khai báo biến ngắn và các biến được khai báo ở đó chỉ hiển thị trong phạm vi của câu lệnh for.

Vòng lặp sẽ ngừng lặp khi điều kiện boolean được đánh giá là sai.

Các câu lệnh init và post là tùy chọn.

Nếu bạn bỏ qua điều kiện lặp, nó sẽ lặp mãi mãi, do đó một vòng lặp vô hạn được diễn đạt một cách ngắn gọn.

```go
sum := 1
for i := 0; i < 10; i++ {
    sum += i
}

// OR
for ; sum < 1000; {
    sum += sum
}

// drop the semicolons
for sum < 1000 {
    sum += sum
}
```

## If

Các biến được khai báo bởi câu lệnh chỉ nằm trong phạm vi cho đến khi kết thúc câu lệnh if.

Các biến được khai báo bên trong câu lệnh if ngắn cũng có sẵn bên trong bất kỳ khối else nào.

```go
if v := math.Pow(x, n); v < lim {
    return v
} else {
    fmt.Printf("%g >= %g\n", v, lim)
}
```

## Switch

Câu lệnh switch là một cách ngắn hơn để viết một chuỗi các câu lệnh if - else. Nó chạy trường hợp đầu tiên có giá trị bằng biểu thức điều kiện.

Switch cases đánh giá từ trên xuống dưới, dừng khi trường hợp thành công.

Switch cases không có điều kiện cũng giống như Switch đúng. Cấu trúc này có thể là một cách rõ ràng để viết các chuỗi if-then-else dài.

```go
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

## Defer

Câu lệnh defer trì hoãn việc thực thi một hàm cho đến khi hàm xung quanh trả về.

- Stacking defers
    - Các lệnh gọi hàm bị trì hoãn được đẩy lên một ngăn xếp. Khi một hàm trả về, các cuộc gọi hoãn lại của nó sẽ được thực hiện theo thứ tự vào trước ra trước. Để tìm hiểu thêm về tuyên bố trì hoãn, hãy đọc bài đăng trên [blog này](https://go.dev/blog/defer-panic-and-recover).
