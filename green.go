package main

import (
	"flag"
	"os"
	"time"
)

var fileName = flag.String("o", "README.md", "output file name")

func main() {
	flag.Parse()

	file, err := os.OpenFile(*fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(time.Now().In(loc).Format("🕰️ Updated on Monday, 02 January 2006 at 15:04 MST"))
	if err != nil {
		panic(err)
	}
}
