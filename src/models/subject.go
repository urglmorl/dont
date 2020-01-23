package models

type SubjectBson struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type Subject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SubjectsBson []SubjectBson
type Subjects []Subject

func (subject *Subject) ToBson() (subjectBson SubjectBson) {
	subjectBson.Id = subject.Id
	subjectBson.Name = subject.Name
	return
}

func (subjectBson *SubjectBson) ToJson() (subject Subject) {
	subject.Id = subjectBson.Id
	subject.Name = subjectBson.Name
	return
}

func (subjectsBson SubjectsBson) ToJson() (subjects Subjects) {
	for _, subjectBson := range subjectsBson {
		subjects = append(subjects, subjectBson.ToJson())
	}
	return
}

func (subjects Subjects) ToBson() (subjectsBson SubjectsBson) {
	for _, subject := range subjects {
		subjectsBson = append(subjectsBson, subject.ToBson())
	}
	return
}
