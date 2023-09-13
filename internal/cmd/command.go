package cmd

import (
	"fmt"
	"os"
)

type Commands map[string]command

type command struct {
	use   string
	short string
	run   func(string, Config)
}

func newCommands() Commands {
	c := Commands{}
	c.addHelpCmd()
	c.addMapCmd()
	c.addMapbCmd()
	c.addExploreCmd()
	c.addCatchCmd()
	c.addInspectCmd()
	c.addPokedexCmd()
	c.addExitCmd()
	return c
}

func (c Commands) Execute(cmd string, arg string, config Config) {
	if _, ok := c[cmd]; ok {
		c[cmd].run(arg, config)
	}
}

func (c Commands) addExitCmd() {
	c["exit"] = command{
		use:   "exit",
		short: "Exit the Pokedex",
		run:   c.runExitCmd,
	}
}

func (c Commands) addHelpCmd() {
	c["help"] = command{
		use:   "help",
		short: "Displays a help message",
		run:   c.runHelpCmd,
	}
}

func (c Commands) runExitCmd(arg string, config Config) {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func (c Commands) runHelpCmd(arg string, config Config) {
	fmt.Println("Available Commands:")
	for _, cmd := range c {
		fmt.Printf("  %s	%s\n", cmd.use, cmd.short)
	}
}
