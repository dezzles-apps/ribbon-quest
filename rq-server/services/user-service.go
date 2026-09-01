package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/db"
	"errors"
	"strings"

	cdb "github.com/dezzles-apps/go-common/db"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	configService *ConfigService
	database      *cdb.Database
}

func NewUserService(
	database *cdb.Database,
	configService *ConfigService,
) *UserService {
	return &UserService{
		configService: configService,
		database:      database,
	}
}

func (us *UserService) LoginUser(input model.AuthInput) error {
	user, err := us.GetUserByUsername(input.Username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return errors.New("invalid username or password")
	}
	return nil
}

func (us *UserService) RegisterUser(input model.AuthInput) error {
	if !us.configService.CanRegisterUsers() {
		return errors.New("user registration is disabled")
	}
	input.Username = strings.ToLower(input.Username)
	existingUser, err := us.GetUserByUsername(input.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = us.database.GetDB().Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", input.Username, hashedPassword)
	if err != nil {
		return err
	}
	return nil

}

func (us *UserService) GetUserByUsername(username string) (*db.User, error) {
	username = strings.ToLower(username)
	var user db.User
	err := us.database.GetDB().QueryRow("SELECT username, password_hash FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
