package main

import (
	"fmt"
	"log" // Log package for logging error messages
	"os"  // OS package used for interacting with the OS such as reading files

	"github.com/golang-jwt/jwt/v5" // Importing the JWT library for Go to handle JSON Web Tokens
)

// JWT token given as string
var tkn = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY5NzYyMzA0NCwiaWF0IjoxNjk3NjIwMDQ0fQ.IRNCYmWQK-7uhlS1uERzCJJScSlpGBdrrO6Ih0K4nJj1v_qvBc_cEZVVMGa3ko6hs_pT03EuhQmKupyMTy7ylkKrZLrLGYhukhMh0eTxX1ifWalGdNwTvgb9ny-beEfaTmX6-dccGvTJCXisEMeap3CRPrI-oYKkRB_1ZKuaJlvgbGgnGTbeDEZhUFhGrC1AfExdIWuB2S8ml1EMtu84UYMuUPcifUEuPPYj0XAKSb92F-_GLSQVYHqopSGWOtL0725s0LzXDY_t9FnsKsY5HN9VsJ3BCsqPTcJ5FpWU4ExT7p7dab6Nbu4ZqU2ysclwTRMiwf249kUcC-EJ-P7hVg`

func main() {
	// Reads the public key from pubkey.pem file
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	var c jwt.RegisteredClaims
	// Parsing the JWT token with the claims
	token, err := jwt.ParseWithClaims(tkn, &c, func(token *jwt.Token) (interface{}, error) {
		// Provides the public key for validating the JWT token
		return publicKey, nil
	})

	if err != nil {
		// If error while parsing the token, print the error and exit
		log.Println("parsing token", err)
		return

	}
	if !token.Valid {
		// If the token is not valid, log the error and exit
		log.Println("invalid token")
		return
	}

	// Print the claims from the token
	fmt.Printf("%+v", c)

}
