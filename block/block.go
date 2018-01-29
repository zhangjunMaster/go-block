package block

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"fmt"
	"math/big"
	"math"
	//"go-block/models"
	// "github.com/satori/go.uuid"
)

//生成区块
type Block struct {
    Timestamp     int64
	//Data          []byte
	
    PrevBlockHash []byte
	Hash          []byte
	Nonce 		  int
}
//bytes.join
//因为有了工作量证明，就不能随便塞hash了
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
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
//区块链就是区块的数组
func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	//result, err := models.Db.Exec(
	//	"INSERT INTO `block`(`name`,`password`) VALUES('tom', 'tom')"
	//)
}

//生成初创快
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) 
}

//生成区块链，*Block{}类型 NewGenesisBlock()生成的数据
func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//工作量认证
const targetBits = 24

//工作量认证的结构体
type ProofOfWork struct {
	block *Block
	target *big.Int
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
			pow.block.Data,
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
	fmt.Printf("\nMining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)	
		hashInt.SetBytes(hash[:]) //将哈希转换成一个大整数
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
func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
 