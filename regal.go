package regal

func RegalEngine(versionHost [][]string, params ...ParamOption) *baseInfo {
	m := NewOrderedMap[string, string]()
	for i := 0; i < len(versionHost); i++ {
		m.Set(versionHost[i][0], versionHost[i][1])
	}

	bi := &baseInfo{
		versionHost: m,
		params:      defaultParams(),
	}

	for _, param := range params {
		param.apply(&bi.params)
	}
	return bi
}
