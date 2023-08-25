package result

type Info struct {
	Name          string
	PinYin        string
	Stroke        int
	Title         string
	Source        string
	Author        string
	SingleExplain []WordExplain // 单个的词语解释
	GroupExplain  string        // 名的词组解释
	GroupWord     string        //字的组合
}

type WordExplain struct {
	Word        string `json:"word"`
	Radicals    string `json:"radicals"`    // 偏旁
	Explanation string `json:"explanation"` // 解释
	More        string `json:"more"`        // 更多释义
}
