package image_form

import (
	"fmt"
)

// get content thumbnail with a content amount
const contentThumbnailHTML = `<div class="content-thumbnail">%s</div>`
const contentThumbnailAmountHTML = `%s<div class="content-thumbnail-image-content-amount">&times%d</div>`

func GetContentThumbnailHTML(contentType, contentId, contentAmount int32) string {
	// attach together with the content thumbnail image to form a whole unit
	imagePtr := GetContentThumbnailImageHTML(contentType, contentId)
	if contentAmount == 0 {
		return fmt.Sprintf(contentThumbnailHTML, *imagePtr)
	}
	// note that this just add it on top, both of these need to be put inside a container for it to work
	return fmt.Sprintf(contentThumbnailHTML, formatter.Sprintf(contentThumbnailAmountHTML, *imagePtr, contentAmount))
}
