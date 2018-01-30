package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
	//"go-block/models"
	// "github.com/satori/go.uuid"
	. "go-block/transaction"
)

//生成区块
type Block struct {
	Timestamp int64
	//Data          []byte
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

//bytes.join
//因为有了工作量证明，就不能随便塞hash了
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	//[]byte{}空的byte类型的数组
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//生成区块链
type Blockchain struct {
	Blocks []*Block
}

//区块链就是区块的数组,将所有的data替换成transactions交易
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	//result, err := models.Db.Exec(
	//	"INSERT INTO `block`(`name`,`password`) VALUES('tom', 'tom')"
	//)
}

//生成初创快
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//生成区块链，*Block{}类型 NewGenesisBlock()生成的数据，区块链必须有地址
func NewBlockchain(address string) *Blockchain {
	genesisCoinbaseData := "这是第一个块"
	cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
	genesis := NewGenesisBlock(cbtx)
	return &Blockchain{[]*Block{genesis}}
}

//工作量认证
const targetBits = 24

//工作量认证的结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	// 获取，每一个交易的hash
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	// 最后获得一个连接后的组合哈希
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

//生成工作量认证
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

//准备数据进行hash
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			[]byte(strconv.FormatInt(pow.block.Timestamp, 16)),
			[]byte(strconv.FormatInt(int64(targetBits), 16)),
			[]byte(strconv.FormatInt(int64(nonce), 16)),
		},
		[]byte{},
	)
	return data
}

//实现 pow算法
//SetBytes,Lsh,Cmp
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var resultHash [32]byte
	nonce := 0
	maxNonce := math.MaxInt64
	//fmt.Printf("\nMining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)
		hashInt.SetBytes(hash[:])          //将哈希转换成一个大整数
		if hashInt.Cmp(pow.target) == -1 { //将这个大整数与目标进行比较
			fmt.Printf("\r%x", hash)
			resultHash = hash
			break
		} else {
			nonce++
		}
	}
	return nonce, resultHash[:]
}

//检验工作量证明
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}

/*
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
}
*/
