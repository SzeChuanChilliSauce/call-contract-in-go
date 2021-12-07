package main

import (
	"call-contract-in-go/cdd"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

var (
	PrivateKey1, _ = crypto.HexToECDSA("your private key")
	FromAddr       = common.HexToAddress("0x89f4931f1E14b3908DBC33C9927d8EfbA3b2fEdf")
	ToAddr         = common.HexToAddress("0xF1C4a341624e4861725B91F78A41ac08CD2Ddd6E")
	ContractCdd    = "0x6d577BF38a6086d48aA2D169C67775fcc901A784"
	BscTestNet     = "https://data-seed-prebsc-2-s2.binance.org:8545/"
)

func main() {
	var txHash string
	client, err := ethclient.Dial(BscTestNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		os.Exit(1)
	}

	txHash, err = cdd.CallContract(client, PrivateKey1, FromAddr, ToAddr, ContractCdd)
	if err != nil {
		os.Exit(2)
	}
	fmt.Println("tx hash: ", txHash)

	txHash, err = cdd.CallContractWithAbi(client, PrivateKey1, FromAddr, ToAddr, ContractCdd)
	if err != nil {
		os.Exit(3)
	}
	fmt.Println("tx hash: ", txHash)
}
