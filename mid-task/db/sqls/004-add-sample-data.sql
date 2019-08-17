INSERT INTO `book`(`name`, `isbn`, `image_url`, `item_url`)
    VALUES ("ルージュ", "9784334777456", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/7456/9784334777456.jpg?_ex=200x200", "https://books.rakuten.co.jp/rb/15696282/");
INSERT INTO `book`(`name`, `isbn`, `image_url`, `item_url`)
    VALUES ("月と太陽の盤", "9784334778736", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8736/9784334778736.jpg?_ex=200x200", "https://books.rakuten.co.jp/rb/15968811/");
INSERT INTO `book`(`name`, `isbn`, `image_url`, `item_url`)
    VALUES ("ノワール", "9784122066762", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/6762/9784122066762.jpg?_ex=200x200", "https://books.rakuten.co.jp/rb/15708413/");
INSERT INTO `book`(`name`, `isbn`, `image_url`, `item_url`)
    VALUES ("凍てつく太陽", "9784344033443", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/3443/9784344033443.jpg?_ex=200x200", "https://books.rakuten.co.jp/rb/15562495/");
INSERT INTO `book`(`name`, `isbn`, `image_url`, `item_url`)
    VALUES ("太陽の子", "9784043520107", "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/0107/9784043520107.jpg?_ex=200x200", "https://books.rakuten.co.jp/rb/979251/");


INSERT INTO `tag`(`name`) VALUES ("泣ける");
INSERT INTO `tag`(`name`) VALUES ("感動");
INSERT INTO `tag`(`name`) VALUES ("笑える");
INSERT INTO `tag`(`name`) VALUES ("萎える");
INSERT INTO `tag`(`name`) VALUES ("燃える");

INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (1,2);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (4,3);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (2,4);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (5,5);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (4,4);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (3,2);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (2,1);
INSERT INTO `tag_book`(`tag_id`, `book_id`) VALUES (1,5);
