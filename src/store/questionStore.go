package store

import (
	"github.com/google/uuid"
	"models"
	"utils"
)

type QuestionStore struct {
}

func (q QuestionStore) create(question models.Question) (models.Question, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	question.Id = utils.UUIDs()

	stmt, err := database.Prepare("INSERT INTO questions(id, name, subjectId, themeId) VALUES(?,?,?,?)")
	if err != nil {
		return models.Question{}, err
	}

	result, err := stmt.Exec(question.Id.String(), question.Name, question.SubjectId, question.ThemeId)
	if err != nil {
		return models.Question{}, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func (q QuestionStore) getById(id uuid.UUID) (models.Question, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	question := models.Question{}

	row := database.QueryRow("SELECT * FROM questions WHERE id=$1", id)

	err = row.Scan(&question.Id, &question.Name, &question.SubjectId, &question.ThemeId)
	if err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func (q QuestionStore) getByThemeId(themeId uuid.UUID) ([]models.Question, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var questions []models.Question

	rows, err := database.Query("SELECT * FROM questions WHERE themeId=$1", themeId.String())
	if err != nil {
		return []models.Question{}, err
	}

	for rows.Next() {
		var question = models.Question{}
		err = rows.Scan(&question.Id, &question.Name, &question.SubjectId, &question.ThemeId)
		if err != nil {
			return []models.Question{}, err
		}
		questions = append(questions, question)
	}

	err = rows.Close()
	if err != nil {
		return []models.Question{}, err
	}

	return questions, nil
}

func (q QuestionStore) getAll() ([]models.Question, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var questions []models.Question

	rows, err := database.Query("SELECT * FROM questions")
	if err != nil {
		return []models.Question{}, err
	}

	for rows.Next() {
		var question = models.Question{}
		err = rows.Scan(&question.Id, &question.Name, &question.SubjectId, &question.ThemeId)
		if err != nil {
			return []models.Question{}, err
		}
		questions = append(questions, question)
	}

	err = rows.Close()
	if err != nil {
		return []models.Question{}, err
	}

	return questions, nil
}

func (q QuestionStore) getOptions(question models.Question) (models.Question, error) {
	var (
		o   OptionStore
		err error
	)

	question.Options, err = o.getByQuestionId(question.Id)
	if err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func (q QuestionStore) update(question models.Question) (models.Question, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("UPDATE questions SET name=?, subjectId=?, themeId=? WHERE id=?")
	if err != nil {
		return models.Question{}, err
	}

	result, err := stmt.Exec(question.Name, question.SubjectId, question.ThemeId)
	if err != nil {
		return models.Question{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return models.Question{}, err
	}

	return question, nil
}

func (q QuestionStore) delete(question models.Question) error {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("DELETE FROM questions WHERE id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(question.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
