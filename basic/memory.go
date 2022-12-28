package basic

const (
	Mem4G   = 4
	Mem16G  = 16
	Mem128G = 128
)

var MemorySize = map[int]int{
	4:   3,
	8:   6,
	16:  12,
	32:  30,
	64:  70,
	128: 150,
}

var memory = MemorySize[Mem16G]

var (
	KeyChanSize         = memory * 1000 //关键流程的通道大小
	KeyCacheSize        = memory * 1000 //关键流程的缓存大小
	KeyAccountCacheSize = memory * 100  //关键账户链流程的缓存大小
	KeyDaemonCacheSize  = 100           //关键守护链流程的缓存大小
	KeyTemporarySize    = memory * 100  //关键流程的暂时性数据大小
	SecondaryChanSize   = memory * 100  //次要流程的通道大小
	SecondaryCacheSize  = memory * 100  //次要流程的缓存大小
)
