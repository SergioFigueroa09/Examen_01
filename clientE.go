package main

import (
	"bufio"
	"fmt"
	"os"
	//"net"
	"net/rpc"

	//"strings"
)

func client() {
	//Variables
	scanner := bufio.NewScanner(os.Stdin) //Scanner para mensajes correctamente
	args := []string{"Anonimo","Hello =)"} //Argumentos para función RPC: Usuario,Mensaje
	var auxlist []string
	var result string //Resultado de RPC
	op := 0

	c, err := rpc.Dial("tcp",":9999")
	if err != nil{
		fmt.Println(err)
		return
	}

	//LOG IN
	fmt.Println("INGRESE SU NOMBRE DE USUARIO: ")
	fmt.Scanln(&args[0])

	err = c.Call("Server.Login",args[0], &result) //Función RPC para Login
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//EMPIEZA CHAT
	for{
				
		fmt.Println("------------------------MENU------------------------")
		fmt.Println("1.-Enviar Mensaje")
		fmt.Println("2.-ver chat")
		fmt.Println("9.-Log Out")

		fmt.Scanln(&op)
		switch op{
		case 1://MENSAJE
			fmt.Println("----ESCRIBA SU MENSAJE:")
			fmt.Printf(">>")
			scanner.Scan()

			
			args[1] = scanner.Text()
			

			
			err = c.Call("Server.Hello",args, &result) //Función RPC para mandar mensaje a servidor
			if err != nil{
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			break;
		case 2: //Convo
			err = c.Call("Server.Enviar",args, &auxlist) //Función RPC para ver el chat
			if err != nil{
				fmt.Println(err)
			} else {
				fmt.Println("-----------------------INICIO DE LA CONVERSACIÓN")
				for i:=0;i<len(auxlist);i++{
					fmt.Println(auxlist[i])
				}
			}
			break;
		case 9://LOGOUT
			err = c.Call("Server.Logout",args[0], &result) //Función RPC para Logout
			if err != nil{
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
			return
		}

		
	}
}

func main(){
	client()
}