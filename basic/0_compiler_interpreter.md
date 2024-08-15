### Trình Biên Dịch vs. Trình Thông Dịch

#### **Trình Biên Dịch (Compiler)**

- **Chức Năng**: Trình biên dịch chuyển đổi mã nguồn từ ngôn ngữ lập trình (như Go) thành mã máy (machine code) mà máy tính có thể hiểu và thực thi. Quy trình này tạo ra một file thực thi độc lập.
- **Ví Dụ**: Go sử dụng trình biên dịch `go build` để chuyển đổi mã nguồn Go thành một file thực thi mà không cần môi trường Go để chạy.

#### **Trình Thông Dịch (Interpreter)**

- **Chức Năng**: Trình thông dịch đọc mã nguồn và thực thi nó trực tiếp, mà không cần biên dịch thành mã máy trước. Thay vì tạo ra file thực thi độc lập, trình thông dịch chạy mã nguồn trực tiếp.
- **Ví Dụ**: Ruby (trong Rails) sử dụng trình thông dịch Ruby để thực thi mã nguồn Ruby. Khi bạn chạy một ứng dụng Rails, trình thông dịch Ruby xử lý mã nguồn Ruby và thực hiện các lệnh trong thời gian thực.

### Cách Chạy Ứng Dụng

#### **Ruby on Rails**

- **Môi Trường**: Rails là một framework viết bằng Ruby, và Ruby là ngôn ngữ thông dịch. Để chạy ứng dụng Rails, bạn cần có Ruby runtime cài đặt trên hệ thống của bạn.
- **Quá Trình Chạy Ứng Dụng**:
  1. **Cài Đặt**: Cài đặt Ruby và các gems (thư viện) cần thiết cho ứng dụng Rails.
  2. **Khởi Chạy**: Sử dụng lệnh `rails server` hoặc `rails s` để khởi động máy chủ web tích hợp. Máy chủ này sử dụng trình thông dịch Ruby để thực thi mã nguồn Ruby.
  3. **Chạy Thực Thi**: Trong khi ứng dụng đang chạy, Ruby runtime xử lý yêu cầu HTTP và điều khiển ứng dụng theo mã nguồn Ruby đã được viết sẵn.

  **Ví dụ**:
  ```bash
  rails server
  ```

  Khi bạn chạy lệnh trên, máy chủ web (như Puma) sẽ được khởi động và sử dụng trình thông dịch Ruby để xử lý các yêu cầu HTTP đến ứng dụng Rails.

#### **Go (Golang)**

- **Môi Trường**: Go là ngôn ngữ biên dịch. Để chạy ứng dụng Go, bạn cần biên dịch mã nguồn thành một file thực thi. File thực thi này có thể chạy độc lập mà không cần môi trường Go.
- **Quá Trình Chạy Ứng Dụng**:
  1. **Cài Đặt**: Cài đặt Go runtime (Go toolchain) để có thể biên dịch mã nguồn Go.
  2. **Biên Dịch**: Sử dụng lệnh `go build` để biên dịch mã nguồn Go thành file thực thi. Lệnh này tạo ra một file nhị phân độc lập từ mã nguồn Go.
  3. **Chạy Thực Thi**: Chạy file thực thi trực tiếp mà không cần Go runtime để chạy ứng dụng.

  **Ví dụ**:
  ```bash
  go build -o myapp main.go
  ./myapp
  ```

  Trong ví dụ trên, lệnh `go build` tạo ra một file thực thi có tên là `myapp`, và bạn có thể chạy ứng dụng bằng cách thực thi file này. Không cần môi trường Go (như các thư viện hay trình biên dịch) sau khi file thực thi đã được tạo ra.

### Tóm Tắt

- **Rails**: Sử dụng Ruby, một ngôn ngữ thông dịch. Bạn cần cài đặt Ruby runtime để chạy ứng dụng. Mã nguồn Ruby được thực thi trực tiếp bởi trình thông dịch.
- **Go**: Sử dụng Go, một ngôn ngữ biên dịch. Bạn cần cài đặt Go toolchain để biên dịch mã nguồn thành file thực thi. File thực thi có thể chạy độc lập mà không cần môi trường Go.

