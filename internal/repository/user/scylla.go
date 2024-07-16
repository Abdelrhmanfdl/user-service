package user

import (
	"log"

	"github.com/Abdelrhmanfdl/user-service/internal/errs"
	"github.com/Abdelrhmanfdl/user-service/internal/models"
	"github.com/gocql/gocql"
)

type scyllaUser struct {
	ID        gocql.UUID `json:"id"`
	Username  string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt int        `json:"createdAt"`
}

type ScyllaUserRepository struct {
	session *gocql.Session
}

func NewScyllaUserRepository(scyllaURL string) *ScyllaUserRepository {
	cluster := gocql.NewCluster(scyllaURL)
	cluster.Keyspace = "chatchatgo"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to ScyllaDB:", err)
	}
	return &ScyllaUserRepository{session: session}
}

func (s *ScyllaUserRepository) CreateUser(user *models.User) (id string, err error) {
	query := `INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`
	id = gocql.TimeUUID().String()
	err = s.session.Query(query, id, user.Username, user.Email, user.Password).Exec()
	return id, classifyError(err)
}

func (s *ScyllaUserRepository) GetUserById(id string) (user *models.User, err error) {
	var userScylla scyllaUser
	query := `SELECT id, username, email FROM users WHERE id = ?`
	if err := s.session.Query(query, id).Scan(&userScylla.ID, &userScylla.Username, &userScylla.Email); err != nil {
		return nil, classifyError(err)
	}
	return &models.User{
		ID:       userScylla.ID.String(),
		Username: userScylla.Username,
		Email:    userScylla.Email,
	}, nil
}

func (s *ScyllaUserRepository) GetUsersByIds(userIds []string) (users []models.User, err error) {
	query := `SELECT id, username, email FROM users WHERE id = ?`
	iter := s.session.Query(query, userIds).Iter()

	var user models.User

	for iter.Scan(&user.ID, &user.Username, &user.Email) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *ScyllaUserRepository) GetUserByEmail(email string) (user *models.User, err error) {
	var userScylla scyllaUser
	query := `SELECT id, username, email, password FROM users WHERE email = ?`
	if err := s.session.Query(query, email).Scan(&userScylla.ID, &userScylla.Username, &userScylla.Email, &userScylla.Password); err != nil {
		return nil, classifyError(err)
	}

	return &models.User{
		ID:       userScylla.ID.String(),
		Username: userScylla.Username,
		Email:    userScylla.Email,
		Password: userScylla.Password,
	}, nil
}

func classifyError(err error) error {
	if err == gocql.ErrNotFound {
		return &errs.NotFoundError{Message: "not found"}
	}
	if err, ok := err.(gocql.RequestError); ok && err.Code() == gocql.ErrCodeUnavailable {
		return &errs.ConnectionError{Message: err.Error()}
	}
	if err, ok := err.(gocql.RequestErrReadTimeout); ok && err.Code() == gocql.ErrCodeReadTimeout {
		return &errs.TimeoutError{Message: err.Error()}
	}
	return err
}
