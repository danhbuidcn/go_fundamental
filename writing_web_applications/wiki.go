package main

import (
	"fmt"
	"os"
	"log"
    "net/http"
	"html/template"
	"regexp"
)

var templates = template.Must(template.ParseFiles("tmpl/index.html", "tmpl/edit.html", "tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

// phương thức gắn liên với kiểu Page. p là reciever, save được gọi trên p
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// tạo struct mới và trả về địa chỉ của struct này, & lấy địa chỉ của struct mới tạo
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func initial_record() {
    // Tạo thư mục data nếu chưa tồn tại
    if err := os.MkdirAll("data", 0755); err != nil {
        log.Fatal(err)
    }

    // Khởi tạo đối tượng Page và gọi phương thức save
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// &Page: lấy địa chỉ của struct mới tạo, thay vì trả về một bản sao của struct thì trả về một con trỏ (*Page)
	if err := p1.save(); err != nil {
        log.Fatal(err) // Xử lý lỗi khi lưu
    }

    // Load page từ file
    p2, err := loadPage("TestPage")
    if err != nil {
        log.Fatal(err) // Xử lý lỗi khi load
    }

    // Kiểm tra con trỏ p2 không phải là nil trước khi sử dụng
    if p2 != nil {
        fmt.Println(string(p2.Body))
    } else {
        log.Fatal("Loaded page is nil")
    }
}

// danh sách record
func indexHandler(w http.ResponseWriter, r *http.Request) {
    files, err := os.ReadDir("data")
    if err != nil {
        http.Error(w, "Unable to read directory", http.StatusInternalServerError)
        return
    }

    titles := []string{}
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        title := file.Name()
        titles = append(titles, title[:len(title)-4]) // Loại bỏ đuôi ".txt"
    }

    renderTemplate(w, "index", titles) // Truyền slice []string vào template
}

// cho phép người dùng xem trang wiki
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

// tải trang hoặc tạo một cấu trúc page trống
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

// submit form
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

// Hàm tạo Handler không cần tham số
func makeIndexHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fn(w, r)
    }
}

func main() {
    initial_record()

    http.HandleFunc("/", makeIndexHandler(indexHandler))
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))

    // Phục vụ các tài nguyên tĩnh từ thư mục "static"
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// $ go build wiki.go
// $ ./wiki
// OR $ go run wiki.go