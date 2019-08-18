<template>
  <v-container grid-list-md>
    <v-layout align-center justify-center>
      <v-flex xs6 class="my-5">
        <v-card tile class="pa-3">
          <v-text-field v-model="inputedText" label="キーワードを入力" required></v-text-field>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="info" @click="searchBooks">検索</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
    <v-spacer></v-spacer>
    <v-layout row wrap>
      <v-flex xs3 v-for="(book_info, index) in books" :key="index">
        <SearchedBookCard :book_info="book_info.Item" />
      </v-flex>
    </v-layout>
    <div class="text-xs-center">
      <v-btn round color="primary" dark @click="moreLoad">もっと見る</v-btn>
    </div>
  </v-container>
</template>

<script>
import SearchedBookCard from '../components/SearchedBookCard';
import store from '../App';

export default {
  components: {
    SearchedBookCard
  },
  data() {
    return {
      books: [],
      inputedText: '',
      pager: 1,
      searchedText: ''
    };
  },
  methods: {
    searchBooks() {
      const url = process.env.VUE_APP_API_URL + '/search';
      const keyword = this.$data.inputedText;
      this.$data.pager = 1;

      this.requestJson(url, keyword, 1).then(json => {
        this.$data.books = json.Items;
        this.$data.searchedText = keyword;
      });
    },
    moreLoad() {
      const url = process.env.VUE_APP_API_URL + '/search';
      this.requestJson(url, this.$data.searchedText, this.$data.pager).then(json => {
        this.$data.books = this.$data.books.concat(json.Items);
      });
    },
    requestJson(url, keyword, pager) {
      return fetch(url + '?keyword=' + keyword + '&page=' + pager)
        .then(response => response.json())
        .then(json => {
          console.log(json);
          this.$data.pager += 1;
          return json;
        });
    }
  }
};
</script>
