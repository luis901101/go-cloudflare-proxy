package dto

type GenericUploadDTO struct {
	UploadURL *string `json:"uploadURL,omitempty"`
	Uid       *string `json:"uid,omitempty"`
}

type GenericUploadResponseDTO = GenericResponseDTO[GenericUploadDTO]
