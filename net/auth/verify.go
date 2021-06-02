package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const sessionServerUri = "https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%s&serverId=%s"

func DecryptLoginResponse(serverPrivateKey *rsa.PrivateKey, serverPublicKey, expectedVerifyToken, verifyToken, sharedSecret []byte, profile *Profile) ([]byte, error) {
	verifyToken, sharedSecret, err := decryptSecrets(serverPrivateKey, verifyToken, sharedSecret)
	if err != nil {
		return nil, err
	}
	if !reflect.DeepEqual(expectedVerifyToken, verifyToken) {
		return nil, errors.New("verify token mismatch")
	}
	return sharedSecret, verifySecret(serverPublicKey, sharedSecret, profile)
}

func decryptSecrets(privateKey *rsa.PrivateKey, verifyToken, sharedSecret []byte) ([]byte, []byte, error) {
	var err error
	verifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, verifyToken)
	sharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, sharedSecret)
	return verifyToken, sharedSecret, err
}

func verifySecret(serverPublicKey, sharedSecret []byte, profile *Profile) error {
	hash := digest(sharedSecret, serverPublicKey)
	return getProfile(profile, hash)
}

func getProfile(profile *Profile, hash string) (err error) {
	uri := fmt.Sprintf(sessionServerUri, profile.Name, hash)

	response, err := http.Get(uri)
	defer response.Body.Close()

	if response.StatusCode == 204 {
		err = errors.New("Verification failed")
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, profile)
	return
}
