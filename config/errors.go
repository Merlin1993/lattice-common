package config

import (
	"fmt"
	"math/big"
)

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a ChainConfig that would alter the past.
type CompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedBlock, newBlock *big.Int) *CompatError {
	var rew *big.Int
	switch {
	case storedBlock == nil:
		rew = newBlock
	case newBlock == nil || storedBlock.Cmp(newBlock) < 0:
		rew = storedBlock
	default:
		rew = newBlock
	}
	err := &CompatError{what, storedBlock, newBlock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *CompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}
