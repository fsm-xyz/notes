# OSS

OSS 不支持直接使用“用户名”和“密码”进行代码上传。

## 权限

### AK访问

```go
package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	// ----------------------------------------------------------------------
	// 1. 配置信息 (建议通过环境变量读取，不要直接写死在代码中)
	// ----------------------------------------------------------------------
	// Endpoint: 请根据你的 Bucket 所在地域填写，例如杭州是 oss-cn-hangzhou.aliyuncs.com
	endpoint := "https://oss-cn-hangzhou.aliyuncs.com"
	
	// AccessKeyId 和 AccessKeySecret: 你的“API 用户名”和“API 密码”
	accessKeyID := "LTAI5txxxxxxxxxxxx"     
	accessKeySecret := "456yyyyyyyyyyyyyyyyyy" 

	// BucketName: 你的存储空间名称
	bucketName := "my-awesome-bucket"

	// ----------------------------------------------------------------------
	// 2. 初始化 Client
	// ----------------------------------------------------------------------
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		handleError(err)
	}

	// ----------------------------------------------------------------------
	// 3. 获取存储空间 (Bucket)
	// ----------------------------------------------------------------------
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}

	// ----------------------------------------------------------------------
	// 4. 上传文件
	// ----------------------------------------------------------------------
	// 参数1 (objectName): 上传到 OSS 后的文件名（包含路径），例如 "images/2023/avatar.jpg"
	// 参数2 (localFile):  本地需要上传的文件路径，例如 "./local_data.jpg"
	objectName := "test/hello_oss.txt"
	localFile := "./local_file.txt"

	// 简单上传：FromFile
	err = bucket.PutObjectFromFile(objectName, localFile)
	if err != nil {
		handleError(err)
	}

	fmt.Printf("文件 %s 成功上传到 OSS 路径 %s\n", localFile, objectName)
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
```
### 签名访问

当需要通过STS访问，需要创建一个虚拟角色，进行授权，生成签名，客户端可以根据这个直接进行访问

代表一个虚拟的 RAM 角色 (Role),读取 被扮演的角色 (Role) 身上的权限策略

为了解决AK的权限过大，虚拟角色进行限制

```go
package main

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// --------------------------------------------------------------------------
// 配置区域
// --------------------------------------------------------------------------
const (
	// 1. 长期身份凭证 (拥有 AliyunSTSAssumeRoleAccess 权限的 RAM 用户)
	AccessKeyID     = "LTAI5txxxxxxxxxx"
	AccessKeySecret = "456yyyyyyyyyyyyyyyyyy"

	// 2. RAM 角色信息 (该角色需拥有 OSS 读写权限)
	// 在 RAM 控制台 -> 角色 -> 点击角色名 -> 基本信息里复制 ARN
	RoleArn = "acs:ram::1234567890123:role/my-oss-upload-role"

	// 3. OSS 设置
	Endpoint   = "https://oss-cn-hangzhou.aliyuncs.com"
	BucketName = "my-awesome-bucket"
)

func main() {
	// ----------------------------------------------------------------------
	// 第一步：模拟【服务端】，向阿里云申请 STS 临时凭证
	// ----------------------------------------------------------------------
	fmt.Println("正在获取 STS 临时凭证...")
	stsCreds, err := getSTSToken()
	if err != nil {
		fmt.Printf("获取 STS 失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("成功获取 STS Token!")
	fmt.Printf("临时 AK: %s...\n", stsCreds.AccessKeyId[0:5])

	// ----------------------------------------------------------------------
	// 第二步：模拟【客户端】，使用刚才拿到的 STS 凭证上传文件
	// ----------------------------------------------------------------------
	fmt.Println("正在使用 STS 凭证连接 OSS...")
	
	// !!! 核心区别：初始化时传入 STS Token !!!
	client, err := oss.New(
		Endpoint,
		stsCreds.AccessKeyId,
		stsCreds.AccessKeySecret,
		oss.SecurityToken(stsCreds.SecurityToken), // 关键参数
	)
	if err != nil {
		handleError(err)
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		handleError(err)
	}

	// 上传文件
	objectName := "test/sts_upload_demo.txt"
	localFile := "./local_file.txt" // 请确保本地有这个文件，或者创建一个
	
	// 为了演示，创建一个临时文件
	createDummyFile(localFile)

	err = bucket.PutObjectFromFile(objectName, localFile)
	if err != nil {
		// 如果这里报 403，说明 RoleArn 对应的角色没有 OSS 权限
		handleError(err)
	}

	fmt.Printf("上传成功！文件已存入: %s/%s\n", BucketName, objectName)
	
	// 清理临时文件
	os.Remove(localFile)
}

// getSTSToken 调用阿里云 STS 服务获取临时凭证
func getSTSToken() (*sts.Credentials, error) {
	// 初始化 STS 客户端 (RegionID 一般写 cn-hangzhou 即可，STS 是全局服务)
	stsClient, err := sts.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 构建请求
	request := sts.CreateAssumeRoleRequest()
	request.RoleArn = RoleArn
	request.RoleSessionName = "go-sdk-upload-session" // 自定义会话名，用于审计
	request.DurationSeconds = "3600"                  // 有效期，单位秒 (900~3600)

	// 发起请求
	response, err := stsClient.AssumeRole(request)
	if err != nil {
		return nil, err
	}

	return &response.Credentials, nil
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func createDummyFile(filename string) {
	f, _ := os.Create(filename)
	f.WriteString("Hello OSS with STS!")
	f.Close()
}

```