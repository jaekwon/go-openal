package main

import (
	"fmt"
	"time"

	"github.com/jaekwon/go-openal/openal"
)

func main() {

	mic := openal.CaptureOpenDevice("", 22050, openal.FormatMono16, 22050/2)
	mic.CaptureStart()

	device := openal.OpenDevice("")
	context := device.CreateContext()
	context.Activate()

	//listener := new(openal.Listener)
	//listener.

	source := openal.NewSource()
	source.SetPitch(1)
	source.SetGain(1)
	source.SetPosition(0, 0, 0)
	source.SetVelocity(0, 0, 0)
	source.SetLooping(false)

	for i := 0; i < 1000000; i++ {
		buf := mic.CaptureSamples(22050 / 2)
		fmt.Printf("%X\n", buf)
		fmt.Println(len(buf))
		buffer := openal.NewBuffer()
		buffer.SetData(openal.FormatMono16, buf, 22050/2)
		source.SetBuffer(buffer)
		source.Play()
		for source.State() == openal.Playing {
			//loop long enough to let the wave file finish
		}
	}
	fmt.Println(source.State())

	source.Pause()
	source.Stop()
	return
	context.Destroy()
	time.Sleep(time.Second)
}
