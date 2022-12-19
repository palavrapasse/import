package entity

type LeakBadActor struct {
	BaId   AutoGenKey
	LeakId AutoGenKey
}

func NewLeakBadActor(baId AutoGenKey, leakId AutoGenKey) LeakBadActor {
	return LeakBadActor{
		BaId:   baId,
		LeakId: leakId,
	}
}
