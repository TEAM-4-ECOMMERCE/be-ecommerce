package helpers

import (
	"strconv"
	"time"
)

func TFCode(numberUnique uint) string {
	y, m, d := time.Now().Date()

	year := strconv.Itoa(y)
	month := strconv.Itoa(int(m))
	day := strconv.Itoa(d)
	unique := strconv.Itoa(int(numberUnique))
	epoch := strconv.Itoa(int(time.Now().Unix()))

	return "TF-" + year + month + day + epoch + unique
}