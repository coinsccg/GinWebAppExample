<template>
  <div>
    <el-tag>{{topic}}</el-tag>
    <el-button
      @click="handleChangeSubscribe"
    >{{topicInfo.subscribeCount}} {{topicInfo.subscribed?"取消订阅":"订阅话题"}}</el-button>
    <el-card :body-style="{ padding: '10px' }">
      <div v-for="(activityInfo,index) in activityInfos" :key="index">
        <ActivityInfo :activityInfo="activityInfo" />
        <el-divider class="divider"></el-divider>
      </div>
      <el-pagination
        :current-page.sync="currentPage"
        :page-size="pageSize"
        layout="total,prev,pager,next,jumper"
        :total="topicInfo.questionInfos.length"
        style="text-align:center;"
      ></el-pagination>
    </el-card>
  </div>
</template>
<script>
import {
  saveSubscribeTopic,
  deleteSubscribeTopic,
  getTopicInfo
} from "@/api/topic";
import ActivityInfo from "../components/ActivityInfo";
export default {
  name: "Topic",
  components: {
    ActivityInfo
  },
  data() {
    return {
      topicInfo: { subscribeCount: 0, subscribed: false, questionInfos: [] },
      pageSize: 5,
      currentPage: 1
    };
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      getTopicInfo({ topic: this.topic }).then(res => {
        this.topicInfo = res;
      });
    },
    handleChangeSubscribe() {
      const params = { topic: this.topic };
      if (this.topicInfo.subscribed) {
        deleteSubscribeTopic(params).then(res => {
          if (res === true) {
            this.topicInfo.subscribed = false;
            --this.topicInfo.subscribeCount;
          }
        });
      } else {
        saveSubscribeTopic(params).then(res => {
          if (res === true) {
            this.topicInfo.subscribed = true;
            ++this.topicInfo.subscribeCount;
          }
        });
      }
    }
  },
  computed: {
    topic() {
      return this.$route.params.topic;
    },
    activityInfos() {
      const activityInfos = [];
      this.topicInfo.questionInfos
        .slice(
          (this.currentPage - 1) * this.pageSize,
          this.currentPage * this.pageSize
        )
        .forEach(questionInfo => {
          activityInfos.push(this.$questionInfo2ActvityInfo(questionInfo));
        });
      return activityInfos;
    }
  },
  watch: {
    topic(val) {
      this.init();
    }
  }
};
</script>
<style scoped>
</style>