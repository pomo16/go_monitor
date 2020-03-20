package base

//AddOneProb 自增频率计数器
type AddOneProb struct {
	WCount map[string]int64 `json:"wcount"`
	WTotal int64            `json:"wtotal"`
}

//NewAddOneProb 创建频率计数器
func NewAddOneProb() *AddOneProb {
	return &AddOneProb{
		WCount: make(map[string]int64),
		WTotal: 0,
	}
}

//Add 累加特定key的计数，key统一多加1是为了抵消训练集词不存在问题
func (p *AddOneProb) Add(key string, val int64) {
	p.WTotal += val
	if _, ok := p.WCount[key]; !ok {
		p.WCount[key] = 1
		p.WTotal++
	}
	p.WCount[key] += val
}

//GetCount 获取key的count数，返回1是为了解决对数inf问题
func (p *AddOneProb) GetCount(key string) int64 {
	if v, ok := p.WCount[key]; ok {
		return v
	}
	return 1
}

//Freq 计算特定key出现的频率
func (p *AddOneProb) Freq(key string) float64 {
	return float64(p.GetCount(key)) / float64(p.WTotal)
}

//Sum 返回总词数
func (p *AddOneProb) Sum() int64 {
	return p.WTotal
}
