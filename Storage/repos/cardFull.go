package repos

import (
	"backend/storage"
)

// Структура для хранения полной информации о карточке
type CardFull struct {
	Id          int64
	Title       string
	Description string
	Images      []byte
	Genres      string
	Rating      int
	Difficulty  int
	Duration    string
	Type        string
}

// Создание новой полной карточки (с комментариями)
func (c *CardFull) CreateCardFull(store *storage.Storage, title string, description string, rating int, duration string, difficulty, variation string, image []byte, genre string) (bool, error) {
	query := "INSERT INTO Card_Quest_Full (Title, Description, Rating, Duration, Difficulty, Type, Images, Genres) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	rows, err := store.Db.Query(query, title, description, rating, duration, difficulty, variation, image, genre)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

// Удаление карточки по ID
func (c *CardFull) DeleteCard(store *storage.Storage, id int64) (int64, error) {
	query := "DELETE FROM Card_Quest_Full WHERE id = ?"
	stmt, err := store.Db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// Модификация полной карточки (с комментариями) по ID

func (c *CardFull) UpdateCard(store *storage.Storage, id int64, title string, description string, rating int, duration string, difficulty int, variation string, image []byte, genre string) (int64, error) {
	query := "UPDATE Card_Quest_Full SET Title=?, Description=?, Rating=?, Duration=?, Difficulty=?, Type=?, Images=?, Genres=? WHERE id = ?"
	stmt, err := store.Db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, title, description, rating, duration, difficulty, variation, image, genre)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// Получение всех полных карточек (с комментариями)
func (c *CardFull) GetAllCardsFull(store *storage.Storage) ([]CardFull, error) {
	query := "SELECT * FROM Card_Quest_Full"
	rows, err := store.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []CardFull
	for rows.Next() {
		var card CardFull
		err := rows.Scan(&card.Id, &card.Title, &card.Description, &card.Rating, &card.Duration, &card.Difficulty, &card.Type, &card.Images, &card.Genres)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

// Получение полной карточки (с комментариями) по ID
func (c *CardFull) GetCardFullByID(store *storage.Storage, id int64) (*CardFull, error) {
	query := "SELECT * FROM Card_Quest_Full WHERE id = ?"
	rows, err := store.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var card CardFull
	err = rows.Scan(&card.Id, &card.Title, &card.Description, &card.Rating, &card.Duration, &card.Difficulty, &card.Type, &card.Images, &card.Genres)
	if err != nil {
		return nil, err
	}
	return &card, nil
}
