package globalidentity

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"fmt"
)

const (
	contentJson = "application/json"
	validateApplicationSuffix = "/api/authorization/validateapplication"
	authenticateUserSuffix = "/api/authorization/authenticate"
	isUserInRolesSuffix = "/api/authorization/isuserinroles"
	validateTokenSuffix = "/api/authorization/validatetokenresponse"
	renewTokenSuffix = "/api/authorization/renewtoken"
)

type GlobalIdentityError []string

func (e GlobalIdentityError) Error() string {
	return fmt.Sprintf("%#v", []string(e))
}


type GlobalIdentityManager interface {
	AuthenticateUser(email string, password string, expirationInMinutes ... int) (string, error)
	ValidateToken(token string) (bool, error)
	IsUserInRoles(userKey string, roles ... string) (bool, error)
	RenewToken(token string) (string, error)
	ValidateApplication(clientApplicationKey string, clientSecretKey string, resources string) (bool, error)
}

type globalIdentityManager struct {
	applicationKey     string
	globalIdentityHost string
}

func New(applicationKey string, globalIdentityHost string) GlobalIdentityManager {
	return &globalIdentityManager{applicationKey,
		globalIdentityHost,
	}
}

func (gim *globalIdentityManager) AuthenticateUser(email string, password string, expirationInMinutes ... int) (string, error) {
	expirationInMinutes = append(expirationInMinutes, 15)
	request := &authenticateUserRequest{
		ApplicationKey: gim.applicationKey,
		TokenExpirationInMinutes: expirationInMinutes[0],
		Email: email,
		Password: password,
	}
	json, err := toJson(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(gim.globalIdentityHost + authenticateUserSuffix, contentJson, json)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response authenticateUserResponse

	err = fromJson(&response, resp.Body)
	if err != nil {
		return "", err
	}
	if !response.Success {
		err = GlobalIdentityError([]string{"Invalid email or password"})
	}
	return response.AuthenticationToken, err
}

func (gim *globalIdentityManager) ValidateToken(token string) (bool, error) {
	request := &validateTokenRequest{
		ApplicationKey:gim.applicationKey,
		Token: token,
	}
	json, err := toJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost + validateTokenSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response validateTokenResponse

	err = fromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}

func (gim *globalIdentityManager) IsUserInRoles(userKey string, roles ... string) (bool, error) {
	request := &isUserInHolesRequest{
		ApplicationKey:gim.applicationKey,
		UserKey: userKey,
		RoleCollection: roles,
	}
	json, err := toJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost + isUserInRolesSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response isUserInRoleResponse

	err = fromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}

func (gim *globalIdentityManager) RenewToken(token string) (string, error) {
	request := &renewTokenRequest{
		ApplicationKey:gim.applicationKey,
		Token: token,
	}
	json, err := toJson(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(gim.globalIdentityHost + renewTokenSuffix, contentJson, json)
	if err != nil {
		return "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response renewTokenResponse

	err = fromJson(&response, resp.Body)
	if err != nil {
		return "", err
	}
	if !response.Success {
		err = GlobalIdentityError(response.OperationReport)
	}
	return response.NewToken, err
}

func (gim *globalIdentityManager) ValidateApplication(clientApplicationKey string, clientSecretKey string, resources string) (bool, error) {
	hmac512 := hmac.New(sha512.New, []byte(clientSecretKey))
	_, err := hmac512.Write([]byte(resources))
	if err != nil {
		return false, err
	}
	encryptedSecretKey := hex.EncodeToString(hmac512.Sum(nil))
	request := &validateApplicationRequest{
		ApplicationKey:gim.applicationKey,
		ClientApplicationKey: clientApplicationKey,
		RawData: resources,
		EncryptedData: encryptedSecretKey,

	}
	json, err := toJson(request)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(gim.globalIdentityHost + validateApplicationSuffix, contentJson, json)
	if err != nil {
		return false, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, GlobalIdentityError([]string{fmt.Sprintf("%v", resp.StatusCode)})
	}

	var response validateApplicationResponse

	err = fromJson(&response, resp.Body)
	if err != nil {
		return false, err
	}
	if !response.Success {
		err = GlobalIdentityError(response.OperationReport)
	}
	return response.Success, err
}