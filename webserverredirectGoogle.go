//web Server 

package main

import (
    //"fmt"
"log"
"net/http"
)

func main() {
    http.HandleFunc("/", handler) //cada petición amnda llamar a un hadler..
    log.Fatal(http.ListenAndServe("localhost:8000",nil)) //especifica que escuche en el puerto 8000.
}

//handler: echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request){
    //fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    http.Redirect(w,r,"http://www.google.com",http.StatusMovedPermanently)
}
