package env

const (
	AppEnv        = "APP_ENV"
	ServerName    = "SERVER_NAME"
	ServerPort    = "SERVER_PORT"
	ServerTimeout = "SERVER_TIMEOUT"
)

type Env struct {
	AppEnv        string
	ServerName    string
	ServerPort    string
	ServerTimeout string
}

func NewDataEnv() *Env {
	return &Env{
		AppEnv:        GetEnv(AppEnv),
		ServerName:    GetSecureEnv(ServerName),
		ServerPort:    GetSecureEnv(ServerPort),
		ServerTimeout: GetSecureEnv(ServerTimeout),
	}
}
