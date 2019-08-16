package model

type Book struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ISBN     int64  `db:"isbn" json:"isbn"`
	ImageURL string `db:"image_url" json:"image_url"`
}
type SearchedBooks struct {
	Count     int `json:"count"`
	Page      int `json:"page"`
	First     int `json:"first"`
	Last      int `json:"last"`
	Hits      int `json:"hits"`
	Carrier   int `json:"carrier"`
	PageCount int `json:"pageCount"`
	Items     []struct {
		Item struct {
			Title          string `json:"title"`
			TitleKana      string `json:"titleKana"`
			SubTitle       string `json:"subTitle"`
			SubTitleKana   string `json:"subTitleKana"`
			SeriesName     string `json:"seriesName"`
			SeriesNameKana string `json:"seriesNameKana"`
			Contents       string `json:"contents"`
			Author         string `json:"author"`
			AuthorKana     string `json:"authorKana"`
			PublisherName  string `json:"publisherName"`
			Size           string `json:"size"`
			Isbn           string `json:"isbn"`
			ItemCaption    string `json:"itemCaption"`
			SalesDate      string `json:"salesDate"`
			ItemPrice      int    `json:"itemPrice"`
			ListPrice      int    `json:"listPrice"`
			DiscountRate   int    `json:"discountRate"`
			DiscountPrice  int    `json:"discountPrice"`
			ItemURL        string `json:"itemUrl"`
			AffiliateURL   string `json:"affiliateUrl"`
			SmallImageURL  string `json:"smallImageUrl"`
			MediumImageURL string `json:"mediumImageUrl"`
			LargeImageURL  string `json:"largeImageUrl"`
			ChirayomiURL   string `json:"chirayomiUrl"`
			Availability   string `json:"availability"`
			PostageFlag    int    `json:"postageFlag"`
			LimitedFlag    int    `json:"limitedFlag"`
			ReviewCount    int    `json:"reviewCount"`
			ReviewAverage  string `json:"reviewAverage"`
			BooksGenreID   string `json:"booksGenreId"`
		} `json:"Item"`
	} `json:"Items"`
	GenreInformation []interface{} `json:"GenreInformation"`
}
