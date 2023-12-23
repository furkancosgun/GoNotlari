package models

// Create a struct to read the username and password from the request body
type Credential struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type CredentialSlice []Credential

func (creds CredentialSlice) FindCredentialByUsername(username string) Credential {
	findedCred := Credential{}

	for _, v := range creds {
		if v.Username == username {
			findedCred = v
			break
		}
	}

	return findedCred
}
