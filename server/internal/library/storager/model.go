package storager

// FileMeta 文件元数据
type FileMeta struct {
	Filename  string // 文件名称
	Size      int64  // 文件大小
	Kind      string // 文件所属分类
	MetaType  string // 文件类型
	NaiveType string // NaiveUI类型
	Ext       string // 文件后缀名
	Md5       string // 文件hash
}
