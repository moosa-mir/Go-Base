package cmd

func Execute() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(payOutCmd)
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
