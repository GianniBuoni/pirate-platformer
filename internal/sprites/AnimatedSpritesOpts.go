package sprites

type AnimatedOpts struct {
	// TODO add fields to add speed and other things
	frameSize float32
	flipH     bool
	flipV     bool
}

func WithFrameSize(fs float32) func(*AnimatedOpts) {
	return func(ao *AnimatedOpts) {
		ao.frameSize = fs
	}
}
