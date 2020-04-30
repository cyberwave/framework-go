package crypto

import (
	"fmt"
	"framework-go/crypto/classic"
	"framework-go/crypto/framework"
	"framework-go/utils/base58"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 7:06 下午
 */

func TestSHA256(t *testing.T) {
	data := []byte("imuge")
	function := GetHashFunctionByName(classic.SHA256_ALGORITHM.Name)
	hash := function.Hash(data)
	fmt.Println("hash: " + hash.ToBase58())
	require.True(t, function.Verify(hash, data))

	// hash from JD Cahin
	jdHash, _ := base58.Decode("j5vkSRxmUjJzo9KBX79cTMRwD8Aw3J7Ke2JnPzS1eq4fH1")
	require.True(t, function.Verify(framework.ParseHashDigest(jdHash), data))
}

func TestRSA(t *testing.T) {
	function := GetCryptoFunctionByName(classic.RSA_ALGORITHM.Name)

	keypair := (function.(framework.AsymmetricKeypairGenerator)).GenerateKeypair()
	fmt.Println("pub: " + keypair.PubKey.ToBase58())
	fmt.Println("priv: " + keypair.PrivKey.ToBase58())
	data := []byte("imuge")

	// sign
	f1 := function.(framework.SignatureFunction)
	sign := f1.Sign(keypair.PrivKey, data)
	fmt.Println("sign: " + sign.ToBase58())
	require.True(t, f1.Verify(keypair.PubKey, data, sign))
	/**
		sign from JD Chain
	pub: Lch98AycNJd1foLA2fzF5ePBSupgUQdDGiXkhuHpG6kHcDpoKWCMBkUN1bCD7Yu32stE2qA49QqCENoakqzUVutLmzg2gYx7wQ6FEU9r1oazSUhFiXd6A3VKjfyvRM7ozaXnjSFu18FKhUg4JgXdtbsaJtAWjaWRhQzDns45rUdak2eWUXyhhhwT3Y5cFQZTHkDrjQjVCpCR9UYv2kx6rV57SWkUywFsnxva6V4oNbFQ39B8G4nsXHvnYCNAuAt2WyLMYQNPQWbhBb7v5DgPqDVKv913Mfnh6yZbEN96UssMVqsgJsqqZ8djDnipVFx9DdXuGMwj8EWrrWYCXKURViexQPz88tmKhcnjiQ
	priv: 5N2f7eVJwBJSvCbpMh3v8fHXFiKyqUkM3kJbwizchLkw3WHep1vwshkTxRndCsGRVAiQqmRxchBDf5Lva8X8pcAmscSAqBSSMur7TQxwRcwpyrJM1iQS1tYEbiBkq3Swnq7NbrjgMqDVVNPLaKH3tgQHsd4czh8GK5i3KC6D93MmTpPRxX6VDMsRzj58jJEHS95ybrspwvPVyoV3gQSwBPTDQxDNQsDx9ov4YDwBFVAXviEihhhTHY1LXcVwnNidLHBryWn5cLGiGL8PEPSDaj2fWDmhEFRfMCLrE7xVo354BKd6sdtE8vbwUh83sGXQsZD9QJLJkQAyFZLJ6Hr7e1qFibzT3Jir8x9dkaQ6hirzzpGevnkKvLrN1bvYDiBGuf3DrJayMyMAeSfHnAkDLnZxkTBqMjtXVLAaKaJqy4z7hnZoQX79tGARywVSKX1YkmWnL4uXzfs79pP124zgJQawQpKns35NM7jnerUc1rBVpcgTPqtLTcZ3zhN5oVeZSiELSN6oQqCaRGnzJGUESsTw6HhiLN3o1H5YVX7mTSjyR2vEo4QxQJMyEd9tamM1fxrwcpFtW1wBZAdwDEE4xhktwEBBeRSH9MGVnoaD6zv6MXh2BmivjDoDM27tvGaD1gj9E3Q22CVMD6bLJGX4umj8cH5ScA3qFMsPYoSobNDNcyK31VDwvF8HwBw2diHnbihUezu9SapKVZdkQoisGXziA5CdKpzQrbCFdtjxvPer4mThGURAnSiddWZwuDiGqtbUcBg85nUWwkEb9GHRAqM4RubFQuu22HDQhPtYNWmNMWPp6ti4zKmfiSGaDDfDs2Xnx94QfCzPNMrDu85sJYcdYTG8GS621Pe1Wu8wAUYe1o1thpf86epNFMEwwfHmK6nZdLmcxnJP7zt6xHocLf5hsK4jQz9Nf7SHmU5cP1HvQ9c2F7Lb2rgW9j76euK4ut6nvMKmCaEdMt37UNfqRXi5Awh8WiQv9BYh3wap4oapsbbxa9JQR3E8LAwFoBr55QRwGPzi8J2QsLekR2M2C4gzWQhY3SGArz3eDx6WqEGWoYwKXGdMs6GeujJFqHmRYWpxmb4NrBT4yxAqqcu1gm6CbenMGunt6o9KbB87UotwYhhpAMemaQGb7BNC1ijd78AtH5FT5xPrY8VoCM5wLXMxvCzJ3AVnWe56UvJH3AhrLfF4fVcpcwjr31Zm1NJEUWRqgWrfYZ6S8rWSUvymDz5D4oa8w4awi1MJPWUaDvnGsrNfcLErJqaxcisoBHy3v8FztGuEpEfAjajBMDat8LGxLdhghvfVMFCzdoXtBmR3PiGJJk6ZpfdWc4SWv5wHm2M8S5qvEELD6SWNVH2JkFEKKZyjNgzXJq2fy5c3zEA1Ayr82Bm2J35EPN2UXDhFi9dgqHiU313bJncuCTUXgHZEz9Eavpy5kjuWPAmscTxDooKPGkdLJAS4waF2d7SDX9x42aKc7pmkSYaE3xxPDeAs6hyYUm5zh2GgBJimmmszRhpRUMd14HuVaFkaTJJLAqqVcj7BMkXLacaBKJqwYwhBkNvTbfhjSqkFZiYxrtTQ9h
	digest: 3zs6kjs99v5Ws7K1ZFite5TaudUhEgK8EWAkZ2FnJ29ZQJ5QAenxX4K52igvRbosTpbxQbAJcQjdie1BhVrMjV4Fes8eqKchVMegnB9LUkSYgpLS6qWF2rngo3piPJivr5o66iSXjqnT2LvEtUhhVhof4ctbk1WegqMETuytsQkZyoQanwU2NjM1MfDosLqCNpfHkNACzuLhBwsWJ7Kgu1K2Yuct2rL1Sr6jLAwRfzD8q22rMFw2PuMuU32e4fX69nuciDtv4cpeRrLE5GKHa2kndACDta1LUqZznLGZGweQEbXJXhJLug2RQao675mySutriqFAsrKqrgKNzS3W5n6yBd6wRsBDD
	*/
	privBytes, _ := base58.Decode("5N2f7eVJwBJSvCbpMh3v8fHXFiKyqUkM3kJbwizchLkw3WHep1vwshkTxRndCsGRVAiQqmRxchBDf5Lva8X8pcAmscSAqBSSMur7TQxwRcwpyrJM1iQS1tYEbiBkq3Swnq7NbrjgMqDVVNPLaKH3tgQHsd4czh8GK5i3KC6D93MmTpPRxX6VDMsRzj58jJEHS95ybrspwvPVyoV3gQSwBPTDQxDNQsDx9ov4YDwBFVAXviEihhhTHY1LXcVwnNidLHBryWn5cLGiGL8PEPSDaj2fWDmhEFRfMCLrE7xVo354BKd6sdtE8vbwUh83sGXQsZD9QJLJkQAyFZLJ6Hr7e1qFibzT3Jir8x9dkaQ6hirzzpGevnkKvLrN1bvYDiBGuf3DrJayMyMAeSfHnAkDLnZxkTBqMjtXVLAaKaJqy4z7hnZoQX79tGARywVSKX1YkmWnL4uXzfs79pP124zgJQawQpKns35NM7jnerUc1rBVpcgTPqtLTcZ3zhN5oVeZSiELSN6oQqCaRGnzJGUESsTw6HhiLN3o1H5YVX7mTSjyR2vEo4QxQJMyEd9tamM1fxrwcpFtW1wBZAdwDEE4xhktwEBBeRSH9MGVnoaD6zv6MXh2BmivjDoDM27tvGaD1gj9E3Q22CVMD6bLJGX4umj8cH5ScA3qFMsPYoSobNDNcyK31VDwvF8HwBw2diHnbihUezu9SapKVZdkQoisGXziA5CdKpzQrbCFdtjxvPer4mThGURAnSiddWZwuDiGqtbUcBg85nUWwkEb9GHRAqM4RubFQuu22HDQhPtYNWmNMWPp6ti4zKmfiSGaDDfDs2Xnx94QfCzPNMrDu85sJYcdYTG8GS621Pe1Wu8wAUYe1o1thpf86epNFMEwwfHmK6nZdLmcxnJP7zt6xHocLf5hsK4jQz9Nf7SHmU5cP1HvQ9c2F7Lb2rgW9j76euK4ut6nvMKmCaEdMt37UNfqRXi5Awh8WiQv9BYh3wap4oapsbbxa9JQR3E8LAwFoBr55QRwGPzi8J2QsLekR2M2C4gzWQhY3SGArz3eDx6WqEGWoYwKXGdMs6GeujJFqHmRYWpxmb4NrBT4yxAqqcu1gm6CbenMGunt6o9KbB87UotwYhhpAMemaQGb7BNC1ijd78AtH5FT5xPrY8VoCM5wLXMxvCzJ3AVnWe56UvJH3AhrLfF4fVcpcwjr31Zm1NJEUWRqgWrfYZ6S8rWSUvymDz5D4oa8w4awi1MJPWUaDvnGsrNfcLErJqaxcisoBHy3v8FztGuEpEfAjajBMDat8LGxLdhghvfVMFCzdoXtBmR3PiGJJk6ZpfdWc4SWv5wHm2M8S5qvEELD6SWNVH2JkFEKKZyjNgzXJq2fy5c3zEA1Ayr82Bm2J35EPN2UXDhFi9dgqHiU313bJncuCTUXgHZEz9Eavpy5kjuWPAmscTxDooKPGkdLJAS4waF2d7SDX9x42aKc7pmkSYaE3xxPDeAs6hyYUm5zh2GgBJimmmszRhpRUMd14HuVaFkaTJJLAqqVcj7BMkXLacaBKJqwYwhBkNvTbfhjSqkFZiYxrtTQ9h")
	jdPriv := framework.ParsePrivKey(privBytes)
	pubBytes, _ := base58.Decode("Lch98AycNJd1foLA2fzF5ePBSupgUQdDGiXkhuHpG6kHcDpoKWCMBkUN1bCD7Yu32stE2qA49QqCENoakqzUVutLmzg2gYx7wQ6FEU9r1oazSUhFiXd6A3VKjfyvRM7ozaXnjSFu18FKhUg4JgXdtbsaJtAWjaWRhQzDns45rUdak2eWUXyhhhwT3Y5cFQZTHkDrjQjVCpCR9UYv2kx6rV57SWkUywFsnxva6V4oNbFQ39B8G4nsXHvnYCNAuAt2WyLMYQNPQWbhBb7v5DgPqDVKv913Mfnh6yZbEN96UssMVqsgJsqqZ8djDnipVFx9DdXuGMwj8EWrrWYCXKURViexQPz88tmKhcnjiQ")
	jdPub := framework.ParsePubKey(pubBytes)
	digestBytes, _ := base58.Decode("3zs6kjs99v5Ws7K1ZFite5TaudUhEgK8EWAkZ2FnJ29ZQJ5QAenxX4K52igvRbosTpbxQbAJcQjdie1BhVrMjV4Fes8eqKchVMegnB9LUkSYgpLS6qWF2rngo3piPJivr5o66iSXjqnT2LvEtUhhVhof4ctbk1WegqMETuytsQkZyoQanwU2NjM1MfDosLqCNpfHkNACzuLhBwsWJ7Kgu1K2Yuct2rL1Sr6jLAwRfzD8q22rMFw2PuMuU32e4fX69nuciDtv4cpeRrLE5GKHa2kndACDta1LUqZznLGZGweQEbXJXhJLug2RQao675mySutriqFAsrKqrgKNzS3W5n6yBd6wRsBDD")
	digest := framework.ParseSignatureDigest(digestBytes)
	require.True(t, f1.Verify(jdPub, data, digest))
	require.Equal(t, digest, f1.Sign(jdPriv, data))

	// encrypt
	f2 := function.(framework.AsymmetricEncryptionFunction)
	encrypt := f2.Encrypt(keypair.PubKey, data)
	require.Equal(t, data, f2.Decrypt(keypair.PrivKey, encrypt))
	/**
		encrypt from JD Chain
	priv: 5N2f6jgYA7jZ6MbQB5U3nkyDFmkWdCzFtLeefSWGbrUsUuJkwdhfznRqr8aExikERFdbEsrkv4P5i7UYSF9SYucr7wNnDdGSAKTLfJnrBjGaWTNjL45XiU2nvnUs3apScaMwgcXrYdP9p89sa1NNWsMcFYiWYg96EwnSMtvNNq9NuNfgcqUwNFceSU3nwbPwmgixGuZXqFsiFw2Ggqetmtanv6gup9Uk78TfvAtyjt9HDZAEGE6rN3yZLWGt2pESpbMMGwsLKUhY1zmU32fzZMq5jp4eZ9wevZkWdi4rd5HuUy2V1NwyiXxHnKutVw8em4N2dHMPjkmptCjkDSacF69XosDX2UaZ7jxcguqkeoPKbRqS6n96uQvSkjYRHG83VwLVB2EdP89QrFPeh4wFm3GheTzakNQW6ypRyHoTuMqm4GSh2UshtHmymsKiSLTvpcwKBJhPxbFmuReDFAnAf7cPC1TLXLBUm3fNQJ9LaeLQCHzs3RNG6cHsQKKBrAMBBZUCERPqc7stMs61NANNAAtQeTuXzi9SMFVXpxMUDkjURseroZN5qUM2oXM12eaHYAN8PHBz7jQirSJP5vQAbQHXDaZLyDD6dNT94RDok7PNXB2pw36ARGDcrK415hphGwpydbHN7jefKm94ncrGHK64ipAEFBLZmrLbHoLndNur1pGPfPgntReZ5JDovoJMRK7cr5KqFw8nNtmrpKcbYrNSx9HgKRZmkKrbPi5ZQ6soKyrmriFMPW87XiipBM3vZq3wpShZ7CXgHqv5fqJ2CwtAGbLbTUY6if4Em3c5jWD5Vc5Y4NbxVWBrMPXTuirvimvPDKmzdD44xyUC6TcEAfKWD3L24USCYGpNQMKRMMoffufz9XAZbm2gVW63WEW7CqVjyVopEVZn4FF3BkEnwRCV5CRQm5w44KyRUC5GLkkVFUYPjRqi9pfhjo6K5RhoHvcsEM5vX6EvcZo2zCN15Y7b3EHRmGmGUk1YPLLBK3xzt4epaBVAb6j2s4En8P1NJQFWQowXRchcrprPkut5yjwPu7MbuR4eEKpiFGCqjN5WW5UdrZhSp3ao7BP3vZmXbQNNybjw6PLfAuj4AoeGiqkXGhQ7E7qPxCUx8V96bD6rwp6BrV6tiw1e2feQbj5TXd9A6xB4RGvtfcF5ohPjDJygQ2Vyq6Bq2c48K9a3SUEXRM2hQJp1BQVWzzrGhHXBmsiu45sWVQacgiJFsfDjQET3BPQFW9ZV34XxG1WxdJhGHvSXrwMg5X9FRLc94BfcAocaUN5BKFoPS26W9H4cF37RCHqCLP8LjSZsHUaKHW41TGeTUzpmzYNCPijtWunnYXfrf7SzryYpLeDdj1UKb7jy2jUDiC74K4QpY7YxvzpLoz2YYrV4yHRZ7knZxiJWXszh2EU2QkMMtfqfu2hJNyyZrp2sX67iXah82eCJ93BZXvocYZvAwkeHca7p7a37q2QBVhLtJpfcz1xPZjR7NxFw1RVKo9pei3X4aVLK3xWs1sssF6bWTrosjJCiZaERCMnWiYo3d1pVGrQPcMJ4bmQuRb9QEGUGkJB7KyBSqF6xrQ
	pub: Lch945YUHdhZh6Y6rbuJboysVDngAPZuoTSnjoztrb2zX7XhAXpPgwu53BuYdoLhpNwZGwPGjihy2S1bPJu7Kx9f2mG962pgiqz3Rn9oLjL1ZiDgx4BRThafTee5szN1kc3egfpB1AHLY4w5uEadCkEVxPMCncVzokunPaQ7SfyoMHeq8vNrpRh76vPMPqHv2gwqoVhex781XNm1hvCqinLS8ZDDErN7TS3zfbb9gFKy9ugEwHVj47Y8SbdMHjLiBvzoAWWV5gDdRN6m6zAKB1DwwaHMKL5Vx5CtZQhmKta4fC2nPNF8oM7vnD5wLuVR76G9WMeK1Adrdj6J19nFu7Jq4NsfRruJMWzka4
	encrypt: 3zs4tugVQsMNvkjZs8dgFznkh6P35DMPHQqQxAKbcCf4v7vMNh6Wj2vcczH6Udr2E2kUp8f7vBN4RbAkQ8pFaeh5kPgN6T8WqnCUZD26dmwvFeQPwET4nkyR7T2wxQbHJHnX9Dqc5U3eVevDg4v7DXtF6epELVvrMozU6zD7f17gjjhvxgBCXMD18ho9X7y8vmsbo1mndHhyfTSuadRy3EDhfZYc1EMyjWmu3J1BSi1HSKQmAQMtQhmtEnaYD41sn2gZMQhVRUoFthoC2zCDJ9RFqyXaeqoLh7yQ97jcEjLVW2chsE1gGBAFUvEP6y93NDdgnEKwkchN36EtDid6r8QbwgRJmZLSh
	*/
	privBytes, _ = base58.Decode("5N2f6jgYA7jZ6MbQB5U3nkyDFmkWdCzFtLeefSWGbrUsUuJkwdhfznRqr8aExikERFdbEsrkv4P5i7UYSF9SYucr7wNnDdGSAKTLfJnrBjGaWTNjL45XiU2nvnUs3apScaMwgcXrYdP9p89sa1NNWsMcFYiWYg96EwnSMtvNNq9NuNfgcqUwNFceSU3nwbPwmgixGuZXqFsiFw2Ggqetmtanv6gup9Uk78TfvAtyjt9HDZAEGE6rN3yZLWGt2pESpbMMGwsLKUhY1zmU32fzZMq5jp4eZ9wevZkWdi4rd5HuUy2V1NwyiXxHnKutVw8em4N2dHMPjkmptCjkDSacF69XosDX2UaZ7jxcguqkeoPKbRqS6n96uQvSkjYRHG83VwLVB2EdP89QrFPeh4wFm3GheTzakNQW6ypRyHoTuMqm4GSh2UshtHmymsKiSLTvpcwKBJhPxbFmuReDFAnAf7cPC1TLXLBUm3fNQJ9LaeLQCHzs3RNG6cHsQKKBrAMBBZUCERPqc7stMs61NANNAAtQeTuXzi9SMFVXpxMUDkjURseroZN5qUM2oXM12eaHYAN8PHBz7jQirSJP5vQAbQHXDaZLyDD6dNT94RDok7PNXB2pw36ARGDcrK415hphGwpydbHN7jefKm94ncrGHK64ipAEFBLZmrLbHoLndNur1pGPfPgntReZ5JDovoJMRK7cr5KqFw8nNtmrpKcbYrNSx9HgKRZmkKrbPi5ZQ6soKyrmriFMPW87XiipBM3vZq3wpShZ7CXgHqv5fqJ2CwtAGbLbTUY6if4Em3c5jWD5Vc5Y4NbxVWBrMPXTuirvimvPDKmzdD44xyUC6TcEAfKWD3L24USCYGpNQMKRMMoffufz9XAZbm2gVW63WEW7CqVjyVopEVZn4FF3BkEnwRCV5CRQm5w44KyRUC5GLkkVFUYPjRqi9pfhjo6K5RhoHvcsEM5vX6EvcZo2zCN15Y7b3EHRmGmGUk1YPLLBK3xzt4epaBVAb6j2s4En8P1NJQFWQowXRchcrprPkut5yjwPu7MbuR4eEKpiFGCqjN5WW5UdrZhSp3ao7BP3vZmXbQNNybjw6PLfAuj4AoeGiqkXGhQ7E7qPxCUx8V96bD6rwp6BrV6tiw1e2feQbj5TXd9A6xB4RGvtfcF5ohPjDJygQ2Vyq6Bq2c48K9a3SUEXRM2hQJp1BQVWzzrGhHXBmsiu45sWVQacgiJFsfDjQET3BPQFW9ZV34XxG1WxdJhGHvSXrwMg5X9FRLc94BfcAocaUN5BKFoPS26W9H4cF37RCHqCLP8LjSZsHUaKHW41TGeTUzpmzYNCPijtWunnYXfrf7SzryYpLeDdj1UKb7jy2jUDiC74K4QpY7YxvzpLoz2YYrV4yHRZ7knZxiJWXszh2EU2QkMMtfqfu2hJNyyZrp2sX67iXah82eCJ93BZXvocYZvAwkeHca7p7a37q2QBVhLtJpfcz1xPZjR7NxFw1RVKo9pei3X4aVLK3xWs1sssF6bWTrosjJCiZaERCMnWiYo3d1pVGrQPcMJ4bmQuRb9QEGUGkJB7KyBSqF6xrQ")
	jdPriv = framework.ParsePrivKey(privBytes)
	pubBytes, _ = base58.Decode("Lch945YUHdhZh6Y6rbuJboysVDngAPZuoTSnjoztrb2zX7XhAXpPgwu53BuYdoLhpNwZGwPGjihy2S1bPJu7Kx9f2mG962pgiqz3Rn9oLjL1ZiDgx4BRThafTee5szN1kc3egfpB1AHLY4w5uEadCkEVxPMCncVzokunPaQ7SfyoMHeq8vNrpRh76vPMPqHv2gwqoVhex781XNm1hvCqinLS8ZDDErN7TS3zfbb9gFKy9ugEwHVj47Y8SbdMHjLiBvzoAWWV5gDdRN6m6zAKB1DwwaHMKL5Vx5CtZQhmKta4fC2nPNF8oM7vnD5wLuVR76G9WMeK1Adrdj6J19nFu7Jq4NsfRruJMWzka4")
	jdPub = framework.ParsePubKey(pubBytes)
	encryptBytes, _ := base58.Decode("3zs4tugVQsMNvkjZs8dgFznkh6P35DMPHQqQxAKbcCf4v7vMNh6Wj2vcczH6Udr2E2kUp8f7vBN4RbAkQ8pFaeh5kPgN6T8WqnCUZD26dmwvFeQPwET4nkyR7T2wxQbHJHnX9Dqc5U3eVevDg4v7DXtF6epELVvrMozU6zD7f17gjjhvxgBCXMD18ho9X7y8vmsbo1mndHhyfTSuadRy3EDhfZYc1EMyjWmu3J1BSi1HSKQmAQMtQhmtEnaYD41sn2gZMQhVRUoFthoC2zCDJ9RFqyXaeqoLh7yQ97jcEjLVW2chsE1gGBAFUvEP6y93NDdgnEKwkchN36EtDid6r8QbwgRJmZLSh")
	jdEncrypt := framework.ParseAsymmetricCiphertext(encryptBytes)
	require.Equal(t, data, f2.Decrypt(jdPriv, jdEncrypt))

}
