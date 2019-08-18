<template>
  <v-chip v-model="exist" close>{{ tag_info.name }}</v-chip>
</template>

<script>
export default {
  props: ['tag_info'],
  data() {
    return {
      exist: true
    };
  },
  watch: {
    exist() {
      this.$emit('deleteTag');
    }
  }
};
</script>

<style>
.card-text {
  min-height: 50px;
}
</style>
