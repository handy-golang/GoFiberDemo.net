package ready

import (
	"time"

	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	mCycle.New(mCycle.Opt{
		Func:      GetData,
		SleepTime: time.Hour * 4,
	}).Start()
}

func GetData() {
	// something....
}
