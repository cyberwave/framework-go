package ledger_model

import "framework-go/crypto/framework"

/*
 * Author: imuge
 * Date: 2020/5/27 下午1:22
 */

type LedgerInfo struct {
	Hash framework.HashDigest
	LatestBlockHash framework.HashDigest
	LatestBlockHeight int64
}
