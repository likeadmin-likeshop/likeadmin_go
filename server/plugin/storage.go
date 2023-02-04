package plugin

import (
	"fmt"
	"io"
	"likeadmin/config"
	"likeadmin/core"
	"likeadmin/core/response"
	"likeadmin/util"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var StorageDriver = storageDriver{}

//UploadFile 文件对象
type UploadFile struct {
	Name string // 文件名称
	Type int    // 文件类型
	Size int64  // 文件大小
	Ext  string // 文件扩展
	Uri  string // 文件路径
	Path string // 访问地址
}

//storageDriver 存储引擎
type storageDriver struct{}

//Upload 根据引擎类型上传文件
func (sd storageDriver) Upload(file *multipart.FileHeader, folder string, fileType int) (uf *UploadFile, e error) {
	// TODO: engine默认local
	if e = sd.checkFile(file, fileType); e != nil {
		return
	}
	key := sd.buildSaveName(file)
	engine := "local"
	if engine == "local" {
		if e = sd.localUpload(file, key, folder); e != nil {
			return
		}
	} else {
		core.Logger.Errorf("storageDriver.Upload engine err: err=[unsupported engine]")
		return nil, response.Failed.Make(fmt.Sprintf("engine:%s 暂时不支持", engine))
	}
	fileRelPath := path.Join(folder, key)
	return &UploadFile{
		Name: file.Filename,
		Type: fileType,
		Size: file.Size,
		Ext:  strings.ToLower(strings.Replace(path.Ext(file.Filename), ".", "", 1)),
		Uri:  fileRelPath,
		Path: util.UrlUtil.ToAbsoluteUrl(fileRelPath),
	}, nil
}

//localUpload 本地上传 (临时方法)
func (sd storageDriver) localUpload(file *multipart.FileHeader, key string, folder string) (e error) {
	// TODO: 临时方法，后续调整
	// 映射目录
	directory := config.Config.UploadDirectory
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		core.Logger.Errorf("storageDriver.localUpload Open err: err=[%+v]", err)
		return response.Failed.Make("打开文件失败!")
	}
	defer src.Close()
	// 文件信息
	savePath := path.Join(directory, folder, path.Dir(key))
	saveFilePath := path.Join(directory, folder, key)
	// 创建目录
	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		core.Logger.Errorf(
			"storageDriver.localUpload MkdirAll err: path=[%s], err=[%+v]", savePath, err)
		return response.Failed.Make("创建上传目录失败!")
	}
	// 创建目标文件
	out, err := os.Create(saveFilePath)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Create err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.Make("创建文件失败!")
	}
	defer out.Close()
	// 写入目标文件
	_, err = io.Copy(out, src)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		return response.Failed.Make("上传文件失败: " + err.Error())
	}
	return nil
}

//checkFile 生成文件名称
func (sd storageDriver) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	date := time.Now().Format("20060201")
	return path.Join(date, util.ToolsUtil.MakeUuid()+ext)
}

//checkFile 文件验证
func (sd storageDriver) checkFile(file *multipart.FileHeader, fileType int) (e error) {
	fileName := file.Filename
	fileExt := strings.ToLower(strings.Replace(path.Ext(fileName), ".", "", 1))
	fileSize := file.Size
	if fileType == 10 {
		// 图片文件
		if !util.ToolsUtil.Contains(config.Config.UploadImageExt, fileExt) {
			return response.Failed.Make("不被支持的图片扩展: " + fileExt)
		}
		if fileSize > config.Config.UploadImageSize {
			return response.Failed.Make("上传图片不能超出限制: " + strconv.FormatInt(config.Config.UploadImageSize/1024/1024, 10) + "M")
		}
	} else if fileType == 20 {
		// 视频文件
		if !util.ToolsUtil.Contains(config.Config.UploadVideoExt, fileExt) {
			return response.Failed.Make("不被支持的视频扩展: " + fileExt)
		}
		if fileSize > config.Config.UploadVideoSize {
			return response.Failed.Make("上传视频不能超出限制: " + strconv.FormatInt(config.Config.UploadVideoSize/1024/1024, 10) + "M")
		}
	} else {
		core.Logger.Errorf("storageDriver.checkFile fileType err: err=[unsupported fileType]")
		return response.Failed.Make("上传文件类型错误")
	}
	return nil
}
