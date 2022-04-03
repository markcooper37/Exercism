package ledger

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	var entriesCopy []Entry
	// Simplified appending of entries to copy
	entriesCopy = append(entriesCopy, entries...)
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}
	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy
	for len(es) > 1 {
		first, rest := es[0], es[1:]
		success := false
		for !success {
			success = true
			for i, e := range rest {
				if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
					m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
					m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
					es[0], es[i+1] = es[i+1], es[0]
					success = false
				}
			}
		}
		es = es[1:]
	}

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
			if len(entry.Date) != 10 {
				co <- Data{e: errors.New("")}
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' || d4 != '-' {
				co <- Data{e: errors.New("")}
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
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
				if locale == "nl-NL" {
					for i := len(parts) - 1; i >= 0; i-- {
						a += parts[i] + "."
					}
					a = a[:len(a)-1] + "," + centsStr[len(centsStr)-2:]
					if negative {
						a += "-"
					} else {
						a += " "
					}
				} else if locale == "en-US" {
					for i := len(parts) - 1; i >= 0; i-- {
						a += parts[i] + ","
					}
					a = a[:len(a)-1] + "." + centsStr[len(centsStr)-2:]
					if negative {
						a += ")"
					} else {
						a += " "
					}
				}
			}
			var al int
			for range a {
				al++
			}
			co <- Data{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
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
