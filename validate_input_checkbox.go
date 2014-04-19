package form

func (va validate) bInputCheckbox() {
	value := va.value.Bool()

	mandatory, ok := va.getBool("Mandatory")
	if ok {
		if mandatory {
			if !value {
				manErr, ok := va.getStr("MandatoryErr")
				if ok {
					va.setErr(FormError(manErr))
				} else {
					va.setErr(FormError(va.i18n.Key(ErrMandatoryCheckbox)))
				}
				return
			}
		}
	}

	va.callExt()
}
