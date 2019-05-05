package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type profile_item struct {
	entity_number   string
	entity_name     string
	incorp_date     string
	due_date        string
	report_date     string
	nature_business string
}

type NameSorter []*profile_item

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].entity_number < a[j].entity_number }

func main() {
	f, ok := os.OpenFile("./profiles.txt", os.O_RDONLY, 0)

	if ok != nil {
		fmt.Println("Open file failed")
		return
	}

	defer f.Close()

	profiles := make([]*profile_item, 50)

	i := 0
	j := 0
	items := make([]string, 6)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			parts := strings.Split(line, ": ")
			items[j] = parts[1]
			j++
		} else {
			profile := profile_item{items[0], items[1], items[2], items[3], items[4], items[5]}
			profiles[i] = &profile

			i++
			j = 0
		}

	}

	sort.Sort(NameSorter(profiles))

	for m := range profiles {
		if profiles[m] != nil {
			fmt.Printf("%d,%s,%s,%s\n", m, profiles[m].entity_number, profiles[m].entity_name, profiles[m].due_date)
		}
	}
}
