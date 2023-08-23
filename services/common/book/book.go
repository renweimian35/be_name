package book

type Info struct {
	Book    string    `json:"book"`    // 源：如楚辞、诗经
	SerNum  int       `json:"serNum"`  // 序号
	Content []Content `json:"context"` //具体内容
}

type Content struct {
	Title      string   `json:"title"`      // 标题，如：短歌行
	Author     string   `json:"author"`     // 作者
	Paragraphs []string `json:"paragraphs"` // 段落内容
}
