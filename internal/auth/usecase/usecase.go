package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/auth"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/models"
	"github.com/mafzaidi/elog/pkg/authorizer"
	"github.com/mafzaidi/elog/pkg/authorizer/masterkey"
	"github.com/mafzaidi/elog/pkg/authorizer/pwd"
	"github.com/mafzaidi/elog/pkg/authorizer/token"
)

type UserToken struct {
	User   *models.User
	Token  string
	Claims *authorizer.Claims
}

type AuthUC struct {
	repo auth.Repository
}

func NewAuthUseCase(repo auth.Repository) auth.UseCase {
	return &AuthUC{
		repo: repo,
	}
}

func (u *AuthUC) Register(pl *auth.RegisterPayload) error {
	if len(pl.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	_, err := u.repo.FindByUsername(pl.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := pwd.Hash(pl.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	masterKey, err := masterkey.Generate()
	if err != nil {
		return errors.New("failed to generate master key")
	}

	encrypted, err := masterkey.Encrypt(masterKey, pl.Password)
	if err != nil {
		return errors.New("failed to encrypt  master key")
	}

	user := &entities.User{
		Username:     pl.Username,
		Fullname:     pl.FullName,
		PhoneNumber:  pl.PhoneNumber,
		Password:     hashedPassword,
		Email:        pl.Email,
		Group:        "user",
		MasterKeyEnc: encrypted.EncodedCipher,
		Salt:         encrypted.EncodedSalt,
	}

	return u.repo.Create(user)
}

func (u *AuthUC) Login(email, password, validToken string, cfg *config.Config) (*auth.UserToken, error) {

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user, err := u.repo.FindByEmail(email)
	if err != nil || !pwd.CheckHash(user.Password, password) {
		return nil, errors.New("email or password is invalid")
	}

	var claims *authorizer.Claims
	if validToken != "" {
		claims, _ = token.Validate(validToken, cfg.JWT.Secret)
	}

	if claims != nil && claims.Email == email {
		ut := &auth.UserToken{
			User:   user,
			Token:  validToken,
			Claims: claims,
		}
		return ut, nil
	}

	if _, err := masterkey.Decrypt(user.MasterKeyEnc, password, user.Salt); err != nil {
		return nil, errors.New("failed to decrypt master key")
	}

	claims = &authorizer.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.App.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Group:    user.Group,
	}

	g := &token.JWTGen{
		Secret: cfg.JWT.Secret,
		Claims: claims,
	}

	token, err := token.Generate(g)

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("failed to generate jwt")
	}

	ut := &auth.UserToken{
		User:   user,
		Token:  token,
		Claims: claims,
	}

	return ut, nil
}

func (u *AuthUC) ConfirmPassword(userID string, password string) error {
	return nil
}
