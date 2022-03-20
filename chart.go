package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/zserge/lorca"
)

var (
	buf9  [512]byte
	buf10 [1024]byte
	buf11 [2048]byte
	buf12 [4096]byte

	curved    chan *CurveD
	delimiter string
	typ       string
)

func init() {
	curved = make(chan *CurveD, 3)

	flag.StringVar(&delimiter, "d", " ", "-d ,  \ndata delimiter")
	flag.StringVar(&typ, "t", "timeline", "-t line  \nnote: chart type")
}

func app() {
	ui, err := lorca.New("", "", 2200, 1200)
	if err != nil {
		panic(err)
	}

	ui.Bind("curve", Curve)
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// go http.Serve(ln, http.FileServer(FS))
	go http.Serve(ln, http.FileServer(http.Dir("./tpl")))
	fmt.Println(ln.Addr())
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
}

func main() {
	flag.Parse()
	//先取程序的标准输入属性信息
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	// 判断标准输入设备属性 os.ModeCharDevice 是否设置
	// 同时判断是否有数据输入
	if (info.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
		bs, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		parseData(bs)
	}
	app()
}

func parseData(input []byte) {
	if typ == "line" {
		parseLineData(input)
		return
	}
	if typ == "timeline" {
		parseTimeLineJsonData(input)
		return
	}
}

func parseLineData(input []byte) {
	datas := strings.Split(string(input), delimiter)
	data := make([]float64, 0, len(datas))
	for _, it := range datas {
		v, err := strconv.ParseFloat(strings.Trim(it, "\n"), 64)
		if err != nil {
			logrus.Error(err)
			continue
		}
		data = append(data, v)
	}
	fmt.Fprintf(os.Stdout, "data: %+v\n", data)

	curved <- &CurveD{
		Type:   "line",
		Title:  "title",
		Data:   data,
		Avg:    0.0,
		PosAvg: 0.0,
		NegAvg: 0.0,
	}
}

type Tline struct {
	T int64
	D float64
}

func parseTimeLineJsonData(input []byte) {
	metrics, err := decodeTimeDatas(input)
	if err != nil {
		panic(err)
	}
	data := parseTimeData(metrics)

	curved <- &CurveD{
		Type:   "timeline",
		Title:  "title",
		Data:   data,
		Avg:    0.0,
		PosAvg: 0.0,
		NegAvg: 0.0,
	}
}

type CurveD struct {
	Type   string      `json:"type"`
	Title  string      `json:"title"`
	Data   interface{} `json:"data"`
	Avg    float64     `json:"avg"`
	PosAvg float64     `json:"posavg"`
	NegAvg float64     `json:"negavg"`
}

func Curve() string {
	cd := <-curved
	defer func() {
		curved <- cd
	}()
	bs, err := json.Marshal(cd)
	if err != nil {
		return err.Error()
	}
	return string(bs)
}
