package lattice_common

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddress_Base58CheckSum(t *testing.T) {
	account := HexToAddress("0x86AF44a0bF984240Ae39C23193CDc4D037aa6Eb0")
	t.Log(account)
	t.Log(account.String())
}

func TestBase58ToAddress(t *testing.T) {
	account := HexToAddress("0x0C44deB6521E7C780040F35c7d148dF7d15B6ba8")
	t.Log(account.String())
	addr, _ := Base58ToAddress(account.String())
	t.Log(addr.String())
}

type JSONAccount struct {
	Acmap map[Address]Address `json:"acmap"`
}

func TestAddressJson(t *testing.T) {
	// zltc_QLbz7JHiBTspS962RLKV8GndWFwjA5K66 hex() 0x0000000000000000000000000000000000000000
	ac1, _ := Base58ToAddress("zltc_efmSKsZ4h8hFAScTC84D3Ct6BYSM9iSFU")
	ac2, _ := Base58ToAddress("zltc_nYEf8ykcGnQLmxxA1K2mopzJxErc5raYL")
	ja := JSONAccount{}
	ja.Acmap = make(map[Address]Address)
	ja.Acmap[ac1] = ac2
	println(fmt.Sprint(ja.Acmap))
	jaJsonB, err := json.Marshal(ja)
	if err != nil {
		println("marshal err：", err.Error())
		return
	}
	println("marshal succ：", string(jaJsonB))
	ja2 := JSONAccount{}
	err = json.Unmarshal(jaJsonB, &ja2)
	if err != nil {
		println("Unmarshal err：", err.Error())
		return
	}
	println(fmt.Sprint(ja2.Acmap))
}

func TestAddress(t *testing.T) {
	addr, err := Base58ToAddress("zltc_RANnDSq8XoYXfGgnHdHdV11aNqZZkinam")
	t.Log(err)
	t.Log(addr)
}

func TestHash(t *testing.T) {
	hash := Hash{}
	hasher := sha256.New()
	//hasher.Write(Hex2Bytes("f8bf"))
	hasher.Write(Hex2Bytes("f8bf1f02a0cb7a4161752480aa9c74d05e47626bddaf0e9be8d264481afe53cc6b081dd1a4a00000000000000000000000000000000000000000000000000000000000000000a0f3f6da93f1bdf7b49c5bba4cb59c6141326a703293d2901ddcc1cbe840170572a00000000000000000000000000000000000000000000000000000000000000000941c7f6c3afde36de5d0aabee1d0eec6dfd26b020194e2c2ecdedeaeb569e92b6d9daee5103b67300ce6830f42408085316139303280018080"))
	hasher.Sum(hash[:0])
	t.Log(hash.String())
}

func TestAddressToHash(t *testing.T) {
	ac1, _ := Base58ToAddress("zltc_obJkpdtj9uYFF5fDSx7dGfYq87Tq99rgk")
	ac2, _ := Base58ToAddress("zltc_f7kaD5c4PJUNAfJywvoXu9axb5MJpL8WV")
	println(ac1.Hex())
	println(ac2.Hex())
}

func TestTransferAddress(t *testing.T) {
	IdentityContractAddress := "IdentityContractAddress"
	identity := BytesToAddress([]byte(IdentityContractAddress))
	t.Log("identity:", identity.Hex())    //0x6e74697479436F6e747261637441646472657373
	t.Log("identity:", identity.String()) //zltc_aQdmesGLjoJ5FJ65t2F7Nf9tTAT2C3dxA

	OracleContractAddress := "OracleContractAddress"
	oracle := BytesToAddress([]byte(OracleContractAddress))
	t.Log("oracle:", oracle.Hex())    //0x7261636C65436f6E747261637441646472657373
	t.Log("oracle:", oracle.String()) //zltc_amPge82fy3fJsLD1eSerqCBEfgjpU43S4

	ac3, _ := Base58ToAddress("zltc_eNN8KCZXwVUQenNC9WCmVQMTHRsQUQv53")
	t.Log(ac3.Hex())
	ac4, _ := Base58ToAddress("zltc_h62ubytmWH51A6oPSHGFmA2g8zUuG9mjV")
	t.Log(ac4.Hex())
}

func TestPrecompileAccount(t *testing.T) {
	CredibilityAddress := "CredibilityAddress"
	address := BytesToAddress([]byte(CredibilityAddress))
	fmt.Println(address.String())
	fmt.Println(address.Hex())
}
