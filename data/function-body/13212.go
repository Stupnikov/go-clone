{
	out.Normf("\n")
	out.Normf("   _      ______ __________   \n")
	out.Normf("  | | /| / / __ `/ ___/ __ \\  \n")
	out.Normf("  | |/ |/ / /_/ / /  / /_/ /  \n")
	out.Normf("  |__/|__/\\__,_/_/  / .___/   \n")
	out.Normf("                   /_/        ")
	out.Boldf("v%s\n", warp.Version)
	out.Normf("\n")
	out.Normf("  Secure terminal sharing with one simple command.\n")
	out.Normf("\n")
	out.Normf("Usage:\n")
	out.Boldf("  warp <command> [<args> ...]\n")
	out.Normf("\n")
	out.Normf("Commands:\n")
	out.Boldf("  help <command>\n")
	out.Normf("    Show help for a specific command.\n")
	out.Valuf("    warp help open\n")
	out.Normf("\n")
	out.Boldf("  open [<id>]\n")
	out.Normf("    Creates a new warp.\n")
	out.Valuf("    warp open\n")
	out.Normf("\n")
	out.Boldf("  connect <id>\n")
	out.Normf("    Connects to an existing warp.\n")
	out.Valuf("    warp connect goofy-dev\n")
	out.Normf("\n")
	out.Boldf("  state\n")
	out.Normf("    Displays the state of the current warp (in-warp only).\n")
	out.Valuf("    warp state\n")
	out.Normf("\n")
	out.Boldf("  authorize <username_or_token>\n")
	out.Normf("    Grants write access to a client (in-warp only).\n")
	out.Valuf("    warp authorize goofy\n")
	out.Normf("\n")
	out.Boldf("  revoke [<username_or_token>]\n")
	out.Normf("    Revokes write access to one or all clients (in-warp only).\n")
	out.Valuf("    warp revoke\n")
	out.Normf("\n")
}