package rollingBuffer

// rolling buffer

type RollingBuf struct {
	values     []interface{}
	max        int
	addPos     int
	currentPos int
	startPos   int
	len        int
}

func New(max int) *RollingBuf {
	return &RollingBuf{
		values:   make([]interface{}, max),
		max:      max,
		addPos:   0,
		startPos: 0,
		len:      0,
	}
}
func (rBuf *RollingBuf) Add(value interface{}) {

	if rBuf.addPos == rBuf.max {
		// buffer has been filled, start overwriting at the first location
		rBuf.addPos = 0
	}
	if rBuf.len == rBuf.max {
		// advance where data now starts
		rBuf.startPos++
		if rBuf.startPos == rBuf.max {
			// startPos has reached the end, point at the first location
			rBuf.startPos = 0
		}
	}
	rBuf.currentPos = rBuf.addPos
	rBuf.values[rBuf.currentPos] = value
	rBuf.addPos++
	if rBuf.len < rBuf.max {
		// new addition, increase the length
		rBuf.len++
	}

}
func (rBuf *RollingBuf) Current() interface{} {
	return rBuf.values[rBuf.currentPos]
}

func (rBuf *RollingBuf) All() []interface{} {
	// create a new slice to return
	rBufRet := make([]interface{}, rBuf.max)
	for count := 0; count < rBuf.max; count++ {
		readPos := rBuf.startPos + count
		if readPos >= rBuf.max {
			readPos = rBuf.startPos + count - rBuf.max
		}
		rBufRet[count] = rBuf.values[readPos]
	}
	return rBufRet
}
