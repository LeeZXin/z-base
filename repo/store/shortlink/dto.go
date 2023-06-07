package shortlink

type InsertLinkReqDTO struct {
	ShortLink string
	LongLink  string
}

type InsertLinkRespDTO struct {
	Success bool
}

type GetLongLinkByShortLinkReqDTO struct {
	ShortLink string
}

type GetLongLinkByShortLinkRespDTO struct {
	Exists   bool
	LongLink string
}
