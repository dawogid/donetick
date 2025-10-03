package utils

// Centralized date/time layout definitions for the application.
// Date (no time): DD/MM/YYYY
// DateTime: DD/MM/YYYY 24-hour clock (UTC unless otherwise documented)
const (
    DateLayout         = "02/01/2006"
    DateTimeLayout     = "02/01/2006 15:04:05"
    FileDateTimeLayout = "02-01-2006_15-04-05" // filesystem safe
)
