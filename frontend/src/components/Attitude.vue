<template>
  <span style="height: 40px;margin-top:4px;">
    <el-button
      size="small"
      @click="handleUpdateAttitude(true)"
      icon="el-icon-caret-top"
      type="primary"
      plain
    >{{getCountString(agree)}} {{attituded===true?"已赞同":''}}</el-button>
    <el-button
      size="small"
      @click="handleUpdateAttitude(false)"
      icon="el-icon-caret-bottom"
      style="margin-left:0px;"
      type="danger"
      plain
    >{{getCountString(against)}} {{attituded===false?"已反对":''}}</el-button>
    <el-button
      type="text"
      :loading="reviewLoading"
      @click="handleQueryReviews"
      icon="el-icon-chat-round"
      style="color:gray;margin-left:10px;"
    >{{reviewCount}}</el-button>
    <el-dialog :visible.sync="showReview" class="review_dialog">
      <div slot="title" class="review_title">
        <b>共{{reviewCount}}条评论</b>
      </div>
      <el-card :body-style="{ padding: '10px' }" v-if="showReview">
        <div style="height:600px;overflow:auto;">
          <template v-for="reviewInfo in reviews">
            <div :key="reviewInfo.review.reviewId">
              <review
                :reviewInfo="reviewInfo"
                :reply_userInfo="getUserInfo(reviewInfo.review.replyId)"
                @reply="reply"
              ></review>
            </div>
          </template>
        </div>
        <div style="margin-top:10px;">
          <el-input style="width:420px;" v-model="revi" :placeholder="ReplyInfo"></el-input>
          <el-button
            style="float:right;margin-bottom:10px;"
            type="primary"
            @click="handleSaveReview"
            plain
          >评论</el-button>
          <el-button
            type="danger"
            @click="replyRId=null"
            v-show="replyRId!=null"
            style="float:right;margin:0 10px 0 0;"
            plain
          >取消</el-button>
        </div>
      </el-card>
    </el-dialog>
  </span>
</template>
<script>
import review from "../components/Review";
import { updateAttitude, deleteAttitude } from "@/api/answer";
import { queryReviews, saveReview } from "@/api/review";
export default {
  name: "attitude",
  components: {
    review
  },
  computed: {
    ReplyInfo() {
      var notReply = this.replyRId === null;
      return (
        (notReply ? "评论" : "回复") +
        "给" +
        (notReply
          ? this.answerInfo.userInfo.user.nickname
          : this.getUserInfo(this.replyRId).user.nickname)
      );
    },
    reviewSet() {
      var set = {};
      this.reviews.forEach(element => {
        set[element.review.reviewId] = element.userInfo;
      });
      return set;
    }
  },
  props: ["answerInfo"],
  data() {
    return {
      attituded: this.answerInfo.attituded,
      agree: this.answerInfo.agree,
      against: this.answerInfo.against,
      showReview: false,
      reviews: [],
      revi: "",
      replyRId: null,
      reviewLoading: false,
      reviewCount: this.answerInfo.reviewCount,
      center: true
    };
  },
  methods: {
    handleQueryReviews() {
      this.showReview = !this.showReview;
      if (this.showReview === false) return;
      this.reviewLoading = true;
      queryReviews({
        answerId: this.answerInfo.answer.answerId
      }).then(res => {
        this.reviews = res;
        this.reviewLoading = false;
      });
    },
    handleSaveReview() {
      var reviewForm = {
        answerId: this.answerInfo.answer.answerId,
        replyRId: this.replyRId,
        revi: this.revi
      };
      saveReview(reviewForm).then(res => {
        var reviewId = Number(res);
        if (reviewId > 0) {
          this.revi = "";
          reviewForm["reviewId"] = reviewId;
          reviewForm["gmtCreate"] = this.$nowTimestamp();
          var reviewInfo = {
            review: reviewForm,
            userInfo: this.$userInfo(false),
            anonymous: true,
            approved: false,
            approveCount: 0,
            questioned: this.questioned && !this.answered,
            answered: this.answered
          };
          this.reviews.push(reviewInfo);
        }
      });
    },
    reply(replyRId) {
      this.replyRId = replyRId;
    },
    getUserInfo(reviewId) {
      if (reviewId != null) return this.reviewSet[reviewId];
      else return null;
    },
    getCountString(count) {
      return this.$getCountString(count);
    },
    handleUpdateAttitude(atti) {
      var params = {
        answerId: this.answerInfo.answer.answerId
      };
      if (this.attituded === atti) {
        deleteAttitude(params).then(res => {
          if (res === true) {
            this.attituded = null;
            if (atti) {
              --this.agree;
            } else {
              --this.against;
            }
          }
        });
      } else {
        params["atti"] = atti;
        updateAttitude(params).then(res => {
          if (res === true) {
            if (atti) {
              if (this.attituded === false) --this.against;
              ++this.agree;
              this.attituded = true;
            } else {
              if (this.attituded === true) --this.agree;
              ++this.against;
              this.attituded = false;
            }
          }
        });
      }
    }
  }
};
</script>
<style>
.el-card__body > ::-webkit-scrollbar {
  width: 0;
  height: 0;
  background-color: transparent;
}
.review_title {
  height: 56px;
  line-height: 56px;
  text-align: center;
}
.review_dialog > .el-dialog > * {
  padding: 0;
}
.review_dialog > div > .el-dialog__header {
  height: 56px;
}
</style>