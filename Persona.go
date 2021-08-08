package main
import "fmt"


type Persona struct{
    edad int 
    nombre string
    cedula int 
    nac int 
}

func main(){
    josue := Persona{20,"Josue",208130585,2001}
    fmt.Println(josue)
}
