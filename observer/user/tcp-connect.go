package user

import (
	"fmt"
	"net"
)

func ConectarAoServidorDoNatan(data any) {
	fmt.Println("Enviando dados para o servidor do Natan...")

	var user = data.(*User)
	var message = fmt.Sprintf("Usuário %s foi criado!", user.Name)

	conn, _ := net.Dial("tcp", "192.168.137.184:8080")

	conn.Write([]byte(message))
}
