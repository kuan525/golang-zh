// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

/*
type GIF struct {
	Image     []*image.Paletted // 连续的图片
	Delay     []int             // 连续的延迟时间，每一帧单位都是百分之一秒，delay中数值表示其两个图像动态展示的时间间隔
	LoopCount int               // 循环次数，如果为0则一直循环。
}
*/

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0xFF, 0, 0xFF}}

// color.RGBA{0, 0xFF, 0, 0xFF}} green

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2 // green color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.

	// 设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	rand.Seed(time.Now().UTC().UnixNano())
	// rand.Intn(100) (0 <= x < n)

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	//一大堆，函数内常量，不可修改的常量
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
