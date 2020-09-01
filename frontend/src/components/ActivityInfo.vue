<template>
  <div style="text-align:left;">
    <div style="height:30px;line-height:30px;">
      <userInfoSmall v-if="userInfo" :userInfo="userInfo" style="margin-right:10px;"></userInfoSmall>
      <b style="float:left;">{{actionStr}}</b>
      <span
        class="gmt_create"
        style="float:right;"
      >{{activityInfo.activity.gmtCreate}}</span>
    </div>
    <div v-if="(!userInfo)&&(hasPeople||hasAnswer||hasQuestion)" style="height:30px;">
      <userInfoSmall v-if="hasPeople" :userInfo="activityInfo.info"></userInfoSmall>
      <userInfoSmall v-else-if="hasAnswer" :userInfo="activityInfo.info.defaultAnswer.userInfo"></userInfoSmall>
      <userInfoSmall v-else :userInfo="activityInfo.info.userInfo"></userInfoSmall>
    </div>
    <div
      v-if="hasQuestion"
      @click="toQuestion"
      class="question"
      style="margin-top:10px;"
    >{{activityInfo.info.caption}}</div>
    <AnswerInfo
      v-if="hasAnswer"
      style="margin-top:10px;"
      :answerInfo="activityInfo.info.defaultAnswer"
    />
  </div>
</template>
<script>
import userInfoSmall from "../components/UserInfoSmall";
import AnswerInfo from "../components/AnswerInfo";
export default {
  name: "ActivityInfo",
  components: {
    AnswerInfo,
    userInfoSmall
  },
  data() {
    return {
      actionStrs: [
        "关注了用户",
        "提出了问题",
        "关注了问题",
        "关注了话题",
        "回答了问题",
        "赞同了回答"
      ]
    };
  },
  props: ["activityInfo"],
  methods: {
    toQuestion() {
      window.open(
        this.$router.resolve({
          path: "/question",
          query: {
            question_id: this.activityInfo.info.question_id,
            answer_id: this.answer_id
          }
        }).href,
        "_blank"
      );
    }
  },
  computed: {
    actionStr() {
      return this.actionStrs[this.activityInfo.activity.act];
    },
    hasQuestion() {
      return [1, 2, 4, 5].indexOf(this.activityInfo.activity.act) != -1;
    },
    hasAnswer() {
      return [4, 5].indexOf(this.activityInfo.activity.act) != -1;
    },
    hasPeople() {
      return this.activityInfo.activity.act === 0;
    },
    userInfo() {
      return this.activityInfo.userInfo;
    },
    answerId() {
      return this.hasAnswer
        ? this.activityInfo.info.defaultAnswer.answer.answerId
        : null;
    }
  }
};
</script>
