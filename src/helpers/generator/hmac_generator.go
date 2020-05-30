package generator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

// GenerateHMAC is generator for hmac
func GenerateHMAC(data interface{}) string {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}

	expirationTime := time.Now().AddDate(0, 0, 30)

	salt := GenerateSalt()
	hash := hmac.New(sha256.New, []byte(os.Getenv("SECRET_KEY_HMAC")))
	io.WriteString(hash, fmt.Sprintf("%v%s", data, salt))
	sha := hex.EncodeToString(hash.Sum(nil))

	return fmt.Sprintf("%s.%d", sha, expirationTime.Unix())
}
