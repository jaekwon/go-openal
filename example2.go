package main

import (
	"fmt"
	"time"

	"github.com/jaekwon/go-openal/openal"
)

func main() {

	mic := openal.CaptureOpenDevice("", 22050, openal.FormatMono16, 22050*2)
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

	buffersAll := openal.NewBuffers(10)
	buffersFree := make([]openal.Buffer, len(buffersAll))
	copy(buffersFree, buffersAll)

	for {
		// Get any free buffers
		buffersProcessed := source.BuffersProcessed()
		if buffersProcessed > 0 {
			buffersNewFree := make([]openal.Buffer, buffersProcessed)
			source.UnqueueBuffers(buffersNewFree)
			buffersFree = append(buffersFree, buffersNewFree...)
		}
		if len(buffersFree) == 0 {
			continue
		}

		//fmt.Println("queued:", source.BuffersQueued())
		//fmt.Println("processed:", source.BuffersProcessed())
		//fmt.Println("captured:", mic.CapturedSamples())

		captureSize := uint32(512)
		if mic.CapturedSamples() >= captureSize {
			inputBytes := mic.CaptureSamples(captureSize)
			buffer := buffersFree[len(buffersFree)-1]
			buffersFree = buffersFree[:len(buffersFree)-1]
			buffer.SetData(openal.FormatMono16, inputBytes, 22050)
			source.QueueBuffer(buffer)

			// If we have enough buffers, start playing
			if source.State() != openal.Playing {
				if source.BuffersQueued() > 2 {
					fmt.Println("Start playing")
					source.Play()
				}
			}
		}
	}
	fmt.Println(source.State())

	source.Pause()
	source.Stop()
	return
	context.Destroy()
	time.Sleep(time.Second)
}
