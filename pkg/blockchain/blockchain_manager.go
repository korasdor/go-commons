package blockchain

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainManager interface {
	GetBalance(addrHex string) (int, error)
	Stop()
}

type Manager struct {
	client *ethclient.Client
}

func NewManager(uri string) (*Manager, error) {
	client, err := ethclient.DialContext(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	return &Manager{
		client: client,
	}, nil
}

func (m *Manager) GetBalance(addrHex string) (int, error) {
	addr := common.HexToAddress(addrHex)
	balance, err := m.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return 0, fmt.Errorf("error to get a balance %w", err)
	}

	convertedBalance := m.convertBalanceToEther(balance)

	return convertedBalance, nil
}

func (m *Manager) convertBalanceToEther(balance *big.Int) int {
	// convert balance to ehter (1 ehter = 10^18 wei)
	floatBalane := &big.Float{}
	floatBalane.SetString(balance.String())

	balanceEther := &big.Float{}
	balanceEther.Quo(floatBalane, big.NewFloat(math.Pow10(18)))

	fmt.Println(balanceEther)

	return 0
}

func (m *Manager) Stop() {
	m.client.Close()
}

// client, err := ethclient.DialContext(context.Background(), alchemyUrl)
// 	if err != nil {
// 		log.Fatalf("Error to create a ether client %v", err)
// 	}
// 	defer client.Close()

// 	//get last block
// 	block, err := client.BlockByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatalf("Error to get a block %v", err)
// 	}
// 	fmt.Println("The block number:", block.Number())

// 	//get balance
// 	addrHex := "0xE94f1fa4F27D9d288FFeA234bB62E1fBC086CA0c"
// 	addr := common.HexToAddress(addrHex)

// 	balance, err := client.BalanceAt(context.Background(), addr, nil)
// 	if err != nil {
// 		log.Fatalf("Error to get a balance %v", err)
// 	}
// 	fmt.Println("The balance:", balance)

// 	//convert balance to ehter (1 ehter = 10^18 wei)
// 	fBalane := &big.Float{}
// 	fBalane.SetString(balance.String())

// 	balanceEther := &big.Float{}
// 	balanceEther.Quo(fBalane, big.NewFloat(math.Pow10(18)))
// 	fmt.Println(balanceEther)
