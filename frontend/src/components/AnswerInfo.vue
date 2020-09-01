<template>
  <div>
    <div
      ref="answer"
      class="questionInfo_answer"
      :style="answerStyle"
      v-html="answerInfo.answer.ans"
    ></div>
    <div v-show="!isShowAll" class="hide_all">
      <el-button @click="showAll" type="text" class="showAll_btn" icon="el-icon-arrow-down">查看全文</el-button>
    </div>
    <div style="margin-top:10px;">
      <attitude :answerInfo="answerInfo"></attitude>
    </div>
  </div>
</template>
<script>
import attitude from "../components/Attitude";
export default {
  name: "AnswerInfo",
  components: {
    attitude
  },
  props: ["answerInfo"],
  data() {
    return {
      isShowAll: false,
      answerStyle: {
        maxHeight: "200px"
      }
    };
  },
  mounted() {
    var obj = this.$refs.answer;
    if (obj.scrollHeight < 200) {
      this.isShowAll = true;
    }
  },
  methods: {
    showAll() {
      this.isShowAll = true;
      var obj = this.$refs.answer;
      this.answerStyle.maxHeight = obj.scrollHeight + "px";
    }
  }
};
</script>
<style scoped>
.questionInfo_answer {
  clear: both;
  overflow: hidden;
  text-overflow: ellipsis;
}
.showAll_btn {
  width: 100%;
  color: black;
  text-align: center;
}
.hide_all {
  position: relative;
  height: 50px;
  width: 100%;
  margin-top: -50px;
  margin-left: 0px;
  background-image: linear-gradient(
    -180deg,
    rgba(255, 255, 255, 0) 0%,
    #fff 70%
  );
}
</style>