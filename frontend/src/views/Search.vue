<template>
  <div>
    <div class="div_input">
      <el-input v-model="content" placeholder="搜索感兴趣的内容吧" style="float:left;width:90%">
        <el-select v-model="select" slot="prepend" style="width:80px;">
          <el-option label="问题" value="1"></el-option>
          <el-option label="用户" value="2"></el-option>
        </el-select>
        <i
          slot="suffix"
          class="el-input__icon el-icon-search"
          @click="handleSearch"
          style="cursor:pointer;"
        ></i>
      </el-input>
      <el-button style="float:right;" @click="showQuestion = true">提问</el-button>
    </div>

    <div v-if="searched" class="search_card">
      <el-card v-if="select==='1'" :body-style="{ padding: '10px' }">
        <div>
          <div v-for="questionInfo in questionInfoDatas" :key="questionInfo.question.questionId">
            <div
              v-html="redFont(questionInfo.question.title)"
              class="question"
              @click="toQuestion(questionInfo.question.questionId)"
            ></div>
            <div v-html="redFont(questionInfo.question.detail)" style="margin-top:10px;"></div>
            <el-divider class="divider"></el-divider>
          </div>
        </div>
        <el-pagination
          :current-page.sync="currentPage"
          :page-size="pageSize"
          layout="total,prev, pager, next, jumper"
          :total="length"
          style="text-align:center;"
        ></el-pagination>
      </el-card>
      <UserInfoTable v-if="select==='2'" :userInfos="userInfos" />
    </div>
    <el-dialog :visible.sync="showQuestion">
      <SaveQuestion v-if="showQuestion" />
    </el-dialog>
  </div>
</template>
<script>
import UserInfo from "../components/UserInfo";
import SaveQuestion from "../components/SaveQuestion";
import UserInfoTable from "../components/UserInfoTable";
import { searchUserInfos, searchQuestionInfos } from "@/api/search";
export default {
  data() {
    return {
      content: "",
      select: "1",
      questionInfos: [],
      userInfos: [],
      searched: false,
      loading: false,
      currentPage: 1,
      pageSize: 10,
      template: "",
      templates: [],
      showQuestion: false
    };
  },
  methods: {
    handleSearch() {
      this.searched = false;
      this.template = this.content.trim().toLowerCase();
      if (this.template === "") return;
      this.currentPage = 1;
      this.loading = true;
      if (this.select === "1") {
        searchQuestionInfos({
          title: this.template,
          userId: this.userId
        }).then(res => {
          this.questionInfos = res.questionInfos;
          this.templates = res.templates;
          this.searched = true;
          this.loading = false;
        });
      } else {
        searchUserInfos({
          nickname: this.template,
          userId: this.userId
        }).then(res => {
          this.userInfos = res;
          this.searched = true;
          this.loading = false;
        });
      }
    },
    redFont(raw_str) {
      if (raw_str == null) return null;
      var str = raw_str.toLowerCase();
      var index_last = 0;
      var res = "";
      while (true) {
        let template = null;
        let index_this = -1;
        this.templates.forEach(element => {
          let index_element = str.indexOf(element, index_last);
          if (
            index_this === -1 ||
            (index_element !== -1 && index_element < index_this)
          ) {
            index_this = index_element;
            template = element;
          }
        });
        if (index_this !== -1) {
          res += "<span>" + raw_str.slice(index_last, index_this) + "</span>";
          res +=
            '<span class="font_red">' +
            raw_str.slice(index_this, index_this + template.length) +
            "</span>";
          index_last = index_this + template.length;
        } else {
          res += "<span>" + raw_str.slice(index_last, str.length) + "</span>";
          break;
        }
      }
      return res;
    }
  },
  computed: {
    questionInfoDatas() {
      return this.questionInfos.slice(
        (this.currentPage - 1) * this.pageSize,
        this.currentPage * this.pageSize
      );
    },
    length() {
      if (this.select === "1") return this.questionInfos.length;
      else if (this.select === "2") return this.userInfos.length;
    },
    userId() {
      return this.$store.userID;
    }
  },
  components: {
    UserInfo,
    SaveQuestion,
    UserInfoTable
  }
};
</script>
<style>
.div_input {
  width: 800px;
  height: 40px;
}
.el-input-group__prepend {
  background-color: #ffffff;
}
.font_red {
  color: red;
}
.search_card {
  text-align: left;
  margin-top: 10px;
  clear: both;
}
.div_span_detail {
  font-size: 15px;
}
</style>
