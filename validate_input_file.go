package form

import (
	"fmt"
	"mime/multipart"
	"path"
	"strings"
)

func (va validateValue) fileInputFile(value *multipart.FileHeader) {
	if value == nil {
		return
	}

	// Mandatory

	manErr := ""
	mandatory := false

	va.fieldsFns.Call("mandatory", map[string]interface{}{
		"mandatory": &mandatory,
		"err":       &manErr,
	})

	if mandatory && value == nil {
		if manErr == "" {
			manErr = va.form.T("ErrFileRequired")
		}
		*(va.err) = fmt.Errorf(manErr)
		return
	}

	// Size and Mime

	size := int64(-1)
	sizeErr := ""
	var mimes []string
	mimesErr := ""

	va.fieldsFns.Call("file", map[string]interface{}{
		"size":      &size,
		"sizeErr":   &sizeErr,
		"accept":    &mimes,
		"acceptErr": &mimesErr,
	})

	if size <= -1 {
		goto mime_check
	}

	if fileSize(value) > size {
		if sizeErr == "" {
			sizeErr = va.form.T("ErrFileSize", map[string]interface{}{
				"Size": size,
			})
		}
		*(va.err) = fmt.Errorf(sizeErr)
		return
	}

mime_check:

	if mimes == nil {
		return
	}

	fileType := fileContentType(value)

	for _, mime := range mimes {
		matched, err := path.Match(mime, fileType)
		if err != nil {
			continue
		}
		if matched {
			return
		}
	}

	if mimesErr == "" {
		mimeStr := ""
		if len(mimes) == 1 {
			mimeStr = mimes[0]
		} else {
			mimeStr = strings.Join(mimes[:len(mimes)-1], ", ") + " & " + mimes[len(mimes)-1]
		}
		mimesErr = va.form.T("ErrMimeCheck", len(mimes), map[string]interface{}{
			"Mimes": mimeStr,
		})
	}

	*(va.err) = fmt.Errorf(mimesErr)
}
