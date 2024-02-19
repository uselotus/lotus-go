package lotus

//// StructToMap converts a struct value to map. Field keys are marked by 'form' tag
//func StructToMap(obj interface{}) map[string]string {
//	var (
//		m   = make(map[string]string)
//		val = reflect.ValueOf(obj)
//		typ = val.Type()
//	)
//
//	for i := 0; i < val.NumField(); i++ {
//		var (
//			field = typ.Field(i)
//			tag   = field.Tag.Get("form")
//			key   = strings.TrimSpace(strings.TrimSuffix(tag, ",omitempty"))
//			fv    = val.Field(i)
//		)
//
//		if key == "-" {
//			continue
//		}
//		if key == "" {
//			key = field.Name
//		}
//
//		if fv.Kind() == reflect.Ptr && fv.IsNil() {
//			// When field is pointer and nil, skip it
//			continue
//		}
//		omitempty := strings.Contains(tag, "omitempty")
//		if omitempty && isEmptyValue(fv) {
//			// When field is empty, skip it
//			continue
//		}
//
//		if fv.Kind() == reflect.Ptr {
//			// Get field real value from pointer
//			fv = fv.Elem()
//		}
//
//
//
//		if fv.Kind() == reflect.Slice {
//			// Handle slice type
//			svs := make([]string, fv.Len())
//			for j := 0; j < fv.Len(); j++ {
//				svs[j] = fmt.Sprintf("%v", fv.Index(j))
//			}
//			m[tag] = strings.Join(svs, ",")
//		} else if fv.Kind() == reflect.String {
//			m[tag] = fv.String()
//		}
//	}
//
//	return m
//}
//
//func isEmptyValue(val reflect.Value) bool {
//	switch val.Kind() {
//	case reflect.String, reflect.Slice, reflect.Map:
//		return val.Len() == 0
//	case reflect.Ptr:
//		return val.IsNil()
//	default:
//		zero := reflect.Zero(val.Type())
//		return reflect.DeepEqual(val.Interface(), zero.Interface())
//	}
//}
