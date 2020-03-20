package tokenizer

//Tokenizer 分词器接口
type Tokenizer interface {
	Tokenize(s string) []string
	TokenSimpleDoc(s string) []string
	TokenDocFile(fpath string) ([][]string, error)
	ExtractTopKWords(s string, topk int) []string
}
