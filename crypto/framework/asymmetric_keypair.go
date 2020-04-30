package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:17 下午
 */

// 非对称秘钥
type AsymmetricKeypair struct {
	PubKey  PubKey
	PrivKey PrivKey
}

func (A AsymmetricKeypair) getAlgorithm() int16 {
	return A.PubKey.GetAlgorithm()
}

func NewAsymmetricKeypair(pubKey PubKey, privKey PrivKey) AsymmetricKeypair {
	if pubKey.GetAlgorithm() != privKey.GetAlgorithm() {
		panic("The algorithms of PubKey and PrivKey don't match!")
	}

	return AsymmetricKeypair{
		pubKey, privKey,
	}
}
