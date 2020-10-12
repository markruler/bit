package cmd

import "github.com/spf13/cobra"

func AllGitSubCommands() []*cobra.Command {
	return []*cobra.Command{
		{Use: "add", Short: "Add file contents to the index"},
		{Use: "am", Short: "Apply a series of patches from a mailbox"},
		{Use: "archive", Short: "Create an archive of files from a named tree"},
		{Use: "bisect", Short: "Use binary search to find the commit that introduced a bug"},
		{Use: "branch", Short: "List, create, or delete branches"},
		{Use: "bundle", Short: "Move objects and refs by archive"},
		{Use: "checkout", Short: "Switch branches or restore working tree files"},
		{Use: "cherry-pick", Short: "Apply the changes introduced by some existing commits"},
		{Use: "citool", Short: "Graphical alternative to git-commit"},
		{Use: "clean", Short: "Remove untracked files from the working tree"},
		{Use: "clone", Short: "Clone a repository into a new directory"},
		{Use: "commit", Short: "Record changes to the repository"},
		{Use: "describe", Short: "Give an object a human readable name based on an available ref"},
		{Use: "diff", Short: "Show changes between commits, commit and working tree, etc"},
		{Use: "fetch", Short: "Download objects and refs from another repository"},
		{Use: "format-patch", Short: "Prepare patches for e-mail submission"},
		{Use: "gc", Short: "Cleanup unnecessary files and optimize the local repository"},
		{Use: "gitk", Short: "The Git repository browser"},
		{Use: "grep", Short: "Print lines matching a pattern"},
		{Use: "gui", Short: "A portable graphical interface to Git"},
		{Use: "init", Short: "Create an empty Git repository or reinitialize an existing one"},
		{Use: "log", Short: "Show commit logs"},
		{Use: "merge", Short: "Join two or more development histories together"},
		{Use: "mv", Short: "Move or rename a file, a directory, or a symlink"},
		{Use: "notes", Short: "Add or inspect object notes"},
		{Use: "pull", Short: "Fetch from and integrate with another repository or a local branch"},
		{Use: "push", Short: "Update remote refs along with associated objects"},
		{Use: "range-diff", Short: "Compare two commit ranges (e.g. two versions of a branch)"},
		{Use: "rebase", Short: "Reapply commits on top of another base tip"},
		{Use: "reset", Short: "Reset current HEAD to the specified state"},
		{Use: "restore", Short: "Restore working tree files"},
		{Use: "revert", Short: "Revert some existing commits"},
		{Use: "rm", Short: "Remove files from the working tree and from the index"},
		{Use: "shortlog", Short: "Summarize 'git log' output"},
		{Use: "show", Short: "Show various types of objects"},
		{Use: "stash", Short: "Stash the changes in a dirty working directory away"},
		{Use: "status", Short: "Show the working tree status"},
		{Use: "submodule", Short: "Initialize, update or inspect submodules"},
		{Use: "switch", Short: "Switch branches"},
		{Use: "tag", Short: "Create, list, delete or verify a tag object signed with GPG"},
		{Use: "worktree", Short: "Manage multiple working trees"},
		{Use: "config", Short: "Get and set repository or global options"},
		{Use: "fast-import", Short: "Backend for fast Git data importers"},
		{Use: "filter-branch", Short: "Rewrite branches"},
		{Use: "mergetool", Short: "Run merge conflict resolution tools to resolve merge conflicts"},
		{Use: "pack-refs", Short: "Pack heads and tags for efficient repository access"},
		{Use: "prune", Short: "Prune all unreachable objects from the object database"},
		{Use: "reflog", Short: "Manage reflog information"},
		{Use: "remote", Short: "Manage set of tracked repositories"},
		{Use: "repack", Short: "Pack unpacked objects in a repository"},
		{Use: "replace", Short: "Create, list, delete refs to replace objects"},
		{Use: "annotate", Short: "Annotate file lines with commit information"},
		{Use: "blame", Short: "Show what revision and author last modified each line of a file"},
		{Use: "count-objects", Short: "Count unpacked number of objects and their disk consumption"},
		{Use: "difftool", Short: "Show changes using common diff tools"},
		{Use: "fsck", Short: "Verifies the connectivity and validity of the objects in the database"},
		{Use: "gitweb", Short: "Git web interface (web frontend to Git repositories)"},
		{Use: "help", Short: "Display help information about Git"},
		{Use: "instaweb", Short: "Instantly browse your working repository in gitweb"},
		{Use: "merge-tree", Short: "Show three-way merge without touching index"},
		{Use: "rerere", Short: "Reuse recorded resolution of conflicted merges"},
		{Use: "show-branch", Short: "Show branches and their commits"},
		{Use: "verify-commit", Short: "Check the GPG signature of commits"},
		{Use: "verify-tag", Short: "Check the GPG signature of tags"},
		{Use: "whatchanged", Short: "Show logs with difference each commit introduces"},
		{Use: "archimport", Short: "Import a GNU Arch repository into Git"},
		{Use: "cvsexportcommit", Short: "Export a single commit to a CVS checkout"},
		{Use: "cvsimport", Short: "Salvage your data out of another SCM people love to hate"},
		{Use: "cvsserver", Short: "A CVS server emulator for Git"},
		{Use: "imap-send", Short: "Send a collection of patches from stdin to an IMAP folder"},
		{Use: "p4", Short: "Import from and submit to Perforce repositories"},
		{Use: "fast-export", Short: "Git data exporter"},
	}
}
