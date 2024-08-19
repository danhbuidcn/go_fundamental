# Writing Web Applications

https://go.dev/doc/articles/wiki/

### Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Data Structures](#data-structures)
- <a href="#introducing-the-net-http-package-an-interlude">Introducing the net/http package (an interlude)</a>
- <a href="#using-net/http-to-serve-wiki-pages">Using net/http to serve wiki pages</a>
- [Editing Pages](#editing-pages)
- <a href="#the-html/template-package">The html/template package</a>
- [Handling non-existent pages](#handling-non-existent-pages)
- [Saving Pages](#saving-pages)
- [Error handling](#error-handling)
- [Template caching](#template-caching)
- [Validation](#validation)
- [Introducing Function Literals and Closures](#introducing-function-literals-and-closures)
- [Try it out!](#try-it-out)
- [Other tasks](#other-tasks)

### Introduction

- Trong hướng dẫn này, chúng ta sẽ:
    - Tạo một cấu trúc dữ liệu với các phương thức load và save
    - Sử dụng gói `net/http` để xây dựng ứng dụng web
    - Sử dụng gói `html/template` để xử lý các template HTML
    - Sử dụng gói `regexp` để xác thực đầu vào của người dùng
    - Sử dụng closures

- **Kiến thức cần có:**
    - Kinh nghiệm lập trình
    - Hiểu biết về các công nghệ web cơ bản (HTTP, HTML)
    - Một chút kiến thức về dòng lệnh UNIX/DOS

### Getting Started

Hiện tại, bạn cần có máy tính chạy FreeBSD, Linux, macOS, hoặc Windows để chạy Go. Chúng tôi sẽ sử dụng ký hiệu `$` để đại diện cho dấu nhắc lệnh.

Cài đặt Go (xem [Hướng dẫn Cài đặt](https://go.dev/doc/install)).

Tạo một thư mục mới cho hướng dẫn này trong GOPATH của bạn và chuyển đến đó:

```sh
$ mkdir gowiki
$ cd gowiki
```

Tạo một tập tin có tên `wiki.go`, mở nó trong trình soạn thảo yêu thích của bạn và thêm các dòng sau:

```go
package main

import (
    "fmt"
    "os"
)
```

Chúng tôi nhập các gói `fmt` và `os` từ thư viện chuẩn của Go. Sau này, khi triển khai thêm chức năng, chúng tôi sẽ thêm nhiều gói hơn vào khai báo nhập này.

### Data Structures

Hãy bắt đầu bằng cách định nghĩa các cấu trúc dữ liệu. Một wiki bao gồm một chuỗi các trang liên kết với nhau, mỗi trang có tiêu đề và nội dung (nội dung trang). Ở đây, chúng tôi định nghĩa `Page` là một cấu trúc với hai trường đại diện cho tiêu đề và nội dung.

```go
type Page struct {
    Title string
    Body  []byte
}
```

Loại `[]byte` có nghĩa là "mảng byte". (Xem [Slices: usage and internals](https://go.dev/blog/slices-intro) để biết thêm về slices). Trường `Body` là `[]byte` thay vì `string` vì đó là loại mà các thư viện io mà chúng tôi sẽ sử dụng mong đợi, như bạn sẽ thấy dưới đây.

Cấu trúc `Page` mô tả cách dữ liệu trang sẽ được lưu trữ trong bộ nhớ. Nhưng còn lưu trữ vĩnh viễn thì sao? Chúng tôi có thể giải quyết vấn đề đó bằng cách tạo một phương thức `save` trên `Page`:

```go
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}
```

Chữ ký của phương thức này đọc như sau: "Đây là một phương thức tên là `save` nhận `p` là một con trỏ đến `Page`. Nó không nhận tham số nào và trả về giá trị của loại `error`."

Phương thức này sẽ lưu nội dung của `Page` vào một tập tin văn bản. Để đơn giản, chúng tôi sẽ sử dụng `Title` làm tên tập tin.

Phương thức `save` trả về một giá trị lỗi vì đó là kiểu trả về của `WriteFile` (một hàm thư viện chuẩn ghi một mảng byte vào một tập tin). Phương thức `save` trả về giá trị lỗi, để ứng dụng có thể xử lý nếu có bất kỳ điều gì sai trong quá trình ghi tập tin. Nếu mọi thứ diễn ra suôn sẻ, `Page.save()` sẽ trả về `nil` (giá trị zero cho các con trỏ, interfaces và một số loại khác).

Hằng số số nguyên bát phân `0600`, được truyền làm tham số thứ ba cho `WriteFile`, chỉ thị rằng tập tin nên được tạo với quyền đọc-ghi cho người dùng hiện tại. (Xem Unix man page open(2) để biết chi tiết.)

Ngoài việc lưu trang, chúng ta cũng muốn tải trang:

```go
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
```

Hàm `loadPage` tạo tên tập tin từ tham số tiêu đề, đọc nội dung của tập tin vào một biến mới `body`, và trả về một con trỏ đến một literal `Page` được xây dựng với các giá trị tiêu đề và nội dung phù hợp.

Các hàm có thể trả về nhiều giá trị. Hàm thư viện chuẩn `os.ReadFile` trả về `[]byte` và `error`. Trong `loadPage`, lỗi chưa được xử lý; "blank identifier" được đại diện bởi ký hiệu gạch dưới (`_`) được sử dụng để loại bỏ giá trị lỗi trả về (về cơ bản là gán giá trị cho không có gì).

Nhưng điều gì sẽ xảy ra nếu `ReadFile` gặp lỗi? Ví dụ, tập tin có thể không tồn tại. Chúng ta không nên bỏ qua các lỗi như vậy. Hãy sửa đổi hàm để trả về `*Page` và `error`.

```go
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
```

Người gọi hàm này bây giờ có thể kiểm tra tham số thứ hai; nếu nó là `nil` thì hàm đã thành công trong việc tải một `Page`. Nếu không, nó sẽ là một lỗi có thể được người gọi xử lý (xem [language specification](https://go.dev/ref/spec#Errors) để biết chi tiết).

Lúc này, chúng ta đã có một cấu trúc dữ liệu đơn giản và khả năng lưu và tải từ tập tin. Hãy viết một hàm `main` để kiểm tra những gì chúng tôi đã viết:

```go
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
```

Sau khi biên dịch và thực thi mã này, một tập tin có tên `TestPage.txt` sẽ được tạo ra, chứa nội dung của `p1`. Tập tin sau đó sẽ được đọc vào cấu trúc `p2`, và phần tử `Body` của nó sẽ được in ra màn hình.

Bạn có thể biên dịch và chạy chương trình như sau:

```sh
$ go build wiki.go
$ ./wiki
```

(Nếu bạn đang sử dụng Windows, bạn phải gõ "wiki" mà không có "./" để chạy chương trình.)

[Nhấp vào đây để xem mã mà chúng tôi đã viết cho đến nay](https://go.dev/doc/articles/wiki/part1.go).

### <p id="introducing-the-net-http-package-an-interlude">Introducing the net/http package (an interlude)</p>

Dưới đây là ví dụ đầy đủ về một máy chủ web đơn giản:

```go
//go:build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Hàm `main` bắt đầu bằng việc gọi `http.HandleFunc`, thông báo cho gói `http` rằng tất cả các yêu cầu đến gốc của web ("/") sẽ được xử lý bởi hàm `handler`.

Sau đó, nó gọi `http.ListenAndServe`, chỉ định rằng nó sẽ lắng nghe trên cổng 8080 trên bất kỳ giao diện mạng nào (":8080"). (Bạn không cần phải lo lắng về tham số thứ hai, `nil`, ngay bây giờ.) Hàm này sẽ chặn (block) cho đến khi chương trình bị kết thúc.

`ListenAndServe` luôn trả về lỗi, vì nó chỉ trả về khi xảy ra lỗi không mong muốn. Để ghi lại lỗi đó, chúng tôi bọc hàm gọi với `log.Fatal`.

Hàm `handler` có kiểu `http.HandlerFunc`. Nó nhận hai đối số là `http.ResponseWriter` và `http.Request`.

- Một giá trị `http.ResponseWriter` sẽ tổng hợp phản hồi của máy chủ HTTP; bằng cách ghi vào nó, chúng ta gửi dữ liệu đến khách hàng HTTP.
- Một `http.Request` là một cấu trúc dữ liệu đại diện cho yêu cầu HTTP từ khách hàng. `r.URL.Path` là thành phần đường dẫn của URL yêu cầu. Cú pháp `[1:]` có nghĩa là "tạo một con slice con của `Path` từ ký tự thứ nhất đến hết." Điều này loại bỏ ký tự "/" đầu tiên khỏi tên đường dẫn.

Nếu bạn chạy chương trình này và truy cập URL:

```
http://localhost:8080/monkeys
```

chương trình sẽ hiển thị một trang với nội dung:

```
Hi there, I love monkeys!
```

### <p id="using-net/http-to-serve-wiki-pages">Using net/http to serve wiki pages</p>

Để sử dụng gói `net/http`, nó phải được nhập vào:

```go
import (
    "fmt"
    "os"
    "log"
    "net/http"
)
```

Hãy tạo một trình xử lý, `viewHandler`, cho phép người dùng xem một trang wiki. Nó sẽ xử lý các URL có tiền tố "/view/".

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
```

Một lần nữa, lưu ý việc sử dụng `_` để bỏ qua giá trị lỗi trả về từ `loadPage`. Điều này được thực hiện ở đây vì đơn giản và thường được coi là thực hành không tốt. Chúng ta sẽ giải quyết điều này sau.

Trước tiên, hàm này trích xuất tiêu đề trang từ `r.URL.Path`, thành phần đường dẫn của URL yêu cầu. `Path` được cắt lại với `[len("/view/"):]` để loại bỏ thành phần "/view/" đầu tiên của đường dẫn yêu cầu. Điều này là vì đường dẫn sẽ luôn bắt đầu bằng "/view/", điều này không phải là một phần của tiêu đề trang.

Hàm sau đó tải dữ liệu trang, định dạng trang bằng một chuỗi HTML đơn giản, và ghi nó vào `w`, `http.ResponseWriter`.

Để sử dụng trình xử lý này, chúng ta sửa đổi hàm `main` để khởi tạo `http` bằng cách sử dụng `viewHandler` để xử lý bất kỳ yêu cầu nào dưới đường dẫn `/view/`.

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

[Nhấp vào đây để xem mã code mà chúng ta đã viết cho đến nay.](https://go.dev/doc/articles/wiki/part2.go)

Hãy tạo một số dữ liệu trang (như `test.txt`), biên dịch mã của chúng ta, và thử phục vụ một trang wiki.

Mở tệp `test.txt` trong trình soạn thảo của bạn và lưu chuỗi "Hello world" (không có dấu ngoặc kép) vào nó.

```bash
$ go build wiki.go
$ ./wiki
```

(Nếu bạn đang sử dụng Windows, bạn phải gõ "wiki" mà không có "./" để chạy chương trình.)

Với máy chủ web đang chạy, việc truy cập vào `http://localhost:8080/view/test` sẽ hiển thị một trang có tiêu đề "test" chứa các từ "Hello world".

### Editing Pages

Một wiki không phải là một wiki nếu không có khả năng chỉnh sửa các trang. Hãy tạo hai trình xử lý mới: một tên là `editHandler` để hiển thị một biểu mẫu 'chỉnh sửa trang', và một tên khác là `saveHandler` để lưu dữ liệu nhập từ biểu mẫu.

Trước tiên, chúng ta thêm chúng vào `main()`:

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Hàm `editHandler` tải trang (hoặc, nếu nó không tồn tại, tạo một cấu trúc `Page` trống) và hiển thị một biểu mẫu HTML.

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}
```

Hàm này sẽ hoạt động tốt, nhưng toàn bộ HTML được mã hóa cứng là khá xấu. Tất nhiên, có một cách tốt hơn.

### <p id="the-html/template-package">The html/template package</p>

Gói `html/template` là một phần của thư viện chuẩn Go. Chúng ta có thể sử dụng `html/template` để giữ mã HTML trong một tệp riêng biệt, cho phép chúng ta thay đổi giao diện của trang chỉnh sửa mà không cần phải sửa đổi mã Go cơ bản.

Đầu tiên, chúng ta cần thêm `html/template` vào danh sách các gói nhập khẩu. Chúng ta cũng sẽ không sử dụng `fmt` nữa, vì vậy phải loại bỏ nó.

```go
import (
    "html/template"
    "os"
    "net/http"
)
```

Hãy tạo một tệp mẫu chứa biểu mẫu HTML. Mở một tệp mới có tên `edit.html` và thêm các dòng sau:

```html
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
```

Chỉnh sửa `editHandler` để sử dụng mẫu, thay vì HTML mã hóa cứng:

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}
```

Hàm `template.ParseFiles` sẽ đọc nội dung của `edit.html` và trả về một `*template.Template`.

Phương thức `t.Execute` thực thi mẫu, ghi HTML đã tạo vào `http.ResponseWriter`. Các chỉ thị `{{.Title}}` và `{{.Body}}` liên quan đến `p.Title` và `p.Body`.

Các chỉ thị mẫu được bao quanh bởi dấu ngoặc nhọn kép. Lệnh `printf "%s" .Body` là một cuộc gọi hàm xuất `.Body` dưới dạng chuỗi thay vì luồng byte, tương tự như cuộc gọi `fmt.Printf`. Gói `html/template` giúp đảm bảo rằng chỉ HTML an toàn và đúng định dạng được tạo ra bởi các hành động mẫu. Ví dụ, nó tự động thoát ký tự lớn hơn (>) bằng cách thay thế nó bằng `&gt;`, để đảm bảo dữ liệu người dùng không làm hỏng HTML của biểu mẫu.

Vì chúng ta đang làm việc với các mẫu, hãy tạo một mẫu cho `viewHandler` có tên là `view.html`:

```html
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>

<div>{{printf "%s" .Body}}</div>
```

Chỉnh sửa `viewHandler` cho phù hợp:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}
```

Lưu ý rằng chúng ta đã sử dụng mã mẫu gần như chính xác trong cả hai trình xử lý. Hãy loại bỏ sự trùng lặp này bằng cách di chuyển mã mẫu vào một hàm riêng:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}
```

Và chỉnh sửa các trình xử lý để sử dụng hàm đó:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}
```

Nếu chúng ta loại bỏ đăng ký của trình xử lý `save` chưa được thực hiện trong `main`, chúng ta có thể biên dịch và thử nghiệm chương trình của chúng ta một lần nữa. [Nhấp vào đây để xem mã code mà chúng ta đã viết cho đến nay.](https://go.dev/doc/articles/wiki/part3.go)

### Handling non-existent pages

Nếu bạn truy cập `/view/APageThatDoesntExist`, bạn sẽ thấy một trang chứa HTML. Điều này xảy ra vì nó bỏ qua giá trị lỗi trả về từ `loadPage` và tiếp tục cố gắng điền dữ liệu vào mẫu mà không có dữ liệu. Thay vào đó, nếu trang yêu cầu không tồn tại, nó nên chuyển hướng khách hàng đến trang chỉnh sửa để nội dung có thể được tạo ra:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}
```

Hàm `http.Redirect` thêm mã trạng thái HTTP `http.StatusFound` (302) và tiêu đề Location vào phản hồi HTTP.

### Saving Pages

Hàm `saveHandler` sẽ xử lý việc gửi các biểu mẫu nằm trên các trang chỉnh sửa. Sau khi bỏ chú thích dòng liên quan trong `main`, hãy triển khai trình xử lý:

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Tiêu đề trang (được cung cấp trong URL) và trường duy nhất của biểu mẫu, `Body`, được lưu vào một `Page` mới. Phương thức `save()` sau đó được gọi để ghi dữ liệu vào tệp, và khách hàng được chuyển hướng đến trang `/view/`.

Giá trị trả về bởi `FormValue` có kiểu `string`. Chúng ta phải chuyển đổi giá trị đó thành `[]byte` trước khi nó phù hợp với cấu trúc `Page`. Chúng ta sử dụng `[]byte(body)` để thực hiện việc chuyển đổi.

### Error handling

Có nhiều điểm trong chương trình của chúng ta nơi lỗi bị bỏ qua. Đây là một thực hành kém, vì khi xảy ra lỗi, chương trình sẽ có hành vi không mong muốn. Một giải pháp tốt hơn là xử lý các lỗi và trả về thông báo lỗi cho người dùng. Bằng cách đó, nếu có điều gì đó sai, máy chủ sẽ hoạt động chính xác theo cách chúng ta mong muốn và người dùng có thể được thông báo.

Đầu tiên, hãy xử lý lỗi trong hàm `renderTemplate`:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

Hàm `http.Error` gửi một mã phản hồi HTTP đã chỉ định (trong trường hợp này là "Lỗi máy chủ nội bộ") và thông báo lỗi. Quyết định đưa điều này vào một hàm riêng đang chứng minh giá trị của nó.

Tiếp theo, hãy sửa chữa `saveHandler`:

```go
func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/save/"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

Bất kỳ lỗi nào xảy ra trong `p.save()` sẽ được thông báo cho người dùng.

### Template caching

Có một sự không hiệu quả trong mã này: `renderTemplate` gọi `ParseFiles` mỗi khi một trang được kết xuất. Một cách tiếp cận tốt hơn là gọi `ParseFiles` một lần tại thời điểm khởi tạo chương trình, phân tích tất cả các mẫu thành một `*Template` duy nhất. Sau đó, chúng ta có thể sử dụng phương thức [ExecuteTemplate](https://pkg.go.dev/html/template#Template.ExecuteTemplate) để kết xuất một mẫu cụ thể.

Đầu tiên, tạo một biến toàn cục tên là `templates` và khởi tạo nó bằng `ParseFiles`.

```go
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
```

Hàm `template.Must` là một trình bao bọc tiện ích sẽ panic khi nhận một giá trị lỗi không phải là nil, và trả về `*Template` không thay đổi. Một panic là hợp lý ở đây; nếu các mẫu không thể được tải, điều hợp lý duy nhất là thoát chương trình.

Hàm `ParseFiles` nhận bất kỳ số lượng chuỗi nào xác định các tệp mẫu của chúng ta và phân tích các tệp đó thành các mẫu được đặt tên theo tên tệp cơ sở. Nếu chúng ta thêm nhiều mẫu hơn vào chương trình, chúng ta sẽ thêm tên của chúng vào các tham số của `ParseFiles`.

Sau đó, sửa đổi hàm `renderTemplate` để gọi phương thức `templates.ExecuteTemplate` với tên của mẫu phù hợp:

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

Lưu ý rằng tên mẫu là tên tệp mẫu, vì vậy chúng ta phải nối thêm ".html" vào tham số `tmpl`.

### Validation

Như bạn có thể đã quan sát, chương trình này có một lỗ hổng bảo mật nghiêm trọng: người dùng có thể cung cấp một đường dẫn tùy ý để được đọc/ghi trên máy chủ. Để giảm thiểu điều này, chúng ta có thể viết một hàm để xác thực tiêu đề với một biểu thức chính quy.

Đầu tiên, thêm `"regexp"` vào danh sách `import`. Sau đó, chúng ta có thể tạo một biến toàn cục để lưu trữ biểu thức xác thực của chúng ta:

```go
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
```

Hàm `regexp.MustCompile` sẽ phân tích và biên dịch biểu thức chính quy, và trả về một `regexp.Regexp`. `MustCompile` khác với `Compile` ở chỗ nó sẽ panic nếu việc biên dịch biểu thức gặp lỗi, trong khi `Compile` trả về lỗi như một tham số thứ hai.

Bây giờ, hãy viết một hàm sử dụng biểu thức `validPath` để xác thực đường dẫn và trích xuất tiêu đề trang:

```go
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // Tiêu đề là biểu thức con thứ hai.
}
```

Nếu tiêu đề hợp lệ, nó sẽ được trả về cùng với giá trị lỗi nil. Nếu tiêu đề không hợp lệ, hàm sẽ ghi một lỗi "404 Not Found" vào kết nối HTTP và trả về một lỗi cho trình xử lý. Để tạo một lỗi mới, chúng ta cần nhập gói `errors`.

Hãy gọi `getTitle` trong mỗi trình xử lý:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
    if err != nil {
        return
    }
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

### Introducing Function Literals and Closures

Việc kiểm tra lỗi trong mỗi trình xử lý tạo ra rất nhiều mã trùng lặp. Nếu chúng ta có thể bọc mỗi trình xử lý trong một hàm thực hiện xác thực và kiểm tra lỗi này? Các hàm `literals` của Go cung cấp một phương tiện mạnh mẽ để trừu tượng hóa chức năng có thể giúp chúng ta ở đây.

Trước tiên, chúng ta sửa lại định nghĩa hàm của từng trình xử lý để nhận một chuỗi tiêu đề:

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)
```

Bây giờ, hãy định nghĩa một hàm bọc nhận một hàm kiểu trên và trả về một hàm loại `http.HandlerFunc` (thích hợp để được truyền cho hàm `http.HandleFunc`):

```go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Ở đây chúng ta sẽ trích xuất tiêu đề trang từ Request,
        // và gọi trình xử lý đã cung cấp 'fn'
    }
}
```

Hàm trả về được gọi là một đóng gói vì nó bao bọc các giá trị được xác định bên ngoài nó. Trong trường hợp này, biến `fn` (tham số duy nhất của `makeHandler`) được bao bọc bởi đóng gói. Biến `fn` sẽ là một trong các trình xử lý `save`, `edit` hoặc `view`.

Bây giờ chúng ta có thể lấy mã từ `getTitle` và sử dụng nó ở đây (với một số sửa đổi nhỏ):

```go
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
```

Đóng gói được trả về bởi `makeHandler` là một hàm nhận `http.ResponseWriter` và `http.Request` (nói cách khác, một `http.HandlerFunc`). Đóng gói trích xuất tiêu đề từ đường dẫn yêu cầu và xác thực nó bằng biểu thức chính quy `validPath`. Nếu tiêu đề không hợp lệ, một lỗi sẽ được viết vào `ResponseWriter` bằng cách sử dụng hàm `http.NotFound`. Nếu tiêu đề hợp lệ, hàm trình xử lý bao bọc

 `fn` sẽ được gọi với `ResponseWriter`, `Request`, và tiêu đề làm tham số.

Bây giờ chúng ta có thể bao bọc các hàm trình xử lý với `makeHandler` trong `main`, trước khi chúng được đăng ký với gói `http`:

```go
func main() {
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Cuối cùng, hãy loại bỏ các cuộc gọi đến `getTitle` từ các hàm xử lý, làm cho chúng đơn giản hơn nhiều:

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

### Try it out!

[Nhấn vào đây để xem mã nguồn cuối cùng.](https://go.dev/doc/articles/wiki/final.go)

Biên dịch lại mã và chạy ứng dụng:

```bash
$ go build wiki.go
$ ./wiki
```

Truy cập `http://localhost:8080/view/ANewPage` sẽ hiển thị cho bạn mẫu trang chỉnh sửa. Bạn có thể nhập một số văn bản, nhấp vào 'Lưu', và sẽ được chuyển hướng đến trang mới được tạo.

### Other tasks

Dưới đây là một số nhiệm vụ đơn giản mà bạn có thể muốn thử:

- Lưu các mẫu HTML vào thư mục `tmpl/` và dữ liệu trang vào thư mục `data/`.
- Thêm một trình xử lý để làm cho root web chuyển hướng đến `/view/FrontPage`.
- Làm đẹp các mẫu trang bằng cách đảm bảo chúng là HTML hợp lệ và thêm một số quy tắc CSS.
- Triển khai liên kết giữa các trang bằng cách chuyển đổi các trường hợp của `[PageName]` thành `<a href="/view/PageName">PageName</a>`. (Gợi ý: bạn có thể sử dụng `regexp.ReplaceAllFunc` để thực hiện điều này)
