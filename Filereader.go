package main

import (
	"bytes"
	"fmt"
	"log"
	"encoding/json"
  	"io"
	"strings"
	"os"
)

var (
	target string = "5.101.118.148"
	total int = 0
)

func popLine(f *os.File) ([]byte, error) {
    	fi, err := f.Stat()
    	if err != nil {
       		 return nil, err
    	}
    	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

    	_, err = f.Seek(0, os.SEEK_SET)
    	if err != nil {
      		  return nil, err
   	 }
    	_, err = io.Copy(buf, f)
    	if err != nil {
      		  return nil, err
   	}
   	line, err := buf.ReadString('\n')
    	if err != nil && err != io.EOF {
        	return nil, err
    	}

    	_, err = f.Seek(0, os.SEEK_SET)
    	if err != nil {
        	return nil, err
    	}
    	nw, err := io.Copy(f, buf)
    	if err != nil {
        	return nil, err
    	}
    	err = f.Truncate(nw)
    	if err != nil {
        	return nil, err
    	}
    	err = f.Sync()
    	if err != nil {
        	return nil, err
    	}

    	_, err = f.Seek(0, os.SEEK_SET)
    	if err != nil {
    	    	return nil, err
    	}
    	return []byte(line), nil
}

func decode() int{
	const jsonStream = `
		
  		
	`
	type Message struct {
		ip string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", m.ip)
		
		if m.ip == target {
			total ++
		}
	}
	return total
}


func main() {
	 fname := `20150817-090000-926.txt`
   	 f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
   	 if err != nil {
     	   fmt.Println(err)
     	   return
   	 }
  	  defer f.Close()
   	 line, err := popLine(f)
   	 if err != nil {
        	fmt.Println(err)
       	 return
    	}
    	fmt.Println("pop:", string(line))
	//print("The total amount of ips matching 5.101.118.148 is: " , decode())

}