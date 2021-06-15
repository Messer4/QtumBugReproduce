package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/qtumproject/janus/pkg/qtum"
	"log"
	"math/big"
	"strings"
)

const cAbi = `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "asset",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "sendAssets",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

const defaultGasLimit = 2500000

const (
	rpc = "http://qtum:testpasswd@localhost:3889"
	senderContractAddr = ""
	tokenAddress = ""
)

func main() {
	c, err := qtum.NewClient(false, rpc)
	if err != nil {
		panic(err)
	}

	 m := &qtum.Method{c}

	parsedABI, err := abi.JSON(strings.NewReader(cAbi))
	if err != nil {
		panic(err)
	}

	contractCallData,err := parsedABI.Pack("sendAssets",common.HexToAddress(tokenAddress), big.NewInt(1000000005))
	if err != nil {
		panic(err)
	}

	resp, err := m.CallContract(&qtum.CallContractRequest{
		From:     "",
		To:       senderContractAddr,
		Data:     hex.EncodeToString(contractCallData),
		GasLimit: big.NewInt(defaultGasLimit),
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v",*resp)
}
