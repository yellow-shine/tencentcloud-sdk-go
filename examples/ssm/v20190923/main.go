package main

import (
	"errors"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

func main() {
	// provider, err := common.DefaultTkeOIDCRoleArnProvider()
	// if err != nil {
	// 	panic(err)
	// }

	durationSeconds := int64(10)
	slog.Info("durationSeconds", "durationSeconds", durationSeconds)

	provider, err := NewOIDCRoleArnProvider(durationSeconds)
	if err != nil {
		panic(err)
	}
	credential, err := provider.GetCredential()
	if err != nil {
		panic(err)
	}
	cre, ok := credential.(*common.RoleArnCredential)
	if !ok {
		panic("credential is not a RoleArnCredential")
	}
	for {
		cre.GetCredential()
		slog.Info("get credential", "SecretId", cre.GetSecretId(), "SecretKey", cre.GetSecretKey(), "Token", cre.GetToken())
		// fmt.Printf("SecretId: %s, SecretKey: %s, Token: %s\n", cre.GetSecretId(), cre.GetSecretKey(), cre.GetToken())
		time.Sleep(1 * time.Second)
	}

}

func NewOIDCRoleArnProvider(durationSeconds int64) (*common.OIDCRoleArnProvider, error) {
	reg := os.Getenv("TKE_REGION")
	if reg == "" {
		return nil, errors.New("env TKE_REGION not exist")
	}

	providerId := os.Getenv("TKE_PROVIDER_ID")
	if providerId == "" {
		return nil, errors.New("env TKE_PROVIDER_ID not exist")
	}

	tokenFile := os.Getenv("TKE_WEB_IDENTITY_TOKEN_FILE")
	if tokenFile == "" {
		return nil, errors.New("env TKE_WEB_IDENTITY_TOKEN_FILE not exist")
	}
	tokenBytes, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, err
	}

	roleArn := os.Getenv("TKE_ROLE_ARN")
	if roleArn == "" {
		return nil, errors.New("env TKE_ROLE_ARN not exist")
	}

	sessionName := "test-session" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)

	// provider.region = region
	// provider.providerId = providerId
	// provider.webIdentityToken = string(tokenBytes)
	// provider.roleArn = roleArn
	// provider.roleSessionName = sessionName
	return common.NewOIDCRoleArnProvider(
		reg,
		providerId,
		string(tokenBytes),
		roleArn,
		sessionName,
		durationSeconds,
	), nil
}
