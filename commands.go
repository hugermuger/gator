package main

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandmap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	err := c.commandmap[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandmap[name] = f
}
