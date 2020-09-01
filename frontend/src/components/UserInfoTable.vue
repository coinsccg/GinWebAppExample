<template>
  <div>
    <el-card :body-style="{ padding: '10px' }">
      <div v-for="userInfo in userInfoDatas" :key="userInfo.user.userId">
        <userInfo :userInfo="userInfo"></userInfo>
        <el-divider class="divider"></el-divider>
      </div>
      <el-pagination
        :current-page.sync="currentPage"
        :page-size="pageSize"
        layout="total,prev, pager, next, jumper"
        :total="userInfos.length"
        style="text-align:center;"
      ></el-pagination>
    </el-card>
  </div>
</template>
<script>
import UserInfo from "../components/UserInfo";
import { getUserInfo } from "@/api/user";
export default {
  name: "UserInfoTable",
  props: ["userInfos"],
  data() {
    return {
      pageSize: 10,
      currentPage: 1
    };
  },
  created() {
    console.log(this.userInfos);
  },
  computed: {
    userInfoDatas() {
      return this.userInfos.slice(
        (this.currentPage - 1) * this.pageSize,
        this.currentPage * this.pageSize
      );
    }
  },
  components: {
    UserInfo
  }
};
</script>
<style scoped>
</style>