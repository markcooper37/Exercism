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

// Data struct clarifies the data to be passes along the channel
type Data struct {
	index int
	entry string
	err   error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	// Sort entries
	entriesCopy := append([]Entry{}, entries...)
	sort.Slice(entriesCopy, func(i, j int) bool {
		return entriesCopy[i].Date < entriesCopy[j].Date ||
			(entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description < entriesCopy[j].Description) ||
			(entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description == entriesCopy[j].Description && entriesCopy[i].Change < entriesCopy[j].Change)
	})

	// Write column headers
	var first, second, third string
	if locale == "nl-NL" {
		first, second, third = "Datum", "Omschrijving", "Verandering"
	} else if locale == "en-US" {
		first, second, third = "Date", "Description", "Change"
	} else {
		return "", errors.New("")
	}
	ledger := fmt.Sprintf("%-10s | %-25s | %s\n", first, second, third)

	// Format entries using parallelism
	co := make(chan Data)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			entryString, err := FormatEntry(currency, locale, entry)
			co <- Data{index: i, entry: entryString, err: err}
		}(i, et)
	}

	// Collect formatted entries and add them to the ledger
	formattedEntries := make([]string, len(entriesCopy))
	for range entriesCopy {
		formattedEntry := <-co
		if formattedEntry.err != nil {
			return "", formattedEntry.err
		}
		formattedEntries[formattedEntry.index] = formattedEntry.entry
	}
	for i := 0; i < len(entriesCopy); i++ {
		ledger += formattedEntries[i]
	}
	return ledger, nil
}

func FormatEntry(currency string, locale string, entry Entry) (string, error) {
	date, err := FormatDate(currency, locale, entry)
	if err != nil {
		return "", err
	}
	description := FormatDescription(currency, locale, entry)
	change, err := FormatChange(currency, locale, entry)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%-10s | %s | %13s\n", date, description, change), nil
}

func FormatDate(currency string, locale string, entry Entry) (string, error) {
	if len(entry.Date) != 10 || entry.Date[4] != '-' || entry.Date[7] != '-' {
		return "", errors.New("invalid date")
	}
	year, month, day := entry.Date[0:4], entry.Date[5:7], entry.Date[8:10]
	var date string
	if locale == "nl-NL" {
		date = day + "-" + month + "-" + year
	} else if locale == "en-US" {
		date = month + "/" + day + "/" + year
	}
	return date, nil
}

func FormatDescription(currency string, locale string, entry Entry) string {
	description := entry.Description
	if len(description) > 25 {
		description = description[:22] + "..."
	} else {
		description = fmt.Sprintf("%-25s", description)
	}
	return description
}

func FormatChange(currency string, locale string, entry Entry) (string, error) {
	cents := entry.Change
	negative := cents < 0
	if cents < 0 {
		cents = cents * -1
	}
	var change string
	if locale != "nl-NL" && locale != "en-US" {
		return "", errors.New("invalid locale")
	} else {
		if negative && locale == "en-US" {
			change += "("
		}
		if currency == "EUR" {
			change += "€"
		} else if currency == "USD" {
			change += "$"
		} else {
			return "", errors.New("invalid currency")
		}
		if locale == "nl-NL" {
			change += " "
		}
		centsStr := fmt.Sprintf("%03s", strconv.Itoa(cents))
		var p1, p2 string
		if locale == "nl-NL" {
			p1, p2 = ".", ","
		} else if locale == "en-US" {
			p1, p2 = ",", "."
		}
		money := p2 + centsStr[len(centsStr)-2:]
		for i := len(centsStr) - 2; i > 0; i -= 3 {
			if i != len(centsStr)-2 {
				money = p1 + money
			}
			var p string
			if i >= 3 {
				p = centsStr[i-3 : i]
			} else {
				p = centsStr[0:i]
			}
			money = p + money
		}
		change += money
		if negative && locale == "nl-NL" {
			change += "-"
		} else if negative && locale == "en-US" {
			change += ")"
		} else {
			change += " "
		}
	}
	return change, nil
}
