package ctxutil

import "context"

const (
	TrackIdKey = "track_id"
)

func WithTrackId(ctx context.Context, trackId string) context.Context {
	return context.WithValue(ctx, TrackIdKey, trackId)
}

func GetTrackId(ctx context.Context) string {
	if ctx.Value(TrackIdKey) != "" {
		s, ok := ctx.Value(TrackIdKey).(string)
		if ok {
			return s
		}
	}
	return ""
}
