package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type profile_item map[string]string

func main() {
	f, ok := os.OpenFile("./profiles.txt", os.O_RDONLY, 0)

	if ok != nil {
		fmt.Println("Open resd file failed")
		return
	}

	defer f.Close()

	jfile, err := os.OpenFile("./profiles.json", os.O_CREATE|os.O_RDWR, 0)

	if err != nil {
		fmt.Println("Open write file failed")
		return
	}

	defer jfile.Close()

	profiles := make(map[string]*profile_item, 50)

	j := 0
	keys := make([]string, 6)
	items := make([]string, 6)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			parts := strings.Split(line, ": ")
			keys[j] = parts[0]
			items[j] = parts[1]
			j++
		} else {
			profile := profile_item{keys[0]: items[0], keys[1]: items[1], keys[2]: items[2], keys[3]: items[3], keys[4]: items[4], keys[5]: items[5]}
			profiles[items[0]] = &profile
			j = 0
		}

	}

	for key, value := range profiles {
		fmt.Printf("%s\n", key)
		fmt.Println(*value)
	}

	jsonString, _ := json.Marshal(profiles)
	fmt.Println(string(jsonString))

	jfile.Write(jsonString)
}
