package handlers

import (
	"fmt"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/signintech/gopdf"
)

type MangaRequest struct {
	MangaURL string `json:"manga_url"`
	MangaName string `json:"manga_name"`
	ChapterNo int `json:"chapter_no"`
}

func MangaToPdf(c *gin.Context){
	var request MangaRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message":"invalid request body",
		})
		c.Abort()
		return
	}


	doc, err := fetchHTML(request.MangaURL)
	if err != nil {
		c.JSON(500, gin.H{
			"message":"failed to fetch html",
		})
		c.Abort()
		return
	}

	imageUrls, err := extractImageUrl(doc)
	fmt.Println("image urls ", imageUrls)
	if err != nil  {
		fmt.Println("error: ", err)
		c.JSON(500, gin.H{
			"message":"failed to fetch urls 47",
		})
		c.Abort()
		return
	}

	err = createPDF(imageUrls, request.MangaName, request.ChapterNo)
	fmt.Println("error ", err)
	if err != nil {
		c.JSON(500, gin.H{
			"message":"failed to fetch urls",
		})
			c.Abort()
		return
	}
c.JSON(http.StatusOK, gin.H{"message": "PDF created successfully!"})
	
}

func fetchHTML(url string)(*goquery.Document, error){
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func extractImageUrl(doc *goquery.Document)([]string, error){
	var imageUrls []string

	doc.Find("img").Each(func(index int, element *goquery.Selection) {
		src, exists := element.Attr("src")
		if exists {
			src = strings.TrimSpace(src)
			imageUrls = append(imageUrls, src)
		}
	})
	return imageUrls, nil
}

func createPDF(imageUrls []string, mangaName string, chapterNumber int) error {
	err := os.MkdirAll("downloads", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create downloads directory: %v", err)
	}

	outputFileName := fmt.Sprintf("downloads/%s_%d.pdf", mangaName, chapterNumber)

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})


	imagePaths := []string{}

	for _, imageUrl := range imageUrls {
		imgPath, err := downloadImage(imageUrl)
		if err != nil {
			return fmt.Errorf("failed to download image %s: %v", imageUrl, err)
		}

		pdf.AddPage()

		imgFile, err := os.Open(imgPath)
		if err != nil {
			return fmt.Errorf("failed to open image file %s: %v", imgPath, err)
		}

		defer imgFile.Close()

		img, err := jpeg.Decode(imgFile)
		if err != nil {
			return fmt.Errorf("failed to decode image %s: %v", imgPath, err)
		}

		bounds := img.Bounds()
		width := float64(bounds.Dx())
		height := float64(bounds.Dy())

		pageWidth, pageHeight := gopdf.PageSizeA4.W, gopdf.PageSizeA4.H
		scaleFactor := pageWidth / width
		if height*scaleFactor > pageHeight {
			scaleFactor = pageHeight / height
		}

		pdf.Image(imgPath, 0, 0, &gopdf.Rect{
			W: width * scaleFactor,
			H: height * scaleFactor,
		})

		imagePaths = append(imagePaths, imgPath)
	}

	err = pdf.WritePdf(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to write PDF file: %v", err)
	}

	// go func(paths []string){
	// 	for _, path := range paths {
	// 		err := os.Remove(path)
	// 		if err == nil {
	// 			break
	// 		}
	// 		fmt.Printf("failed to delete image file %s: %v\n", path, err)
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }(imagePaths)

		go func(paths []string) {
		for _, path := range paths {
			for i := 0; i < 3; i++ { 
				err := os.Remove(path)
				if err == nil {
					break
				}
				fmt.Printf("failed to delete image file %s (attempt %d): %v\n", path, i+1, err)
				time.Sleep(500 * time.Millisecond) 
			}
		}
	}(imagePaths)

	return nil
}



func downloadImage(url string) (string, error) {
	err := os.MkdirAll("downloads", os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create downloads directory: %v", err)
	}

	imageFileName := filepath.Join("downloads", filepath.Base(url))
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get image from URL %s: %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get image from URL %s: status code %d", url, res.StatusCode)
	}

	// imageFileName := filepath.Base(url)
	imgFile, err := os.Create(imageFileName)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %v", imageFileName, err)
	}
	defer imgFile.Close()

	_, err = io.Copy(imgFile, res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image to file %s: %v", imageFileName, err)
	}

	return imageFileName, nil
}

