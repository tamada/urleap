package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GenerateCompletion(flag *pflag.FlagSet) error {
	command := &cobra.Command{
		Use: "completions",
	}
	command.Flags().AddFlagSet(flag)
	os.Mkdir("completions/", 0755)
	os.Mkdir("completions/bash", 0755)
	os.Mkdir("completions/zsh", 0755)
	os.Mkdir("completions/fish", 0755)
	os.Mkdir("completions/powershell", 0755)
	command.GenBashCompletionFileV2("completions/bash/urleap", true)
	command.GenZshCompletionFile("completions/zsh/urleap")
	command.GenFishCompletionFile("completions/fish/urleap", true)
	command.GenPowerShellCompletionFile("completions/ps1/urleap")
	return nil
}
