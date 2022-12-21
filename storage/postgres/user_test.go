package postgres_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/bxcodec/faker/v4"
// 	"github.com/stretchr/testify/require"
// 	"github.com/mirasildev/blog/storage/repo"
// )

// func createUser(t *testing.T) *repo.User {
// 	u := createUser(t)
// 	user, err := strg.User().Create(&repo.User{
// 		FirstName: faker.FirstName(),
// 		LastName: faker.LastName(),
// 		Email: faker.Email(),
// 		Password: faker.Password(),
// 		Username: faker.Username(),
// 		Type: "user",
// 		CreatedAt:   time.Now(),
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)

// 	return user
// }

// func deleteNote(id int64, t *testing.T) {
// 	err := strg.User().DeleteUser(id)
// 	require.NoError(t, err)
// }

// func updateNote(t *testing.T) {
// 	u := createUser(t)
// 	note, err := strg.User().UpdateUser(&repo.User{
// 		UserID:      u.ID,
// 		Title:       faker.Sentence(),
// 		Description: faker.Sentence(),
// 		UpdatedAt:   time.Now(),
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, note)

// 	deleteNote(note.ID, t)
// }
// func TestCreateNote(t *testing.T) {
// 	note := createNote(t)
// 	deleteUser(note.ID, t)
// 	require.NotEmpty(t, note)
// }

// func TestGetNote(t *testing.T) {
// 	c := createNote(t)

// 	Note, err := strg.Note().Get(c.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, Note)
// }

// func TestUpdateNote(t *testing.T) {
// 	updateNote(t)
// }

// func TestGetAllNote(t *testing.T) {

// 	Notes, err := strg.Note().GetAllNotes(&repo.GetAllNotesParams{
// 		Limit: 10,
// 		Page:  1,
// 	})
// 	require.NoError(t, err)
// 	require.NotEmpty(t, Notes)

// }
