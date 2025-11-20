package image_form

import (
	"fmt"
)

// render an trade item by html
// Looks something like this
// -------------
// |stock:     |
// -------------
// |thumbnail  |
// |asset      |
// |           |
// -------------
// |icon  price|
// -------------
//

// format this string with 5 strings:
// - the stock amount: text
// - the content thumbnail: html element
// - the source item asset thumbnail: html element
// - the price: text
// the html element have helpers
const tradeHTMLFormatString = `<div class="trade-item">
	<div class="trade-item-stock-amount-area">
		<div class="trade-item-stock-amount-text">Stock:	%s</div>
	</div>
	<div class="trade-item-thumbnail-area">
		%s
	</div>
	<div class="trade-item-thumbnail-price-separation-area"></div>
	<div class="trade-item-price-area">
		<div class="trade-item-price-icon-area">
			%s
		</div>
		<div class="trade-item-price-text"> %s </div>
	</div>
</div>`

// for this API, if content amount is 0, then it's not displayed
// even if it is 1, it will be displayed
func GetTradeItemHTML(contentType, contentId, contentAmount, stock, sourceContentId, sourceContentAmount int32) string {
	return fmt.Sprintf(tradeHTMLFormatString,
		formatter.Sprintf("%d", stock),
		GetContentThumbnailHTML(contentType, contentId, contentAmount),
		*GetExchangeEventPointIconAssetPath(sourceContentId),
		formatter.Sprintf("%d", sourceContentAmount),
	)
}
