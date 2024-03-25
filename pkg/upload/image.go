package upload

import (
	"fmt"
	"github.com/kingsill/gin-example/pkg/file"
	"github.com/kingsill/gin-example/pkg/logging"
	"github.com/kingsill/gin-example/pkg/setting"
	"github.com/kingsill/gin-example/pkg/util"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// GetImageName 计算MD5加密之后的图片名
func GetImageName(name string) string {
	//将图片的名字剥离扩展名
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)

	//对单纯的图片名进行MD5加密
	fileName = util.EncodeMD5(fileName)

	//将MD5加密后的图片名和后缀返回
	return fileName + ext
}

// GetImagePath 包装文件路径 upload/images/
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageFullPath 拼凑完整路径 runtime/+upload/images/
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt 检查图片格式是否正确
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		//都大写进行对比
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckImageSize 检查图片的大小是否小于规定的最大值 5M
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)

	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	//检查图片目录
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	//检查访问权限
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
