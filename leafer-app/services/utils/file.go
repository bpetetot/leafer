package utils

// IsHidden return true if the file name starts with a dot
func IsHidden(filename string) bool {
	return filename[0:1] == "."
}
