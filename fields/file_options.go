package fields

func FileSuffix(suffix *string) func(*File) {
	return func(f *File) {
		f.Suffix = suffix
	}
}

func FileRequired(required bool, errKey string) func(*File) {
	return func(f *File) {
		f.Required = required
		f.RequiredErrKey = errKey
	}
}

func FileMime(errKey string, mime ...string) func(*File) {
	return func(f *File) {
		f.MimeErrKey = errKey
		f.Mime = mime
	}
}

func FileSizeInByte(sizeInByte int64, errKey string) func(*File) {
	return func(f *File) {
		f.SizeInByte = sizeInByte
		f.SizeInByteErrKey = errKey
	}
}

func FileExtra(extra func()) func(*File) {
	return func(f *File) {
		f.Extra = extra
	}
}
