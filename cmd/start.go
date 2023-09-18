package cmd

import (
	"context"

	"github.com/dose-na-nuvem/vehicles/pkg/service"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia o microserviço",
	Long: `Permite gerenciar os veículos e as tags de customers:

Será possível adicionar carros com multiplas tags e relacionar os carros com clientes.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		vehicles := service.New(cfg)

		err := vehicles.Start(ctx)
		if err != nil {
			cfg.Logger.Error("falha ao iniciar o serviço", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
