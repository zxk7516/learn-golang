package model

type Student struct {
	Name  string
	Score float64
}
type student struct {
	Name  string
	score float64
}

func New_student(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

func (stu *student) Get_score() float64 {
	return stu.score
}
