package helper

import (
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz1234567890"
const length = 36

var screatkey = []byte("Si kepo hahaha")

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type Helper struct {
}


func (u *Helper) GetTimeNow() string {
	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05.999999"))
}

func (u *Helper) SaveUploadedFile(file multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}