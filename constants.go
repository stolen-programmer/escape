package escape

// ESC CSI
const (
	ESC = "\u001b"
	CSI = ESC + "["
)

// Cursor
const (
	Up       = "A" // 上
	Down     = "B" // 下
	Forward  = "C" // 前进
	Backward = "D" // 退格
	NextLine = "E" // 下一行
	PrevLine = "F" // 上一行
	HAP      = "G" // current cursor horizontal abstract position
	Move     = "H" // direct position cursor
)

// SGR Color
const (
	SGR       = "m"
	Normal    = CSI + "0" + SGR
	Underline = "4" // 下划线

	// 30 - 37 foreground color 前景色
	FBlack        = "30"
	FRed          = "31"
	FGreen        = "32"
	FBlue         = "34"
	ForegroundExt = "38;2" // 扩展 rgb

	// 40 - 47 background color 背景色
	BBlack        = "40"
	BRed          = "41"
	BGreen        = "42"
	BBlue         = "44"
	BackgroundExt = "48;2"
)

// ReportPos
const (
	ReportPos = "6n"
)

