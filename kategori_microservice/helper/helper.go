package helper

import (
	"math/rand"
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
