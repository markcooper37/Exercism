package school

import (
	"sort"
)

// Define the Grade and School types here.
type Grade struct {
	grade    int
	students []string
}

type School []Grade

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	for index, schoolGrade := range *s {
		if schoolGrade.grade == grade {
			(*s)[index].students = append(schoolGrade.students, student)
			return
		}
	}
	*s = append(*s, Grade{grade: grade, students: []string{student}})
}

func (s *School) Grade(level int) []string {
	for _, grade := range *s {
		if grade.grade == level {
			return grade.students
		}
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	sort.Slice(*s, func(i, j int) bool { return (*s)[i].grade < (*s)[j].grade })
	for _, grade := range *s {
		sort.Strings(grade.students)
	}
	return *s
}
