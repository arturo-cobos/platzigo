// Range

package main

import "fmt"

func main() {
    //arreglo - slice - tipo compuesto
    numeros := []int {2,4,6}
    suma := 0
    
    for _ , numero := range numeros {
        suma += numero 
        
    }
    fmt.Println("suma:", suma)

    for i, numero := range numeros {
        if numero == 3   {
            fmt.Println("index:", i)
        }  
    }


    algo := map [string]string{"a":"auto","b":"bebe"}

    for key, value := range algo {
        fmt.Printf("letra: %s item: %s\n", key, value)
    }
}
