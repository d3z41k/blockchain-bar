package main

import (
	"context"
	"d3z41k/blockchain-bar/database"
	"d3z41k/blockchain-bar/node"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var migrateCmd = func() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrates the blockchain database according to new business rules.",
		Run: func(cmd *cobra.Command, args []string) {
			// make miner configurable via a CLI flag on boot-up
			miner, _ := cmd.Flags().GetString(flagMiner)
			ip, _ := cmd.Flags().GetString(flagIP)
			port, _ := cmd.Flags().GetUint64(flagPort)

			peer := node.NewPeerNode(
				"127.0.0.1",
				8080,
				true,
				database.NewAccount(miner),
				false,
			)

			n := node.New(getDataDirFromCmd(cmd), ip, port, database.NewAccount(miner), peer)

			n.AddPendingTX(database.NewTx("andrej", "andrej", 3, ""), peer)
			n.AddPendingTX(database.NewTx("andrej", "babayaga", 2000, ""), peer)
			n.AddPendingTX(database.NewTx("babayaga", "andrej", 1, ""), peer)
			n.AddPendingTX(database.NewTx("babayaga", "caesar", 1000, ""), peer)
			n.AddPendingTX(database.NewTx("babayaga", "andrej", 50, ""), peer)

			ctx, closeNode := context.WithTimeout(context.Background(), time.Minute*15)

			go func() {
				ticker := time.NewTicker(time.Second * 10)

				for {
					select {
					case <-ticker.C:
						if !n.LatestBlockHash().IsEmpty() {
							closeNode()
							return
						}
					}
				}
			}()

			err := n.Run(ctx)
			if err != nil {
				fmt.Println(err)
			}

		},
	}

	addDefaultRequiredFlags(migrateCmd)

	return migrateCmd
}
