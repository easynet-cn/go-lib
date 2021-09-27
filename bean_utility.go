package golib

import (
	"fmt"
	"reflect"
	"time"
)

type BeanUpdateLog struct {
	Changed    bool                `json:"changed"`
	UpdateLogs []BeanUpdateLogItem `json:"updateLogs"`
}

type BeanUpdateLogItem struct {
	PropertyName     string      `json:"propertyName"`
	OldProperyValue  interface{} `json:"oldProperyValue"`
	NewPropertyValue interface{} `json:"newPropertyValue"`
}

type BeanUtiltiy struct {
	datetimeFormat     string
	propertyNameFormat func(string) string
}

func NewBeanUtiltiy(datetimeFormat string, propertyNameFormat func(string) string) *BeanUtiltiy {
	return &BeanUtiltiy{datetimeFormat: datetimeFormat, propertyNameFormat: propertyNameFormat}
}

func (s *BeanUtiltiy) Copy(source interface{}, dest interface{}, ignoreFields ...string) ([]string, error) {
	destType := reflect.TypeOf(dest)
	destValue := reflect.ValueOf(dest)
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	if destType.Kind() != reflect.Ptr {
		return make([]string, 0), fmt.Errorf("a must be a struct pointer")
	}

	destValue = reflect.ValueOf(destValue.Interface())

	m := StringSliceToMap(ignoreFields)
	fileds := make([]string, 0)

	for i := 0; i < sourceValue.NumField(); i++ {
		name := sourceType.Field(i).Name

		if _, ok := m[name]; !ok {
			fileds = append(fileds, name)
		}
	}

	if len(fileds) == 0 {
		return make([]string, 0), nil
	}

	modified := make([]string, 0, len(fileds))

	for i := 0; i < len(fileds); i++ {
		name := fileds[i]
		sourceFieldValue := sourceValue.FieldByName(name)
		destFieldValue := destValue.Elem().FieldByName(name)

		if destFieldValue.IsValid() {
			sourceFieldValueTypeStr := sourceFieldValue.Type().String()
			destFieldValueTypeStr := destFieldValue.Type().String()

			if (sourceFieldValueTypeStr == "time.Time" ||
				sourceFieldValueTypeStr == "*time.Time") &&
				(destFieldValueTypeStr == "string" ||
					destFieldValueTypeStr == "*string") {
				timeStr := ""

				if sourceFieldValueTypeStr == "time.Time" {
					t := (sourceFieldValue.Interface()).(time.Time)

					if !t.IsZero() {
						timeStr = t.Format(s.datetimeFormat)
					}

				} else if sourceFieldValueTypeStr == "*time.Time" {
					timePtr := (sourceFieldValue.Interface()).(*time.Time)

					if timePtr != nil && !timePtr.IsZero() {
						timeStr = (*timePtr).Format(s.datetimeFormat)
					}
				}

				if destFieldValueTypeStr == "string" &&
					!reflect.DeepEqual(destFieldValue.Interface(), timeStr) {

					destFieldValue.Set(reflect.ValueOf(timeStr))

					modified = append(modified, name)
				} else if destFieldValueTypeStr == "*string" {
					strPtr := (destFieldValue.Interface()).(*string)

					if strPtr != nil && !reflect.DeepEqual(*strPtr, timeStr) {
						destFieldValue.Set(reflect.ValueOf(&timeStr))

						modified = append(modified, name)
					} else {
						destFieldValue.Set(reflect.ValueOf(&timeStr))

						modified = append(modified, name)
					}
				}
			} else if (sourceFieldValueTypeStr == "string" || sourceFieldValueTypeStr == "*string") &&
				(destFieldValueTypeStr == "time.Time" || destFieldValueTypeStr == "*time.Time") {
				timeStr := ""

				if sourceFieldValueTypeStr == "string" {
					timeStr = (sourceFieldValue.Interface()).(string)
				} else if sourceFieldValueTypeStr == "*string" {
					timePtr := (sourceFieldValue.Interface()).(*string)

					if timePtr != nil {
						timeStr = *timePtr
					}
				}

				if t, err := time.ParseInLocation(s.datetimeFormat, timeStr, time.Local); err == nil {
					if destFieldValueTypeStr == "time.Time" && !reflect.DeepEqual(t, destFieldValue.Interface()) {
						destFieldValue.Set(reflect.ValueOf(t))

						modified = append(modified, name)
					} else if destFieldValueTypeStr == "*time.Time" {
						timePtr := (destFieldValue.Interface()).(*time.Time)

						if timePtr != nil && !reflect.DeepEqual((*timePtr).Format(s.datetimeFormat), timeStr) {
							destFieldValue.Set(reflect.ValueOf(&t))

							modified = append(modified, name)
						} else if timePtr == nil {
							destFieldValue.Set(reflect.ValueOf(&t))

							modified = append(modified, name)
						}
					}

				}
			} else if sourceFieldValueTypeStr[0] == '*' &&
				sourceFieldValueTypeStr[1:] == destFieldValueTypeStr &&
				!sourceFieldValue.IsNil() && !reflect.DeepEqual(sourceFieldValue.Elem().Interface(), destFieldValue.Interface()) {
				destFieldValue.Set(reflect.ValueOf(sourceFieldValue.Elem().Interface()))

				modified = append(modified, name)
			} else if destFieldValueTypeStr[0] == '*' &&
				destFieldValueTypeStr[1:] == sourceFieldValueTypeStr &&
				((!destFieldValue.IsZero() && !reflect.DeepEqual(sourceFieldValue.Interface(), destFieldValue.Elem().Interface())) || (destFieldValue.IsNil())) {
				switch sourceFieldValue.Kind() {
				case reflect.Bool:
					v := sourceFieldValue.Interface().(bool)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int:
					v := sourceFieldValue.Interface().(int)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int8:
					v := sourceFieldValue.Interface().(int8)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int16:
					v := sourceFieldValue.Interface().(int16)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int32:
					v := sourceFieldValue.Interface().(int32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Int64:
					v := sourceFieldValue.Interface().(int64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint:
					v := sourceFieldValue.Interface().(uint)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint8:
					v := sourceFieldValue.Interface().(uint8)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint16:
					v := sourceFieldValue.Interface().(uint16)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint32:
					v := sourceFieldValue.Interface().(uint32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Uint64:
					v := sourceFieldValue.Interface().(uint64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Float32:
					v := sourceFieldValue.Interface().(float32)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.Float64:
					v := sourceFieldValue.Interface().(float64)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				case reflect.String:
					v := sourceFieldValue.Interface().(string)
					destFieldValue.Set(reflect.ValueOf(&v))

					modified = append(modified, name)
				}

			} else if destFieldValue.Kind() == sourceFieldValue.Kind() &&
				!reflect.DeepEqual(destFieldValue.Interface(), sourceFieldValue.Interface()) {
				destFieldValue.Set(sourceFieldValue)

				modified = append(modified, name)
			}
		}

	}

	return modified, nil
}

func (s *BeanUtiltiy) CopyAndLogProperties(source interface{}, dest interface{}, ignoreFields ...string) (BeanUpdateLog, error) {
	beanUpdateLogs := &BeanUpdateLog{UpdateLogs: make([]BeanUpdateLogItem, 0)}
	destType := reflect.TypeOf(dest)
	destValue := reflect.ValueOf(dest)
	sourceType := reflect.TypeOf(source)
	sourceValue := reflect.ValueOf(source)

	if destType.Kind() != reflect.Ptr {
		return *beanUpdateLogs, fmt.Errorf("dest must be dest struct pointer")
	}

	destValue = reflect.ValueOf(destValue.Interface())

	m := StringSliceToMap(ignoreFields)
	fileds := make([]string, 0)

	for i := 0; i < sourceValue.NumField(); i++ {
		name := sourceType.Field(i).Name

		if _, ok := m[name]; !ok {
			fileds = append(fileds, name)
		}
	}

	if len(fileds) == 0 {
		return *beanUpdateLogs, nil
	}

	for i := 0; i < len(fileds); i++ {
		name := fileds[i]
		sourceFieldValue := sourceValue.FieldByName(name)

		destFieldValue := destValue.Elem().FieldByName(name)

		if destFieldValue.IsValid() {
			lname := s.propertyNameFormat(name)
			destFieldValueTypeStr := destFieldValue.Type().String()
			sourceFieldValueTypeStr := sourceFieldValue.Type().String()

			if (sourceFieldValueTypeStr == "time.Time" ||
				sourceFieldValueTypeStr == "*time.Time") &&
				(destFieldValueTypeStr == "string" ||
					destFieldValueTypeStr == "*string") {
				timeStr := ""

				if sourceFieldValueTypeStr == "time.Time" {
					t := (sourceFieldValue.Interface()).(time.Time)

					if !t.IsZero() {
						timeStr = t.Format(s.datetimeFormat)
					}

				} else if sourceFieldValueTypeStr == "*time.Time" {
					timePtr := (sourceFieldValue.Interface()).(*time.Time)

					if timePtr != nil && !timePtr.IsZero() {
						timeStr = (*timePtr).Format(s.datetimeFormat)
					}
				}

				if destFieldValueTypeStr == "string" &&
					!reflect.DeepEqual(destFieldValue.Interface(), timeStr) {
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: timeStr}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					destFieldValue.Set(reflect.ValueOf(timeStr))
				} else if destFieldValueTypeStr == "*string" {
					strPtr := (destFieldValue.Interface()).(*string)

					if strPtr != nil && !reflect.DeepEqual(*strPtr, timeStr) {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(&timeStr))
					} else {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(&timeStr))
					}
				}
			} else if (sourceFieldValueTypeStr == "string" || sourceFieldValueTypeStr == "*string") &&
				(destFieldValueTypeStr == "time.Time" || destFieldValueTypeStr == "*time.Time") {
				timeStr := ""

				if sourceFieldValueTypeStr == "string" {
					timeStr = (sourceFieldValue.Interface()).(string)
				} else if sourceFieldValueTypeStr == "*string" {
					timePtr := (sourceFieldValue.Interface()).(*string)

					if timePtr != nil {
						timeStr = *timePtr
					}
				}

				if t, err := time.ParseInLocation(s.datetimeFormat, timeStr, time.Local); err == nil {
					if destFieldValueTypeStr == "time.Time" && !reflect.DeepEqual(t, destFieldValue.Interface()) {
						beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: t}
						beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

						destFieldValue.Set(reflect.ValueOf(t))
					} else if destFieldValueTypeStr == "*time.Time" {
						timePtr := (destFieldValue.Interface()).(*time.Time)

						if timePtr != nil && !reflect.DeepEqual((*timePtr).Format(s.datetimeFormat), timeStr) {
							beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &t}
							beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

							destFieldValue.Set(reflect.ValueOf(&t))
						} else if timePtr == nil {
							beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: &timeStr}
							beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

							destFieldValue.Set(reflect.ValueOf(&t))
						}
					}

				}
			} else if sourceFieldValueTypeStr[0] == '*' &&
				sourceFieldValueTypeStr[1:] == destFieldValueTypeStr &&
				!sourceFieldValue.IsNil() && !reflect.DeepEqual(sourceFieldValue.Elem().Interface(), destFieldValue.Interface()) {

				beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
				beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

				destFieldValue.Set(reflect.ValueOf(sourceFieldValue.Elem().Interface()))
			} else if destFieldValueTypeStr[0] == '*' &&
				destFieldValueTypeStr[1:] == sourceFieldValueTypeStr &&
				((!destFieldValue.IsZero() && !reflect.DeepEqual(sourceFieldValue.Interface(), destFieldValue.Elem().Interface())) || (destFieldValue.IsNil())) {
				switch sourceFieldValue.Kind() {
				case reflect.Bool:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(bool)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int8:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int8)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int16:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int16)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Int64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(int64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint8:
					v := sourceFieldValue.Interface().(uint8)
					destFieldValue.Set(reflect.ValueOf(&v))

					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)
				case reflect.Uint16:
					v := sourceFieldValue.Interface().(uint16)
					destFieldValue.Set(reflect.ValueOf(&v))

					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)
				case reflect.Uint32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Uint64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(uint64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Float32:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(float32)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.Float64:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(float64)
					destFieldValue.Set(reflect.ValueOf(&v))
				case reflect.String:
					beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
					beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

					v := sourceFieldValue.Interface().(string)
					destFieldValue.Set(reflect.ValueOf(&v))
				}

			} else if destFieldValue.Kind() == sourceFieldValue.Kind() &&
				!reflect.DeepEqual(destFieldValue.Interface(), sourceFieldValue.Interface()) {
				beanUpdateLogItem := &BeanUpdateLogItem{PropertyName: lname, OldProperyValue: destFieldValue.Interface(), NewPropertyValue: sourceFieldValue.Interface()}
				beanUpdateLogs.UpdateLogs = append(beanUpdateLogs.UpdateLogs, *beanUpdateLogItem)

				destFieldValue.Set(sourceFieldValue)
			}
		}

	}

	beanUpdateLogs.Changed = len(beanUpdateLogs.UpdateLogs) > 0

	return *beanUpdateLogs, nil
}
