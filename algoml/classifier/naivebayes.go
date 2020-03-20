package classifier

import (
	"encoding/json"
	"gowatcher/go_monitor/algoml/base"
	"io/ioutil"
	"math"
	"strconv"

	"github.com/tidwall/gjson"
)

//Class 代表数据类型
type Class string

//NaiveBayes 朴素贝叶斯分类器
type NaiveBayes struct {
	Total int64                      `json:"total"`
	Data  map[Class]*base.AddOneProb `json:"data"`
}

//NewNaiveBayes 创建朴素贝叶斯模型
func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		Total: 0,
		Data:  make(map[Class]*base.AddOneProb),
	}
}

//Load 加载数据文件
func (p *NaiveBayes) Load(fpath string) error {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	root := string(data)
	p.Total, _ = strconv.ParseInt(gjson.Get(root, "total").String(), 10, 64)
	d := gjson.Get(root, "data")
	d.ForEach(func(key, val gjson.Result) bool {
		wTotal, _ := strconv.ParseInt(val.Get("wtotal").String(), 10, 64)
		polarity := Class(key.String())
		p.Data[polarity] = &base.AddOneProb{
			WCount: make(map[string]int64),
			WTotal: wTotal,
		}
		c := val.Get("wcount")
		c.ForEach(func(key, val gjson.Result) bool {
			cnt, _ := strconv.ParseInt(val.String(), 10, 64)
			p.Data[polarity].WCount[key.String()] = cnt
			return true
		})
		return true
	})

	return nil
}

//Save 存储数据文件
func (p *NaiveBayes) Save(fpath string) error {
	data, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fpath, data, 0644); err != nil {
		return err
	}
	return nil
}

//Train 贝叶斯训练过程
func (p *NaiveBayes) Train(docs [][]string, class Class) {
	for _, d := range docs {
		p.trainOne(d, class)
	}
	p.calTotalWords()
}

//trainOne 文档训练原子过程
func (p *NaiveBayes) trainOne(doc []string, class Class) {
	if _, exist := p.Data[class]; !exist {
		p.Data[class] = &base.AddOneProb{
			WCount: make(map[string]int64),
			WTotal: 0,
		}
	}

	for _, s := range doc {
		p.Data[class].Add(s, 1)
	}
}

//calTotalWords 计算训练集总词数
func (p *NaiveBayes) calTotalWords() {
	var tmpTotal int64
	for _, v := range p.Data {
		tmpTotal += v.Sum()
	}
	p.Total = tmpTotal
}

//Classify 数据识别
func (p *NaiveBayes) Classify(doc []string) (string, float64) {
	tmpD := make(map[Class]float64)
	for k, v := range p.Data {
		tmpD[k] += math.Log(float64(v.Sum())) - math.Log(float64(p.Total))
		for _, word := range doc {
			tmpD[k] += math.Log(v.Freq(word))
		}
	}

	class, prob := "", 0.0
	for k := range p.Data {
		now := 0.0
		for otherk := range p.Data {
			now += math.Exp(tmpD[otherk] - tmpD[k])
		}
		now = 1 / now
		if now > prob {
			class, prob = string(k), now
		}
	}

	return class, prob
}
