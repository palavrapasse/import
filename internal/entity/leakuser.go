package entity

type LeakUser struct {
	UserId AutoGenKey
	LeakId AutoGenKey
}

func NewLeakUser(userId AutoGenKey, leakId AutoGenKey) LeakUser {
	return LeakUser{
		UserId: userId,
		LeakId: leakId,
	}
}
