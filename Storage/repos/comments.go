package repos

// // Структура для хранения информации о комментарии
// type Comment struct {
// 	Id     int64
// 	Author string
// 	Rating int
// 	Text   string
// }

// // Создание нового комментария для карточки
// func (c *Comment) CreateComment(author string, rate int, text string) (bool, error) {
// 	query := "INSERT INTO Comments (Author, Rating, Text) VALUES (?, ?, ?)"
// 	rows, err := c.Storage.Db.Query(query, author, rate, text)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer rows.Close()

// 	return true, nil
// }

// // Удаление комментария по ID
// func (c *Comment) DeleteComment(id int64) (int64, error) {
// 	query := "DELETE FROM Comments WHERE id = ?"
// 	stmt, err := c.Storage.Db.Prepare(query)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rowsAffected, nil
// }

// // Получение всех комментариев к карточке по ID
// func (c *Comment) GetCommentByCardID(cardID int64) ([]Comment, error) {
// 	query := "SELECT * FROM Comments WHERE card_id = ?"
// 	rows, err := c.Storage.Db.Query(query, cardID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var comments []Comment
// 	for rows.Next() {
// 		var comment Comment
// 		err := rows.Scan(&comment.Id, &comment.Author, &comment.Rating, &comment.Text)
// 		if err != nil {
// 			return nil, err
// 		}
// 		comments = append(comments, comment)
// 	}
// 	return comments, nil
// }

// // Удаление всех комментариев к карточке по ID
// func (c *Comment) DeleteComments(cardID int64) (int64, error) {
// 	query := "DELETE FROM comments WHERE card_id = ?"
// 	stmt, err := c.Storage.Db.Prepare(query)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(cardID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rowsAffected, nil
// }
