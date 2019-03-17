package dummy

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	wallet := New()
	encryptedKey, _ := wallet.EncryptKey(`/some/path/here`, []byte(`some key here...`), []byte(`some nonce here`), `a key name`)
	decryptedKey, _ := wallet.DecryptKey(`/some/path/here`, encryptedKey, []byte(`some nonce here`), `a key name`)

	if string(decryptedKey) != `some key here...` {
		t.Errorf("Invalid key decrypted: %v", decryptedKey)
	}
}
