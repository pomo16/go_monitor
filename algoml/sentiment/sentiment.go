package sentiment

import (
	"path"

	"gowatcher/go_monitor/algoml/classifier"
	"gowatcher/go_monitor/algoml/config"
	"gowatcher/go_monitor/algoml/tokenizer"
)

//Sentiment 情感分析
type Sentiment struct {
	Tokenizer  tokenizer.Tokenizer
	Classifier *classifier.NaiveBayes
}

//NewSentiment 创建情感分析器
func NewSentiment(tokenizer tokenizer.Tokenizer) *Sentiment {
	sentiment := &Sentiment{
		Tokenizer:  tokenizer,
		Classifier: classifier.NewNaiveBayes(),
	}
	modelfile := path.Join(config.ModelDir, "sentiment/sentiment.json")
	sentiment.load(modelfile)
	return sentiment
}

//Classify 对输入样本做正负向判定
func (p *Sentiment) Classify(sample string) float64 {
	doc := p.Tokenizer.TokenSimpleDoc(sample)
	polarity, score := p.Classifier.Classify(doc)
	if polarity == "pos" {
		return score
	}
	return 1 - score
}

//Train 根据正负向语料训练集执行训练
func (p *Sentiment) Train(fpos string, fneg string) error {
	posSet, err := p.Tokenizer.TokenDocFile(fpos)
	if err != nil {
		return err
	}

	negSet, err := p.Tokenizer.TokenDocFile(fneg)
	if err != nil {
		return err
	}

	p.Classifier.Train(posSet, "pos")
	p.Classifier.Train(negSet, "neg")
	p.save(path.Join(config.ModelDir, "sentiment.json"))
	return nil
}

//load 加载模型数据文件
func (p *Sentiment) load(fpath string) {
	p.Classifier.Load(fpath)
}

//saveDataModel 存储模型数据文件
func (p *Sentiment) save(fpath string) {
	p.Classifier.Save(fpath)
}
