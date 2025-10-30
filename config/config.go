package config

// Aqui você NÃO declara DB novamente
var logger *Logger

func Init() error {
	// Inicializa o banco de dados
	InitDatabase()

	// Inicializa o logger (opcional)
	logger = GetLogger("sqlite")

	return nil
}

func GetLogger(p string) *Logger {
	if logger == nil {
		logger = NewLogger(p)
	}
	return logger
}
