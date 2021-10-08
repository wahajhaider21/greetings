package assignment01IBC

type BlockData struct {
Transactions []string
}
type Block struct {
Data        BlockData
PrevPointer *Block
PrevHash    string
CurrentHash string
}
func CalculateHash(inputBlock *Block) string {
...
}
func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
...
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
...
}
func ListBlocks(chainHead *Block) {
...
}
func VerifyChain(chainHead *Block) {
...
}