package model

type GiffyResponse struct {
	Data []struct {
		Type             string `json:"type"`
		ID               string `json:"id"`
		URL              string `json:"url"`
		Slug             string `json:"slug"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Title            string `json:"title"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		Images           struct {
			Original               original        `json:"original"`
			Downsized              sizeurl         `json:"downsized"`
			DownsizedLarge         sizeurl         `json:"downsized_large"`
			DownsizedMedium        sizeurl         `json:"downsized_medium"`
			DownsizedSmall         sizeurlmp4      `json:"downsized_small"`
			DownsizedStill         sizeurl         `json:"downsized_still"`
			FixedHeight            sizeMp4Fix      `json:"fixed_height"`
			FixedHeightDownsampled sizeFix         `json:"fixed_height_downsampled"`
			FixedHeightSmall       sizeMp4Fix      `json:"fixed_height_small"`
			FixedHeightSmallStill  sizeurl         `json:"fixed_height_small_still"`
			FixedHeightStill       sizeurl         `json:"fixed_height_still"`
			FixedWidth             sizeMp4Fix      `json:"fixed_width"`
			FixedWidthDownsampled  sizeFix         `json:"fixed_width_downsampled"`
			FixedWidthSmall        sizeMp4Fix      `json:"fixed_width_small"`
			FixedWidthSmallStill   sizeurl         `json:"fixed_width_small_still"`
			FixedWidthStill        sizeurl         `json:"fixed_width_still"`
			Looping                sizeurlmp4Short `json:"looping"`
			OriginalStill          sizeurl         `json:"original_still"`
			OriginalMp4            sizeurlmp4      `json:"original_mp4"`
			Preview                sizeurlmp4      `json:"preview"`
			PreviewGif             sizeurl         `json:"preview_gif"`
			PreviewWebp            sizeurl         `json:"preview_webp"`
			Four80WStill           sizeurl         `json:"480w_still"`
		} `json:"images"`
		User                     user      `json:"user"`
		AnalyticsResponsePayload string    `json:"analytics_response_payload"`
		Analytics                analistic `json:"analytics"`
	} `json:"data"`
	Pagination pagination `json:"pagination"`
	Meta       meta       `json:"meta"`
}

type original struct {
	Height   string `json:"height"`
	Width    string `json:"width"`
	Size     string `json:"size"`
	URL      string `json:"url"`
	Mp4Size  string `json:"mp4_size"`
	Mp4      string `json:"mp4"`
	WebpSize string `json:"webp_size"`
	Webp     string `json:"webp"`
	Frames   string `json:"frames"`
	Hash     string `json:"hash"`
}

type sizeFix struct {
	Height   string `json:"height"`
	Width    string `json:"width"`
	Size     string `json:"size"`
	URL      string `json:"url"`
	WebpSize string `json:"webp_size"`
	Webp     string `json:"webp"`
}

type sizeMp4Fix struct {
	Height   string `json:"height"`
	Width    string `json:"width"`
	Size     string `json:"size"`
	URL      string `json:"url"`
	Mp4Size  string `json:"mp4_size"`
	Mp4      string `json:"mp4"`
	WebpSize string `json:"webp_size"`
	Webp     string `json:"webp"`
}

type sizeurlmp4Short struct {
	Mp4Size string `json:"mp4_size"`
	Mp4     string `json:"mp4"`
}

type sizeurlmp4 struct {
	Height  string `json:"height"`
	Width   string `json:"width"`
	Mp4Size string `json:"mp4_size"`
	Mp4     string `json:"mp4"`
}

type sizeurl struct {
	Height string `json:"height"`
	Width  string `json:"width"`
	Size   string `json:"size"`
	URL    string `json:"url"`
}

type user struct {
	AvatarURL    string `json:"avatar_url"`
	BannerImage  string `json:"banner_image"`
	BannerURL    string `json:"banner_url"`
	ProfileURL   string `json:"profile_url"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	InstagramURL string `json:"instagram_url"`
	WebsiteURL   string `json:"website_url"`
	IsVerified   bool   `json:"is_verified"`
}

type url struct {
	URL string `json:"url"`
}

type analistic struct {
	Onload  url `json:"onload"`
	Onclick url `json:"onclick"`
	Onsent  url `json:"onsent"`
}

type pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

type meta struct {
	Status     int    `json:"status"`
	Msg        string `json:"msg"`
	ResponseID string `json:"response_id"`
}
