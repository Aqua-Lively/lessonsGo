package main

import (
	"errors"
	"log"
	"sync"
)

type playerWallet struct {
	coins int64
	mu    sync.Mutex
}

func (w *playerWallet) spendCoins(amount int64) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.coins-amount < 0 {
		return errors.New("erros!")
	}

	w.coins -= amount
	log.Printf("spend %d coin(s), balance: %d", amount, w.coins)
	return nil
}

func payForUnitsAndBuildings(w *playerWallet) {
	if err := w.spendCoins(3); err != nil {
		log.Println(err)
	}
}

func buyAnotherUnitInShop(w *playerWallet) {
	if err := w.spendCoins(4); err != nil {
		log.Println(err)
	}
}
