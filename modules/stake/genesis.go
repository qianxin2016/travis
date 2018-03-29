package stake

import (
	crypto "github.com/tendermint/go-crypto"
	"github.com/ethereum/go-ethereum/common"
)

/**** code to parse accounts from genesis docs ***/

// GenesisValidator - genesis validator parameters
type genesisValidator struct {
	Address common.Address 	`json:"address"`
	PubKey  crypto.PubKey 	`json:"pub_key"`
	Power 	uint64         	`json:"power"`
	Name    string          `json:"name"`
}

// ToAccount - GenesisAccount struct to a basecoin Account
//func (g GenesisAccount) ToAccount() Account {
//	return Account{
//		Coins: g.Balance,
//	}
//}
//
//// GetAddr - Get the address of the genesis account
//func (g GenesisAccount) GetAddr() ([]byte, error) {
//	noAddr, noPk := len(g.Address) == 0, g.PubKey.Empty()
//
//	if noAddr {
//		if noPk {
//			return nil, errors.New("No address given")
//		}
//		return g.PubKey.Address(), nil
//	}
//	if noPk { // but is addr...
//		return g.Address, nil
//	}
//	// now, we have both, make sure they check out
//	if bytes.Equal(g.Address, g.PubKey.Address()) {
//		return g.Address, nil
//	}
//	return nil, errors.New("Address and pubkey don't match")
//}
