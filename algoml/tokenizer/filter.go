package tokenizer

import (
	"bufio"
	"os"
	"path"
	"strings"

	"gowatcher/go_monitor/algoml/base"
	"gowatcher/go_monitor/algoml/config"
)

//POSFilter 词性过滤词典
var POSFilter = map[string]struct{}{
	"l": struct{}{},
	"i": struct{}{},
	"c": struct{}{},
	"d": struct{}{},
	"m": struct{}{},
	"r": struct{}{},
	"t": struct{}{},
	"u": struct{}{},
	"e": struct{}{},
	"y": struct{}{},
}

//IsHitPOSFilter 是否命中
func IsHitPOSFilter(pos string) bool {
	if _, ok := POSFilter[pos]; ok {
		return true
	}
	return false
}

//StopWords 停止词字典，用map的key进行查找
type StopWords struct {
	Dict map[string]struct{}
}

//NewStopWords 创建停止词词典
func NewStopWords() *StopWords {
	filter := &StopWords{
		Dict: make(map[string]struct{}),
	}
	wpath := path.Join(config.DictDir, "stop_words.utf8")
	if err := filter.LoadStopWordsFile(wpath); err != nil {
		return nil
	}
	return filter
}

//LoadStopWordsFile 加载停止词文件
func (p *StopWords) LoadStopWordsFile(fpath string) error {
	swfile, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer swfile.Close()

	scanner := bufio.NewScanner(swfile)
	for scanner.Scan() {
		p.Add(scanner.Text())
	}
	return nil
}

//Exist 判定停止词是否存在
func (p *StopWords) Exist(word string) bool {
	if _, exist := p.Dict[word]; exist {
		return true
	}
	return false
}

//Add 增加停止词
func (p *StopWords) Add(word string) {
	if !p.Exist(word) {
		p.Dict[word] = struct{}{}
	}
}

//Del 删除停止词
func (p *StopWords) Del(word string) {
	if p.Exist(word) {
		delete(p.Dict, word)
	}
}

//FilterStopWords 过滤分词后的停止词
func (p *StopWords) FilterStopWords(raw []string) []string {
	result := []string{}
	for _, w := range raw {
		if p.Exist(w) || base.IsDigitForm(w) {
			continue
		}
		result = append(result, w)
	}
	return result
}

//FilterWordsByPOS 过滤无关词性的词
func (p *StopWords) FilterWordsByPOS(words []string, pos []string) []string {
	posMap := make(map[string]struct{})
	for _, p := range pos {
		tmp := strings.Split(p, "/")
		if len(tmp) < 2 || len(tmp[1]) < 1 {
			continue
		}

		if IsHitPOSFilter(tmp[1][0:1]) {
			continue
		}

		posMap[tmp[0]] = struct{}{}
	}

	result := []string{}
	for _, w := range words {
		if _, ok := posMap[w]; ok {
			result = append(result, w)
		}
	}

	return result
}

//FilterWords 词汇过滤
func (p *StopWords) FilterWords(raw []string, pos []string) []string {
	result := p.FilterStopWords(raw)
	result = p.FilterWordsByPOS(result, pos)
	return result
}
