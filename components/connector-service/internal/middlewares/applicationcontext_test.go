package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kyma-project/kyma/components/connector-service/internal/httpcontext"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testApplication = "test-app"
)

func TestApplicationContextMiddleware_Middleware(t *testing.T) {

	t.Run("should set context based on header", func(t *testing.T) {
		// given
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			clusterCtx, ok := ctx.Value(httpcontext.ApplicationContextKey).(httpcontext.ApplicationContext)
			require.True(t, ok)

			assert.Equal(t, testApplication, clusterCtx.Application)
			w.WriteHeader(http.StatusOK)
		})

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)
		req.Header.Set(httpcontext.ApplicationHeader, testApplication)
		req.Header.Set(httpcontext.TenantHeader, testTenant)
		req.Header.Set(httpcontext.GroupHeader, testGroup)

		rr := httptest.NewRecorder()

		middleware := NewApplicationContextMiddleware()

		// when
		resultHandler := middleware.Middleware(handler)
		resultHandler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 400 if no application header provided", func(t *testing.T) {
		// given
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)
		req.Header.Set(httpcontext.TenantHeader, testTenant)
		req.Header.Set(httpcontext.GroupHeader, testGroup)

		rr := httptest.NewRecorder()

		middleware := NewApplicationContextMiddleware()

		// when
		resultHandler := middleware.Middleware(handler)
		resultHandler.ServeHTTP(rr, req)

		// then
		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
