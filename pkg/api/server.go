package api

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strings"
	"sync/atomic"
	"time"

	proto "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/generated/financial_integration_service_api/v1"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plaid/plaid-go/v14/plaid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	"github.com/SimifiniiCTO/simfiny-core-lib/instrumentation"
	taskprocessor "github.com/SimifiniiCTO/simfiny-core-lib/task-processor"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/plaidhandler"
	database "github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/postgres-database"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/internal/secrets"
	_ "github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/api/docs"
	"github.com/SimifiniiCTO/simfiny-financial-integration-service/pkg/fscache"
)

// @title FIS-Service API
// @version 2.0
// @description Go microservice template for Kubernetes.

// @contact.name Source Code
// @contact.url https://github.com/SimifiniiCTO/simfiny-financial-integration-service

// @license.name MIT License
// @license.url https://github.com/SimifiniiCTO/simfiny-financial-integration-service/blob/master/LICENSE

// @host localhost:9898
// @BasePath /
// @schemes http https

var (
	healthy int32
	ready   int32
	watcher *fscache.Watcher
)

type Config struct {
	HttpClientTimeout              time.Duration `mapstructure:"http-client-timeout"`
	HttpServerTimeout              time.Duration `mapstructure:"http-server-timeout"`
	HttpServerShutdownTimeout      time.Duration `mapstructure:"http-server-shutdown-timeout"`
	BackendURL                     []string      `mapstructure:"backend-url"`
	UILogo                         string        `mapstructure:"ui-logo"`
	UIMessage                      string        `mapstructure:"ui-message"`
	UIColor                        string        `mapstructure:"ui-color"`
	UIPath                         string        `mapstructure:"ui-path"`
	DataPath                       string        `mapstructure:"data-path"`
	ConfigPath                     string        `mapstructure:"config-path"`
	CertPath                       string        `mapstructure:"cert-path"`
	Host                           string        `mapstructure:"host"`
	Port                           string        `mapstructure:"port"`
	SecurePort                     string        `mapstructure:"secure-port"`
	PortMetrics                    int           `mapstructure:"port-metrics"`
	Hostname                       string        `mapstructure:"hostname"`
	H2C                            bool          `mapstructure:"h2c"`
	RandomDelay                    bool          `mapstructure:"random-delay"`
	RandomDelayUnit                string        `mapstructure:"random-delay-unit"`
	RandomDelayMin                 int           `mapstructure:"random-delay-min"`
	RandomDelayMax                 int           `mapstructure:"random-delay-max"`
	RandomError                    bool          `mapstructure:"random-error"`
	Unhealthy                      bool          `mapstructure:"unhealthy"`
	Unready                        bool          `mapstructure:"unready"`
	JWTSecret                      string        `mapstructure:"jwt-secret"`
	CacheServer                    string        `mapstructure:"cache-server"`
	GrpcServiceEndpoint            string        `mapstructure:"grpc-service-endpoint"`
	StripeEndpointSigningSecretKey string        `mapstructure:"stripe-endpoint-signing-key"`
	Environment                    string        `mapstructure:"service-environment"`
	RpcTimeout                     time.Duration `mapstructure:"rpc-timeout"`
	StripeApiKey                   string        `mapstructure:"stripe-api-key"`
	PlaidClientID                  string        `mapstructure:"plaid-client-id"`
	PlaidSecretKey                 string        `mapstructure:"plaid-secret-key"`
	PlaidEnv                       string        `mapstructure:"plaid-env"`
	PlaidOauthDomain               string        `mapstructure:"plaid-oauth-domain"`
	PlaidWebhooksEnabled           bool          `mapstructure:"plaid-webhooks-enabled"`
	PlaidWebhookOauthDomain        string        `mapstructure:"plaid-webhook-oauth-domain"`
	AwsAccessKeyID                 string        `mapstructure:"aws-access-key-id"`
	AwsRegion                      string        `mapstructure:"aws-region"`
	AwsSecretAccessKey             string        `mapstructure:"aws-secret-access-key"`
	AwsKmsKeyID                    string        `mapstructure:"aws-kms-key-id"`
	MaxPlaidLinks                  int           `mapstructure:"max-plaid-links"`
	BillingEnabled                 bool          `mapstructure:"stripe-enabled"`
}

type Server struct {
	router                      *mux.Router
	logger                      *zap.Logger
	config                      *Config
	pool                        *redis.Pool
	handler                     http.Handler
	instrumentation             *instrumentation.Client
	plaidWrapper                *plaidhandler.PlaidWrapper
	inMemoryWebhookVerification plaidhandler.WebhookVerification
	conn                        *database.Db
	plaidClient                 *plaid.APIClient
	grpcGw                      *runtime.ServeMux
	kms                         secrets.KeyManagement
	taskprocessor               *taskprocessor.TaskProcessor
}

