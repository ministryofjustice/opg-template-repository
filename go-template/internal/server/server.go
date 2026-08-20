package server

import (
	"html/template"
	"io"
	"log/slog"
	"net/http"

	"github.com/ministryofjustice/opg-go-common/securityheaders"
	"github.com/ministryofjustice/opg-go-common/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Template interface {
	Execute(wr io.Writer, data any) error
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

func New(logger *slog.Logger, templates map[string]*template.Template, envVars EnvironmentVars) http.Handler {
	mux := http.NewServeMux()
	home := homeHandler(logger, templates["home.gotmpl"])

	static := http.FileServer(http.Dir(envVars.WebDir + "/static"))
	mux.Handle("/assets/", static)
	mux.Handle("/javascript/", static)
	mux.Handle("/stylesheets/", static)
	mux.HandleFunc("/", home)

	return otelhttp.NewHandler(http.StripPrefix(envVars.Prefix, telemetry.Middleware(logger)(securityheaders.Use(mux))), "{{SERVICE_NAME}}")
}

func homeHandler(logger *slog.Logger, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if tmpl == nil {
			_, _ = w.Write([]byte("Hello world...!"))
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.ExecuteTemplate(w, "page", nil); err != nil {
			logger.Error("failed to render home template", slog.Any("err", err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
