package file

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

// GetSize multipart.file用于处理HTTP请求中文件上传到类型   os.file则主要是本地文件的操作
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)
	return len(content), err
}

// GetExt 获取文件扩展名
func GetExt(filename string) string {
	return path.Ext(filename)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) bool {
	//os.stat用于获取文件的相关信息
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission 检查访问文件的权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	//检查是否有访问文件的权限
	return os.IsPermission(err)
}

// IsNotExistMkDir 检查是否存在目录，不存在则创建目录
func IsNotExistMkDir(src string) error {

	if notExist := CheckExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 创建目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm) //权限0777，权限拉满
	if err != nil {
		return err
	}

	return nil
}

// Open 算是简单包装os.openfile
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
