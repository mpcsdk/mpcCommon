package kms

import (
	"fmt"
	"net/http"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	csms "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1"

	// "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	// "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	// "github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
	csmsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/region"
	// "https://iam.cn-north-4.myhuaweicloud.com"
	// type KMShuawei struct {
	// 	projectId string
	// 	region    string
	// 	ak        string
	// 	sk        string
	// 	keyName   string
	// 	keyVal    string
	// }
	// func NewKMS(procjectId string, region string, ak string, sk string, key string, val string) *KMShuawei {
	// 	s := &KMShuawei{
	// 		projectId: procjectId,
	// 		region:    region,
	// 		ak:        ak,
	// 		sk:        sk,
	// 		keyName:   key,
	// 		keyVal:    val,
	// 	}
	// 	s.init()
	// }
	// func (s *KMShuawei)
	// func kms() {
	// 	// 配置认证信息
	// 	auth, err := basic.NewCredentialsBuilder().
	// 		// 可通过环境变量等方式配置认证信息，参考2.4认证信息管理章节
	// 		WithAk("WSX7DMRH7TRFIFJX7WRL").
	// 		WithSk("lYVb8VpPWmDagh5WZYjZ2e5tIBBiBvqEZM0ja7lS").
	// 		// 如果未填写ProjectId，SDK会自动调用IAM服务查询所在region对应的项目id
	// 		WithProjectId("8e89f41d12d0409fbfcf27ec30b5cb4a").
	// 		// 配置SDK内置的IAM服务地址，默认为https://iam.myhuaweicloud.com
	// 		WithIamEndpointOverride("https://iam.cn-north-4.myhuaweicloud.com").
	// 		SafeBuild()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// 使用默认配置
	// 	httpConfig := config.DefaultHttpConfig()
	// 	// 配置是否忽略SSL证书校验， 默认不忽略
	// 	httpConfig.WithIgnoreSSLVerification(true)
	// 	// 默认超时时间为120秒，可根据需要配置
	// 	httpConfig.WithTimeout(120 * time.Second)
	// 	// 根据需要配置网络代理
	// 	// proxy := config.NewProxy().
	// 	// 	// 请根据实际情况替换示例中的代理协议、地址和端口号
	// 	// 	WithSchema("http").
	// 	// 	WithHost("proxy.huaweicloud.com").
	// 	// 	WithPort(80).
	// 	// 	// 如果代理需要认证，请配置用户名和密码
	// 	// 	WithUsername(os.Getenv("PROXY_USERNAME")).
	// 	// 	WithPassword(os.Getenv("PROXY_PASSWORD"))
	// 	// httpConfig.WithProxy(proxy)
	// 	// 根据需要配置如何创建网络连接
	// 	// dialContext := func(ctx context.Context, network string, addr string) (net.Conn, error) {
	// 	// 	// 此处需自行实现
	// 	// }
	// 	// httpConfig.WithDialContext(dialContext)
	// 	// 配置HTTP监听器进行调试，请勿用于生产环境
	// 	requestHandler := func(request http.Request) {
	// 		fmt.Println("request:", request)
	// 	}
	// 	responseHandler := func(response http.Response) {
	// 		fmt.Println("resp:", response)
	// 	}
	// 	httpHandler := httphandler.NewHttpHandler().AddRequestHandler(requestHandler).AddResponseHandler(responseHandler)
	// 	httpConfig.WithHttpHandler(httpHandler)
	// 	// 获取可用地区
	// 	region, err := csmsRegion.SafeValueOf("ap-southeast-3")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// 创建服务客户端
	// 	hcClient, err := csms.CsmsClientBuilder().
	// 		WithRegion(region).
	// 		WithCredential(auth).
	// 		WithHttpConfig(httpConfig).
	// 		SafeBuild()
	// 	// hcClient, err := vpc.VpcClientBuilder().
	// 	// 	// 配置地区, 如果地区不存在会导致panic
	// 	// 	WithRegion(region).
	// 	// 	// 配置认证信息
	// 	// 	WithCredential(auth).
	// 	// 	// HTTP配置
	// 	// 	WithHttpConfig(httpConfig).
	// 	// 	SafeBuild()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	client := csms.NewCsmsClient(hcClient)
	// 	// 创建请求
	// 	request := &model.ListSecretsRequest{}
	// 	// request := &vpcModel.ListVpcsRequest{}
	// 	// // 配置每页返回的个数
	// 	// limit := int32(1)
	// 	// request.Limit = &limit
	// 	// 发送请求并获取响应
	// 	response, err := client.ListSecrets(request)
	// 	// 处理异常，打印响应信息
	// 	if err == nil {
	// 		fmt.Printf("%+v\n", response)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// }
	// ///
	// csms "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1"
	// "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/model"
	// csmsRegion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/csms/v1/region"
)

