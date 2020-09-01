<template>
  <div style="text-align:left;">
    <div style="float:left;">
      <TopicTag v-for="topic in questionInfo.question.topics" :key="topic" :topic="topic" />
    </div>
    <div style="float:right;">
      <el-button
        v-if="questionInfo.subscribed"
        type="text"
        @click="handleSubscribe"
        icon="el-icon-star-on"
      ></el-button>
      <el-button v-else type="text" @click="handleSubscribe" icon="el-icon-star-off"></el-button>
      <el-button
        type="text"
        class="subscribe_btn"
        style="margin-left:0"
      >{{questionInfo.subscribeCount}}</el-button>
      <el-divider direction="vertical"></el-divider>
      <i class="el-icon-view"></i>
      <span class="viewed">{{questionInfo.viewCount}}</span>
      <el-divider direction="vertical"></el-divider>

      <el-button type="text" @click="showAnswerDrawer=true" icon="el-icon-edit">回答</el-button>
    </div>
    <h1 style="clear:both;">{{questionInfo.question.caption}}</h1>
    <div style="color:gray;font-size:10px;margin-top:10px;">问题描述：</div>
    <div v-html="questionInfo.question.content"></div>
    <el-card
      v-if="questionInfo.defaultAnswer != null||questionInfo.userAnswer!= null"
      style="margin-top:10px;"
      :body-style="{ padding: '10px' }"
    >
      <div v-if="questionInfo.defaultAnswer != null">
        <answer
          :answerInfo="questionInfo.defaultAnswer"
          :answered="false"
          :questioned="questionInfo.questioned"
          @handleDeleteAnswer="handleDeleteAnswer"
        ></answer>
      </div>
      <div v-if="questionInfo.userAnswer!= null">
        <answer
          :answerInfo="questionInfo.userAnswer"
          :answered="true"
          :questioned="questionInfo.questioned"
          @handleDeleteAnswer="handleDeleteAnswer"
          @preAnswer="showAnswerDrawer=true"
        ></answer>
      </div>
    </el-card>
    <el-card :body-style="{ padding: '10px' }">
      <el-divider
        class="divider"
      >{{answerInfos.length===0?'暂无其他回答':('更多'+answerInfos.length+'个回答')}}</el-divider>
      <template v-for="answerInfo in answerInfos">
        <div :key="answerInfo.answer.answerId" style="text-align:left;">
          <answer :answerInfo="answerInfo" :answered="false" :questioned="questionInfo.questioned" />
        </div>
      </template>
    </el-card>

    <el-drawer
      :visible.sync="showAnswerDrawer"
      size="70%"
      :direction="'btt'"
      style="text-align:center;"
    >
      <el-row type="flex" justify="center">
        <el-col :span="20">
          <el-input
            type="textarea"
            v-model="answer.ans"
            placeholder="请输入回答"
            :autosize="{ minRows: 8, maxRows:10}"
          ></el-input>
        </el-col>
      </el-row>
      <el-checkbox v-model="answer.anonymous">开启匿名</el-checkbox>
      <el-button type="primary" @click="handleAnswer">提交</el-button>
    </el-drawer>
    <el-dialog :visible.sync="showQuestion">
      <SaveQuestion
        v-if="showQuestion"
        :question="questionInfo.question"
        @updateQuestion="updateQuestion"
      />
    </el-dialog>
  </div>
