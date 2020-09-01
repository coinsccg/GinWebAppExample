<template>
  <el-tag class="topic" @click="toTopic">{{topic}}</el-tag>
</template>
<script>
export default {
  name: "TopicTag",
  props: ["topic"],
  methods: {
    toTopic() {
      this.$router.push({ path: "/topic/" + this.topic });
    }
  }
};
</script>
<style scoped>
.topic {
  margin-right: 10px;
  cursor: pointer;
}
</style>