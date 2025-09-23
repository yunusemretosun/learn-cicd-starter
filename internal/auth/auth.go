package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
// bu fonksiyon geriye string ve error döner
// string: API anahtarı
// error: hata durumu
// fonksiyonun amacı, HTTP istek başlıklarından API anahtarını çıkarmaktır
// eğer başlıklar içinde "Authorization" başlığı yoksa, fonksiyon ErrNoAuthHeaderIncluded hatasını döner
// eğer "Authorization" başlığı varsa, başlığın değeri boş mu diye kontrol eder
// eğer boşsa, "malformed authorization header" hatasını döner
// eğer başlık değiri boşsa, değeri boşluk karakterine göre böler
// eğer bölünmüş parçaların sayısı 2'den azsa veya ilk parça "ApiKey" değilse, "malformed authorization header" hatasını döner
// eğer tüm kontroller geçilirse, ikinci parçayı (yani API anahtarını) döner
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
