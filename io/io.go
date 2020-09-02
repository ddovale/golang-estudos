package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// readFile()
	// readFileWithSmallChunks()
	//readFilePerLine()
	// createFile()
}

func readFile() {

	//This way only will find the file if this go file was run from the source folder of test.txt
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Content of file:", string(file))

	//Another approachs - https://golangbot.com/read-files/
	//https://www.it-swarm.dev/pt/go/golang-como-ler-um-arquivo-de-texto/823313613/#:~:text=Para%20obter%20o%20conte%C3%BAdo%20do,de%20um%20descritor%20de%20arquivo.&text=Caso%20contr%C3%A1rio%2C%20voc%C3%AA%20tamb%C3%A9m%20pode,seu%20arquivo%20em%20tokens%3A%20separator.
}

func readFileWithSmallChunks() {

	//This way only will find the file if this go file was run from the source folder of test.txt
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//Read 3 characters per round
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}

func readFilePerLine() {
	//This way only will find the file if this go file was run from the source folder of test.txt
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func createFile() {
	//The way to format dates in GO is so strange...
	f, err := os.Create(time.Now().Format("20060102150405") + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Println(l, "bytes written successfully")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
