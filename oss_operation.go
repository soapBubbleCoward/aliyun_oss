package aliyun_oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOssClient struct {
	Domain string
	OriginalFileName bool
	Client *oss.Bucket
}

//推送文件到oss
//params:  ossDir string  `oss dir [要推送到的oss目录]`  example: test/20201121/
//params:  file interface `upload file resource [文件资源]`
//return string  `oss file accessible uri [可访问地址]`
func (client *AliOssClient) Put(ossDir string, file interface{}) string {
	//file to []byte
	//文件转字节流
	uploadFile := &OssFile{
		File: file,
	}

	ossFile,err := uploadFile.FileTypeTransForm()

	if err != nil {
		panic("transfer file failed" + err.Error())
	}

	//ossPath = oss dir + upload file name
	//example: oss dir is diy ==== test/20201121/
	//time.Now().Format("20060102")
	//ossPath := path + fileName
	var ossPath string

	//judge is use origin file name if false fileName = fileNewName (is a only name) else file init name
	if client.OriginalFileName == false {
		ossPath = ossDir + ossFile.FileNewName
	} else {
		ossPath = ossDir + ossFile.FileOldName
	}

	//upload file to oss
	err = client.Client.PutObject(ossPath,bytes.NewReader(ossFile.FileByte))

	if err != nil {
		panic("put file to oss failed:" + err.Error())
	}

	return client.Domain + "/" + ossPath
}

//校验文件是否已经存在
//check file already exists in oss server
//params: ossFilePath	string 	`file oss path [文件的oss的路径]`
func (client *AliOssClient) HasExists(ossFilePath string) bool {

	//oss check fun
	isExists,err := client.Client.IsObjectExist(ossFilePath)

	if err != nil {
		panic("check file in oss is exists failed:" + err.Error())
	}

	return isExists
}

//删除文件-单文件删除
//delete one file in oss
//params ossPath string `file oss path [文件的oss路径]`
//return bool
func (client *AliOssClient) Delete(ossFilePath string) bool {

	//oss delete one file fun
	err := client.Client.DeleteObject(ossFilePath)

	if err != nil {
		panic("delete file "+ ossFilePath +" failed:" + err.Error())
	}

	return true
}

//删除文件-多文件删除
//delete more file in oss
//params ossPath []string `file oss path array [文件的oss路径数组]`
//return bool
func (client *AliOssClient) DeleteMore(ossFilePath []string) bool {
	//oss delete more file fun
	_,err := client.Client.DeleteObjects(ossFilePath)

	if err != nil {
		panic("delete more file in oss failed:" + err.Error())
	}

	return true
}