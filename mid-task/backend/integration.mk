export GO111MODULE := on

UID := demo
PORT := 5000
HOST := localhost

BOOK_ID:=1
BOOK_NAME:=GACKTの勝ち方
BOOK_ISBN:=9784800297662
BOOK_IMAGE:=https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/7662/9784800297662.jpg?_ex=200x200
BOOK_ITEM:=https://books.rakuten.co.jp/rb/16006950/

req-books-index:
	curl -v -XGET $(HOST):$(PORT)/books

req-books-post:
	curl -v -XPOST $(HOST):$(PORT)/books -d '{"name": "$(BOOK_NAME)", "isbn": "$(BOOK_ISBN)", "image_url": "$(BOOK_IMAGE)", "item_url": "$(BOOK_ITEM)"}'

req-books-delete:
	curl -v -XDELETE $(HOST):$(PORT)/books/$(BOOK_ID)

req-public:
	curl -v $(HOST):$(PORT)/public
