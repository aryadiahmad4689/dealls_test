package usecase

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/aryadiahmad4689/dealls_test/src/modules/user/entity"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (us *UseCase) SignIn(ctx context.Context, req entity.User) (string, error) {
	// Ambil user berdasarkan email
	user, err := us.repo.User.GetUserByEmail(ctx, req.Email)
	if err == sql.ErrNoRows {
		// Jika user tidak ditemukan atau error lainnya
		return "", errors.New("email not found")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// Jika password tidak sesuai
		return "", errors.New("invalid password")
	}

	// Set waktu kedaluwarsa token menjadi satu menit dari waktu sekarang
	expirationTime := time.Now().Add(10 * time.Minute)

	// Buat klaim token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   user.Id,
	}

	// Buat token baru
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_KEY")))
	if err != nil {
		return "", errors.New("token failed created")
	}

	return tokenString, nil
}
