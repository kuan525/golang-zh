// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int
var cycles float64
var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0xFF, 0, 0xFF}}

// color.RGBA{0, 0xFF, 0, 0xFF}} green

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2 // green color in palette
)

func main() {
	cycles = 5
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/cycles=") == true {
			str := strings.TrimPrefix(r.URL.Path, "/cycles=")
			tmp, err := strconv.Atoi(str)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			} else {
				cycles = float64(tmp)
			}
		}
		fmt.Fprintf(os.Stdout, "cycles %.2f\n", cycles)
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	//一大堆，函数内常量，不可修改的常量
	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	//随机生成一个数值
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	//将GIF的LoopCount参数设置位nframes（64） 循环次数
	anim := gif.GIF{LoopCount: nframes}

	//相位差
	phase := 0.0 // phase difference

	//循环次数，每一次搞一下
	for i := 0; i < nframes; i++ {
		// Rect表示图片边界， size（100）
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)

		//新建一个色板 palette是一个切片
		img := image.NewPaletted(rect, palette)

		// cycles（5） math.Pi（PI） res(0.001)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)

			// freq是上面获取的rand.Float64()随机值， phase是相位差
			y := math.Sin(t*freq + phase)

			// 设置颜色 位置
			// img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			if i%2 == 1 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			}
		}

		// 相位差 ++
		phase += 0.1

		// gif anim后面追加
		anim.Delay = append(anim.Delay, delay) // delay（8）每一帧之间的间隔时间
		anim.Image = append(anim.Image, img)   //img
	}

	//保存到文件（out）中
	// 将图片按照帧与帧之间指定的循环次数和时延写入out中
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
