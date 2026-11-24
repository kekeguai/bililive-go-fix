package build

import (
	"os"
	"strings"
	"testing"
)

func TestGetBuildFlagsPreservesSemverBuildMetadata(t *testing.T) {
	previousVersion, hadVersion := os.LookupEnv("APP_VERSION")
	previousTelemetry, hadTelemetry := os.LookupEnv("NO_TELEMETRY")
	t.Cleanup(func() {
		if hadVersion {
			_ = os.Setenv("APP_VERSION", previousVersion)
		} else {
			_ = os.Unsetenv("APP_VERSION")
		}
		if hadTelemetry {
			_ = os.Setenv("NO_TELEMETRY", previousTelemetry)
		} else {
			_ = os.Unsetenv("NO_TELEMETRY")
		}
	})

	_ = os.Setenv("APP_VERSION", "v0.8.4+missevan.fix")
	_ = os.Unsetenv("NO_TELEMETRY")

	flags := GetBuildFlags(false)
	if !strings.Contains(flags.LdFlags, ".AppVersion=v0.8.4+missevan.fix") {
		t.Fatalf("build flags did not preserve SemVer metadata: %s", flags.LdFlags)
	}
	if strings.Contains(flags.LdFlags, "&#43;") {
		t.Fatalf("build flags HTML-escaped '+' in SemVer metadata: %s", flags.LdFlags)
	}
}
