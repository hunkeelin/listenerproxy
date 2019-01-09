package main
import (
    "github.com/hunkeelin/mtls/klinserver"
    "net/http"
    "net/http/httputil"
    "net/url"
    "fmt"
)
func listen (w http.ResponseWriter, r *http.Request) {
    rurl := "http://localhost:80"
    url,err := url.Parse(rurl)
    if err != nil {
        panic(err)
    }
    proxy := httputil.NewSingleHostReverseProxy(url)
    r.URL.Host = url.Host
    r.URL.Scheme = url.Scheme
    r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
    r.Host = url.Host
    proxy.ServeHTTP(w,r)
}
func serv(){
    fmt.Println("starting")
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
