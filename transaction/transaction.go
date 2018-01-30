package transaction

import (
	"fmt"

	"github.com/satori/go.uuid"
)

//Vout是该输出在那笔交易中所有输出的索引
//ScriptSig 是一个脚本，signature提供了可解锁输出结构里面 ScriptPubKey 字段的数据
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

func (self *Transaction) SetID() {
	u1 := uuid.Must(uuid.NewV4()).String()
	self.ID = []byte(u1)
}

//coinbase交易，这是第一次交易
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	subsidy := 50 //挖矿奖励
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}
