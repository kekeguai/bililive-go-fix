package update

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{name: "修复版构建元数据与上游正式版等价", v1: "v0.8.4+missevan.fix", v2: "v0.8.4", want: 0},
		{name: "正式版高于对应预发布版", v1: "v0.8.4-missevan-fix", v2: "v0.8.4", want: -1},
		{name: "较新的正式版本仍高于修复构建", v1: "v0.8.4+missevan.fix", v2: "v0.8.5", want: -1},
		{name: "较旧的正式版本低于修复构建", v1: "v0.8.4+missevan.fix", v2: "v0.8.3", want: 1},
		{name: "忽略构建元数据差异", v1: "0.8.4+local.1", v2: "v0.8.4+local.2", want: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CompareVersions(test.v1, test.v2)
			if err != nil {
				t.Fatalf("CompareVersions() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("CompareVersions(%q, %q) = %d, want %d", test.v1, test.v2, got, test.want)
			}
		})
	}
}

func TestCompareVersionsRejectsInvalidVersion(t *testing.T) {
	if _, err := CompareVersions("v0.8.4+missevan.fix", "not-a-version"); err == nil {
		t.Fatal("CompareVersions() error = nil, want an invalid-version error")
	}
}
