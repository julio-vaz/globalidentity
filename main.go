package main

import(
	. "./globalidentity"
	"fmt"
)

func main(){ 
	applicationKey := "302c257e-984c-4773-940a-fd2eb3ba030a"
	expirationInMinutes := 65
	email := "scar0000@gmail.com"
	password := "stone2016"
	a := AuthenticateUser(applicationKey, email, password, expirationInMinutes)
	fmt.Println(a)

	token := "CE2A9D089704BC379A1E0528AB1454EA49B55122FA1911257DF85EC4747B9AA8863229F3DAC324AD0AC16D1C14FAF06C481C69F91734B916FE10443933A736198E376B856591899F67644B5D618A0AA4"
	b := ValidateToken(applicationKey, token)
	fmt.Println(b)


	clientApplicationKey := "3154c332-d3ab-4f4f-9d2f-e181858dd72d"
	clientSecretKey := "A3D3E8E1048B2E5F9FFD50A815910FD6447B8A31"
	resources := "teste3"
	c := ValidateApplication(applicationKey, clientApplicationKey, clientSecretKey, resources)
	fmt.Println(c)
}