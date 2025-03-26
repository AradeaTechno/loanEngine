package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var defaultTagMessages = map[string]string{
	"required": "{field} is required.",
	"email":    "{field} must contain a valid email address.",
	"number":   "{field} must be in numeric format.",
	"min":      "{field} length must be at least {param} characters long.",
	"max":      "{field} length must be no more than {param} characters long.",
}

func init() {
	// Initialize the global validate instance
	validate = validator.New()
	validate.RegisterValidation("role_validator", CustomRoleValidator)
}

func CustomRoleValidator(fl validator.FieldLevel) bool {
	validRoles := []string{"staff", "field_officer"}
	value := fl.Field().String()
	for _, role := range validRoles {
		if value == role {
			return true
		}
	}
	return false
}

func GetValidator() *validator.Validate {
	return validate
}

func GetCustomMessage(field, tag, param string) string {
	if msgTemplate, exists := defaultTagMessages[tag]; exists {
		msg := strings.ReplaceAll(msgTemplate, "{field}", field)
		msg = strings.ReplaceAll(msg, "{param}", param)
		return msg
	}
	return fmt.Sprintf("%s is invalid.", field)
}

func ValidateStruct(s any) []string {
	validate := GetValidator() // Use the initialized validator instance
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	var messages []string
	structType := reflect.TypeOf(s)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()
		// FIND THE json TAG
		if structField, found := structType.FieldByName(field); found {
			jsonTag := structField.Tag.Get("json")
			if jsonTag != "" {
				field = jsonTag
			}
		}
		message := GetCustomMessage(field, tag, param)
		messages = append(messages, message)
	}
	return messages
}

func ValidatePassword(password string) (bool, error) {
	// CHECK LENGTH
	if len(password) < 8 {
		return false, errors.New("password must be at least 8 characters long")
	}
	// DEFINED REGEX PATTERN
	var (
		lowerCase    = regexp.MustCompile(`[a-z]`)
		upperCase    = regexp.MustCompile(`[A-Z]`)
		number       = regexp.MustCompile(`[0-9]`)
		specialChars = regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};':",.<>\/?\\|` + "`" + `~]`)
	)

	if !lowerCase.MatchString(password) {
		return false, errors.New("password must contain at least one lowercase letter")
	}
	if !upperCase.MatchString(password) {
		return false, errors.New("password must contain at least one uppercase letter")
	}
	if !number.MatchString(password) {
		return false, errors.New("password must contain at least one number")
	}
	if !specialChars.MatchString(password) {
		return false, errors.New("password must contain at least one special character")
	}

	return true, nil
}

func ParseFormValueToInt64(value string, fieldName string) (int64, error) {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %s", fieldName, value)
	}
	return parsedValue, nil
}

// CHECK IS ALLOWED IMAGE
func IsAllowedImage(file multipart.File) error {
	// Read a small portion of the file header to determine the MIME type
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return err
	}
	// Reset file pointer to beginning after reading
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	// Detect MIME type from file header
	mimeType := http.DetectContentType(buffer)
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return fmt.Errorf("unsupported file type: %s. Only JPG, JPEG, or PNG are allowed", mimeType)
	}

	return nil
}

// Helper function to check if the file is a CSV
func IsAllowedFile(fileType string, file *multipart.FileHeader) bool {
	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	switch fileType {
	case "pdf":
		return fileExtension == ".pdf"
	default:
		return false
	}
}
