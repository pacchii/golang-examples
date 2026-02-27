package _8_interfaces

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type student struct {
	firstName string
}

func (s student) writeOut(w io.Writer) {
	w.Write([]byte(s.firstName))
}

func WriterExample() {
	SimpleWrite()
	BufferedWriter()
	WriteToFile()
}

func WriteToFile() {

	s := student{"Prashant"}

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var b bytes.Buffer

	s.writeOut(f)
	s.writeOut(&b)

	fmt.Println(b.String()) //Prashant
}

func BufferedWriter() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String()) //Hello!
	b.WriteString("World!")

	fmt.Println(b.String()) //Hello World!

	//Reset
	b.Reset()
	b.WriteString("Its Friday Today!!!!!!")
	fmt.Println(b.String()) //Its Friday Today!!!!!!
}

func SimpleWrite() {

	f, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	s := []byte("Hello World")

	_, er := f.Write(s)
	if er != nil {
		log.Fatal(er)
	}
}
