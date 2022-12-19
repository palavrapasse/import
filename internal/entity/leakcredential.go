package entity

type LeakCredentials struct {
	CredId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakCredentials(credId AutoGenKey, leakId AutoGenKey) LeakCredentials {
	return LeakCredentials{
		CredId: credId,
		LeakId: leakId,
	}
}
