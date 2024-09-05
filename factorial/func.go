package main

import (
	"context"
	"encoding/json"
	//"fmt"
	"io"
	"log"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(myHandler))
	//fmt.Println(factorialNumber(6))
}

type Fact struct {
	Number int `json:"Number"`
}

func myHandler(ctx context.Context, in io.Reader, out io.Writer) {
	f := &Fact{Number: 0}
	json.NewDecoder(in).Decode(f)

	//num := 	
	msg := struct {
		Msg int `json:"factorial"`
	}{
		//Msg: fmt.Sprintf("Hello %s", p.Number),

		
		Msg: factorialNumber(f.Number),
	}
	log.Print("Inside Go factorial function")
	json.NewEncoder(out).Encode(&msg)
}

func factorialNumber(number int) int {

	if number == 1{
		return 1
	}

	factOfNumber := number * factorialNumber(number-1)

	return factOfNumber
}

