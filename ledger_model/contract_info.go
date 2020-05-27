package ledger_model

import binary_proto "framework-go/binary-proto"

/*
 * Author: imuge
 * Date: 2020/5/27 下午4:00
 */

var _ binary_proto.DataContract = (*ContractInfo)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(ContractInfo{})
}

type ContractInfo struct {
	BlockchainIdentity
	MerkleSnapshot
	ChainCode []byte `primitiveType:"BYTES"`
}

func (c ContractInfo) ContractCode() int32 {
	return binary_proto.CONTRACT_INFO
}

func (c ContractInfo) ContractName() string {
	return "ContractInfo"
}

func (c ContractInfo) Description() string {
	return ""
}
