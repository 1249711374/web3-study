package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	visit "metanode/contract"
)

var (
	PrivateKeyStr = "我的私钥"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=========== 任务1 ===================>")
	// 查询特定区块信息
	err = blockInfo(client, 8780641)
	if err != nil {
		log.Printf("find block info error: %v", err)
		return
	}

	toAddress := "0xE66A7F2AB811C5c22Ba0C4802095fB6FD716fa22"
	txHash, err := transfer(client, common.HexToAddress(toAddress), big.NewInt(2_000_000))
	if err != nil {
		log.Printf("transfer error: %v", err)
	}

	log.Printf("txHash: %s", txHash)

	fmt.Println("=========== 任务2 ===================>")

	contractAddr, err := deployContract(client)
	if err != nil {
		log.Fatal(err)
	}

	u, err := Visit(contractAddr, client)
	log.Println(u, err)

	u, err = Visit(contractAddr, client)
	log.Println(u, err)

	u, err = Visit(contractAddr, client)
	log.Println(u, err)
}

func blockInfo(cli *ethclient.Client, blockNumber int64) error {
	block, err := cli.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		err = fmt.Errorf("error getting block by number: %w", err)
		return err
	}

	log.Printf("block number: %s, tranaction count: %d, block.time: %d \n", block.Hash().String(), len(block.Transactions()), block.Time())

	return nil
}

func transfer(cli *ethclient.Client, to common.Address, value *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(PrivateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, err := cli.SuggestGasPrice(context.Background())

	chainID, err := cli.ChainID(context.Background())

	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      uint64(21000),
		GasPrice: gasPrice,
		Data:     data,
	})

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.SendTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
	}

	return tx.Hash().Hex(), nil
}

func deployContract(client *ethclient.Client) (string, error) {
	privateKey, err := crypto.HexToECDSA(PrivateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units

	// 说明 之前测试部署过合约 ,再次部署提示 （replacement transaction underpriced）错误，故增加gas 确保成功🏅
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100)) // 增加 20%
	auth.GasPrice = gasPrice

	address, tx, _, err := visit.DeployVisit(auth, client)
	if err != nil {
		log.Fatalf("部署Visit合约失败: %v", err)
	}

	log.Printf("合约地址: %s, 交易hash: %s  \n", address.Hex(), tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("合约部署 交易确认失败: %v", err)
	}

	if receipt.Status != 1 {
		log.Fatalf("交易未完成 receipt.status: %v", receipt.Status)
	} else {
		log.Printf("部署交易确认✅: %s, receipt.status: %d \n", tx.Hash().Hex(), receipt.Status)
	}

	return address.Hex(), nil
}

func Visit(contract string, client *ethclient.Client) (uint64, error) {
	contractAddr := common.HexToAddress(contract)
	newVisit, err := visit.NewVisit(contractAddr, client)
	if err != nil {
		return 0, err
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	// 设置 Gas 限制和 Gas 价格
	auth.GasLimit = uint64(100000) // 根据 visit() 函数调整 Gas 限制
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, fmt.Errorf("获取 Gas 价格失败: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(0) // visit() 函数不发送 ETH

	tx, err := newVisit.Visit(auth)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	log.Printf("交易hash: %s \n", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("访问visit 交易确认失败: %v", err)
	}

	if receipt.Status != 1 {
		log.Fatalf("交易未完成 receipt.status: %v", receipt.Status)
	} else {
		log.Printf("交易确认✅: %s, receipt.status: %d \n", tx.Hash().Hex(), receipt.Status)
	}

	// 7. 查询 count 变量
	count, err := newVisit.Count(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return 0, fmt.Errorf("查询 count 失败: %v", err)
	}
	log.Printf("当前 count 值: %d", count.Uint64())

	return count.Uint64(), nil
}
