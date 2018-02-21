struct Book {
   char  title[50];
   char  author[50];
};

Book b;	// a mess
memset(&b, 0, sizeof(Book));
strcpy(b.title, "Golang");
