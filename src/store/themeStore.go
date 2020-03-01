package store

import (
	"github.com/google/uuid"
	"models"
	"utils"
)

type ThemeStore struct {
}

func (t ThemeStore) create(theme models.Theme) (models.Theme, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	theme.Id = utils.UUIDs()

	stmt, err := database.Prepare("INSERT INTO themes(id, name, subjectId) VALUES(?,?,?)")
	if err != nil {
		return models.Theme{}, err
	}

	result, err := stmt.Exec(theme.Id.String(), theme.Name, theme.SubjectId)
	if err != nil {
		return models.Theme{}, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return models.Theme{}, err
	}

	return theme, nil
}

func (t ThemeStore) getById(id uuid.UUID) (models.Theme, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	theme := models.Theme{}

	row := database.QueryRow("SELECT * FROM themes WHERE id=$1", id)

	err = row.Scan(&theme.Id, &theme.Name, &theme.SubjectId)
	if err != nil {
		return models.Theme{}, err
	}

	return theme, nil
}

func (t ThemeStore) getBySubjectId(id uuid.UUID) ([]models.Theme, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var themes []models.Theme

	rows, err := database.Query("SELECT * FROM themes WHERE subjectId=$1", id)
	if err != nil {
		return []models.Theme{}, err
	}

	for rows.Next() {
		var theme = models.Theme{}
		err = rows.Scan(&theme.Id, &theme.Name, &theme.SubjectId)
		if err != nil {
			return []models.Theme{}, err
		}
		themes = append(themes, theme)
	}

	err = rows.Close()
	if err != nil {
		return []models.Theme{}, err
	}

	return themes, nil
}

func (t ThemeStore) getAll() ([]models.Theme, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var themes []models.Theme

	rows, err := database.Query("SELECT * FROM themes")
	if err != nil {
		return []models.Theme{}, err
	}

	for rows.Next() {
		var theme = models.Theme{}
		err = rows.Scan(&theme.Id, &theme.Name, &theme.SubjectId)
		if err != nil {
			return []models.Theme{}, err
		}
		themes = append(themes, theme)
	}

	err = rows.Close()
	if err != nil {
		return []models.Theme{}, err
	}

	return themes, nil
}

func (t ThemeStore) getQuestionsByThemeId(theme models.Theme) (models.Theme, error) {
	var (
		q   QuestionStore
		err error
	)

	theme.Questions, err = q.getByThemeId(theme.Id)
	if err != nil {
		return models.Theme{}, err
	}

	return theme, nil
}

func (t ThemeStore) update(theme models.Theme) (models.Theme, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("UPDATE themes SET name=?, subjectId=? WHERE id=?")
	if err != nil {
		return models.Theme{}, err
	}

	result, err := stmt.Exec(theme.Name, theme.SubjectId)
	if err != nil {
		return models.Theme{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return models.Theme{}, err
	}

	return theme, nil
}

func (t ThemeStore) delete(theme models.Theme) error {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("DELETE FROM themes WHERE id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(theme.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
