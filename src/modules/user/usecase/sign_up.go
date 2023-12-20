package usecase

import (
	"context"
	"os"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (us *UseCase) SignUp(ctx context.Context, user entity.User) (string, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	newUser, err := us.repo.User.StoreUser(ctx, user)
	if err != nil {
		return "", err
	}

	// Set waktu kedaluwarsa token menjadi satu menit dari waktu sekarang
	expirationTime := time.Now().Add(10 * time.Minute)

	// Buat klaim token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   newUser.Id,
	}

	// Buat token baru
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key (ganti "your_secret_key" dengan key sebenarnya)
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
