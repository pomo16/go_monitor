package tokenizer

import (
	"bufio"
	"os"
	"path"

	"gowatcher/go_monitor/algoml/config"

	"github.com/yanyiwu/gojieba"
)

//JiebaTokenizer 中文分词器
type JiebaTokenizer struct {
	jieba    *gojieba.Jieba
	filter   *StopWords
	useHMM   bool
	filterOn bool
}

//GetJiebaDictPath 获取jieba的词典路径
func GetJiebaDictPath() []string {
	dict := path.Join(config.DictDir, "jieba.dict.utf8")
	hmm := path.Join(config.DictDir, "hmm_model.utf8")
	user := path.Join(config.DictDir, "user.dict.utf8")
	idf := path.Join(config.DictDir, "idf.utf8")
	stopWords := path.Join(config.DictDir, "stop_words.utf8")
	return []string{dict, hmm, user, idf, stopWords}
}

//NewJiebaTokenizer 初始化结巴分词器
func NewJiebaTokenizer(loadStopWords bool) *JiebaTokenizer {
	path := GetJiebaDictPath()
	tokenizer := &JiebaTokenizer{
		jieba:    gojieba.NewJieba(path[0], path[1], path[2], path[3], path[4]),
		useHMM:   true,
		filterOn: false,
	}
	//加载停用词
	if loadStopWords {
		tokenizer.filterOn = true
		tokenizer.filter = NewStopWords()
	}
	return tokenizer
}

//SwitchWordFilter 控制停止词开关
func (x *JiebaTokenizer) SwitchWordFilter(status bool) {
	x.filterOn = status
}

//Tokenize jieba默认采取精准分词
func (x *JiebaTokenizer) Tokenize(s string) []string {
	raw := x.jieba.Cut(s, x.useHMM)
	if x.filterOn {
		return x.filter.FilterStopWords(raw)
	}
	return raw
}

//Free 分词器用完后把相关资源释放掉
func (x *JiebaTokenizer) Free() {
	x.jieba.Free()
}

//TokenDocFile 用分词器取文本语料特征
func (x *JiebaTokenizer) TokenDocFile(fpath string) ([][]string, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tokenSet := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := x.Tokenize(scanner.Text())
		tokenSet = append(tokenSet, tokens)
	}

	return tokenSet, nil
}

//TokenSimpleDoc 取样本语料特征
func (x *JiebaTokenizer) TokenSimpleDoc(sample string) []string {
	return x.Tokenize(sample)
}

//ExtractTopKWords 利用TF-IDF提取TopK关键词
func (x *JiebaTokenizer) ExtractTopKWords(sample string, topk int) []string {
	rankWords := x.jieba.Extract(sample, topk)
	posWords := x.jieba.Tag(sample)
	result := x.filter.FilterWords(rankWords, posWords)
	return result
}
