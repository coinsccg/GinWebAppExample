<template>
  <img src="https://pic4.zhimg.com/aadd7b895_xs.jpg" @click="toPeople" style="cursor: pointer;" />
</template>
<script>
export default {
  name: "avater",
  props: ["userId", "avater"],
  computed: {
    avaterUrl() {
      return this.$avaterUrl(this.avater, this.userId);
    }
  },
  methods: {
    toPeople() {
      if (this.userId != null) {
        this.$router.push({
          path: "/people/" + this.userId
        });
      }
    }
  }
};
</script>
