package main
import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

//Otra forma de hacer el reto pero con sobrescritura de url, No funciona con google.com
//Pero se ve curioso tener el portal de movistar con tu localhost ? http://localhost:8005/web/guest
//También funciona con otros sitios que hagan sobrescritura  de url.

func main() {
    targetUrl, _ := url.Parse("http://movistar.com.co",)   
    log.Fatal(http.ListenAndServe(":80002", httputil.NewSingleHostReverseProxy(targetUrl)))
}
