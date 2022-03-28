package main

import (
	"fmt"
	"time"
	"context"
	"path/filepath"
	hostagentclient "github.com/lima-vm/lima/pkg/hostagent/api/client"
	"github.com/lima-vm/lima/pkg/store"
	"github.com/lima-vm/lima/pkg/store/filenames"
	"github.com/spf13/cobra"
)

func newRemountCommand() *cobra.Command {
	var remountCmd = &cobra.Command{
		Use:               "remount INSTANCE",
		Short:             "Re-mount all mounts on instance",
		Args:              cobra.MaximumNArgs(1),
		RunE:              remountAction,
		ValidArgsFunction: remountBashComplete,
	}

	return remountCmd
}

func remountAction(cmd *cobra.Command, args []string) error {
	instName := DefaultInstanceName
	if len(args) > 0 {
		instName = args[0]
	}

	inst, err := store.Inspect(instName)
	if err != nil {
		return err
	}
	haSock := filepath.Join(inst.Dir, filenames.HostAgentSock)
	haClient, err := hostagentclient.NewHostAgentClient(haSock)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	err := haClient.ReloadMounts(ctx)
	if err != nil {
		return err
	}
	return err
}


func remountBashComplete(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return bashCompleteInstanceNames(cmd)
}
