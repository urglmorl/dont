package store

import (
	"constants"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"models"
	"utils"
)

var (
	subjectStore  = SubjectStore{}
	themeStore    = ThemeStore{}
	questionStore = QuestionStore{}
	optionStore   = OptionStore{}
)

func init() {
	// Проверка работоспособности базы данных при запуске
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)
}

func openDatabase() (*sql.DB, error) {
	database, err := sql.Open(constants.DatabaseDriverName, constants.DatabasePath)
	if err != nil {
		return &sql.DB{}, err
	}
	return database, nil
}

func closeDatabase(database *sql.DB) {
	err := database.Close()
	utils.CheckError(err)
}

func CreateSubject(subject models.Subject) (models.Subject, error) {
	return subjectStore.create(subject)
}

func CreateTheme(theme models.Theme) (models.Theme, error) {
	return themeStore.create(theme)
}

func CreateQuestion(question models.Question) (models.Question, error) {
	return questionStore.create(question)
}

func CreateOption(option models.Option, questionId uuid.UUID) (models.Option, error) {
	return optionStore.create(option, questionId)
}

func GetSubjects() ([]models.Subject, error) {
	return subjectStore.getAll()
}

func GetThemes() ([]models.Theme, error) {
	return themeStore.getAll()
}

func GetQuestions() ([]models.Question, error) {
	return questionStore.getAll()
}

func GetOptions() ([]models.Option, error) {
	return optionStore.getAll()
}

func GetSubjectById(id uuid.UUID) (models.Subject, error) {
	return subjectStore.getById(id)
}

func GetThemeById(id uuid.UUID) (models.Theme, error) {
	return themeStore.getById(id)
}

func GetQuestionById(id uuid.UUID) (models.Question, error) {
	return questionStore.getById(id)
}

func GetOptionById(id uuid.UUID) (models.Option, error) {
	return optionStore.getById(id)
}

func GetThemesBySubjectId(id uuid.UUID) ([]models.Theme, error) {
	return themeStore.getBySubjectId(id)
}

func UpdateSubject(subject models.Subject) (models.Subject, error) {
	return subjectStore.update(subject)
}

func UpdateTheme(theme models.Theme) (models.Theme, error) {
	return themeStore.update(theme)
}

func UpdateQuestion(question models.Question) (models.Question, error) {
	return questionStore.update(question)
}

func UpdateOption(option models.Option, questionId uuid.UUID) (models.Option, error) {
	return optionStore.update(option, questionId)
}

func DeleteSubject(subject models.Subject) error {
	return subjectStore.delete(subject)
}

func DeleteTheme(theme models.Theme) error {
	return themeStore.delete(theme)
}

func DeleteQuestion(question models.Question) error {
	return questionStore.delete(question)
}

func DeleteOption(option models.Option) error {
	return optionStore.delete(option)
}
