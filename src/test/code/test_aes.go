package code

import (
	"crypto/rand"

	"github.com/starter-go/application"
	"github.com/starter-go/security/keys"
	"github.com/starter-go/vlog"
)

// TestAES ...
type TestAES struct {

	//starter:component

	KeysMan keys.Manager //starter:inject("#")
	KeysSer keys.Service //starter:inject("#")
	Logger  vlog.Logger  //starter:inject("#")
}

// Life ...
func (inst *TestAES) Life() *application.Life {
	return &application.Life{
		Order:  99,
		OnLoop: inst.run,
	}
}

func (inst *TestAES) run() error {

	alg, err := inst.KeysSer.GetSecretKeyAlgorithm("aes", nil)
	if err != nil {
		return err
	}

	params := &keys.SecretKeyParams{Size: 192}
	// params.Size = 128
	sk, err := alg.GetGenerator().Generate(params)
	if err != nil {
		return err
	}

	sk1d := &keys.SecretKeyData{}
	sk2d, err := sk.Export(sk1d)
	if err != nil {
		return err
	}

	sk2, err := alg.GetLoader().Load(sk2d)
	if err != nil {
		return err
	}

	iv := make([]byte, sk2.BlockSize())
	opt := &keys.CryptOptions{
		Block:   keys.CBC,
		Padding: keys.PKCS7Padding,
		IV:      iv,
	}
	rand.Read(opt.IV)

	data1 := make([]byte, 10)
	data1 = []byte("hello,aes")
	data2, _, err := sk2.NewEncrypter(opt).Encrypt(data1, opt)

	data3, _, err := sk2.NewDecrypter(opt).Decrypt(data2, opt)

	vlog.Debug("todo: %d", len(data3))

	return err
}
