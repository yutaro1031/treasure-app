<template>
  <v-card class="pt-3">
    <v-img :src="book_info.image_url" height="170px" width="120px" class="mx-auto"> </v-img>

    <v-card-text>
      <h3 class="card-text">{{ book_info['name'] }}</h3>
    </v-card-text>

    <v-card-actions>
      <v-btn class="ml-2" color="info" v-bind:href="book_info.item_url" target="_blank">
        詳細をみる
      </v-btn>
      <v-spacer></v-spacer>
      <v-btn flat icon color="red" @click="deleteBook">
        <v-icon>fas fa-trash-alt</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  props: ['book_info'],
  methods: {
    deleteBook() {
      const url = process.env.VUE_APP_API_URL;
      const id = this.$props.book_info.id;
      fetch(url + '/books/' + id, { method: 'DELETE' }).then(this.$emit('delete'));
    }
  }
};
</script>

<style>
.card-text {
  min-height: 50px;
}
</style>
