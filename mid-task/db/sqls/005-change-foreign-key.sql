ALTER TABLE tag_book DROP FOREIGN KEY tag_book_ibfk_2;
ALTER TABLE tag_book ADD FOREIGN KEY tag_book_ibfk_2(book_id) REFERENCES book (id) ON DELETE CASCADE;
ALTER TABLE tag_book DROP FOREIGN KEY tag_book_ibfk_1;
ALTER TABLE tag_book ADD FOREIGN KEY tag_book_ibfk_1(tag_id) REFERENCES tag (id) ON DELETE CASCADE;