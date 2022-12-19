package entity

type LeakPlatform struct {
	PlatId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakPlatform(platId AutoGenKey, leakId AutoGenKey) LeakPlatform {
	return LeakPlatform{
		PlatId: platId,
		LeakId: leakId,
	}
}