func NewServer(config *Config, logger *zap.Logger, telemetry *instrumentation.Client, db *database.Db, router *mux.Router, plaidClient *plaidhandler.PlaidWrapper, kmsClient secrets.KeyManagement, tp *taskprocessor.TaskProcessor) (*Server, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	gw := runtime.NewServeMux()
	ctx := context.Background()
	if err := proto.RegisterFinancialServiceHandlerFromEndpoint(ctx, gw, config.GrpcServiceEndpoint, opts); err != nil {
		return nil, err
	}

	var muxRouter *mux.Router
	if router != nil {
		muxRouter = router
	} else {
		muxRouter = mux.NewRouter()
	}

	srv := &Server{
		router:                      muxRouter,
		logger:                      logger,
		config:                      config,
		instrumentation:             telemetry,
		conn:                        db,
		grpcGw:                      gw,
		plaidWrapper:                plaidClient,
		inMemoryWebhookVerification: plaidhandler.NewInMemoryWebhookVerification(5*time.Minute, logger, plaidClient),
		kms:                         kmsClient,
		taskprocessor:               tp,
	}

	return srv, nil
}

// RegisterHandlers registers all the handlers for the server
func (s *Server) RegisterHandlers() {
	s.router.Handle("/grpc-gateway", s.grpcGw)
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	s.router.HandleFunc("/", s.indexHandler).HeadersRegexp("User-Agent", "^Mozilla.*").Methods("GET")
	s.router.HandleFunc("/", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/version", s.versionHandler).Methods("GET")
	s.router.HandleFunc("/echo", s.echoHandler).Methods("POST")
	s.router.HandleFunc("/env", s.envHandler).Methods("GET", "POST")
	s.router.HandleFunc("/headers", s.echoHeadersHandler).Methods("GET", "POST")
	s.router.HandleFunc("/delay/{wait:[0-9]+}", s.delayHandler).Methods("GET").Name("delay")
	s.router.HandleFunc("/healthz", s.healthzHandler).Methods("GET")
	s.router.HandleFunc("/readyz", s.readyzHandler).Methods("GET")
	s.router.HandleFunc("/readyz/enable", s.enableReadyHandler).Methods("POST")
	s.router.HandleFunc("/readyz/disable", s.disableReadyHandler).Methods("POST")
	s.router.HandleFunc("/panic", s.panicHandler).Methods("GET")
	s.router.HandleFunc("/status/{code:[0-9]+}", s.statusHandler).Methods("GET", "POST", "PUT").Name("status")
	s.router.HandleFunc("/store", s.storeWriteHandler).Methods("POST", "PUT")
	s.router.HandleFunc("/store/{hash}", s.storeReadHandler).Methods("GET").Name("store")
	s.router.HandleFunc("/cache/{key}", s.cacheWriteHandler).Methods("POST", "PUT")
	s.router.HandleFunc("/cache/{key}", s.cacheDeleteHandler).Methods("DELETE")
	s.router.HandleFunc("/cache/{key}", s.cacheReadHandler).Methods("GET").Name("cache")
	s.router.HandleFunc("/configs", s.configReadHandler).Methods("GET")
	s.router.HandleFunc("/token", s.tokenGenerateHandler).Methods("POST")
	s.router.HandleFunc("/token/validate", s.tokenValidateHandler).Methods("GET")
	s.router.HandleFunc("/api/info", s.infoHandler).Methods("GET")
	s.router.HandleFunc("/api/echo", s.echoHandler).Methods("POST")
	s.router.HandleFunc("/ws/echo", s.echoWsHandler)
	s.router.HandleFunc("/chunked", s.chunkedHandler)
	s.router.HandleFunc("/chunked/{wait:[0-9]+}", s.chunkedHandler)
	s.router.HandleFunc("/stripe/webhook", s.handleStripeWebhook).Methods("POST")
	s.router.HandleFunc("/plaid/webhook", s.handlerPlaidWebhook).Methods("POST")
	s.router.HandleFunc("/stripe/checkout", s.handleCheckoutSession).Methods("POST")

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	s.router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		doc, err := swag.ReadDoc()
		if err != nil {
			s.logger.Error("swagger error", zap.Error(err), zap.String("path", "/swagger.json"))
		}
		w.Write([]byte(doc))
	})
}

