<template>
  <v-card class="pt-3">
    <v-img :src="book_info.image_url" height="170px" width="120px" class="mx-auto"> </v-img>

    <v-card-text>
      <h3 class="card-text">{{ book_info['name'] }}</h3>
      <div class="hoge" v-for="(tag_info, index) in tagDatas" :key="index">
        <Tag :tag_info="tag_info" @deleteTag="deleteTag(index)" />
      </div>
      <v-select :items="tags" item-value="id" item-text="name" label="タグを追加" @change="addTag"></v-select>
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
import Tag from './Tag';

export default {
  components: {
    Tag
  },
  props: ['book_info', 'tags'],
  data() {
    return {
      tagDatas: this.$props.book_info.tags
    };
  },
  methods: {
    deleteBook() {
      const url = process.env.VUE_APP_API_URL;
      const id = this.$props.book_info.id;
      fetch(url + '/books/' + id, { method: 'DELETE' }).then(this.$emit('deleteBook'));
    },
    deleteTag(index) {
      const url = process.env.VUE_APP_API_URL;
      const id = this.$props.book_info.tags[index].id;
      fetch(url + '/tag_books/' + id, { method: 'DELETE' }).then(console.log('Success deleted tag.'));
    },
    addTag(tagID, name) {
      console.log(tagID);
      const url = process.env.VUE_APP_API_URL;
      const bookID = this.$props.book_info.id;
      const body = JSON.stringify({
        book_id: bookID,
        tag_id: tagID
      });
      const headers = {
        'Content-Type': 'application/json'
      };
      fetch(url + '/tag_books', { method: 'POST', headers, body })
        .then(response => response.json())
        .then(json => {
          const tagInfos = this.$props.tags;
          for (let i = 0, len = tagInfos.length; i < len; i += 1) {
            if (json.tag_id === tagInfos[i].id) this.$data.tagDatas.push({ id: json.id, name: tagInfos[i].name });
          }
        });
    }
  }
};
</script>

<style>
.card-text {
  min-height: 50px;
}
</style>
