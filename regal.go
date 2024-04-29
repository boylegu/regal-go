package regal

func RegalEngine(versionHost map[string]string, params ...ParamOption) *baseInfo {
	bi := &baseInfo{
		versionHost: versionHost,
		params:      defaultParams(),
	}

	for _, param := range params {
		param.apply(&bi.params)
	}
	return bi
}
