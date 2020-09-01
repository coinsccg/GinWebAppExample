<template>
  <div>
    <div>
      <userInfo :userInfo="answerInfo.userInfo"></userInfo>
    </div>
    <div v-html="answer.ans" class="answer"></div>
    <div
      class="gmt_create"
      style="margin-top:10px;"
    >回答于 {{$getTimeString(answer.gmtModify!=null?answer.gmtModify:answer.gmtCreate)}}</div>
    <div style="margin-top:10px">
      <attitude :answerInfo="answerInfo"></attitude>
      <span style="margin-left:10px;">
        <el-button
          type="text"
          @click="preAnswer"
          v-if="answered"
          icon="el-icon-edit-outline"
          style="color:gray;"
        >修改</el-button>
        <el-button
          type="text"
          @click="handleDeleteAnswer"
          v-if="answered||isAdmin"
          icon="el-icon-delete"
          style="color:red;"
        >删除</el-button>
      </span>
    </div>
    <el-divider class="divider"></el-divider>
  </div>
</template>
<script>
import userInfo from "../components/UserInfo";
import attitude from "../components/Attitude";
export default {
  name: "answer",
  components: {
    userInfo,
    attitude
  },
  data() {
    return {
      answer: this.answerInfo.answer
    };
  },
  created() {},
  methods: {
    handleDeleteAnswer() {
      this.$emit("handleDeleteAnswer");
    },
    preAnswer() {
      this.$emit("preAnswer");
    }
  },
  computed: {
    isAdmin() {
      return this.$store.getters.authority > 1;
    }
  },
  props: ["answerInfo", "answered", "questioned"]
};
</script>

<style>
.answer {
  clear: both;
  text-align: left;
  min-height: 200px;
  overflow: hidden;
  margin-top: 10px;
  /* color: brown;
  text-decoration: line-through; */
}
</style>