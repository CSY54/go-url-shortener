package url

func ToUrl(uploadUrlDTO UploadUrlDTO) Url {
	return Url{
		Url: uploadUrlDTO.Url,
		ExpireAt: uploadUrlDTO.ExpireAt,
	}
}

func ToResponse(url Url) ResponseUrlDTO {
	id, _ := Uint32ToB64(uint32(url.ID))
	return ResponseUrlDTO{
		ID: id,
		// TODO
		ShortUrl: "http://localhost/" + id,
	}
}
