<template>
  <v-container grid-list-md>
    <v-layout align-center id="text-layout">
      <v-flex xs12>
        <h1 class="text-xs-center" id="heading">これが俺の本だ。</h1>
      </v-flex>
    </v-layout>
    <v-spacer></v-spacer>
    <v-layout row wrap>
      <v-flex xs3 v-for="(book_info, index) in books" :key="index">
        <BookCard :book_info="book_info" :tags="tags" @deleteBook="deleteBook(index)" />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import BookCard from '../components/BookCard';

export default {
  components: {
    BookCard
  },
  data() {
    return {
      books: [],
      tags: []
    };
  },
  created() {
    const url = process.env.VUE_APP_API_URL;
    this.requestJson(url + '/tags').then(json => (this.$data.tags = json));
    this.requestJson(url + '/books').then(json => (this.$data.books = json));
  },
  methods: {
    deleteBook(index) {
      this.$data.books.splice(index, 1);
    },
    requestJson(url) {
      return fetch(url)
        .then(response => response.json())
        .catch(err => console.error(err));
    }
  }
};
</script>
