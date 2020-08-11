package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	listAll("../../", 0)
}

func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tempHier := curHier; tempHier > 0; tempHier-- {
				fmt.Print("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tempHier := curHier; tempHier > 0; tempHier-- {
				fmt.Print("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
