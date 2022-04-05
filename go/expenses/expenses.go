package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
    for _, record := range in {
        if predicate(record) {
            out = append(out, record)
        }
    }
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
        return record.Day >= p.From && record.Day <= p.To
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
        return record.Category == c
    }
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    total := 0.0
	filteredRecords := Filter(in, ByDaysPeriod(p))
    for _, record := range filteredRecords {
        total += record.Amount
    }
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	total := 0.0
	categoryRecords := Filter(in, ByCategory(c))
    if len(categoryRecords) == 0 {
        return 0, errors.New("no records match the given category")
    }
	filteredRecords := Filter(categoryRecords, ByDaysPeriod(p))
    for _, record := range filteredRecords {
        total += record.Amount
    }
	return total, nil
}
