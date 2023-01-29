package service

func NewCommands() map[string]Cli {
	commands := make(map[string]Cli)
	commands["ModifySecurityGroupRule"] = &ModifySecurityGroupRule{}

	return commands
}
