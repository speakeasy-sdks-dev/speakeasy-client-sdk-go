// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type APIKeyDetails struct {
	WorkspaceID string `json:"workspace_id"`
}

func (o *APIKeyDetails) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}