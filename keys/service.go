package keys

// Service ...
type Service interface {
	GetManager() Manager

	GetKeyPairAlgorithm(algorithm string, selector func(reg *Registration) bool) (KeyPairAlgorithm, error)

	GetSecretKeyAlgorithm(algorithm string, selector func(reg *Registration) bool) (SecretKeyAlgorithm, error)
}
