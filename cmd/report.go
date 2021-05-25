package cmd

import (
	"context"
	"github.com/turbot/steampipe/steampipeconfig/modconfig"
	"gopkg.in/olahol/melody.v1"

	reportserver2 "github.com/turbot/steampipe/report/reportserver"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/logging"
	"github.com/turbot/steampipe/cmdconfig"
	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/db"
	"github.com/turbot/steampipe/utils"
	"github.com/turbot/steampipe/workspace"
)

// ReportCmd :: represents the report command
func ReportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "report [report]",
		TraverseChildren: true,
		Args:             cobra.ArbitraryArgs,
		Run:              runReportCmd,
		Short:            "Run a report",
		Long:             `Run a report...TODO better description!`,
	}

	cmdconfig.
		OnCmd(cmd)
	return cmd
}

func runReportCmd(cmd *cobra.Command, args []string) {
	logging.LogTime("runReportCmd start")
	cmdconfig.Viper().Set(constants.ConfigKeyShowInteractiveOutput, false)

	defer func() {
		logging.LogTime("runReportCmd end")
		if r := recover(); r != nil {
			utils.ShowError(helpers.ToError(r))
		}
	}()

	_, cancel := context.WithCancel(context.Background())
	startCancelHandler(cancel)

	// start db if necessary
	err := db.EnsureDbAndStartService(db.InvokerReport)
	utils.FailOnErrorWithMessage(err, "failed to start service")
	defer db.Shutdown(nil, db.InvokerReport)

	// load the workspace
	workspace, err := workspace.Load(viper.GetString(constants.ArgWorkspace))
	utils.FailOnErrorWithMessage(err, "failed to load workspace")
	defer workspace.Close()

	// TODO remove this
	workspace.ReportMap = make(map[string]*modconfig.Report)
	workspace.ReportMap["simple"] = &modconfig.Report{ShortName: "simple", Title: "Simple"}
	workspace.ReportMap["two-panel"] = &modconfig.Report{ShortName: "two-panel", Title: "Two Panel"}
	workspace.ReportMap["bar-chart"] = &modconfig.Report{ShortName: "bar-chart", Title: "Bar Chart"}

	webSocket := melody.New()
	server := reportserver2.Server{webSocket, workspace}

	// TODO add this in when Kai exposes it, mock for now
	// workspace.registerUpdateHandler(server.HandleWorkspaceUpdate)
	//go reportevents.GenerateReportEvents(mockReport, server.HandleWorkspaceUpdate)

	server.Start()
}