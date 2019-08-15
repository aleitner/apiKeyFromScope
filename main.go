package main

import (
	"fmt"
	"os"
	"storj.io/storj/lib/uplink"
)

func main()  {
	if len(os.Args) <= 1 {
		fmt.Println("Must specify Scope.")
		return
	}

	scopeb58 := os.Args[1]
	scope, err := uplink.ParseScope(scopeb58)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(scope.APIKey.Serialize())
}
