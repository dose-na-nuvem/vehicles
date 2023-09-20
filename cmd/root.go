package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/dose-na-nuvem/vehicles/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Configura o nome padrão do arquivo de configuração.
var (
	configFile     string
	cfg            = config.New()
)

// rootCmd representa o comando base quando chamado sem nenhum subcomando.
var rootCmd = &cobra.Command{
	Use: "vehicles",

	Short: "Controle de veículos e tags",

	Long: `Permite gerenciar os veículos e as tags de customers:
Será possível adicionar carros com multiplas tags e relacionar os carros com os donos atuais.`,
}

// Execute adiciona todos os comandos filhos ao comando raiz e configura as bandeiras apropriadamente.
// Isso é chamado por main.main(). Isso só precisa acontecer uma vez para o rootCmd.
func Execute() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml",
		"Define o arquivo de configuração a utilizar.")
	
	startCmd.Flags().StringVar(&cfg.Server.TLS.CertFile, "server.tls.certfile", "", "caminho do certificado.")

	startCmd.Flags().StringVar(&cfg.Server.TLS.CertKeyFile, "server.tls.certkeyfile", "",
		"caminho da chave privada do certificado.")

	startCmd.Flags().BoolVar(&cfg.Server.TLS.Insecure, "server.tls.insecure", false, "Força o modo inseguro.")

	startCmd.Flags().StringVar(&cfg.Server.GRPC.Endpoint, "server.grpc.endpoint", "127.0.0.1:56533",
		"Endereço gRPC onde o serviço vai servir requisições.")

	// Associa o Viper as flags
	if err := viper.BindPFlags(startCmd.Flags()); err != nil {
		cfg.Logger.Error("falha ao ligar as flags", zap.Error(err))
	}
	
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	// Configura o nome padrão do arquivo de configuração, sem a extensão.
	viper.SetConfigFile(configFile)

	// Tenta ler o arquivo de configuração, ignorando erros caso o mesmo não seja encontrado
	// retorna um erro se não conseguirmos analisar o arquivo de configuração encontrado.
	if err := viper.ReadInConfig(); err != nil {
		// Não há problems se não existir um arquivo de configuração.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			cfg.Logger.Error("arquivo não encontrado",
				zap.String("arquivo", configFile),
				zap.Error(err),
			)

			return
		}

		cfg.Logger.Error("falha na leitura do arquivo de configuração", zap.Error(err))
	} else {
		cfg.Logger.Info("arquivo de configuração lido", zap.String("config", configFile))
	}

	// Converter o estado interno do Viper em nosso objeto de configuração
	if err := viper.Unmarshal(&cfg); err != nil {
		cfg.Logger.Error("falhou ao converter o arquivo de configuração", zap.Error(err))

		return
	}
}
