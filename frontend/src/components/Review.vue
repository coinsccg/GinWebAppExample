<template>
  <div>
    <div style="height:30px;">
      <userInfoSmall style="float:left;" :userInfo="reviewInfo.userInfo"></userInfoSmall>
      <span v-if="reply_userInfo!=null" style="float:left;">
        <i class="el-icon-arrow-right reply_i"></i>
        <userInfoSmall style="margin-left:20px;" :userInfo="reply_userInfo"></userInfoSmall>
      </span>
      <div class="gmt_create review_time">
        <span>{{reviewInfo.questioned?"(提问者)":" "}}</span>
        <span>{{reviewInfo.answered?"(回答者)":" "}}</span>
        <span>{{(review.replyId!=null?'回复':'评论')+'于 '+$getTimeString(this.review.gmtCreate)}}</span>
      </div>
    </div>
    <div
      style="clear:left;margin-top:10px;"
      :class="{'review_deleted':review.revi===null}"
    >{{review.revi!=null?review.revi:'该评论已删除'}}</div>

    <div>
      <el-button
        type="text"
        @click="handleUpdateApprove"
        :class="{'approved':!reviewInfo.approved}"
        :loading="approveLoading"
      >赞{{reviewInfo.approveCount}}</el-button>
      <el-button type="text" @click="reply" style="color:gray;" icon="el-icon-chat-dot-round">回复</el-button>
      <el-button
        type="text"
        @click="handleDeleteReview"
        v-if="canDelete"
        style="color:red;"
        icon="el-icon-delete"
      >删除</el-button>
    </div>
    <el-divider class="review_divider" content-position="right"></el-divider>
  </div>
</template>
<script>
import { deleteReview, updateApprove } from "@/api/review";
import userInfoSmall from "../components/UserInfoSmall";
export default {
  name: "review",
  data() {
    return {
      review: this.reviewInfo.review,
      approveLoading: false
    };
  },
  props: ["reviewInfo", "reply_userInfo"],
  computed: {
    reviewed() {
      return this.userId != null && this.userId == this.review.userId;
    },
    canDelete() {
      return (this.reviewed || this.isAdmin) && !this.review.deleted;
    },
    userId() {
      return this.$store.getters.userId;
    },
    isAdmin() {
      return this.$store.getters.authority > 1;
    }
  },
  created() {},
  methods: {
    reply() {
      this.$emit("reply", this.review.reviewId);
    },
    handleDeleteReview() {
      var params = {
        reviewId: this.review.reviewId
      };
      deleteReview(params).then(res => {
        if (res === true) {
          this.review.revi = null;
          this.review.deleted = true;
        }
      });
    },
    handleUpdateApprove() {
      this.approveLoading = true;
      var approve = !this.reviewInfo.approved;
      var params = {
        reviewId: this.review.reviewId,
        approve: approve
      };
      updateApprove(params).then(res => {
        if (res === true) {
          this.reviewInfo.approved = !this.reviewInfo.approved;
          if (approve) ++this.reviewInfo.approveCount;
          else --this.reviewInfo.approveCount;
        }
        this.approveLoading = false;
      });
    }
  },
  components: {
    userInfoSmall
  }
};
</script>
<style scoped>
.review_deleted {
  color: gray;
}
.approved {
  color: gray;
}
.review_time {
  float: right;
  text-align: left;
  line-height: 30px;
}
.review_divider {
  margin: 0 0 10px 0;
}
.reply_i {
  float: left;
  margin-left: 20px;
  line-height: 30px;
  font-weight: bold;
}
</style>