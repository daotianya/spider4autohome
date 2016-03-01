package conf

var DB = make(map[string]string)

const IMGRoot = "/d/images/"
const LOGRoot = "D:/d/log"

func init() {
	DB["user"] = "root"
	DB["pass"] = "123456"
	DB["name"] = "spider"
	DB["host"] = "localhost"
	DB["port"] = "3306"
	DB["char"] = "utf8"
}
