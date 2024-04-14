package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Wallet interface {
	GetID() string
	GetTokens() int
	GetOnlineTime() time.Duration
	GetCoinAge() int
	SetTokens(tokens int)
	SetOnlineTime(onlineTime time.Duration)
	SetCoinAge(coinAge int)
}

type wallet struct {
	id         string
	tokens     int
	onlineTime time.Duration
	coinAge    int
}

func (w *wallet) GetID() string {
	return w.id
}

func (w *wallet) GetTokens() int {
	return w.tokens
}

func (w *wallet) GetOnlineTime() time.Duration {
	return w.onlineTime
}

func (w *wallet) GetCoinAge() int {
	return w.coinAge
}

func (w *wallet) SetTokens(tokens int) {
	w.tokens = tokens
}

func (w *wallet) SetOnlineTime(onlineTime time.Duration) {
	w.onlineTime = onlineTime
}

func (w *wallet) SetCoinAge(coinAge int) {
	w.coinAge = coinAge
}

var Wallets []Wallet

func createWallet(id int) Wallet {
	tokens := rand.Intn(100) + 1 // 随机生成1-100个代币
	return &wallet{
		id:         strconv.Itoa(id),
		tokens:     tokens,
		onlineTime: 0,
	}
}

func distributeTokensAndCalculateCoinAge(wallets []Wallet) {
	for i := 0; i < len(wallets); i++ {
		wallet := wallets[i]
		wallet.SetTokens(rand.Intn(100) + 1)
		onlineTime := time.Duration(rand.Intn(60)+1) * time.Minute // 随机在线时间（1-10分钟）
		wallet.SetOnlineTime(wallet.GetOnlineTime() + onlineTime)
		wallet.SetCoinAge(wallet.GetTokens() * int(wallet.GetOnlineTime().Minutes()))
	}
}

func SelectMaxCoinAgeWallet(wallets []Wallet) int {
	maxCoinAge := 0
	selectedWalletIndex := -1
	for i := 0; i < len(wallets); i++ {
		wallet := wallets[i]
		coinAge := wallet.GetTokens() * int(wallet.GetOnlineTime().Minutes())
		if coinAge > maxCoinAge {
			maxCoinAge = coinAge
			selectedWalletIndex = i
		}
	}
	return selectedWalletIndex
}

func SelectMaxCoinAge(wallets []Wallet) int {
	maxCoinAge := 0
	for i := 0; i < len(wallets); i++ {
		wallet := wallets[i]
		coinAge := wallet.GetTokens() * int(wallet.GetOnlineTime().Minutes())
		if coinAge > maxCoinAge {
			maxCoinAge = coinAge
		}
	}
	return maxCoinAge
}

func resetSelectedWalletCoinAge(wallets []Wallet, selectedWalletIndex int) {
	wallets[selectedWalletIndex].SetCoinAge(0)
}

func Get() int {
	// 初始化钱包
	Wallets = make([]Wallet, 0)

	// 创建多个用户（钱包）
	for i := 1; i <= 10; i++ {
		wallet := createWallet(i)
		Wallets = append(Wallets, wallet)
	}

	// 动态在线，空投随机的代币并统计在线时间
	distributeTokensAndCalculateCoinAge(Wallets)

	// 清空切片
	CoinAgeSlice = CoinAgeSlice[:0]
	// 统计币龄，币龄等于用户钱包的代币乘以该用户持有该代币的在线时间
	for i := 0; i < len(Wallets); i++ {
		wallet := Wallets[i]
		CoinAgeSlice = append(CoinAgeSlice, wallet.GetCoinAge())
		fmt.Printf("用户 %d 的币龄为 %d\n", i+1, wallet.GetCoinAge())
	}
	fmt.Println("---------------")

	// 选出币龄最大值作为记账者，被选出者币龄清零
	selectedWalletIndex := SelectMaxCoinAgeWallet(Wallets)
	//fmt.Printf("被选出的记账者是用户 %d\n", selectedWalletIndex+1)
	SelectedWalletCoinAge = Wallets[selectedWalletIndex].GetCoinAge()
	resetSelectedWalletCoinAge(Wallets, selectedWalletIndex)
	return selectedWalletIndex + 1
}

var SelectedWalletCoinAge int
var CoinAgeSlice []int
var SelectWalletSlice []int

func runPosNet() {
	index := Get()
	//fmt.Printf("新的记账者是用户 %d\n", index)
	SelectWalletSlice = append(SelectWalletSlice, index)
}
