package main
import (
    "github.com/hunkeelin/mtls/klinserver"
    "net/http"
    "fmt"
)
func listen (w http.ResponseWriter, r *http.Request) {
//    fmt.Printf("%#v\n", r)
//    fmt.Println(r.RequestURI)
//    fmt.Println(r.Method)
    fmt.Println("got request from",r.RemoteAddr)
    return

}
func serv(){
    con := http.NewServeMux()
    con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        listen(w, r)
    })
    s := &klinserver.ServerConfig{
        BindPort: "2018",
        ServeMux: con,
    }
    err := klinserver.Server(s)
    if err != nil {
        panic(err)
    }
}
func main(){
    serv()
}
