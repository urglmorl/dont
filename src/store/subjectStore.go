package store

import (
	"github.com/google/uuid"
	"models"
	"utils"
)

type SubjectStore struct {
}

func (s SubjectStore) create(subject models.Subject) (models.Subject, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	subject.Id = utils.UUIDs()

	stmt, err := database.Prepare("INSERT INTO subjects(id, name) VALUES(?,?)")
	if err != nil {
		return models.Subject{}, err
	}

	result, err := stmt.Exec(subject.Id.String(), subject.Name)
	if err != nil {
		return models.Subject{}, err
	}

	// TODO: логировать все изменения базы данных,
	// 	а не использовать заглушки по типу _
	_, err = result.LastInsertId()
	if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func (s SubjectStore) getById(id uuid.UUID) (models.Subject, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	subject := models.Subject{}

	row := database.QueryRow("SELECT * FROM subjects WHERE id=$1", id)

	err = row.Scan(&subject.Id, &subject.Name)
	if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func (s SubjectStore) getAll() ([]models.Subject, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var subjects []models.Subject

	rows, err := database.Query("SELECT * FROM subjects")
	if err != nil {
		return []models.Subject{}, err
	}

	for rows.Next() {
		var subject = models.Subject{}
		err = rows.Scan(&subject.Id, &subject.Name)
		if err != nil {
			return []models.Subject{}, err
		}
		subjects = append(subjects, subject)
	}

	err = rows.Close()
	if err != nil {
		return []models.Subject{}, err
	}

	return subjects, nil
}

func (s SubjectStore) getThemesBySubjectId(subject models.Subject) (models.Subject, error) {
	var (
		t   ThemeStore
		err error
	)

	subject.Themes, err = t.getBySubjectId(subject.Id)
	if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func (s SubjectStore) update(subject models.Subject) (models.Subject, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("UPDATE subjects SET name=? WHERE id=?")
	if err != nil {
		return models.Subject{}, err
	}

	result, err := stmt.Exec(subject.Name)
	if err != nil {
		return models.Subject{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return models.Subject{}, err
	}

	return subject, nil
}

func (s SubjectStore) delete(subject models.Subject) error {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("DELETE FROM subjects WHERE id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(subject.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
