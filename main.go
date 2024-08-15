package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d đang xử lý job %d\n", id, job)
        results <- job * 2 // Xử lý job và gửi kết quả vào channel
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    var wg sync.WaitGroup

    // Tạo 3 worker
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Gửi các công việc vào channel
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs) // Đóng channel jobs khi hoàn tất gửi công việc

    // Đợi tất cả worker hoàn thành
    wg.Wait()
    close(results) // Đóng channel results khi tất cả công việc đã được xử lý

    // Nhận kết quả từ channel results
    for result := range results {
        fmt.Println("Kết quả:", result)
    }
}
