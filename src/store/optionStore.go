package store

import (
	"github.com/google/uuid"
	"models"
	"utils"
)

type OptionStore struct {
}

func (o OptionStore) create(option models.Option, questionId uuid.UUID) (models.Option, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	option.Id = utils.UUIDs()

	stmt, err := database.Prepare("INSERT INTO options(id, name, isRight, questionId) VALUES(?,?,?,?)")
	if err != nil {
		return models.Option{}, err
	}

	result, err := stmt.Exec(option.Id.String(), option.Name, option.IsRight, questionId.String())
	if err != nil {
		return models.Option{}, err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return models.Option{}, err
	}

	return models.Option{}, nil
}

func (o OptionStore) getById(id uuid.UUID) (models.Option, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	option := models.Option{}

	row := database.QueryRow("SELECT * FROM options WHERE id=$1", id)

	err = row.Scan(&option.Id, &option.Name, option.IsRight)
	if err != nil {
		return models.Option{}, err
	}

	return option, nil
}

func (o OptionStore) getByQuestionId(questionId uuid.UUID) ([]models.Option, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var options []models.Option

	rows, err := database.Query("SELECT * FROM options WHERE questionId=$1", questionId.String())
	if err != nil {
		return []models.Option{}, nil
	}

	for rows.Next() {
		var option = models.Option{}
		err = rows.Scan(&option.Id, &option.Name, &option.IsRight)
		if err != nil {
			return []models.Option{}, err
		}
		options = append(options, option)
	}

	err = rows.Close()
	if err != nil {
		return []models.Option{}, err
	}

	return options, nil
}

func (o OptionStore) getAll() ([]models.Option, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	var options []models.Option

	rows, err := database.Query("SELECT * FROM options")
	if err != nil {
		return []models.Option{}, err
	}

	for rows.Next() {
		var option = models.Option{}
		err = rows.Scan(&option.Id, &option.Name, &option.IsRight)
		if err != nil {
			return []models.Option{}, err
		}
		options = append(options, option)
	}

	err = rows.Close()
	if err != nil {
		return []models.Option{}, err
	}

	return options, nil
}

func (o OptionStore) update(option models.Option, questionId uuid.UUID) (models.Option, error) {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("UPDATE options SET name=?, questionId=? WHERE id=?")
	if err != nil {
		return models.Option{}, err
	}

	result, err := stmt.Exec(option.Name, questionId)
	if err != nil {
		return models.Option{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return models.Option{}, err
	}

	return option, nil
}

func (o OptionStore) delete(option models.Option) error {
	database, err := openDatabase()
	utils.CheckError(err)
	defer closeDatabase(database)

	stmt, err := database.Prepare("DELETE FROM options WHERE id=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(option.Id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return err
}
