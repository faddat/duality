package keeper

import (
	"crypto/sha256"
	"bytes"
	//"strings"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SortTokens(ctx sdk.Context, token0 string, token1 string) (string, string, error) {

	// Calculate sha256 Checksum for Token0, Token1
	token0Hash := sha256.Sum256([]byte(token0))
	token1Hash := sha256.Sum256([]byte(token1))

	//calculates an comparison integer for the two hashes
	comparisonInt :=  bytes.Compare(token0Hash[:], token1Hash[:])

	/* If comparisonInt == -1 : token0Hash < token1Hash
	   comparisonInt == 0 token0Hash == token1Hash (return an error)
	   comparisonInt == 1 token0Hash > token1Hash (switch elements)
	*/
	if comparisonInt == -1 {
		return token0, token1, nil
	} else if comparisonInt == 0 {
		return "", "",  sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else {
		return token1, token0, nil
	}
	

}

func (k Keeper) SortTokensDeposit(ctx sdk.Context, token0 string, token1 string, amounts0 []sdk.Dec, amounts1 []sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec, error) {

	// Calculate sha256 Checksum for Token0, Token1
	token0Hash := sha256.Sum256([]byte(token0))
	token1Hash := sha256.Sum256([]byte(token1))

	//calculates an comparison integer for the two hashes
	comparisonInt :=  bytes.Compare(token0Hash[:], token1Hash[:])

	/* If comparisonInt == -1 : token0Hash < token1Hash (returns parameters as given)
	   comparisonInt == 0 token0Hash == token1Hash (return an error)
	   comparisonInt == 1 token0Hash > token1Hash (switch elements and amount arrays)
	*/ 
	if comparisonInt == -1 {
		return token0, token1, amounts0, amounts1, nil
	} else if comparisonInt == 0 {
		return "", "", nil, nil,  sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else {
		return token1, token0, amounts1, amounts0, nil
	}

}
