package app

import (
	"fmt"
	"strings"

	"github.com/yann0917/dedao/services"
	"github.com/yann0917/dedao/utils"
)

// EbookDetail 电子书详情
func EbookDetail(id int) (detail *services.EbookDetail, err error) {
	courseDetail, err := CourseDetail(CateEbook, id)
	if err != nil {
		return
	}
	enID := courseDetail["enid"].(string)
	detail, err = getService().EbookDetail(enID)

	return
}

func EbookDetailByEnid(enID string) (detail *services.EbookDetail, err error) {
	detail, err = getService().EbookDetail(enID)
	return
}

// EbookCommentList get ebook 评分&书评
// sort like_count
func EbookCommentList(enID, sort string, page, limit int) (list *services.EbookCommentList, err error) {
	list, err = getService().EbookCommentList(enID, sort, page, limit)
	return
}

func EbookInfo(enID string) (info *services.EbookInfo, err error) {
	token, err1 := getService().EbookReadToken(enID)
	if err1 != nil {
		err = err1
		return
	}

	info, err = getService().EbookInfo(token.Token)
	return
}

func EbookPage(enID string) (info *services.EbookInfo, svgContent utils.SvgContents, err error) {
	token, err1 := getService().EbookReadToken(enID)
	if err1 != nil {
		err = err1
		return
	}

	info, err = getService().EbookInfo(token.Token)
	if err != nil {
		return
	}
	// fmt.Printf("%#v\n", info.BookInfo.Pages)
	// fmt.Printf("%#v\n", info.BookInfo.EbookBlock)
	// fmt.Printf("%#v\n", info.BookInfo.Toc)
	// fmt.Printf("%#v\n", info.BookInfo.Orders)
	wgp := utils.NewWaitGroupPool(10)
	for i, order := range info.BookInfo.Orders {
		wgp.Add()
		go func(i int, order services.EbookOrders) {
			defer func() {
				wgp.Done()
			}()
			index, count, offset := 0, 20, 0
			svgList, err1 := generateEbookPages(order.ChapterID, token.Token, index, count, offset)
			if err1 != nil {
				err = err1
				return
			}

			href, text, level := "", "", 0
			for _, ebookToc := range info.BookInfo.Toc {
				if strings.Contains(ebookToc.Href, order.ChapterID) {
					href = ebookToc.Href
					level = ebookToc.Level
					text = ebookToc.Text
					break
				}
			}

			svgContent = append(svgContent, &utils.SvgContent{
				Contents:   svgList,
				ChapterID:  order.ChapterID,
				PathInEpub: order.PathInEpub,
				TocLevel:   level,
				TocHref:    href,
				TocText:    text,
				OrderIndex: i,
			})
		}(i, order)
	}
	wgp.Wait()
	return
}

func generateEbookPages(chapterID, token string, index, count, offset int) (svgList []string, err error) {
	fmt.Printf("chapterID:%#v\n", chapterID)
	pageList, err := getService().EbookPages(chapterID, token, index, count, offset)
	if err != nil {
		return
	}

	for _, item := range pageList.Pages {
		svgList = append(svgList, item.Svg)
	}
	// fmt.Printf("IsEnd:%#v\n", pageList.IsEnd)
	if !pageList.IsEnd {
		index = count
		count += 20
		list, err1 := generateEbookPages(chapterID, token, index, count, offset)
		if err1 != nil {
			err = err1
			return
		}

		svgList = append(svgList, list...)
	}
	// FIXME: debug
	// err = utils.SaveFile(chapterID, "", strings.Join(svgList, "\n"))
	return
}
