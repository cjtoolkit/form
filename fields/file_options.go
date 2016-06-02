package fields

type FileOption func(*File)

func FileSuffix(suffix *string) FileOption {
	return func(f *File) {
		f.Suffix = suffix
	}
}

func FileRequired(errKey string) FileOption {
	return func(f *File) {
		f.Required = true
		f.RequiredErrKey = errKey
	}
}

func FileMime(errKey string, mime ...string) FileOption {
	return func(f *File) {
		f.MimeErrKey = errKey
		f.Mime = mime
	}
}

func FileSizeInByte(sizeInByte int64, errKey string) FileOption {
	return func(f *File) {
		f.SizeInByte = sizeInByte
		f.SizeInByteErrKey = errKey
	}
}

func FileExtra(extra func()) FileOption {
	return func(f *File) {
		f.Extra = extra
	}
}
