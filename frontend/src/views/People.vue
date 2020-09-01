<template>
  <div v-if="userInfo!=null" style="text-align: left;">
    <div class="user_card" style="font-size:15px;">
      <div class="user_avater_border">
        <img :src="avaterUrl" class="user_avater" />
      </div>
      <div style="width:612px;float:left;padding: 10px 0 10px;">
        <el-dropdown :show-timeout="0" trigger="click" style="float: right;margin-right:20px;">
          <span>
            <i class="el-icon-more" style="color:gray;cursor: pointer;line-height:50px;"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item
              v-if="isSuperAdmin&&userInfo.user.authority===1"
              @click.native="handleUpdateAuthority(0)"
            >取消社区管理员</el-dropdown-item>
            <el-dropdown-item
              v-if="isSuperAdmin&&userInfo.user.authority===0"
              @click.native="handleUpdateAuthority(1)"
            >设为社区管理员</el-dropdown-item>
            <el-dropdown-item @click.native="handleFollow">{{userInfo.followed?'已':''}}关注</el-dropdown-item>
            <el-dropdown-item @click.native="handleOpenChat">发消息</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <div>
          <div style="float:right;text-align:center;width:200px;">
            <div class="follow_div" style="width:99px;border-right: 1px solid #ebebeb;">
              <div>关注{{sexStr}}</div>
              <el-tooltip :content="''+userInfo.followersCount" placement="left-start">
                <b class="follow_div_b">{{userInfo.followersCount}}</b>
              </el-tooltip>
            </div>
            <div class="follow_div" style="width:100px;">
              <div>{{sexStr}}关注</div>
              <el-tooltip :content="''+userInfo.followCount" placement="right-start">
                <b class="follow_div_b">{{userInfo.followCount}}</b>
              </el-tooltip>
            </div>
          </div>
          <b style="line-height:50px;height:50px;font-size:30px;">{{userInfo.user.nickname}}</b>
          <div class="user_detail">
            <div class="user_detail_value">{{userInfo.user.intro}}</div>
          </div>
        </div>
      </div>
    </div>
    <div class="activities_div_div user_card">
      <div style="text-align:center;border-bottom: 1px solid #ebebeb;">
        <el-menu :default-active="tabindex" mode="horizontal" @select="handleSelect">
          <el-menu-item v-for="(item,index) in menu" :key="index" :index="''+index">
            <span>{{item}}</span>
          </el-menu-item>
        </el-menu>
      </div>
      <div
        class="activities_div"
        v-show="tabindex==='0'"
        v-infinite-scroll="load"
        infinite-scroll-disabled="loading"
      >
        <div v-for="(activityInfo,index) in activityInfos" :key="index">
          <ActivityInfo :activityInfo="activityInfo" />
          <el-divider class="divider"></el-divider>
        </div>
      </div>
      <div v-if="tabindex==='1'||tabindex==='2'" class="activities_div">
        <div v-for="(activityInfo,index) in questionActivityInfos" :key="index">
          <ActivityInfo :activityInfo="activityInfo" />
          <el-divider class="divider"></el-divider>
        </div>
      </div>
      <div v-if="tabindex==='3'" class="activities_div">
        <div v-for="(activityInfo,index) in subscribeQuestions" :key="index">
          <ActivityInfo :activityInfo="activityInfo" />
          <el-divider class="divider"></el-divider>
        </div>
      </div>
      <div v-if="tabindex==='4'||tabindex==='5'">
        <UserInfoTable :userInfos="userInfos" />
      </div>
    </div>
    <el-dialog :visible.sync="showChat" :before-close="handleCloseChat" class="chat_dialog">
      <Chat v-if="showChat" :peopleUserInfo="userInfo" />
    </el-dialog>
  </div>
