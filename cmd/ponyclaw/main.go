// PonyClaw - Ultra-lightweight personal AI agent
// Forked from PicoClaw: https://github.com/sipeed/picoclaw
// Originally inspired by nanobot: https://github.com/HKUDS/nanobot
// License: MIT
//
// Copyright (c) 2026 PicoClaw contributors
// Copyright (c) 2026 vanyongqi

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/agent"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/auth"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/cron"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/gateway"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/migrate"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/onboard"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/skills"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/status"
	"github.com/lansely/ponyclaw/cmd/ponyclaw/internal/version"
)

func NewPonyclawCommand() *cobra.Command {
	short := fmt.Sprintf("%s ponyclaw - Personal AI Assistant v%s\n\n", internal.Logo, internal.GetVersion())

	cmd := &cobra.Command{
		Use:     "ponyclaw",
		Short:   short,
		Example: "ponyclaw list",
	}

	cmd.AddCommand(
		onboard.NewOnboardCommand(),
		agent.NewAgentCommand(),
		auth.NewAuthCommand(),
		gateway.NewGatewayCommand(),
		status.NewStatusCommand(),
		cron.NewCronCommand(),
		migrate.NewMigrateCommand(),
		skills.NewSkillsCommand(),
		version.NewVersionCommand(),
	)

	return cmd
}

func main() {
	cmd := NewPonyclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
