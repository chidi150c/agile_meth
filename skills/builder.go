package skills

import "strings"

func ExtractCommand(message string) string {
    // Find the index of "command:" in the message
    startIndex := strings.Index(message, "command:")
    if startIndex == -1 {
        return "" // "command:" not found in the message
    }

    // Extract the substring starting from "command:"
    commandSubstring := message[startIndex+len("command:"):]
    
    // Trim any leading or trailing whitespace
    command := strings.TrimSpace(commandSubstring)

    return command
}
