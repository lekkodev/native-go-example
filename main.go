package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lekkodev/native-go-example/internal/lekko"
)

var contents = map[string]string{
	"/home":      "welcome!",
	"/protected": "for your eyes only",
}

func main() {
	endpoint := os.Args[1]
	isUserAdmin := os.Args[2] == "true"

	ctx := context.Background()
	lc := lekko.NewLekkoClient(ctx)

	// Because we have named our namespace "default", we have a public field "Default" on the generated client
	canAccess := lc.Default.GetCanAccess(endpoint, isUserAdmin)
	if !canAccess {
		fmt.Println("access denied!")
		return
	}
	resp, ok := contents[endpoint]
	if !ok {
		fmt.Println("not found")
		return
	}
	fmt.Println(resp)
}
