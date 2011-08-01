package main

import (
	"fmt"
	"syscall"
)

const DATA_DIR = "data"
const DELIMITER = ":"
const FILENAME = DATA_DIR + "/wal"
// Q: Why won't := work here?
var data = make(map[string] string)

func get(key string) string{
	return data[key]
}


// TODO: This should return an error
func put(key string, val string){
	// Update the in-memory data	
	data[key] = val
	// Write to the WAL

	//TODO: Create the dirs and data file if they don't already exist
	
	fd, err := syscall.Open(DATA_DIR + "/wal", 2, 0600)
	if err != 0{
		fmt.Printf("Cannot open your db.data file\n")
	}
	
	_, err2 := syscall.Write(fd, []byte(key + DELIMITER + val))
	if err2 != 0{
		fmt.Printf("Error writing to your file %v\n", err)
	}
}

//Loads the WAL and makes the in-memory data store consistent
func loadDB() {
	fd, err := os.Open(FILENAME)

	if err != 0{
		fmt.Printf("Cannot open your db.data file\n")
	}
	
	_, err2 := syscall.Read(fd, 
}

func main(){

	put("foo", "bar")
	put("neutral", "milk")
	fmt.Printf("%v\n", get("foo"))
	fmt.Printf("%v\n", get("neutral"))
	//fmt.Printf("%v\n", data["a"])
}