package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"spider/utils"
)

func SaveImage(src string, filePath string) (name string) {
	res, err := http.Get(src)
	if err != nil {
		fmt.Printf("download image ERROR: %s, image src: %s", err, src)
		return ""
	}
	//创建目录
	if !isDirExist(filePath) {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			fmt.Printf("make dir ERROR: %s, filepath: %s", err, filePath)
		}
	}

	filename := filepath.Base(src)
	newFileName := utils.MD5(filename) + "." + utils.GetFileFormat(filename)
	fileFull := filePath + newFileName

	dst, err := os.Create(fileFull)
	defer dst.Close()

	if err != nil {
		fmt.Println("create file ERROR: %s, image src: %s", err, src)
		return ""
	}

	if _, err := io.Copy(dst, res.Body); err != nil {
		return ""
	} else {
		return newFileName
	}
}

func isDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}
