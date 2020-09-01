<template>
  <div>
    <el-tabs type="border-card" v-model="tabIndex" style="margin-bottom:10px;">
      <el-tab-pane label="热榜" name="0">
        <div
          style="overflow:auto;"
          v-infinite-scroll="load0"
          infinite-scroll-disabled="loading0"
          :style="tabStyle"
        >
          <div v-for="questionInfo in questionList" :key="questionInfo.question_id">
            <ActivityInfo :activityInfo="$questionInfo2ActvityInfo(questionInfo)"></ActivityInfo>
            <el-divider class="divider"></el-divider>
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane label="关注" name="1">
        <div>
          关注~~
        </div>
      </el-tab-pane>
      <el-tab-pane label="话题" name="2">
        <div>
         话题~~
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import { queryQuestions } from "@/api/question";
import ActivityInfo from "../components/ActivityInfo";
export default {
  data() {
    return {
      limit: 10,
      tabIndex: "0",
      tabStyle: {
        maxHeight: window.innerHeight - 131 + "px"
      },
      loading0: false,
      questionList: [],
      page0: 0
    };
  },
   components: {
    ActivityInfo
  },
  methods: {
    load0() {
      if (this.loading0) return;
      this.loading0 = true;
      var params = {
        page: this.page0,
        limit: this.limit
      };
      queryQuestions(params).then(res => {
        console.log(res.data)
        this.questionList = res.data.question_list;
        ++this.page0;
        if (res.data.total_count < this.limit) {
          this.$message({
            message: "暂时没有更多了"
          });
          setTimeout(() => {
            this.loading0 = false;
            console.log("冷却完成");
          }, 1000);
        } else this.loading0 = false;
      });
    }
  },
  mounted() {
    window.onresize = () => {
      this.tabStyle.maxHeight = window.innerHeight - 131 + "px";
    };
  },
  computed: {}
};
</script>
<style>
.el-tabs--border-card > .el-tabs__content > * ::-webkit-scrollbar {
  width: 0;
  height: 0;
  background-color: transparent;
}
.el-tabs--border-card > .el-tabs__content {
  padding: 10px;
}
.el-tabs--border-card > .el-tabs__header {
  background-color: #ffffff;
}
</style>