</template>
<script>
import {
  getQuestion,
  saveSubscribe,
  deleteSubscribe,
  deleteQuestion
} from "@/api/question";
import {
  saveAnswer,
  updateAnswer,
  deleteAnswer,
  queryAnswers
} from "@/api/answer";
import Answer from "../components/Answer";
import UserInfo from "../components/UserInfo";
import SaveQuestion from "../components/SaveQuestion";
export default {
  data() {
    return {
      questionInfo: {
        questionID: this.$route.query.question_id,
        question: {
          anonymous: false,
          caption: null,
          content: null,
          gmtCreate: 0,
          gmtModify: null,
          userID: null
        },
        subscribed: false,
        subscribeCount: 0,
        viewCount: 0,
        questioned: false,
        defaultAnswer: null,
        userAnswer: null
      },
      answerSet: {},
      answerInfos: [],
      answer: {
        userID: null,
        questionID: null,
        answerID: null,
        ans: null,
        anonymous: false,
        gmtCreate: 0,
        gmtModify: null
      },
      showAnswerDrawer: false,
      showQuestion: false
    };
  },
  components: {
    Answer,
    UserInfo,
    SaveQuestion
  },
  created() {
    var questionID = this.$route.query.question_id;
    console.log(this.$route.query.question_id)
    if (questionID == null) return;
    getQuestion(questionID).then(res => {
      console.log(res.data)
      console.log(this.questionInfo)
      this.questionInfo.question = res.data;
      if (this.questionInfo.userAnswer != null) {
        this.answer = this.questionInfo.userAnswer.answer;
      }
      queryAnswers(this.$route.query.question_id).then(res => {
        console.log(res.data)
        var answerInfos = res.data;
        var old_length = this.answerInfos.total_count;
        for (var i = 0; i < answerInfos.length; ++i) {
          var answerInfo = answerInfos[i];
          if (
            this.questionInfo.defaultAnswer != null &&
            answerInfo.answer.answerId ===
              this.questionInfo.defaultAnswer.answer.answerId
          ) {
            //若已经是默认回答
            this.questionInfo.defaultAnswer = answerInfo;
          } else if (
            this.questionInfo.userAnswer != null &&
            answerInfo.answer.answerId ===
              this.questionInfo.userAnswer.answer.answerId
          ) {
            this.questionInfo.userAnswer = answerInfo;
          } else {
            var answerIdKey = "" + answerInfo.answer.answerId;
            if (this.answerSet.hasOwnProperty(answerIdKey)) {
              //set中已存在这个key，应该替换掉旧的
              this.answerInfos[this.answerSet[answerIdKey]] = answerInfo;
            } else {
              //不存在这个key
              this.answerInfos.push(answerInfo);
              this.answerSet[answerIdKey] = old_length++;
            }
          }
        }
      });
    });
  },
  methods: {
    handleDeleteAnswer() {
      var params = {
        answerId: this.answer.answerId
      };
      deleteAnswer(params).then(res => {
        if (res === true) {
          this.questionInfo.userAnswer = null;
          this.answer.answerId = null;
          this.$message({
            showClose: true,
            message: "删除成功",
            type: "success"
          });
        }
      });
    },
    handleAnswer() {
      this.answer.userId = this.userId;
      this.answer.questionId = this.questionInfo.question.questionId;
      if (this.answer.answerId != null) {
        updateAnswer(this.answer).then(res => {
          if (res === true) {
            this.answer.gmtModify = this.$nowTimestamp();
            this.questionInfo.userAnswer.answer = this.answer;
            this.questionInfo.userAnswer.userInfo = this.$userInfo(
              this.answer.anonymous
            );
            this.showAnswerDrawer = false;
          }
        });
      } else {
        saveAnswer(this.answer).then(res => {
          var answerId = Number(res);
          if (answerId > 0) {
            this.answer.answerId = answerId;
            this.answer.gmtCreate = this.$nowTimestamp();
            var answerInfo = {
              answer: this.answer,
              against: 0,
              agree: 0,
              attituded: null,
              userInfo: this.$userInfo(this.answer.anonymous)
            };
            this.questionInfo.userAnswer = answerInfo;
            this.showAnswerDrawer = false;
          }
        });
      }
    },
    updateQuestion(question) {
      this.questionInfo.question = question;
      this.showQuestion = false;
    },
    handleDeleteQuestion() {
      const params = {
        questionId: this.$route.query.questionId
      };
      deleteQuestion(params).then(res => {
        if (res === true) this.$router.push({ path: "/" });
      });
    }
  },
  computed: {
    userID() {
      return this.$store.getters.userID;
    }
  }
};
</script>

<style>
.title {
  clear: both;
  font-size: 30px;
  font-weight: bold;
}
.subscribe_btn {
  font-size: 17px;
}
.divider {
  margin: 10px 0 10px 0;
}
</style>
