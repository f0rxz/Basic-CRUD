package repos

import (
	"backend/storage"
)

// Структура для хранения информации о карточке
type Card struct {
	Id          int64
	Title       string
	Description string
	Rating      int
	Duration    string
	Difficulty  int
	Type        string
	Image       []byte
	Genres      string
}

// Создание новой карточки
func (c *Card) CreateCard(store *storage.Storage,title string, description string, rating int, duration string, difficulty int, variation string, image []byte, genre string) (bool, error) {
	query := "INSERT INTO Card_Quest (Title, Description, Rating, Duration, Difficulty, Type, Image, Genres) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	rows, err := store.Db.Query(query, title, description, rating, duration, difficulty, variation, image, genre)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return true, nil
}

// Удаление карточки по ID
func (c *Card) DeleteCard(store *storage.Storage,id int64) (int64, error) {
	query := "DELETE FROM Card_Quest WHERE id = ?"
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

// Модификация карточки по ID
func (c *Card) UpdateCard(store *storage.Storage,id int64, title string, description string, rating int, duration string, difficulty int, variation string, image []byte, genre string) (int64, error) {
	query := "UPDATE Card_Quest SET Title=?, Description=?, Rating=?, Duration=?, Difficulty=?, Type=?, Image=?, Genres=? WHERE id = ?"
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

// Получение карточки по ID
func (c *Card) GetCardByID(store *storage.Storage,id int64) (*Card, error) {
	query := "SELECT * FROM Card_Quest WHERE id = ?"
	rows, err := store.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var card Card
	err = rows.Scan(&card.Id, &card.Title, &card.Description, &card.Rating, &card.Duration, &card.Difficulty, &card.Type, &card.Image, &card.Genres)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

// Получение всех карточек
func (c *Card) GetAllCards(store *storage.Storage) ([]Card, error) {
	query := "SELECT * FROM Card_Quest"
	rows, err := store.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []Card
	for rows.Next() {
		var card Card
		err = rows.Scan(&card.Id, &card.Title, &card.Description, &card.Rating, &card.Duration, &card.Difficulty, &card.Type, &card.Image, &card.Genres)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
