package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
	//"bufio"
	//"strconv"
)

type Server struct {
	Lista_Usuarios_Actual []string
	Historial_Usuarios []string
	Historial_Mensajes []string
	Full_Convo []string
	Enviado bool
}

var S Server

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func VerUsuarios(){
	fmt.Println("****Actualmente la lista de usuarios es: ")
	fmt.Println(S.Lista_Usuarios_Actual)
}

func CrearBackup(){
	f, err := os.Create("BACKUP.txt")
    if err != nil {
        fmt.Println(err)
        return
	}
	for i:=0;i<len(S.Historial_Usuarios);i++{
		f.WriteString(">"+S.Historial_Usuarios[i]+": "+S.Historial_Mensajes[i]+"\n")
	}
    f.WriteString("****FIN DE LA CONVERSACIÓN")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println("SE REALIZÓ EL BACKUP!!")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func (this *Server) Enviar(args []string, reply *[]string) error {
	*reply = S.Full_Convo
	return nil
}

func  VerConvo(){
	fmt.Println("-----------------------INICIO DE LA CONVERSACIÓN")
	for i:=0;i<len(S.Historial_Usuarios);i++{
		fmt.Println(">"+S.Historial_Usuarios[i]+": "+S.Historial_Mensajes[i])
	}
	fmt.Println("-----------------------FIN DE LA CONVERSACIÓN")
}

func (this *Server) Hello(args []string, reply *string) error {
	//fmt.Println(args[0] + ": " + args[1])
	S.Historial_Usuarios = append(S.Historial_Usuarios,args[0])
	S.Historial_Mensajes = append(S.Historial_Mensajes,args[1])
	aux := ">>"+args[0]+": "+args[1]
	S.Full_Convo = append(S.Full_Convo,aux)
	*reply = "Se envió su mensaje"

	return nil
}

func (this *Server) Login(usuario string, reply *string) error{
	S.Lista_Usuarios_Actual = append(S.Lista_Usuarios_Actual,usuario)
	//fmt.Println("****"+usuario+" ha iniciado sesión.")
	//fmt.Println("Ahora la lista de usuarios es: ")
	//fmt.Println(S.Lista_Usuarios_Actual)
	
	S.Historial_Usuarios = append(S.Historial_Usuarios,"■■■■■ SERVER")
	S.Historial_Mensajes = append(S.Historial_Mensajes,usuario+" ha iniciado sesión.")

	aux:= "■■■■■ SERVER: "+usuario+" ha iniciado sesión."
	S.Full_Convo = append(S.Full_Convo,aux)

	*reply = "Has iniciado sesión."
	return nil
}

func (this *Server) Logout(usuario string, reply *string) error{
	x:=999
	for i:=0;i<len(S.Lista_Usuarios_Actual);i++{
		if(S.Lista_Usuarios_Actual[i]==usuario){
			x = i
		}
	}
	if(x==999){
		*reply = "****ERROR! Usuario no encontrado"
		return nil
	}else{
		S.Lista_Usuarios_Actual[x] = S.Lista_Usuarios_Actual[len(S.Lista_Usuarios_Actual)-1]
		S.Lista_Usuarios_Actual[len(S.Lista_Usuarios_Actual)-1]=""
		S.Lista_Usuarios_Actual = S.Lista_Usuarios_Actual[:len(S.Lista_Usuarios_Actual)-1]
	}

	/* fmt.Println("****"+usuario+" ha cerrado sesión.")
	fmt.Println("****Ahora la lista de usuarios es: ")
	fmt.Println(S.Lista_Usuarios_Actual) */

	S.Historial_Usuarios = append(S.Historial_Usuarios,"■■■■■ SERVER")
	S.Historial_Mensajes = append(S.Historial_Mensajes,usuario+" ha cerrado sesión.")

	aux:= "■■■■■ SERVER: "+usuario+" ha cerrado sesión."
	S.Full_Convo = append(S.Full_Convo,aux)

	*reply = "Has cerrado sesión. Adios!"
	return nil
	
}

func main() {
	println("STARTING SERVER...")
	//VARIABLES
	S.Enviado = true
	op := 0
	//ECHAR A ANDAR EL SERVER
	go server()
	println("Server up.")
	//MENU
	for{
		fmt.Println("------------------------SERVIDOR------------------------")
		fmt.Println("1.-Ver Conversación completa")
		fmt.Println("2.-Ver usuarios")
		fmt.Println("3.-Crear BACKUP de los mensajes")
		fmt.Println("9.-Cerrar Servidor")
		fmt.Printf("-----Ingrese su opción: ")
		fmt.Scanln(&op)
		switch op{
			case 1:
				VerConvo()
				break;
			case 2:
				VerUsuarios()
				break;
			case 3:
				CrearBackup()
				break;
			case 9:
				fmt.Println("Cerrando Servidor...")
				return
		}
	}
}
