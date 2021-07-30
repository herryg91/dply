package k8s_repository

func toInt32Pointer(in int) *int32 {
	in32 := int32(in)
	return &in32
}
