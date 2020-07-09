package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	go sum(ctx, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func sum(ctx context.Context, digit ...int) int {
	var result int
	fmt.Println("Start sum()")
	for _, zahl := range digit {
		fmt.Println("Hallo")
		result += zahl
		if zahl%2 == 0 {
			time.Sleep(1 + time.Second)
		}
		select {
		case <-ctx.Done():
			result = math.MaxInt32
			fmt.Println("Break in sum")
			return result
		default:
			fmt.Println("Default")
		}
	}
	fmt.Println("Fertig gezählt")
	return result
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
