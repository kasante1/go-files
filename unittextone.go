package main
 
import (
	"log"
	"os"
)
 
func main() {
	emptyFile, err := os.Create("empty1.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()

}