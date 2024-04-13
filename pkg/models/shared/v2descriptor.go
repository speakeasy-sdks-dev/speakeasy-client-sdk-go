// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// V2Descriptor - V2 descriptor
type V2Descriptor struct {
	// Annotations
	Annotations *Annotations `json:"annotations,omitempty"`
	// Digest
	Digest *string `json:"digest,omitempty"`
	// Media type
	MediaType *string `json:"mediaType,omitempty"`
	// Size
	Size *int64 `json:"size,omitempty"`
}

func (o *V2Descriptor) GetAnnotations() *Annotations {
	if o == nil {
		return nil
	}
	return o.Annotations
}

func (o *V2Descriptor) GetDigest() *string {
	if o == nil {
		return nil
	}
	return o.Digest
}

func (o *V2Descriptor) GetMediaType() *string {
	if o == nil {
		return nil
	}
	return o.MediaType
}

func (o *V2Descriptor) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}