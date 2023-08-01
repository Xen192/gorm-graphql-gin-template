package security

import (
	"mygpt/pkg/lib/encrypt"
	"mygpt/pkg/utils"

	"github.com/sirupsen/logrus"
)

type EncryptedJSON struct {
	Content string `json:"content"`
	Hash    string `json:"hash"`
}

func EncryptJSON(content map[string]interface{}) EncryptedJSON {
	marshalledData := string(utils.JSONMarshal(content))
	hash := encrypt.GetHash(marshalledData)

	encryptedData, err := encrypt.Encrypt(marshalledData)
	if err != nil {
		logrus.Error(err)
		return EncryptedJSON{}
	}

	return EncryptedJSON{
		Content: encryptedData,
		Hash:    hash,
	}
}

func DecryptJSON(content EncryptedJSON) map[string]interface{} {
	decryptedData, err := encrypt.Decrypt(content.Content)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	unmarshalledData := utils.SafeDereference(utils.JSONUnMarshal[map[string]interface{}]([]byte(decryptedData)))
	hash := encrypt.GetHash(decryptedData)
	if hash != content.Hash {
		logrus.Error("Hash Mismatch")
		return nil
	}

	return unmarshalledData
}
