package notepad

func getSnapshot(content string) (snaptshot string) {
	if len(content) >= 100 {
		snaptshot = content[0:99] + "..."
	} else {
		snaptshot = content[0:45] + "..."
	}
	return
}
