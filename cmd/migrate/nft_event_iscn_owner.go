package migrate

import (
	"github.com/spf13/cobra"

	"github.com/likecoin/likecoin-chain-tx-indexer/db"
	"github.com/likecoin/likecoin-chain-tx-indexer/db/schema/parallel"
	"github.com/likecoin/likecoin-chain-tx-indexer/logger"
)

var MigrationNftEventIscnOwnerCommand = &cobra.Command{
	Use:   "nft-event-iscn-owner",
	Short: "Setup iscn_owner_at_the_time column in nft_event table",
	RunE: func(cmd *cobra.Command, args []string) error {
		batchSize, err := cmd.Flags().GetUint64(CmdBatchSize)
		if err != nil {
			return err
		}
		pool, err := db.GetConnPoolFromCmdArgs(cmd)
		if err != nil {
			logger.L.Panicw("Cannot initialize database connection pool", "error", err)
		}
		conn, err := db.AcquireFromPool(pool)
		if err != nil {
			logger.L.Panicw("Cannot acquire connection from database connection pool", "error", err)
		}
		defer conn.Release()
		return parallel.MigrateNftEventIscnOwner(conn, batchSize)
	},
}

func init() {
	MigrationNftEventIscnOwnerCommand.PersistentFlags().Uint64(
		CmdBatchSize,
		1000,
		"number of ids in nft_event table to scan each time",
	)
}
