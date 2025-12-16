package utils

import "golang.org/x/crypto/bcrypt"

func HashPassoword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hashed),nil
}

func ComparePassword(hashed string,plain string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(plain))
	return err==nil
}