package Globalidentity

import (
	"github.com/levigross/grequests"
	"fmt"
	"crypto/sha512"
	"crypto/hmac"
	"encoding/hex"
)

var Api = "https://dlpgi.dlp-payments.com/api/"


func AuthenticateUser(applicationKey string, email string, password string, expirationInMinutes int) map[string]interface{} {
	json := map[string]interface{}{
		"Email": email,
		"Password": password,
		"ApplicationKey": applicationKey,
		"TokenExpirationInMinutes": expirationInMinutes,
	}
	resp, _ := grequests.Post(Api + "Authorization/authenticate", &grequests.RequestOptions{
        JSON: json,
    })
	var response map[string]interface{}	
	resp.JSON(&response)

	return response
}

func ValidateToken(applicationKey string, token string) map[string]interface{} {
	json := map[string]interface{}{
		"ApplicationKey": applicationKey,
		"Token": token,
	}
	resp, _ := grequests.Post(Api + "Authorization/ValidateToken", &grequests.RequestOptions{
		JSON: json,
	})
	var response map[string]interface{}
	resp.JSON(&response)

	return response
}

func HasRoles(applicationKey string, userKey string, roles []string) map[string]interface{} {
	json := map[string]interface{}{
		"ApplicationKey": applicationKey,
		"UserKey": userKey,
		"RoleCollection": roles,
	}

	resp, _ := grequests.Post(Api + "Authorization/IsUserInRole", &grequests.RequestOptions{
		JSON: json,
	})
	var response map[string]interface{}
	resp.JSON(&response)

	return response
}

func ValidateApplication(applicationKey string, clientApplicationKey string, clientSecretKey string, resources string) map[string]interface{} {
	hmac512 := hmac.New(sha512.New, []byte(clientSecretKey))
	hmac512.Write([]byte(resources))
	encryptedSecretKey := hex.EncodeToString(hmac512.Sum(nil))
	fmt.Println(encryptedSecretKey)

	json := map[string]interface{}{
		"ApplicationKey": applicationKey,
		"ClientApplicationKey": clientApplicationKey,
		"RawData": resources,
		"EncryptedData": encryptedSecretKey,
	}

	fmt.Println(json)

	resp, _ := grequests.Post(Api + "Authorization/ValidateApplication", &grequests.RequestOptions{
		JSON: json,
	})
	var response map[string]interface{}
	resp.JSON(&response)

	return response
}