</template>
<script>
import {
  getUserInfo,
  queryPeopleActivities,
  saveFollow,
  deleteFollow,
  queryUserInfosByFollowerId,
  queryFollowerInfosByPeopleId,
  queryAnswersByUserId,
  queryQuestionsByUserId,
  updateAuthority,
  querySubscribeQuestionsByUserId
} from "@/api/user";
import Chat from "../components/Chat";
import UserInfoTable from "../components/UserInfoTable";
import ActivityInfo from "../components/ActivityInfo";
export default {
  name: "people",
  data() {
    return {
      userInfo: null,
      loading: false,
      subscribeQuestions: [],
      activityInfos: [],
      questionInfos: [],
      page: 0,
      limit: 10,
      colors: ["#DD0000", "brown", "#0000DD", "#48B753", "#DDA522", "green"],
      userInfos: [],
      tabindex: "0",
      firstLoadUserInfos: false,
      firstLoadFollowerInfos: false,
      showChat: false
    };
  },
  components: {
    ActivityInfo,
    Chat,
    UserInfoTable
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      if (this.peopleId != null) {
        this.userInfo = null;
        this.loading = false;
        this.page = 0;
        this.activityInfos = [];
        this.userInfos = [];
        this.firstLoadUserInfos = false;
        this.firstLoadFollowerInfos = false;
        getUserInfo({
          peopleId: this.peopleId
        }).then(res => {
          this.userInfo = res;
        });
      }
    },
    load() {
      if (this.loading) return;
      this.loading = true;
      var params = {
        peopleId: this.peopleId,
        page: this.page,
        limit: this.limit
      };
      queryPeopleActivities(params).then(res => {
        this.activityInfos = this.activityInfos.concat(res);
        ++this.page;
        if (res.length < this.limit) {
          console.log("暂时没有更多了");
          setTimeout(() => {
            this.loading = false;
            console.log("冷却完成");
          }, 1000);
        } else this.loading = false;
      });
    },
    handleFollow() {
      const params = {
        peopleId: this.peopleId
      };
      if (this.userInfo.followed) {
        deleteFollow(params).then(res => {
          if (res === true) {
            --this.userInfo.followersCount;
            this.userInfo.followed = false;
          }
        });
      } else {
        saveFollow(params).then(res => {
          if (res === true) {
            ++this.userInfo.followersCount;
            this.userInfo.followed = true;
          }
        });
      }
    },
    handleSelect(key, keyPath) {
      this.tabindex = key;
    },
    handleOpenChat() {
      this.$store.commit("readMessage");
      this.showChat = true;
    },
    handleCloseChat() {
      this.$store.commit("closeMessage");
      this.showChat = false;
    },
    handleUpdateAuthority(authority) {
      const data = {
        userId: this.peopleId,
        authority: authority
      };
      updateAuthority(data).then(res => {
        if (res === true) this.userInfo.user.authority = authority;
      });
    }
  },
  computed: {
    questionActivityInfos() {
      const activityInfos = [];
      this.questionInfos.forEach(questionInfo => {
        const activityInfo = this.$questionInfo2ActvityInfo(questionInfo);
        activityInfos.push(activityInfo);
      });
      return activityInfos;
    },
    isSuperAdmin() {
      return this.$store.getters.authority === 2;
    },
    avaterUrl() {
      return this.$avaterUrl(
        this.userInfo.user.avater,
        this.userInfo.user.userId
      );
    },
    userId() {
      return this.$store.getters.userId;
    },
    sexStr() {
      if (this.isMe) return "我";
      switch (this.userInfo.user.sex) {
        case 0:
          return "他";
        case 1:
          return "她";
        default:
          return "ta";
      }
    },
    peopleId() {
      return Number(this.$route.params.peopleId);
    },
    isMe() {
      return this.peopleId === this.userId;
    },
    menu() {
      return [
        "动态",
        "回答",
        "提问",
        "订阅",
        "关注" + this.sexStr + "的",
        this.sexStr + "关注的"
      ];
    }
  },
  watch: {
    tabindex(val) {
      const params = {
        peopleId: this.peopleId
      };
      this.page = 0;
      if (val === "1") {
        queryAnswersByUserId(params).then(res => {
          this.questionInfos = res;
        });
      } else if (val === "2") {
        queryQuestionsByUserId(params).then(res => {
          this.questionInfos = res;
        });
      } else if (val === "3") {
        querySubscribeQuestionsByUserId(params).then(res => {
          this.subscribeQuestions = res;
        });
      } else if (val === "4") {
        queryFollowerInfosByPeopleId(params).then(res => {
          this.userInfos = res;
        });
      } else if (val === "5") {
        queryUserInfosByFollowerId(params).then(res => {
          this.userInfos = res;
        });
      } else {
        this.userInfos = [];
      }
    },
    peopleId(val) {
      this.init();
    }
  }
};
</script>
<style>
.user_avater {
  width: 160px;
  padding: 4px;
  height: 160px;
  border-radius: 5%;
}
.user_avater_border {
  margin: 10px;
  background-color: gray;
  width: 168px;
  height: 168px;
  border-radius: 5%;
  float: left;
}
.user_card {
  background-color: #ffffff;
  float: left;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}
.user_detail {
  float: left;
  text-align: left;
  width: 532px;
}
.user_detail > * {
  float: left;
  clear: left;
  margin-top: 10px;
}
.activities_div {
  max-height: 580px;
  overflow: auto;
  padding: 10px;
}
.activities_div_div {
  width: 800px;
  margin-top: 10px;
  float: left;
}
.activities_div_div > ::-webkit-scrollbar {
  width: 0;
  background-color: transparent;
}
.follow_div {
  float: left;
  height: 50px;
}
.follow_div_b {
  cursor: pointer;
  font-size: 20px;
  height: 30px;
  line-height: 30px;
}
</style>
