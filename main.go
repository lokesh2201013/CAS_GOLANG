package main

import (
	"fmt"

	"github.com/lokesh2201013/CAS_GOLANG/p2p"
)

func main(){
	tr:=p2p.NewTCPTransport(":3000")
	if err :=tr.ListenAndAccept();err!=nil{
		log.Fatal(err)
	}
	
	select{}
 	fmt.Println("lets go")
}