package main

import (
	"context"
	"fmt"
	"log"
)


func main(){
	store,err:=NewMongoStore(context.Background())
	if err!=nil{
		log.Fatalf("Unable to connect to database due to error : %v",err)
	}
	server:=NewApiServer(":5000",store)	
	fmt.Println("Server listening on port ",server.listenAddr)
	server.Start()
	
}