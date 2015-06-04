package form

import (
	"fmt"
	"html"
	"mime/multipart"
	"strings"
	"time"
)

var es = html.EscapeString

func isStructPtr(v interface{}) bool {
	str := fmt.Sprintf("%T", v)
	if len(str) <= 0 || str[0] != '*' {
		return false
	}
	return true
}

func fileSize(fileHeader *multipart.FileHeader) int64 {
	size := int64(0)

	file, err := fileHeader.Open()
	if err != nil {
		return size
	}

	defer file.Close()

	size, err = file.Seek(0, 2)
	if err != nil {
		return 0
	} else {
		file.Seek(0, 0)
	}

	return size
}

func fileContentType(fileHeader *multipart.FileHeader) string {
	return fileHeader.Header.Get("Content-Type")
}

const (
	// All meets html5 specification.
	dateTimeFormat      = "2006-01-02T15:04Z07:00" //RFC3339
	dateTimeLocalFormat = "2006-01-02T15:04"
	dateFormat          = "2006-01-02"
	timeFormat          = "15:04"
	monthFormat         = "2006-01"
	weekFormat          = "%d-W%d" // yyyy-Www
)

// Get number of weeks in a year, can be either 52 or 53!
func WeekInAYear(year int, loc *time.Location) int {
	if loc == nil {
		loc = time.UTC
	}
	// Probably not the ideal way to do it, but it's works perfectly.
	_, weeks := time.Date(year+1, 0, 0, 0, 0, 0, 0, loc).ISOWeek()
	if weeks == 49 {
		return 53
	}
	return 52
}

// Get time with the starting day of that week in that year!
func StartingDayOfWeek(year, week int, loc *time.Location) time.Time {
	if loc == nil {
		loc = time.UTC
	}
	if week < 1 {
		week = 1
	}
	if week == 1 {
		return time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	}
	max := WeekInAYear(year, loc)
	if week > max {
		week = max
	}
	day := 1
	for {
		_, _week := time.Date(year, 1, day, 0, 0, 0, 0, loc).ISOWeek()
		if _week == 2 {
			break
		}
		day++
	}
	day = day + ((week - 2) * 7)
	return time.Date(year, 1, day, 0, 0, 0, 0, loc)
}

// Render Attribute into strings.
func RenderAttr(attr map[string]string) (str string) {
	if attr == nil {
		return
	}

	for key, value := range attr {
		value = strings.TrimSpace(value)
		if value == "" {
			str += fmt.Sprintf(`%s `, es(key))
		} else {
			str += fmt.Sprintf(`%s="%s" `, es(key), es(value))
		}
	}

	return
}
