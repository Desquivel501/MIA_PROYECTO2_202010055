package main  

import "MIA_PROYECTO2_202010055/analizador"

import (
    "bufio"
	"fmt"
	"os"
	"strings"
)


func main() {  
    a := analizador.New("")
    // mkdisk -size=3 -unit=M -path=/home/desquivel/Desktop/Disco2.dsk
    // fdisk -type=P -path=/home/desquivel/Desktop/Disco2.dsk -unit=K -name=Particion1 -size=300
    // fdisk -type=P -path=/home/desquivel/Desktop/Disco2.dsk -unit=K -name=Particion2 -size=500

    reader := bufio.NewReader(os.Stdin)
    for true{
        fmt.Println("||========================================================||")
        fmt.Println("||                    [MIA] PROYECTO 2                    ||")
        fmt.Println("||========================================================||")
        fmt.Print("[MIA]@Proyecto2:~$ ")
        
        comando, _ := reader.ReadString('\n')

        if(strings.Contains(comando, "exit")){
            break
        }
        a.Analizar(comando)
        
    }
    
}