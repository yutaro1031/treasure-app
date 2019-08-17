<template>
  <v-card class="pt-3">
    <v-img :src="book_info.largeImageUrl" height="170px" width="120px" class="mx-auto"> </v-img>

    <v-card-text>
      <div class="card-text">
        <p>{{ book_info['title'] }}</p>
        <span class="grey--text discription-link">
          <a class="not-decolation" v-bind:href="book_info.itemUrl" target="_blank">
            楽天市場で詳細をみる
          </a>
        </span>
      </div>
    </v-card-text>

    <v-card-actions>
      <v-btn v-if="!isRegistered" class="mx-auto" color="info" @click="registBook">
        登録する
      </v-btn>
      <v-btn v-else class="mx-auto" color="blue-grey" disabled>
        登録済み
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ['book_info'],
  data() {
    return {
      isRegistered: false
    };
  },
  methods: {
    registBook() {
      const url = process.env.VUE_APP_API_URL + '/books';
      const bookInfo = this.$props.book_info;
      const body = JSON.stringify({
        name: bookInfo.title,
        isbn: bookInfo.isbn,
        image_url: bookInfo.largeImageUrl,
        item_url: bookInfo.largeImageUrl
      });
      const headers = {
        'Content-Type': 'application/json'
      };
      fetch(url, { method: 'POST', headers, body }).then((this.$data.isRegistered = true));
    }
  }
};
</script>

<style>
.discription-link {
  display: block;
  position: absolute;
  bottom: 0;
  right: 0;
}
.card-text {
  position: relative;
  min-height: 100px;
}
.not-decolation {
  text-decoration: none;
  color: grey;
}
</style>
