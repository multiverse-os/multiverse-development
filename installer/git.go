package install

func SetupGitRepository() error {
	terminal.Run("git --global user.email \"you@example.com\"")
	terminal.Run("git --global user.name \"Your Name\"")

	// TODO: Pull down repository
}
