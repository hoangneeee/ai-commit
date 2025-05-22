package main

import (
	"fmt"
	"strings"
)

// buildPrompt xây dựng prompt phù hợp cho việc tạo commit message
func buildPrompt(diff string) string {
	return `Tôi sẽ cung cấp cho bạn một git diff. Hãy tạo một commit message ngắn gọn, rõ ràng và có ý nghĩa mô tả những thay đổi trong diff này.

Yêu cầu:
- Viết bằng tiếng Anh
- Giới hạn trong 1 dòng, tối đa 72 ký tự
- Sử dụng dạng thức: "<type>(<scope>): <subject>"
- Các loại type thông dụng: feat, fix, docs, style, refactor, test, chore

Ví dụ:
- feat(auth): add login with Google
- fix(api): handle null pointer in user data
- docs(readme): update installation instructions

Git diff:
` + diff
}

// processMessage xử lý message trả về từ API
func processMessage(message string) (string, error) {
	message = strings.TrimSpace(message)

	// Xóa các ký tự đặc biệt nếu có
	message = strings.Trim(message, "`")
	message = strings.TrimSpace(message)

	// Giới hạn độ dài tối đa
	if len(message) > 72 {
		message = message[:69] + "..."
	}

	if message == "" {
		return "", fmt.Errorf("generated message is empty")
	}

	return message, nil
}
