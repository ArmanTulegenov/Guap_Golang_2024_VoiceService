package main

import (
	"fmt"

	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	inputPath  = "./sample1.m4a"
	outputPath = "output.pcm"
)

func main() {
	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		panic(err)
	}

	trans.MediaFile().SetAudioChannels(1)
	trans.MediaFile().SetAudioRate(16000)
	trans.MediaFile().SetOutputFormat("s16le")
	trans.MediaFile().SetAudioCodec("pcm_s16le")

	done := trans.Run(true)
	progress := trans.Output()
	for p := range progress {
		fmt.Println(p)
	}

	fmt.Println(<-done)

	whisper := NewWhisperWrapper(outputPath)
	text, err := whisper.ConvertToText()
	if err != nil {
		panic(err)
	}

	fmt.Println("Converted text:", text)
}