func (s *Server) RegisterMiddlewares() {
	prom := NewPrometheusMiddleware()
	s.router.Use(prom.Handler)
	httpLogger := NewLoggingMiddleware(s.logger)
	s.router.Use(httpLogger.Handler)
	s.router.Use(versionMiddleware)
	if s.config.RandomDelay {
		randomDelayer := NewRandomDelayMiddleware(s.config.RandomDelayMin, s.config.RandomDelayMax, s.config.RandomDelayUnit)
		s.router.Use(randomDelayer.Handler)
	}
	if s.config.RandomError {
		s.router.Use(randomErrorMiddleware)
	}
}

func (s *Server) ListenAndServe() (*http.Server, *http.Server) {
	go s.startMetricsServer()

	s.RegisterHandlers()
	s.RegisterMiddlewares()

	if s.config.H2C {
		s.handler = h2c.NewHandler(s.router, &http2.Server{})
	} else {
		s.handler = s.router
	}

	s.printRoutes()

	// load configs in memory and start watching for changes in the config dir
	if stat, err := os.Stat(s.config.ConfigPath); err == nil && stat.IsDir() {
		var err error
		watcher, err = fscache.NewWatch(s.config.ConfigPath)
		if err != nil {
			s.logger.Error("config watch error", zap.Error(err), zap.String("path", s.config.ConfigPath))
		} else {
			watcher.Watch()
		}
	}

	// start redis connection pool
	ticker := time.NewTicker(30 * time.Second)
	s.startCachePool(ticker)

	// create the http server
	srv := s.StartServer()

	// create the secure server
	secureSrv := s.startSecureServer()

	// signal Kubernetes the server is ready to receive traffic
	if !s.config.Unhealthy {
		atomic.StoreInt32(&healthy, 1)
	}
	if !s.config.Unready {
		atomic.StoreInt32(&ready, 1)
	}

	return srv, secureSrv
}

func (s *Server) Shutdown(ctx context.Context, srv *http.Server, secureSrv *http.Server) {
	// all calls to /healthz and /readyz will fail from now on
	atomic.StoreInt32(&healthy, 0)
	atomic.StoreInt32(&ready, 0)

	// close cache pool
	if s.pool != nil {
		_ = s.pool.Close()
	}

	s.logger.Info("Shutting down HTTP/HTTPS server", zap.Duration("timeout", s.config.HttpServerShutdownTimeout))

	// wait for Kubernetes readiness probe to remove this instance from the load balancer
	// the readiness check interval must be lower than the timeout
	if viper.GetString("level") != "debug" {
		time.Sleep(3 * time.Second)
	}

	// determine if the http server was started
	if srv != nil {
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Warn("HTTP server graceful shutdown failed", zap.Error(err))
		}
	}

	// determine if the secure server was started
	if secureSrv != nil {
		if err := secureSrv.Shutdown(ctx); err != nil {
			s.logger.Warn("HTTPS server graceful shutdown failed", zap.Error(err))
		}
	}
}

func (s *Server) StartServer() *http.Server {

	// determine if the port is specified
	if s.config.Port == "0" {
		// move on immediately
		return nil
	}

	s.logger.Info("Starting HTTP Server.", zap.String("addr", s.config.Host+":"+s.config.Port))

	srv := &http.Server{
		Addr:         s.config.Host + ":" + s.config.Port,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  2 * s.config.HttpServerTimeout,
		Handler:      s.handler,
	}

	// start the server in the background
	go func() {
		s.logger.Info("Starting HTTP Server.", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatal("HTTP server crashed", zap.Error(err))
		}
	}()

	// return the server and routine
	return srv
}

func (s *Server) startSecureServer() *http.Server {

	// determine if the port is specified
	if s.config.SecurePort == "0" {

		// move on immediately
		return nil
	}

	srv := &http.Server{
		Addr:         s.config.Host + ":" + s.config.SecurePort,
		WriteTimeout: s.config.HttpServerTimeout,
		ReadTimeout:  s.config.HttpServerTimeout,
		IdleTimeout:  2 * s.config.HttpServerTimeout,
		Handler:      s.handler,
	}

	cert := path.Join(s.config.CertPath, "tls.crt")
	key := path.Join(s.config.CertPath, "tls.key")

	// start the server in the background
	go func() {
		s.logger.Info("Starting HTTPS Server.", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServeTLS(cert, key); err != http.ErrServerClosed {
			s.logger.Fatal("HTTPS server crashed", zap.Error(err))
		}
	}()

	// return the server
	return srv
}

func (s *Server) startMetricsServer() {
	if s.config.PortMetrics > 0 {
		mux := http.DefaultServeMux
		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%v", s.config.PortMetrics),
			Handler: mux,
		}

		srv.ListenAndServe()
	}
}

func (s *Server) printRoutes() {
	s.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
}

type ArrayResponse []string
type MapResponse map[string]string
