package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// Create new struct type to reduce repetition and increase clarity
type Data struct {
	i int
	s string
	e error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	// Simplified appending of entries to copy
	entriesCopy := append([]Entry{}, entries...)
	sort.Slice(entriesCopy, func(i, j int) bool { 
		return entriesCopy[i].Date < entriesCopy[j].Date ||
		(entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description < entriesCopy[j].Description) ||
		(entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description == entriesCopy[j].Description && entriesCopy[i].Change < entriesCopy[j].Change)
	})

	// Reduced repetition when finding strings that will be used in s
	var first, second, third string
	if locale == "nl-NL" {
		first, second, third = "Datum", "Omschrijving", "Verandering"
	} else if locale == "en-US" {
		first, second, third = "Date", "Description", "Change"
	} else {
		return "", errors.New("")
	}
	// Simplified expression to create s
	s := fmt.Sprintf("%-10s | %-25s | %s\n", first, second, third)

	// Parallelism, always a great idea
	co := make(chan Data)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
				co <- Data{e: errors.New("")}
			}
			d1, d3, d5 := entry.Date[0:4], entry.Date[5:7], entry.Date[8:10]
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = fmt.Sprintf("%-25s", de)
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale != "nl-NL" && locale != "en-US" {
				co <- Data{e: errors.New("")}
			} else {
				if negative && locale == "en-US" {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- Data{e: errors.New("")}
				}
				if locale == "nl-NL" {
					a += " "
				}
				// Added zeros more concisely
				centsStr := fmt.Sprintf("%03s", strconv.Itoa(cents))
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				var p1, p2 string
				if locale == "nl-NL" {
					p1, p2 = ".", ","
				} else if locale == "en-US" {
					p1, p2 = ",", "."
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + p1
				}
				a = a[:len(a)-1] + p2 + centsStr[len(centsStr)-2:]
				if negative && locale == "nl-NL" {
					a += "-"
				} else if negative && locale == "en-US" {
					a += ")"
				} else {
					a += " "
				}
			}
			co <- Data{i: i, s: fmt.Sprintf("%-10s | %s | %13s\n", d, de, a)}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
