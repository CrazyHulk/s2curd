func (s *{{.StructTableName}}) Scan(src interface{}) (err error) {
	*s = {{.StructTableName}}{}
	if src == nil {
		return
	}
	switch srcData := src.(type) {
	case string:
		err = json.Unmarshal([]byte(srcData), &s)
		return
	case []byte:
		err = json.Unmarshal(srcData, &s)
		return
	default:
		err = errors.Errorf("[i18n.String] unknown scan source: %+v", src)
	}
	return
}

// Value 实现 driver.Valuer 接口
func (s {{.StructTableName}}) Value() (v driver.Value, err error) {
	return json.Marshal(s)
}