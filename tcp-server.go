package main 


import "net"
import "fmt"
import "bufio"
import "strings" // only needed for sample processing 

func main(){
	
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port

	conn, _ := ln.Accept()

	//run loop 

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		fmt.Printf("Message Received:", string(message))

		// sample process for string received

		newMessage := strings.ToUpper(message)

		// send new string back to client 

		conn.Write([]byte(newMessage + "\n")) 
	}

}