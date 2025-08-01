package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*", numStarsPerLine)
    return border + "\n" + welcomeMsg + "\n" + border
}


// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    lines := strings.Split(oldMsg, "\n")
    if len(lines) < 2 {
        return ""
    }
    // Pick the middle line (assumed to contain the actual message)
    middle := lines[1]
    middle = strings.Trim(middle, "*")      // remove leading/trailing asterisks
    middle = strings.TrimSpace(middle)      // remove leading/trailing spaces
    return middle
}


