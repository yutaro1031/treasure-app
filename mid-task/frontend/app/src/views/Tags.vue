<template>
  <v-container grid-list-md>
    <v-layout align-center justify-center>
      <v-flex xs6 class="my-5">
        <v-card tile class="pa-3">
          <v-text-field v-model="inputedText" label="タグ名を入力" required></v-text-field>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="info" @click="addTag">追加</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
    <v-layout align-center justify-center>
      <v-flex xs12 class="my-5">
        <v-card tile class="pa-3">
          <span v-for="(tag_info, index) in tags" :key="index">
            <Tag :tag_info="tag_info" @deleteTag="deleteTag(index)" />
          </span>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import store from '../App';
import Tag from '../components/Tag';

export default {
  components: {
    Tag
  },
  data() {
    return {
      tags: [],
      inputedText: ''
    };
  },
  created() {
    const url = process.env.VUE_APP_API_URL + '/tags';
    this.requestJson(url, 'GET').then(json => (this.$data.tags = json));
  },
  methods: {
    requestJson(url, method, body = null) {
      const headers = { 'Content-Type': 'application/json' };
      return fetch(url, { method, headers, body })
        .then(response => response.json())
        .catch(err => console.error(err));
    },
    addTag() {
      const url = process.env.VUE_APP_API_URL + '/tags';
      const tagName = this.$data.inputedText;
      const body = JSON.stringify({ name: tagName });
      this.requestJson(url, 'POST', body).then(json => this.$data.tags.push(json));
    },
    deleteTag(index) {
      const url = process.env.VUE_APP_API_URL + '/tags/';
      const id = this.$data.tags[index].id;
      fetch(url + id, { method: 'DELETE' }).then(console.log('Success deleted tag.'));
    }
  }
};
</script>
