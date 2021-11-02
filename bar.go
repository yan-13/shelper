package shelper

import (
    "fmt"
    "strconv"
    "strings"
)

type Bar struct {
    percent int64  //百分比
    cur     int    //当前进度位置
    total   int    //总进度
    rate    string //进度条
    graph   string //显示符号
}

const (
    defaultGraph = "█"
    maxGraphNum  = 50
)

func (bar *Bar) New(total int) {
    bar.cur = 0
    bar.total = total
    bar.graph = defaultGraph
    bar.percent = bar.getPercent()
    bar.rate = bar.getRate()
}

func (bar *Bar) SetGraph(graph string) {
    if 1 != len(graph) {
        graph = defaultGraph
    }
    bar.graph = graph
}

func (bar *Bar) Play(cur int) {
    bar.cur = cur
    bar.percent = bar.getPercent()
    bar.rate = bar.getRate()
    fmt.Printf("\r[%-"+strconv.Itoa(maxGraphNum)+"s]%3d%%  %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
    fmt.Println()
}

func (bar *Bar) getPercent() int64 {
    return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) getRate() string {
    return strings.Repeat(bar.graph, int(bar.percent*maxGraphNum/100))
}
