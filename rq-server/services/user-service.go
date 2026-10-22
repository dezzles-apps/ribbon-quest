package services

import (
	"database/sql"
	"dezzles-apps/rq-server/model"
	"dezzles-apps/rq-server/model/db"
	"dezzles-apps/rq-server/model/errors"
	"strings"

	cdb "github.com/dezzles-apps/go-common/db"
	cmodel "github.com/dezzles-apps/go-common/model"
	"go.uber.org/zap"
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

func (us *UserService) LoginUser(ctx *cmodel.DAContext, input model.AuthInput) errors.Error {
	user, e := us.GetUserByUsername(ctx, input.Username)
	if e != nil {
		return e
	}
	if user == nil {
		return errors.IncorrectUsernameOrPassword
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return errors.IncorrectUsernameOrPassword
	}
	return nil
}

func (us *UserService) RegisterUser(ctx *cmodel.DAContext, input model.AuthInput) errors.Error {
	if !us.configService.CanRegisterUsers() {
		return errors.UserRegistrationDisabled
	}
	input.Username = strings.ToLower(input.Username)
	existingUser, e := us.GetUserByUsername(ctx, input.Username)
	if e == nil && existingUser != nil {
		return model.InvalidUsernameOrPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Logger.Error("Error occurred generating hashed password", zap.Error(err))
		return model.InternalServerError
	}

	_, err = us.database.GetDB().Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", input.Username, hashedPassword)
	if err != nil {
		ctx.Logger.Error("Error occurred saving user", zap.Error(err))
		return model.InternalServerError
	}
	return nil

}

func (us *UserService) GetUserByUsername(ctx *cmodel.DAContext, username string) (*db.User, errors.Error) {
	username = strings.ToLower(username)
	var user db.User
	err := us.database.GetDB().QueryRow("SELECT username, password_hash FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		ctx.Logger.Error("Error retrieving user", zap.Error(err))
		return nil, model.InternalServerError
	}
	return &user, nil
}
