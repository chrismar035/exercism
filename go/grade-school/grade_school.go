package school

import "sort"

type attendance interface {
	Grade(int) []string
	Enrollment() []Grade
	Add(string, int)
}

type School struct {
	grades map[int][]string
}

type Grade struct {
	level int
	names []string
}

func New() *School {
	return &School{grades: make(map[int][]string)}
}

func (s School) Grade(level int) []string {
	return s.grades[level]
}

func (s School) Enrollment() []Grade {
	var enrollment []Grade
	for level, names := range s.grades {
		enrollment = append(enrollment, Grade{level: level, names: names})
	}
	sort.Sort(ByLevel(enrollment))
	return enrollment
}

func (s School) Add(name string, level int) {
	s.grades[level] = append(s.grades[level], name)
	sort.Strings(s.grades[level])
}

type ByLevel []Grade

func (a ByLevel) Len() int {
	return len(a)
}

func (a ByLevel) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByLevel) Less(i, j int) bool {
	return a[i].level < a[j].level
}
