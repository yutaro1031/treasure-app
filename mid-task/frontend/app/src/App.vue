<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <span>俺の書籍管理</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn flat @click="$router.push({ path: 'search' })" target="_blank">
        <span class="mr-2">俺の書籍を増やす</span>
      </v-btn>
    </v-toolbar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
var store = {
  debug: true,
  state: {
    evaluationData: []
  }
};
export default store;
</script>
