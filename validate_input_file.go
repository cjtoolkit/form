package form

import (
	"fmt"
	"mime/multipart"
	"path"
	"strings"
)

func (va validate) fileInputFile() {
	file := va.value.Interface().(*multipart.FileHeader)
	_fileSize := int64(0)
	_fileContentType := ""
	mandatory, ok := va.getBool("Mandatory")
	if file == nil {
		if ok {
			if mandatory {
				manErr, ok := va.getStr("MandatoryErr")
				if ok {
					va.setErr(FormError(manErr))
				} else {
					va.setErr(FormError(va.i18n.Key(ErrFileRequired)))
				}
				return
			}
		}
	}
	_fileSize = fileSize(file)
	_fileContentType = fileContentType(file)

	sizeErr := ""
	size, ok := va.getInt("Size")
	if !ok {
		goto mime_check
	}

	sizeErr, ok = va.getStr("SizeErr")
	if !ok {
		sizeErr = fmt.Sprintf(va.i18n.Key(ErrFileSize), size)
	}

	if _fileSize > size {
		va.setErr(FormError(sizeErr))
		return
	}

mime_check:

	mimes, ok := va.getStrs("Accept")
	if !ok {
		return
	}

	for _, mime := range mimes {
		matched, err := path.Match(mime, _fileContentType)
		if err != nil {
			continue
		}
		if matched {
			return
		}
	}

	mimeStr := ""
	if len(mimes) == 1 {
		mimeStr = mimes[0]
	} else {
		mimeStr = strings.Join(mimes[:len(mimes)-1], ", ") + " & " + mimes[len(mimes)-1]
	}

	va.setErr(FormError(va.i18n.Key(ErrMimeCheck) + mimeStr))
}
