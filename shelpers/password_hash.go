package shelpers

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"strings"
)

var SECRET = "THIS()IS$(\"A\\##\\HARDCODED3X4MPL3S3CR3T"
var SECRET_SEED = 1337331
var SYN = "%password%defaultvaluesarecool%secret%%password%21"

func Generate_Hash(password string) string {
	tmp := ""
	tmp = strings.Replace(SYN, "%password%", password, -1)
	tmp = strings.Replace(tmp, "%secret%", SECRET, -1)
	tmp = strings.Replace(tmp, "1", "비밀$μυστικό", -1)
	tmp = strings.Replace(tmp, "2", "@비밀", -1)
	tmp = strings.Replace(tmp, "3", "\\", -1)
	tmp = strings.Replace(tmp, "4", "秘密$)μυστικό=", -1)
	tmp = strings.Replace(tmp, "5", "$비밀μυστικό(/%", -1)
	tmp = strings.Replace(tmp, "6", "\",μυστικό%)", -1)
	tmp = strings.Replace(tmp, "7", "$\")\"μυστικό)秘密", -1)
	tmp = strings.Replace(tmp, "8", "''비밀$)μυστικό/=비밀§!$", -1)
	tmp = strings.Replace(tmp, "9", "''비밀$)/秘密fdf§!$", -1)
	r := []rune(tmp)

	for index, char := range r {
		rand.Seed(int64(SECRET_SEED + index))
		r[index] += char * rune(index + int(SECRET[rand.Intn(len(SECRET))]))
	}

	sh := sha512.New()
	sh.Write([]byte(string(r)))
	return hex.EncodeToString(sh.Sum(nil))
}
