package code

import (
	"bytes"
	"crypto"
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/security/keys"
	"github.com/starter-go/vlog"
)

// TestRSA ...
type TestRSA struct {

	//starter:component

	KeysMan keys.Manager //starter:inject("#")
	KeysSer keys.Service //starter:inject("#")
	Logger  vlog.Logger  //starter:inject("#")
}

// Life ...
func (inst *TestRSA) Life() *application.Life {
	return &application.Life{
		Order:  99,
		OnLoop: inst.run,
	}
}

func (inst *TestRSA) run() error {

	alg, err := inst.KeysSer.GetKeyPairAlgorithm("rsa", nil)
	if err != nil {
		return err
	}

	size := 1024 * 2
	pair, err := alg.GetGenerator().Generate(&keys.KeyPairParams{Size: size})
	if err != nil {
		return err
	}

	err = inst.tryCrypt(pair)
	if err != nil {
		return err
	}

	err = inst.trySign(pair)
	if err != nil {
		return err
	}

	return nil
}

func (inst *TestRSA) tryCrypt(pair keys.KeyPair) error {

	opt := &keys.CryptOptions{
		Hash:    crypto.SHA256, // "sha256",
		Padding: keys.OAEP,
	}
	plain := "hello,rsa"
	plain1 := []byte(plain)
	opt.Label = []byte("la12")
	mi, op2, err := pair.PublicKey().NewEncrypter(opt).Encrypt(plain1, opt)
	if err != nil {
		return err
	}

	opt.Label = []byte("la12")
	plain2, op3, err := pair.PrivateKey().NewDecrypter(opt).Decrypt(mi, opt)
	if err != nil {
		return err
	}

	if !bytes.Equal(plain1, plain2) {
		return fmt.Errorf("plain1 != plain2")
	}

	vlog.Debug("%s", op2.Block)
	vlog.Debug("%s", op3.Block)
	return nil
}

func (inst *TestRSA) trySign(pair keys.KeyPair) error {

	opt := &keys.SignatureOptions{
		Hash:    crypto.SHA256,
		Padding: keys.PSS,
	}

	plain := "hello,rsa.sign"
	haha := opt.Hash.New()
	haha.Write([]byte(plain))
	sum := haha.Sum(nil)
	s1 := &keys.Signature{
		Algorithm: "rsa",
		Digest:    sum,
	}

	s2, err := pair.PrivateKey().NewSigner(opt).Sign(s1)
	if err != nil {
		return err
	}

	signAlgName := s2.Algorithm.String()
	signAlg, err := inst.KeysSer.GetSignatureAlgorithm(signAlgName, nil)
	if err != nil {
		return err
	}
	err = signAlg.NewVerifier(pair.PublicKey()).Verify(s2)
	if err != nil {
		return err
	}

	return nil
}
