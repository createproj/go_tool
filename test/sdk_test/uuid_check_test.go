package sdk_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/google/uuid"
)

// TestCreateShortUUID 测试生成最短的UUID（去除横杠）
func TestCreateShortUUID(t *testing.T) {
	// 调用被测函数
	shortUUID := CreateShortUUID()

	// 打印生成结果
	t.Logf("single UUID: %s", shortUUID)

	// 断言：非空
	if shortUUID == "" {
		t.Fatal("CreateShortUUID() returned an empty string")
	}

	// 断言：不含横杠
	if contains(shortUUID, "-") {
		t.Errorf("CreateShortUUID() = %q, should not contain '-'", shortUUID)
	}

	// 断言：长度应为 32（UUID 去除 4 个横杠后的长度）
	if len(shortUUID) != 32 {
		t.Errorf("CreateShortUUID() length = %d, want 32", len(shortUUID))
	}

	// 断言：是合法的十六进制 UUID 格式
	if matched, _ := regexp.MatchString(`^[0-9a-f]{32}$`, shortUUID); !matched {
		t.Errorf("CreateShortUUID() = %q, is not a valid 32-char hex UUID", shortUUID)
	}

	// 断言：多次生成不重复，并逐个打印
	const n = 10
	seen := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		id := CreateShortUUID()
		t.Logf("UUID #%d: %s", i+1, id) // 打印每一个

		if _, exists := seen[id]; exists {
			t.Fatalf("CreateShortUUID() generated duplicate UUID: %q", id)
		}
		seen[id] = struct{}{}
	}
}

// contains 简单判断字符串是否包含子串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || indexOf(s, substr) >= 0)
}

func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// CreateShortUUID 生成不带横杠的最短 UUID（32 位十六进制字符串）
func CreateShortUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
