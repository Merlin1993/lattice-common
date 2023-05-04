package config

import (
	"math/big"
	common "zkjg.com/lattice/common"
)

var (
	//是否是联盟链，联盟链会需要一些证书配置
	IsConsortiumBlockchain = false
	CertificateSalt        = "yEFyJ4Ou9UeP1W4ya29EuWwzPvqbQGAL"
	//联盟链公钥
	ConsortiumCAAccount, _ = common.Base58ToAddress("zltc_UXufZvWwvCzmE7bSnn9Q5W2UV9fppJViF")

	DaemonAddress = common.EmptyAddress
	// 是否是中心化CA
	IsCentralizedCA = true
)

var (
	TestNetLatcConfig = &LatcConfig{
		LatcID:    big.NewInt(2),
		LatcGod:   common.HexToAddress("4Af586Fc61eEdfE7B7094C65Db97CCd24Cf6Fa8A"),
		Epoch:     30000,
		Tokenless: true,
	}
	DefaultEpoch = uint(30000)
)

type LatcConfig struct {
	LatcID               *big.Int         `json:"latcId"`
	LatcGod              common.Address   `json:"latcGod"`
	LatcSaints           []common.Address `json:"latcSaints"`
	Epoch                uint             `json:"epoch" 	gencodec:"required"`
	Tokenless            bool             `json:"tokenless" 	gencodec:"required"`
	Period               uint             `json:"period" 	gencodec:"required"`
	NoEmptyAnchor        bool             `json:"noEmptyAnchor"`
	EmptyAnchorPeriodMul uint64           `json:"emptyAnchorPeriodMul"`
	GM                   bool             `json:"GM"`
	RootPublicKey        string           `json:"rootPublicKey"` // 中心化CA根证书公钥
	IsContractVote       bool             `json:"isContractVote"`
	IsDictatorship       bool             `json:"isDictatorship"`
	DeployRule           uint8            `json:"deployRule"` // 0 代表不需要投票，1代表一票通过，2代表共识通过
}
