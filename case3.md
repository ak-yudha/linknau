# Package Management

## Table of Contents

- [Inisialisasi Modul Go](#Initialize)
- [Import a third-party package](#third-party)
- [Use the package](#use)

## Initialize
1. Inisialisasi Modul Go:

    ```bash
    go mod init myproject
    ```
   dari perintah diatas akan menghasilkan output file baru bernama go.mod dengan isi:
   ```bash
    module myproject

    go 1.20
    ```

## third-party
1. Impor paket dalam kode Go:
    ```bash
    package main

    import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux" // Paket pihak ketiga
    )
    
    func main() {
    router := mux.NewRouter() // Membuat router dengan Gorilla Mux

    // Menetapkan route handler
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Memulai server HTTP
    log.Fatal(http.ListenAndServe(":8080", router))
    }
    ```
2. Unduh paket dengan go get:

    ```bash
    go get github.com/gorilla/mux
    ```
   dari perintah diatas output akan ditambahkan ke go.mod dengan isi:
    ```bash
    module myproject

    go 1.20

    require github.com/gorilla/mux v1.8.0 // Paket pihak ketiga
    ``` 
3. Perbaharui file go.sum
    Secara otomatis jika kita import package dari third party akan di tambahkan juga di file go.sum, contoh:
    ```bash
    github.com/gorilla/mux v1.8.0 h1:somesha256hash...
    github.com/gorilla/mux v1.8.0/go.mod h1:anotherchecksum...
    ``` 
## use
1. Use the package
    ```bash
    package main
    
    import (
    "github.com/gorilla/mux"
    "net/http"
    )
    
    func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
    })
    http.ListenAndServe(":8080", r)
    }
    ``` 
