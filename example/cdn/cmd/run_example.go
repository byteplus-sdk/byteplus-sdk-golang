package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	. "github.com/byteplus-sdk/byteplus-sdk-golang/example/cdn"
)

type example struct {
	name string
	f    func(t *testing.T)
}

var examples = []example{
	{name: "AddCdnDomain", f: AddCdnDomain},
	{name: "ListCdnDomains", f: ListCdnDomains},
	{name: "DescribeCdnData", f: DescribeCdnData},
}

func main() {
	var t = &testing.T{}
	if len(os.Args) <= 1 {
		funs := make([]string, 0)
		for _, item := range examples {
			funs = append(funs, item.name)
		}
		fmt.Printf("./run_example %s\n", strings.Join(funs, " "))
		return
	}

	fmt.Printf("API Endpoint: %s \n", DefaultInstance.Client.ServiceInfo.Host)

	for idx, item := range examples {
		examples[idx].name = strings.ToLower(item.name)
	}

	for _, fun := range os.Args[1:] {
		fun = strings.ToLower(strings.TrimSpace(fun))
		if fun == "" {
			continue
		}
		found := false
		for _, item := range examples {
			if item.name == fun {
				fmt.Printf("running example: %s\n", fun)
				found = true
				item.f(t)
				break
			}
		}
		if found {
			continue
		}
		fmt.Printf("Unknown example name: %s\n", fun)
	}
}
