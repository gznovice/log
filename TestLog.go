package main

import(
	"log"
	"strings"
	"os"
)

func main(){
	log.Println("I am a first msg")
	
	
	log.SetPrefix("Prefix: ")
	
	log.Println("I am a second msg")
	
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	
	myslice := []string{"the", "third", "msg"}
	
	log.Printf("I am %v.\n", strings.Join(myslice, " "))
	
	file, e := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
      log.Fatalln("Failed to open log file")
   }
 
   log.SetOutput(file)
	
   log.Println("I am the 4th msg")
   
   log.Panic("Warning:we lost...")
}