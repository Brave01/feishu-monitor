package main

import "feishu-monitor/internal"

func main() {
	err := internal.AddFsMultiRecords()
	if err != nil {
		panic(err)
	}
}