type KMShuawei struct {
	projectId string
	region    string
	ak        string
	sk        string
	//
	debug bool
	////
	auth *basic.Credentials
	cli  *csms.CsmsClient
}

func NewKMS(
	procjectId string,
	region string,
	ak string,
	sk string,
	debug bool) (*KMShuawei, error) {
	s := &KMShuawei{
		projectId: procjectId,
		region:    region,
		ak:        ak,
		sk:        sk,
		debug:     debug,
	}
	err := s.init()
	if err != nil {
		return nil, err
	}

	return s, nil
}
func (s *KMShuawei) init() error {
	// 配置认证信息
	auth, err := basic.NewCredentialsBuilder().
		// 可通过环境变量等方式配置认证信息，参考2.4认证信息管理章节
		WithAk(s.ak).
		WithSk(s.sk).
		// 如果未填写ProjectId，SDK会自动调用IAM服务查询所在region对应的项目id
		WithProjectId(s.projectId).
		// 配置SDK内置的IAM服务地址，默认为https://iam.myhuaweicloud.com
		// WithIamEndpointOverride("https://iam.cn-north-4.myhuaweicloud.com").
		SafeBuild()
	if err != nil {
		return err
	}
	s.auth = auth
	////
	// 使用默认配置
	httpConfig := config.DefaultHttpConfig()
	// 配置是否忽略SSL证书校验， 默认不忽略
	httpConfig.WithIgnoreSSLVerification(true)
	// 默认超时时间为120秒，可根据需要配置
	httpConfig.WithTimeout(120 * time.Second)

	requestHandler := func(request http.Request) {
		if s.debug {
			fmt.Println("\nrequest:", request)
			fmt.Println()
		}
	}
	responseHandler := func(response http.Response) {
		if s.debug {
			fmt.Println("\nresp:", response)
			fmt.Println()
		}
	}
	httpHandler := httphandler.NewHttpHandler().AddRequestHandler(requestHandler).AddResponseHandler(responseHandler)
	httpConfig.WithHttpHandler(httpHandler)

	// 获取可用地区
	region, err := csmsRegion.SafeValueOf(s.region)
	if err != nil {
		return err
	}
	// 创建服务客户端
	hcClient, err := csms.CsmsClientBuilder().
		WithRegion(region).
		WithCredential(auth).
		WithHttpConfig(httpConfig).
		SafeBuild()
	// hcClient, err := vpc.VpcClientBuilder().
	// 	// 配置地区, 如果地区不存在会导致panic
	// 	WithRegion(region).
	// 	// 配置认证信息
	// 	WithCredential(auth).
	// 	// HTTP配置
	// 	WithHttpConfig(httpConfig).
	// 	SafeBuild()
	if err != nil {
		return err
	}
	client := csms.NewCsmsClient(hcClient)
	s.cli = client
	return nil
}

// //
func (s *KMShuawei) ListKeys(marker string) (*model.ListSecretsResponse, error) {
	request := &model.ListSecretsRequest{
		Marker: &marker,
	}
	// request := &vpcModel.ListVpcsRequest{}
	// // 配置每页返回的个数
	// limit := int32(1)
	// request.Limit = &limit
	// 发送请求并获取响应
	response, err := s.cli.ListSecrets(request)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *KMShuawei) AddKey(key, val string) (*model.CreateSecretResponse, error) {
	request := &model.CreateSecretRequest{
		Body: &model.CreateSecretRequestBody{
			Name:         key,
			SecretString: &val,
		},
	}
	///
	response, err := s.cli.CreateSecret(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *KMShuawei) DelKey(key string) (*model.DeleteSecretResponse, error) {
	request := &model.DeleteSecretRequest{
		SecretName: key,
	}
	///
	response, err := s.cli.DeleteSecret(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (s *KMShuawei) GetKey(key string) (*model.ShowSecretVersionResponse, error) {
	request := &model.ShowSecretVersionRequest{
		SecretName: key,
		VersionId:  "latest",
	}
	///
	response, err := s.cli.ShowSecretVersion(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
