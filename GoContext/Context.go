package GoContext

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Key int

const UserIdctx Key = 0

func ProcessLongTask(ctx context.Context) string {
	id := ctx.Value(UserIdctx).(int)
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("done processing id: %d\n", id)
	case <-ctx.Done():
		slog.Warn("request canceled")
		return ""
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")

	ctx := context.WithValue(r.Context(), UserIdctx, id)

	result := ProcessLongTask(ctx)

	w.Write([]byte(result))
}
