package filer

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"study_room_management_backend/config"
	"time"
)

var (
	cfg = config.C.Storage
)

func GenObjectKey(folder string, uid int64, ext string) string {
	return fmt.Sprintf("%s/%d_%d%s", folder, uid, time.Now().Unix(), ext)
}

func SaveLocal(r io.Reader, key string) (string, error) {
	p := path.Join(cfg.Local.Path, key)
	if err := os.MkdirAll(path.Dir(p), 0755); err != nil {
		return "", err
	}
	f, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	if _, err := io.Copy(f, r); err != nil {
		return "", err
	}
	// 返回可访问的 URL，这里用 nginx 挂静态目录
	return fmt.Sprintf("http://"+viper.GetString("server.port")+"/static/%s", key), nil
}
