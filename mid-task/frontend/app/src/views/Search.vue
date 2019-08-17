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
      inputedText: ''
    };
  },
  methods: {
    searchBooks() {
      const url = process.env.VUE_APP_API_URL + '/search';
      const keyword = this.$data.inputedText;
      console.log(keyword);

      fetch(url + '?keyword=' + keyword)
        .then(response => response.json())
        .then(json => {
          console.log(json.Items);
          this.$data.books = json.Items;
        });
    }
  }
};
</script>
