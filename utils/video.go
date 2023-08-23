package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func IsMP4File(filePath string) bool {
	extension := strings.ToLower(filepath.Ext(filePath))
	return extension == ".mp4"
}

// 读取视频文件中指定帧并将其作为JPEG图像返回
func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:%d.jpg", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		panic(err)
	}
	return buf
}

// 截取视频的第一帧作为封面
func SaveVideoImg(videoPath string, saveImgPath string) error {
	reader := ReadFrameAsJpeg(videoPath, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		return errors.New("图像解码失败")
	}
	err = imaging.Save(img, saveImgPath)
	if err != nil {
		return errors.New("图片保存到本地失败")
	}
	return nil
}

func ReplaceFileExtension(filePath string) string {
	extension := filepath.Ext(filePath)
	newFilePath := strings.TrimSuffix(filePath, extension) + ".jpeg"
	return newFilePath
}

func GetFileUrl(fileName string) (string, error) {
	Cfg := model.Cfg
	ip := Cfg.Section("server").Key("ip").String()
	port, err := Cfg.Section("server").Key("port").Int()
	if err != nil {
		return "", err
	}
	baseUrl := fmt.Sprintf("http://%s:%d/static/%s", ip, port, fileName)
	return baseUrl, nil
}
