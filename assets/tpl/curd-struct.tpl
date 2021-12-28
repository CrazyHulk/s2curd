func (s *{{.Name}}) Scan(src interface{}) (err error) {
	*s = {{.Name}}{}
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
		err = errors.Errorf("[{{.Name}}] unknown scan source: %+v", src)
	}
	return
}

// Value 实现 driver.Valuer 接口
func (s {{.Name}}) Value() (v driver.Value, err error) {
	return json.Marshal(s)
}
