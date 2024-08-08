package util

import (
	"fmt"
	"regexp"
	"strconv"
)

func extractSemanticVersionNumbers(version string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(version, 3)
	if len(matches) < 3 {
		panic(fmt.Sprintf("invalid semantic version '%s': version must start with three numeric components (MAJOR.MINOR.PATCH)", version))
	}
	result := make([]int, 3)
	for i := 0; i < 3; i++ {
		result[i], _ = strconv.Atoi(matches[i])
	}
	return result
}

// CompareSemanticVersions compares two version strings
// returns -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
func CompareSemanticVersions(v1, v2 string) int {
	v1Numbers := extractSemanticVersionNumbers(v1)
	v2Numbers := extractSemanticVersionNumbers(v2)

	for i := 0; i < 3; i++ {
		if v1Numbers[i] < v2Numbers[i] {
			return -1
		}
		if v1Numbers[i] > v2Numbers[i] {
			return 1
		}
	}
	return 0
}

func main() {
	tests := []struct {
		v1       string
		v2       string
		expected int // -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2
		comment  string
	}{
		{"1.2.3", "1.3.0", -1, "正常版本号比较"},
		{"1.2.3 beta", "1.3.0", -1, "带空格的beta版本号比较"},
		{"1.2.3 alpha", "1.3.0", -1, "带空格的alpha版本号比较"},
		{"1.2.3 rc", "1.3.0", -1, "带空格的rc版本号比较"},
		{"1.2.3", "1.2.3", 0, "相同版本号比较"},
		{"1.2.3", "1.2.4", -1, "patch版本号比较"},
		{"1.2.3", "1.3.0", -1, "minor版本号比较"},
		{"1.2.3", "2.0.0", -1, "major版本号比较"},
		{"1.2.3", "1.2.2", 1, "较大版本号比较较小版本号"},
		{"1.2.3", "1.2.3-beta", 1, "正式版本比beta版本大"},
		{"1.2.3-beta", "1.2.3-alpha", 1, "beta比alpha大"},
		{"1.2.3-beta", "1.2.3-rc", -1, "beta比rc小"},
		{"1.2.3-beta", "1.2.3", -1, "beta比正式版小"},
		{"1.2.3-beta", "1.2.4", -1, "beta比高版本号小"},
		{"1.2.3-beta", "1.3.0", -1, "beta比高minor版本小"},
		{"1.2.3-beta", "2.0.0", -1, "beta比高major版本小"},
		{"1.2.3-beta", "1.2.3-beta", 0, "相同beta版本比较"},
		{"1.2.3-beta", "1.2.3-beta.1", -1, "beta比beta.1小"},
		{"1.2.3-beta", "1.2.3-beta.2", -1, "beta比beta.2小"},
		{"1.2.3-beta", "1.2.3-beta.10", -1, "beta比beta.10小"},
		{"1.2.3-beta", "1.2.3-rc", -1, "beta比rc小"},
		{"1.2.3-beta", "1.2.3-rc.1", -1, "beta比rc.1小"},
		{"1.2.3-beta", "1.2.3-rc.10", -1, "beta比rc.10小"},
		{"1.2.3-beta", "1.2.3-rc.100", -1, "beta比rc.100小"},
		{"1.2.3-beta", "1.2.3-rc.1000", -1, "beta比rc.1000小"},
		{"1.2.3-beta", "1.2.3-rc.10000", -1, "beta比rc.10000小"},
	}

	for i, test := range tests {
		result := CompareSemanticVersions(test.v1, test.v2)
		if result != test.expected {
			fmt.Printf("❌ 测试 %d 失败: %s vs %s\n", i+1, test.v1, test.v2)
			fmt.Printf("   期望: %d, 实际: %d\n", test.expected, result)
			fmt.Printf("   说明: %s\n\n", test.comment)
		} else {
			fmt.Printf("✅ 测试 %d 通过: %s vs %s (%s)\n", i+1, test.v1, test.v2, test.comment)
		}
	}
}
