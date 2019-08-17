<template>
  <v-card class="pt-3">
    <v-img :src="book_info.largeImageUrl" height="170px" width="120px" class="mx-auto"> </v-img>

    <v-card-text>
      <h3 class="card-text">{{ book_info['title'] }}</h3>
    </v-card-text>

    <v-card-actions>
      <v-btn class="mx-auto" color="info" v-bind:href="book_info.itemUrl" target="_blank">
        楽天市場で詳細をみる
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ['book_info'],
  // data() {},
  created() {},
  methods: {}
};
</script>

<style>
.card-text {
  min-height: 50px;
}
</style>
