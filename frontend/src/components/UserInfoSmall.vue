<template>
  <span class="userInfoSmall">
    <avater
      :userId="userInfo"
      class="userInfoSmall_avater"
    />
    <b
      class="userInfoSmall_text"
      :class="{'anonymous':userInfo.anonymous}"
    >{{userInfo}}</b>
  </span>
</template>
<script>
import avater from "../components/Avater";
export default {
  components: {
    avater
  },
  data() {
    return {};
  },
  name: "userInfoSmall",
  props: ["userInfo"]
};
</script>
<style>
.userInfoSmall {
  height: 30px;
  float: left;
}
.userInfoSmall_text {
  height: 30px;
  line-height: 30px;
  max-width: 150px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  float: left;
  margin-left: 10px;
}
.userInfoSmall_avater {
  width: 30px;
  border-radius: 50%;
  float: left;
}
</style>
