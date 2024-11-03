package helper

// Helper functions to safely dereference pointers or handle nil values
func DereferenceString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func DereferenceFloat(f *float64) int {
	if f == nil {
		return 0
	}
	return int(*f)
}

func ConvertToNullableFloat64(value *int) *float64 {
	if value == nil {
		return nil
	}
	convertedValue := float64(*value)
	return &convertedValue
}

func DereferenceInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
